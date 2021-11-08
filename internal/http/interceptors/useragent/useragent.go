package useragent

import (
	"net/http"

	ctxpkg "github.com/cs3org/reva/pkg/ctx"
	"google.golang.org/grpc/metadata"
)

// New returns a new HTTP middleware that stores the log
// in the context with request ID information.
func New() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return handler(h)
	}
}

func handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// add user agent to the context
		userAgent := r.UserAgent()
		ctx = ctxpkg.ContextSetUserAgent(ctx, userAgent)
		ctx = metadata.AppendToOutgoingContext(ctx, ctxpkg.UserAgentHeader, userAgent)

		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
