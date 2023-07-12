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

package mentix

import (
	"context"
	"net/http"

	"github.com/cs3org/reva/pkg/appctx"
	"github.com/cs3org/reva/pkg/mentix"
	"github.com/cs3org/reva/pkg/mentix/config"
	"github.com/cs3org/reva/pkg/mentix/exchangers"
	"github.com/cs3org/reva/pkg/mentix/meshdata"
	"github.com/cs3org/reva/pkg/rhttp"
	"github.com/cs3org/reva/pkg/rhttp/mux"
	"github.com/cs3org/reva/pkg/utils/cfg"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func init() {
	rhttp.Register(serviceName, New)
}

type svc struct {
	conf *config.Configuration
	mntx *mentix.Mentix
	log  *zerolog.Logger

	stopSignal chan struct{}
}

const (
	serviceName = "mentix"
)

func (s *svc) Close() error {
	// Trigger and close the stopSignal signal channel to stop Mentix
	s.stopSignal <- struct{}{}
	close(s.stopSignal)

	return nil
}

func (s *svc) Name() string {
	return serviceName
}

func (s *svc) Unprotected() []string {
	// Get all endpoints exposed by the RequestExchangers
	importers := s.mntx.GetRequestImporters()
	exporters := s.mntx.GetRequestExporters()

	getEndpoints := func(exchangers []exchangers.RequestExchanger) []string {
		endpoints := make([]string, 0, len(exchangers))
		for _, exchanger := range exchangers {
			if !exchanger.IsProtectedEndpoint() {
				endpoints = append(endpoints, exchanger.Endpoint())
			}
		}
		return endpoints
	}

	endpoints := make([]string, 0, len(importers)+len(exporters))
	endpoints = append(endpoints, getEndpoints(importers)...)
	endpoints = append(endpoints, getEndpoints(exporters)...)

	return endpoints
}

func (s *svc) Register(r mux.Router) {
	r.Mount("/mentix", http.HandlerFunc(s.mntx.RequestHandler))
}

func (s *svc) startBackgroundService() {
	// Just run Mentix in the background
	go func() {
		if err := s.mntx.Run(s.stopSignal); err != nil {
			s.log.Err(err).Msg("error while running mentix")
		}
	}()
}

func applyInternalConfig(m map[string]interface{}, conf *config.Configuration) {
	getSubsections := func(section string) []string {
		subsections := make([]string, 0, 5)
		if list, ok := m[section].(map[string]interface{}); ok {
			for name := range list {
				subsections = append(subsections, name)
			}
		}
		return subsections
	}

	conf.EnabledConnectors = getSubsections("connectors")
	conf.EnabledImporters = getSubsections("importers")
	conf.EnabledExporters = getSubsections("exporters")
}

func applyDefaultConfig(conf *config.Configuration) {
	// General

	if conf.UpdateInterval == "" {
		conf.UpdateInterval = "1h" // Update once per hour
	}

	// Connectors
	if conf.Connectors.GOCDB.Scope == "" {
		conf.Connectors.GOCDB.Scope = "SM" // TODO(Daniel-WWU-IT): This might change in the future
	}

	// Exporters
	addDefaultConnector := func(enabledList *[]string) {
		if len(*enabledList) == 0 {
			*enabledList = append(*enabledList, "*")
		}
	}

	if conf.Exporters.WebAPI.Endpoint == "" {
		conf.Exporters.WebAPI.Endpoint = "/sites"
	}
	addDefaultConnector(&conf.Exporters.WebAPI.EnabledConnectors)

	if conf.Exporters.CS3API.Endpoint == "" {
		conf.Exporters.CS3API.Endpoint = "/cs3"
	}
	addDefaultConnector(&conf.Exporters.CS3API.EnabledConnectors)
	if len(conf.Exporters.CS3API.ElevatedServiceTypes) == 0 {
		conf.Exporters.CS3API.ElevatedServiceTypes = append(conf.Exporters.CS3API.ElevatedServiceTypes, meshdata.EndpointGateway, meshdata.EndpointOCM, meshdata.EndpointWebdav)
	}

	if conf.Exporters.SiteLocations.Endpoint == "" {
		conf.Exporters.SiteLocations.Endpoint = "/loc"
	}
	addDefaultConnector(&conf.Exporters.SiteLocations.EnabledConnectors)

	addDefaultConnector(&conf.Exporters.PrometheusSD.EnabledConnectors)
	addDefaultConnector(&conf.Exporters.Metrics.EnabledConnectors)
}

// New returns a new Mentix service.
func New(ctx context.Context, m map[string]interface{}) (rhttp.Service, error) {
	var c config.Configuration
	if err := cfg.Decode(m, &c); err != nil {
		return nil, err
	}
	applyInternalConfig(m, &c)
	applyDefaultConfig(&c)

	// Create the Mentix instance
	log := appctx.GetLogger(ctx)
	mntx, err := mentix.New(&c, log)
	if err != nil {
		return nil, errors.Wrap(err, "mentix: error creating Mentix")
	}

	// Create the service and start its background activity
	s := &svc{
		conf:       &c,
		mntx:       mntx,
		log:        log,
		stopSignal: make(chan struct{}),
	}
	s.startBackgroundService()
	return s, nil
}
