package gc

import (
	"testing"
)

func TestPrintGCStats(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go PrintGCStats()
			for i := 0; i < 100000; i++ {
				_ = make([]byte, 1<<20)
			}
		})
	}
}

func TestPrintMemStats(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go PrintMemStats()
			for i := 0; i < 100000; i++ {
				_ = make([]byte, 1<<20)
			}
		})
	}
}
