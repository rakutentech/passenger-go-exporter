// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIT License.
// License that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rakutentech/passenger-go-exporter/logging"
	"github.com/rakutentech/passenger-go-exporter/metric"
	"github.com/rakutentech/passenger-go-exporter/passenger"
)

var (
	listenPort = flag.Int("port", 9768, "Listening port number.")
	logfmt     = flag.String("logfmt", "logfmt", "PromLogFormat[logfmt|json].")
	loglevel   = flag.String("loglevel", "info", "PromLogLevel[debug, info, warn, error].")
	timeout    = flag.Int("timeout", 30, "Maximum time (in seconds) to wait for passenger instance.")
)

func main() {
	flag.Parse()

	logger := logging.NewLogger(*logfmt, *loglevel)

	// Search passenger instance.
	_ = level.Info(logger).Log(logging.Msg("Searching passenger instance."))
	c := passenger.Context{}
	factory := passenger.CreateFactory(c)
	server := factory.FindInstance()

	// Calculate max retries based on timeout (each retry is 200ms)
	maxRetries := (*timeout * 1000) / 200

	for i := 0; i < maxRetries; i++ {
		if server != nil {
			break
		}
		_ = level.Info(logger).Log(logging.Msg("passenger not found. wait 200ms."))
		time.Sleep(time.Millisecond * 200)
		server = factory.FindInstance()
	}
	if server == nil {
		_ = level.Error(logger).Log(logging.Msg("passenger not found."))
		return
	}
	_ = level.Info(logger).Log(logging.Msg("Found passenger instance."))

	// Collector setup.
	collector := metric.NewCollector(server, logger)
	prometheus.MustRegister(collector)
	_ = level.Info(logger).Log(logging.Msgf("Starting passenger-go-exporter[port %d]", *listenPort))

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
		_ = level.Error(logger).Log(logging.Err(&err))
	}
}
