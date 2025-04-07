package channel

import (
	"testing"
)

func Test_notifyQuit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notifyQuit()
		})
	}
}
