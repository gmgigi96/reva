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

package helloworld

import (
	"context"
	"net/http"

	"github.com/cs3org/reva/pkg/appctx"
	"github.com/cs3org/reva/pkg/rhttp"
	"github.com/cs3org/reva/pkg/rhttp/mux"
	"github.com/cs3org/reva/pkg/utils/cfg"
)

const name = "helloworld"

func init() {
	rhttp.Register(name, New)
}

// New returns a new helloworld service.
func New(ctx context.Context, m map[string]interface{}) (rhttp.Service, error) {
	var c config
	if err := cfg.Decode(m, &c); err != nil {
		return nil, err
	}

	return &svc{conf: &c}, nil
}

// Close performs cleanup.
func (s *svc) Close() error {
	return nil
}

func (s *svc) Name() string {
	return name
}

type config struct {
	HelloMessage string `mapstructure:"message"`
}

func (c *config) ApplyDefaults() {
	if c.HelloMessage == "" {
		c.HelloMessage = "Hello World!"
	}
}

type svc struct {
	conf *config
}

func (s *svc) Register(r mux.Router) {
	r.Get("/helloworld", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := appctx.GetLogger(r.Context())
		if _, err := w.Write([]byte(s.conf.HelloMessage)); err != nil {
			log.Err(err).Msg("error writing response")
		}
	}), mux.Unprotected())
}
