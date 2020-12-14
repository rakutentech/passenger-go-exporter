package e2e

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	if os.Getenv("E2E") != "true" {
		t.Skip("This test should be execute after build only.")
	}

	os.Setenv("PASSENGER_INSTANCE_REGISTRY_DIR", "/sock")

	_, cancel := runExporter(t)
	defer cancel()

	callPassengerApp(t)

	time.Sleep(10 * time.Millisecond)
	url := "http://localhost:9768/metrics"
	checkMetrics(t, url)
}

func TestNotFoundPassengerRun(t *testing.T) {
	if os.Getenv("E2E") != "true" {
		t.Skip("This test should be execute after build only.")
	}

	os.Setenv("PASSENGER_INSTANCE_REGISTRY_DIR", "/not-found-dir")

	cmd, _ := runExporter(t)
	assert.Equal(t, true, cmd.ProcessState.Exited())
}

// runExporter executes passenger-go-exporter.
func runExporter(t *testing.T) (*exec.Cmd, context.CancelFunc) {
	statusc := make(chan string) //startup message channel.

	workdir, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwf() failed %s", err.Error())
	}
	homedir := filepath.Dir(filepath.Dir(workdir))
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	exporter := exec.CommandContext(ctx, filepath.Join(homedir, "passenger-go-exporter"))

	// Wait for Startup.
	stderr, err := exporter.StderrPipe()
	if err != nil {
		t.Fatalf("Failed to connect stderr pipe : %s", err.Error())
	}
	errScanner := bufio.NewScanner(stderr)
	go func() {
		for errScanner.Scan() {
			fmt.Println(errScanner.Text())
			if strings.Contains(errScanner.Text(), "Starting passenger-go-exporter") {
				time.Sleep(time.Millisecond * 10)
				statusc <- "starting"
			}
		}
	}()

	err = exporter.Start()
	if err != nil {
		t.Fatalf("Failed to run: %s", err.Error())
	}
	go func() {
		if err = exporter.Wait(); err != nil {
			cancel()
		}
	}()

WAIT_FOR:
	for {
		time.Sleep(time.Millisecond * 10)
		select {
		case <-statusc:
			break WAIT_FOR
		case <-ctx.Done():
			break WAIT_FOR
		}
	}
	return exporter, cancel
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
		if strings.HasPrefix(lineStr, "passenger_go_process_real_memory_bytes") {
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

func getPrometheusFieldsValue(t *testing.T, line string) float64 {
	fields := strings.Fields(line)
	fmt.Println(line)
	v, err := strconv.ParseFloat(fields[len(fields)-1], 64)
	if err != nil {
		t.Fatalf("Parse Error[%s]", line)
	}
	return v
}
