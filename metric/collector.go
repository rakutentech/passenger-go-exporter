// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIt License.
// License that can be found in the LICENSE file.

package metric

import (
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rakutentech/passenger-go-exporter/logging"
	"github.com/rakutentech/passenger-go-exporter/passenger"
)

// Collector implements the prometheus.Collector interface.
//
// For more details, Please check following site.
//
// https://godoc.org/github.com/prometheus/client_golang/prometheus#Collector
type Collector struct {
	server passenger.Server
	logger log.Logger
}

// NewCollector creates new Collector instance.
func NewCollector(server passenger.Server, logger log.Logger) *Collector {
	return &Collector{
		server: server,
		logger: logger,
	}
}

// Describe implements the prometheus.Collector interface.
//
// Send the monitored prometheus.Desc.
func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- WaitListSizeGaugeDesc
	ch <- ProcessCountGaugeDesc
	ch <- ProcessRealMemoryBytesGaugeDesc
	ch <- ProcessProcessedCounterDesc
}

// Collect implements the prometheus.Collector interface.
//
// Get new metrics from passenger,and send them.
func (collector *Collector) Collect(ch chan<- prometheus.Metric) {
	info, err := collector.server.Metrics()
	if err != nil {
		level.Error(collector.logger).Log(logging.Err(&err))
		level.Error(collector.logger).Log(logging.Msg("restart process."))
		panic(err)
	}
	for _, sg := range info.SuperGroups {
		shortName := sg.Group.Name
		if i := strings.LastIndex(shortName, "/"); i > 0 {
			shortName = shortName[i+1:]
		}
		ch <- prometheus.MustNewConstMetric(WaitListSizeGaugeDesc, prometheus.GaugeValue, float64(sg.Group.GetWaitListSize), shortName)
		ch <- prometheus.MustNewConstMetric(ProcessCountGaugeDesc, prometheus.GaugeValue, float64(len(sg.Group.Processes)), shortName)
		for _, p := range sg.Group.Processes {
			ch <- prometheus.MustNewConstMetric(ProcessRealMemoryBytesGaugeDesc, prometheus.GaugeValue, float64(p.RealMemory*1024), shortName, p.PID)
			ch <- prometheus.MustNewConstMetric(ProcessProcessedCounterDesc, prometheus.CounterValue, float64(p.Processed), shortName, p.PID)
		}
	}
}
