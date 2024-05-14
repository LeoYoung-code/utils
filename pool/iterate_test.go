package pool

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestIterateWithStep(t *testing.T) {
	process := func(data []string) error {
		for _, str := range data {
			if str == "error" {
				return errors.New("found error")
			}
		}
		return nil
	}

	tests := []struct {
		name  string
		data  []string
		step  int
		wantErr string
	}{
		{"EmptyData", []string{}, 1, ""},
		{"StepLargerThanLen", []string{"a", "b", "c"}, 5, ""},
		{"StepOneWithData", []string{"a", "b", "c"}, 1, ""},
		{"AllDataProcessed", []string{"a", "b", "c"}, 3, ""},
		{"ErrorInProcess", []string{"a", "b", "error"}, 1, "found error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			err := IterateWithStep(ctx, tt.data, tt.step, process)
			if err != nil {
				if tt.wantErr == "" || !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("IterateWithStep() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if tt.wantErr != "" && err == nil {
				t.Errorf("IterateWithStep() expected error did not occur")
			}
		})
	}
}