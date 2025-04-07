package channel

import (
	"testing"
)

func Test_waitQuit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			waitQuit()
		})
	}
}
