// Copyright 2018-2023 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package ocdav

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strings"

	gatewayv1beta1 "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	userv1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	rpc "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	"github.com/cs3org/reva/pkg/appctx"
	ctxpkg "github.com/cs3org/reva/pkg/ctx"
	"github.com/cs3org/reva/pkg/rgrpc/todo/pool"
	"github.com/cs3org/reva/pkg/rhttp"
	"github.com/cs3org/reva/pkg/rhttp/middlewares"
	"github.com/cs3org/reva/pkg/rhttp/mux"
	"google.golang.org/grpc/metadata"
)

type tokenStatInfoKey struct{}

// DavHandler routes to the different sub handlers.
type DavHandler struct {
	AvatarsHandler      *AvatarsHandler
	FilesHandler        *WebDavHandler
	FilesHomeHandler    *WebDavHandler
	MetaHandler         *MetaHandler
	TrashbinHandler     *TrashbinHandler
	SpacesHandler       *SpacesHandler
	PublicFolderHandler *WebDavHandler
	PublicFileHandler   *PublicFileHandler
	OCMSharesHandler    *WebDavHandler

	router mux.Router
}

func (h *DavHandler) init(s *svc) error {
	c := s.c
	h.AvatarsHandler = new(AvatarsHandler)
	if err := h.AvatarsHandler.init(c); err != nil {
		return err
	}
	h.FilesHandler = new(WebDavHandler)
	if err := h.FilesHandler.init(c.FilesNamespace, false); err != nil {
		return err
	}
	h.FilesHomeHandler = new(WebDavHandler)
	if err := h.FilesHomeHandler.init(c.WebdavNamespace, true); err != nil {
		return err
	}
	h.MetaHandler = new(MetaHandler)
	if err := h.MetaHandler.init(c); err != nil {
		return err
	}
	h.TrashbinHandler = new(TrashbinHandler)

	h.SpacesHandler = new(SpacesHandler)
	if err := h.SpacesHandler.init(c); err != nil {
		return err
	}

	h.PublicFolderHandler = new(WebDavHandler)
	if err := h.PublicFolderHandler.init("public", true); err != nil { // jail public file requests to /public/ prefix
		return err
	}

	h.PublicFileHandler = new(PublicFileHandler)
	if err := h.PublicFileHandler.init("public"); err != nil { // jail public file requests to /public/ prefix
		return err
	}

	h.OCMSharesHandler = new(WebDavHandler)
	if err := h.OCMSharesHandler.init(c.OCMNamespace, false); err != nil {
		return err
	}

	if err := h.TrashbinHandler.init(c); err != nil {
		return err
	}

	h.initRouter(s)
	return nil
}

func (h *DavHandler) initRouter(s *svc) {
	h.router = mux.NewServeMux()
	h.router.Route("/", func(r mux.Router) {
		r.Handle("avatars/*", h.AvatarsHandler.Handler(s))
		r.Handle("files/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			_, r.URL.Path = rhttp.ShiftPath(r.URL.Path)

			var requestUserID string
			var oldPath = r.URL.Path

			// detect and check current user in URL
			requestUserID, r.URL.Path = rhttp.ShiftPath(r.URL.Path)

			// note: some requests like OPTIONS don't forward the user
			contextUser, ok := ctxpkg.ContextGetUser(ctx)
			if ok && isOwner(requestUserID, contextUser) {
				// use home storage handler when user was detected
				base := path.Join(ctx.Value(ctxKeyBaseURI).(string), "files", requestUserID)
				ctx := context.WithValue(ctx, ctxKeyBaseURI, base)
				r = r.WithContext(ctx)

				h.FilesHomeHandler.Handler(s).ServeHTTP(w, r)
			} else {
				r.URL.Path = oldPath
				base := path.Join(ctx.Value(ctxKeyBaseURI).(string), "files")
				ctx := context.WithValue(ctx, ctxKeyBaseURI, base)
				r = r.WithContext(ctx)

				h.FilesHandler.Handler(s).ServeHTTP(w, r)
			}
		}), mux.WithMiddleware(prependUsernameFiles()))
		r.Handle("meta/*", h.MetaHandler.Handler(s), mux.WithMiddleware(keyBase("meta")))
		r.Handle("trash-bin/*", h.TrashbinHandler.Handler(s), mux.WithMiddleware(keyBase("trash-bin")))
		r.Handle("spaces/*", h.SpacesHandler.Handler(s), mux.WithMiddleware(keyBase("spaces")))
		r.Route("ocm", func(r mux.Router) {
			r.Mount("/", h.OCMSharesHandler.Handler(s))
		}, mux.Unprotected(), mux.WithMiddleware(authenticateOCM(s)), mux.WithMiddleware(keyBase("ocm")))
		r.Handle("public-files/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := appctx.GetLogger(ctx)

			_, r.URL.Path = rhttp.ShiftPath(r.URL.Path)

			token, _ := rhttp.ShiftPath(r.URL.Path)
			c, err := pool.GetGatewayServiceClient(pool.Endpoint(s.c.GatewaySvc))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			// the public share manager knew the token, but does the referenced target still exist?
			sRes, err := getTokenStatInfo(ctx, c, token)
			switch {
			case err != nil:
				log.Error().Err(err).Msg("error sending grpc stat request")
				w.WriteHeader(http.StatusInternalServerError)
				return
			case sRes.Status.Code == rpc.Code_CODE_PERMISSION_DENIED:
				fallthrough
			case sRes.Status.Code == rpc.Code_CODE_NOT_FOUND:
				log.Debug().Str("token", token).Interface("status", sRes.Status).Msg("resource not found")
				w.WriteHeader(http.StatusNotFound) // log the difference
				return
			case sRes.Status.Code == rpc.Code_CODE_UNAUTHENTICATED:
				log.Debug().Str("token", token).Interface("status", sRes.Status).Msg("unauthorized")
				w.WriteHeader(http.StatusUnauthorized)
				return
			case sRes.Status.Code != rpc.Code_CODE_OK:
				log.Error().Str("token", token).Interface("status", sRes.Status).Msg("grpc stat request failed")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			log.Debug().Interface("statInfo", sRes.Info).Msg("Stat info from public link token path")

			if sRes.Info.Type != provider.ResourceType_RESOURCE_TYPE_CONTAINER {
				ctx := context.WithValue(ctx, tokenStatInfoKey{}, sRes.Info)
				r = r.WithContext(ctx)
				h.PublicFileHandler.Handler(s).ServeHTTP(w, r)
			} else {
				h.PublicFolderHandler.Handler(s).ServeHTTP(w, r)
			}
		}), mux.Unprotected(), mux.WithMiddleware(authenticatePublicLink(s)), mux.WithMiddleware(keyBase("public-files")))
	})
}

func isOwner(userIDorName string, user *userv1beta1.User) bool {
	return userIDorName != "" && (userIDorName == user.Id.OpaqueId || strings.EqualFold(userIDorName, user.Username))
}

func prependUsernameFiles() middlewares.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := appctx.GetLogger(ctx)

			// if there is no file in the request url we assume the request url is: "/remote.php/dav/files"
			// https://github.com/owncloud/core/blob/18475dac812064b21dabcc50f25ef3ffe55691a5/tests/acceptance/features/apiWebdavOperations/propfind.feature
			if r.URL.Path == "/files" {
				log.Debug().Str("path", r.URL.Path).Msg("method not allowed")
				contextUser, ok := ctxpkg.ContextGetUser(ctx)
				if ok {
					r.URL.Path = path.Join(r.URL.Path, contextUser.Username)
				}

				if r.Header.Get("Depth") == "" {
					w.WriteHeader(http.StatusMethodNotAllowed)
					b, err := Marshal(exception{
						code:    SabredavMethodNotAllowed,
						message: "Listing members of this collection is disabled",
					})
					if err != nil {
						log.Error().Msgf("error marshaling xml response: %s", b)
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
					_, err = w.Write(b)
					if err != nil {
						log.Error().Msgf("error writing xml response: %s", b)
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

func authenticateOCM(s *svc) middlewares.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := appctx.GetLogger(ctx)

			c, err := pool.GetGatewayServiceClient(pool.Endpoint(s.c.GatewaySvc))
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			_, r.URL.Path = rhttp.ShiftPath(r.URL.Path)

			// OC10 and Nextcloud (OCM 1.0) are using basic auth for carrying the
			// shared token.
			var token string

			username, _, ok := r.BasicAuth()
			if ok {
				// OCM 1.0
				token = username
				r.URL.Path, _ = url.JoinPath("/", token, r.URL.Path)
				ctx = context.WithValue(ctx, ctxOCM10, true)
			} else {
				token, _ = rhttp.ShiftPath(r.URL.Path)
				ctx = context.WithValue(ctx, ctxOCM10, false)
			}

			authRes, err := handleOCMAuth(ctx, c, token)
			switch {
			case err != nil:
				log.Error().Err(err).Msg("error during ocm authentication")
				w.WriteHeader(http.StatusInternalServerError)
				return
			case authRes.Status.Code == rpc.Code_CODE_PERMISSION_DENIED:
				log.Debug().Str("token", token).Msg("permission denied")
				fallthrough
			case authRes.Status.Code == rpc.Code_CODE_UNAUTHENTICATED:
				log.Debug().Str("token", token).Msg("unauthorized")
				w.WriteHeader(http.StatusUnauthorized)
				return
			case authRes.Status.Code == rpc.Code_CODE_NOT_FOUND:
				log.Debug().Str("token", token).Msg("not found")
				w.WriteHeader(http.StatusNotFound)
				return
			case authRes.Status.Code != rpc.Code_CODE_OK:
				log.Error().Str("token", token).Interface("status", authRes.Status).Msg("grpc auth request failed")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			ctx = ctxpkg.ContextSetToken(ctx, authRes.Token)
			ctx = ctxpkg.ContextSetUser(ctx, authRes.User)
			ctx = metadata.AppendToOutgoingContext(ctx, ctxpkg.TokenHeader, authRes.Token)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func authenticatePublicLink(s *svc) middlewares.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			c, err := pool.GetGatewayServiceClient(pool.Endpoint(s.c.GatewaySvc))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var res *gatewayv1beta1.AuthenticateResponse
			token, _ := rhttp.ShiftPath(r.URL.Path)
			if _, pass, ok := r.BasicAuth(); ok {
				res, err = handleBasicAuth(r.Context(), c, token, pass)
			} else {
				q := r.URL.Query()
				sig := q.Get("signature")
				expiration := q.Get("expiration")
				// We restrict the pre-signed urls to downloads.
				if sig != "" && expiration != "" && r.Method != http.MethodGet {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				res, err = handleSignatureAuth(r.Context(), c, token, sig, expiration)
			}

			switch {
			case err != nil:
				w.WriteHeader(http.StatusInternalServerError)
				return
			case res.Status.Code == rpc.Code_CODE_PERMISSION_DENIED:
				fallthrough
			case res.Status.Code == rpc.Code_CODE_UNAUTHENTICATED:
				w.WriteHeader(http.StatusUnauthorized)
				return
			case res.Status.Code == rpc.Code_CODE_NOT_FOUND:
				w.WriteHeader(http.StatusNotFound)
				return
			case res.Status.Code != rpc.Code_CODE_OK:
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			ctx = ctxpkg.ContextSetToken(ctx, res.Token)
			ctx = ctxpkg.ContextSetUser(ctx, res.User)
			ctx = metadata.AppendToOutgoingContext(ctx, ctxpkg.TokenHeader, res.Token)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func keyBase(p string) middlewares.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			base, _ := ctx.Value(ctxKeyBaseURI).(string)
			base = path.Join("/", base, p)
			ctx = context.WithValue(ctx, ctxKeyBaseURI, base)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// Handler handles requests.
func (h *DavHandler) Handler() http.Handler {
	return h.router
}

func getTokenStatInfo(ctx context.Context, client gatewayv1beta1.GatewayAPIClient, token string) (*provider.StatResponse, error) {
	return client.Stat(ctx, &provider.StatRequest{Ref: &provider.Reference{Path: path.Join("/public", token)}})
}

func handleBasicAuth(ctx context.Context, c gatewayv1beta1.GatewayAPIClient, token, pw string) (*gatewayv1beta1.AuthenticateResponse, error) {
	authenticateRequest := gatewayv1beta1.AuthenticateRequest{
		Type:         "publicshares",
		ClientId:     token,
		ClientSecret: "password|" + pw,
	}

	return c.Authenticate(ctx, &authenticateRequest)
}

func handleSignatureAuth(ctx context.Context, c gatewayv1beta1.GatewayAPIClient, token, sig, expiration string) (*gatewayv1beta1.AuthenticateResponse, error) {
	authenticateRequest := gatewayv1beta1.AuthenticateRequest{
		Type:         "publicshares",
		ClientId:     token,
		ClientSecret: "signature|" + sig + "|" + expiration,
	}

	return c.Authenticate(ctx, &authenticateRequest)
}

func handleOCMAuth(ctx context.Context, c gatewayv1beta1.GatewayAPIClient, token string) (*gatewayv1beta1.AuthenticateResponse, error) {
	return c.Authenticate(ctx, &gatewayv1beta1.AuthenticateRequest{
		Type:     "ocmshares",
		ClientId: token,
	})
}
