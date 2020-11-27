// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIt License.
// License that can be found in the LICENSE file.

package metric

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rakutentech/passenger-go-exporter/logging"
	"github.com/stretchr/testify/assert"
)

func TestDescribeMock(t *testing.T) {
	server := MockServer{}
	logger := logging.NewLogger("logfmt", "info")

	collector := NewCollector(server, logger)

	metrics := []*prometheus.Desc{
		WaitListSizeGaugeDesc,
		ProcessCountGaugeDesc,
		ProcessRealMemoryBytesGaugeDesc,
		ProcessProcessedCounterDesc,
	}

	descch := make(chan *prometheus.Desc, len(metrics))
	collector.Describe(descch)

	for _, metric := range metrics {
		msg := <-descch
		assert.Equal(t, msg, metric)
	}
}

func TestCollectMock(t *testing.T) {
	server := MockServer{}
	logger := logging.NewLogger("logfmt", "info")

	collector := NewCollector(server, logger)

	metrics := []*prometheus.Desc{
		WaitListSizeGaugeDesc,
		ProcessCountGaugeDesc,
		ProcessRealMemoryBytesGaugeDesc,
		ProcessProcessedCounterDesc,
	}

	metricch := make(chan prometheus.Metric, len(metrics))
	collector.Collect(metricch)

	count := 0
	for range metrics {
		<-metricch
		count++
	}
	assert.Equal(t, 4, count)
}

func TestCollectError(t *testing.T) {
	server := ErrorMockServer{}
	logger := logging.NewLogger("logfmt", "info")

	collector := NewCollector(server, logger)

	metrics := []*prometheus.Desc{
		WaitListSizeGaugeDesc,
		ProcessCountGaugeDesc,
		ProcessRealMemoryBytesGaugeDesc,
		ProcessProcessedCounterDesc,
	}

	metricch := make(chan prometheus.Metric, len(metrics))

	defer func() {
		o := recover()
		v, ok := o.(error)
		if !ok || v.Error() != "mock error" {
			assert.Fail(t, "unkown error", "unkown error [%+v]", o)
		}
	}()

	collector.Collect(metricch)
	assert.Fail(t, "Not catch error", "Not catch error")
}
