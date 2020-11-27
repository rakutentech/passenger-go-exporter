// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIt License.
// License that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"net/http"

	"ghe.rakuten-it.com/SOK/passenger-go-exporter/logging"
	"ghe.rakuten-it.com/SOK/passenger-go-exporter/metric"
	"ghe.rakuten-it.com/SOK/passenger-go-exporter/passenger"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenPort = flag.Int("port", 9768, "Listening port number.")
	logfmt     = flag.String("logfmt", "logfmt", "PromLogFormat[logfmt|json].")
	loglevel   = flag.String("loglevel", "info", "PromLogLevel[debug, info, warn, error].")
)

func main() {
	flag.Parse()

	logger := logging.NewLogger(*logfmt, *loglevel)

	// Search passenge instance.
	level.Info(logger).Log(logging.Msg("Searching passenger instance."))
	c := passenger.Context{}
	factory := passenger.CreateFactory(c)
	server := factory.FindInstance()
	if server == nil {
		level.Error(logger).Log(logging.Msg("passenger not found."))
		return
	}
	level.Info(logger).Log(logging.Msg("Found passenger instance."))

	// Collector setup.
	collector := metric.NewCollector(server, logger)
	prometheus.MustRegister(collector)
	level.Info(logger).Log(logging.Msgf("Starting passenger-go-exporter[port %d]", *listenPort))

	// HTTP Server setup.
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if server.IsEnabled() {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
	err := http.ListenAndServe(fmt.Sprintf(":%d", *listenPort), nil)
	if err != nil {
		level.Error(logger).Log(logging.Err(&err))
	}
}
