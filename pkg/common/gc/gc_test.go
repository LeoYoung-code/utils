package gc

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestPrintGCStats(t *testing.T) {
	// Redirect stdout to capture output
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	// Run GC stats collection
	go printGCStats()
	time.Sleep(2 * time.Second) // Allow time for GC to run

	// Restore stdout
	os.Stdout = old
	_ = w.Close()

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	// Validate JSON format
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(output), &m); err != nil {
		t.Errorf("Invalid JSON output: %v", err)
	}

	// Verify required fields
	requiredFields := []string{"NumGC", "PauseTotal", "LastGC"}
	for _, field := range requiredFields {
		if _, exists := m[field]; !exists {
			t.Errorf("Missing required field: %s", field)
		}
	}
}

func TestPrintMemStats(t *testing.T) {
	// Redirect stdout to capture output
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	// Run mem stats collection
	go printMemStats()
	time.Sleep(2 * time.Second) // Allow time for stats to update

	// Restore stdout
	os.Stdout = old
	_ = w.Close()

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	// Validate JSON format
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(output), &m); err != nil {
		t.Errorf("Invalid JSON output: %v", err)
	}

	// Verify memory stats are positive values
	memoryFields := []string{"Alloc", "TotalAlloc", "HeapAlloc"}
	for _, field := range memoryFields {
		if val, ok := m[field].(float64); !ok || val <= 0 {
			t.Errorf("Invalid value for %s: %v", field, val)
		}
	}
}
