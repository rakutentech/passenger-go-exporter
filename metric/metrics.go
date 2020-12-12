// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIT License.
// License that can be found in the LICENSE file.

package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	namespace = "passenger_go"
	// WaitListSizeGaugeDesc is metrics of "info/supergroups/supergroups/group/get_wait_list_size"
	WaitListSizeGaugeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wait_list_size"),
		"number of requests in queue with each application",
		[]string{"name"},
		nil,
	)

	// ProcessCountGaugeDesc is metrics of "info/supergroups/supergroups/group/processes" length.
	ProcessCountGaugeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "process_count"),
		"number of current application processes with each application",
		[]string{"name"},
		nil,
	)

	// ProcessRealMemoryBytesGaugeDesc is metrics of "info/supergroups/supergroups/group/process/real_memory".
	ProcessRealMemoryBytesGaugeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "process_real_memory_bytes"),
		"memory usage of process in each PID",
		[]string{"name", "pid"},
		nil,
	)

	// ProcessProcessedCounterDesc is metrics of "info/supergroups/supergroups/group/process/processed".
	ProcessProcessedCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "process_processed"),
		"the number of requests handled by each process",
		[]string{"name", "pid"},
		nil,
	)
)
