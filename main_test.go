// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIt License.
// License that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	if os.Getenv("GITHUB_RUN_ID") != "" {
		t.Skip("This test not running at GitHub Action env")
	}
	os.Setenv("PASSENGER_INSTANCE_REGISTRY_DIR", "/sock")

	statusc := make(chan string) //startup message channel.
	runExporter(t, statusc)

	callPassengerApp(t)

	time.Sleep(10 * time.Millisecond)
	url := "http://localhost:9768/health"
	checlHealth(t, url)

	url = "http://localhost:9768/metrics"
	checkMetrics(t, url)
}

func TestRunNotFound(t *testing.T) {
	os.Setenv("PASSENGER_INSTANCE_REGISTRY_DIR", "/tmp")

	statusc := make(chan string, 1) //startup message channel.

	runExporter(t, statusc)

	select {
	case <-statusc:
		assert.True(t, true)
	case <-time.After(3 * time.Second):
		assert.Fail(t, "main is not stopped", "main is not stopped")
	}
}

// runExporter executes passenger-go-exporter.
func runExporter(t *testing.T, ch chan string) {
	go func() {
		main()
		ch <- "FINISH"
	}()
}

// callPassengerApp calls passenger-app.
func callPassengerApp(t *testing.T) {
	url := "http://passenger-app:3000/"
	res, err := http.Get(url)
	if err != nil {
		t.Fatalf("HTTP Call failed[%s]", err.Error())
	}
	defer res.Body.Close()
}

func checkMetrics(t *testing.T, url string) {
	res, err := http.Get(url)
	if err != nil {
		t.Fatalf("HTTP Call failed[%s]", err.Error())
	}
	defer res.Body.Close()

	reader := bufio.NewReaderSize(res.Body, 1024)

	existsProcessCount := false
	existsProcessProcessed := false
	existsProcessRealMemory := false
	existsWaitListSie := false
	for {
		line, _, err := reader.ReadLine()
		lineStr := string(line)
		if strings.HasPrefix(lineStr, "passenger_go_process_count") {
			existsProcessCount = true
			val := getPrometheusFieldsValue(t, lineStr)
			assert.Greater(t, val, float64(0))
		}
		if strings.HasPrefix(lineStr, "passenger_go_process_processed") {
			existsProcessProcessed = true
			val := getPrometheusFieldsValue(t, lineStr)
			assert.Greater(t, val, float64(0))
		}
		if strings.HasPrefix(lineStr, "passenger_go_process_real_memory") {
			existsProcessRealMemory = true
			val := getPrometheusFieldsValue(t, lineStr)
			assert.Greater(t, val, float64(0))
		}
		if strings.HasPrefix(lineStr, "passenger_go_wait_list_size") {
			existsWaitListSie = true
			val := getPrometheusFieldsValue(t, lineStr)
			assert.GreaterOrEqual(t, val, float64(0))
		}

		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatalf("HTTP Call failed[%s]", err.Error())
		}
	}

	if !existsProcessCount {
		t.Fatal("not found passenger_go_process_count")
	} else if !existsProcessProcessed {
		t.Fatal("not found passenger_go_process_processed")
	} else if !existsProcessRealMemory {
		t.Fatal("not found passenger_go_process_real_memory")
	} else if !existsWaitListSie {
		t.Fatal("not found passenger_go_wait_list_size")
	}
}

func checlHealth(t *testing.T, url string) {
	res, err := http.Get(url)
	if err != nil {
		t.Fatalf("HTTP Call failed[%s]", err.Error())
	}
	defer res.Body.Close()

	assert.Equal(t, 200, res.StatusCode)
}

func getPrometheusFieldsValue(t *testing.T, line string) float64 {
	fields := strings.Fields(line)
	fmt.Println(line)
	v, err := strconv.ParseFloat(fields[len(fields)-1], 64)
	if err != nil {
		t.Fatalf("Parse Error[%s]", line)
	}
	return v
}
