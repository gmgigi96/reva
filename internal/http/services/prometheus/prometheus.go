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

package prometheus

import (
	"context"
	"net/http"

	"contrib.go.opencensus.io/exporter/prometheus"
	"github.com/cs3org/reva/pkg/rhttp/global"
	"github.com/cs3org/reva/pkg/rhttp/mux"
	"github.com/pkg/errors"
	"go.opencensus.io/stats/view"
)

const name = "prometheus"

func init() {
	global.Register(name, New)
}

// New returns a new prometheus service.
func New(ctx context.Context, _ map[string]interface{}) (global.Service, error) {
	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "revad",
	})
	if err != nil {
		return nil, errors.Wrap(err, "prometheus: error creating exporter")
	}

	view.RegisterExporter(pe)
	return &svc{h: pe}, nil
}

func (s *svc) Name() string {
	return name
}

type svc struct {
	h http.Handler
}

func (s *svc) Register(r mux.Router) {
	// TODO(labkode): all prometheus endpoints are public?
	r.With("/metrics", mux.Unprotected())
	r.Mount("/metrics", s.h)
}

func (s *svc) Close() error {
	return nil
}
