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

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/cs3org/reva/pkg/errtypes"
	"github.com/cs3org/reva/pkg/rhttp"
	"github.com/pkg/errors"
)

// ErrTokenInvalid is the error returned by the invite-accepted
// endpoint when the token is not valid.
var ErrTokenInvalid = errors.New("the invitation token is invalid")

// ErrServiceNotTrusted is the error returned by the invite-accepted
// endpoint when the service is not trusted to accept invitations.
var ErrServiceNotTrusted = errors.New("service is not trusted to accept invitations")

// ErrUserAlreadyAccepted is the error returned by the invite-accepted
// endpoint when a user is already know by the remote cloud.
var ErrUserAlreadyAccepted = errors.New("user already accepted an invitation token")

// ErrTokenNotFound is the error returned by the invite-accepted
// endpoint when the request is done using a not existing token.
var ErrTokenNotFound = errors.New("token not found")

// OCMClient is the client for an OCM provider.
type OCMClient struct {
	client *http.Client
}

// Config is the configuration to be used for the OCMClient.
type Config struct {
	Timeout  time.Duration
	Insecure bool
}

// New returns a new OCMClient.
func New(c *Config) *OCMClient {
	return &OCMClient{
		client: rhttp.GetHTTPClient(
			rhttp.Timeout(c.Timeout),
			rhttp.Insecure(c.Insecure),
		),
	}
}

// InviteAcceptedRequest contains the parameters for accepting
// an invitation.
type InviteAcceptedRequest struct {
	UserID            string `json:"userID"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	RecipientProvider string `json:"recipientProvider"`
	Token             string `json:"token"`
}

// User contains the remote user's information when accepting
// an invitation.
type User struct {
	UserID string `json:"userID"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

func (r *InviteAcceptedRequest) toJSON() (io.Reader, error) {
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(r); err != nil {
		return nil, err
	}
	return &b, nil
}

// InviteAccepted informs the sender that the invitation was accepted to start sharing
// https://cs3org.github.io/OCM-API/docs.html?branch=develop&repo=OCM-API&user=cs3org#/paths/~1invite-accepted/post
func (c *OCMClient) InviteAccepted(ctx context.Context, endpoint string, r *InviteAcceptedRequest) (*User, error) {
	url, err := url.JoinPath(endpoint, "invite-accepted")
	if err != nil {
		return nil, err
	}

	body, err := r.toJSON()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error doing request")
	}
	defer resp.Body.Close()

	return c.parseInviteAcceptedResponse(resp)
}

func (c *OCMClient) parseInviteAcceptedResponse(r *http.Response) (*User, error) {
	switch r.StatusCode {
	case http.StatusOK:
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			return nil, errors.Wrap(err, "error decoding response body")
		}
		return &u, nil
	case http.StatusBadRequest:
		return nil, ErrTokenInvalid
	case http.StatusNotFound:
		return nil, ErrTokenNotFound
	case http.StatusConflict:
		return nil, ErrUserAlreadyAccepted
	case http.StatusForbidden:
		return nil, ErrServiceNotTrusted
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding response body")
	}
	return nil, errtypes.InternalError(string(body))
}
