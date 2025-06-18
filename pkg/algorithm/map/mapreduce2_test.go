package _map

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSequenceTraversalMap(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
	}{
		{
			name: "empty map",
			m:    map[string]int{},
		},
		{
			name: "single element",
			m:    map[string]int{"hello": 1},
		},
		{
			name: "multiple elements",
			m:    map[string]int{"hello": 1, "go": 2, "land": 3},
		},
		{
			name: "special characters in keys",
			m:    map[string]int{"!@#$": 1, "%^&*": 2, "()*": 3},
		},
		{
			name: "numeric keys",
			m:    map[string]int{"1": 1, "2": 2, "3": 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout to verify the function runs without error
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call the function
			SequenceTraversalMap(tt.m)

			// Restore stdout
			w.Close()
			os.Stdout = old

			// Read what was captured
			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			// For empty map, we expect no output
			if len(tt.m) == 0 {
				if output != "" {
					t.Errorf("Expected no output for empty map, got: %s", output)
				}
			} else {
				// For non-empty maps, we just verify output was produced
				if output == "" {
					t.Errorf("Expected output for non-empty map, got empty string")
				}
				// Verify that all keys appear in output
				for key := range tt.m {
					if !strings.Contains(output, key) {
						t.Errorf("Expected key %s to appear in output", key)
					}
				}
			}
		})
	}
}
