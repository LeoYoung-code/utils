package utils

import (
	"testing"
)

func Test_sqlInSliceInt64(t *testing.T) {
	tests := []struct {
		name  string
		opt   string
		field string
		s     []int64
		want  string
	}{
		{"11", "and", "aa", []int64{1, 2, 3}, " and `aa` IN (1,2,3)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SqlInSliceInt64(tt.opt, tt.field, tt.s); got != tt.want {
				t.Errorf("SqlInSliceInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlInSliceString(t *testing.T) {
	tests := []struct {
		name  string
		opt   string
		field string
		s     []string
		want  string
	}{
		{"22", "and", "aa", []string{"1", "2", "3"}, " and `aa` IN ('1','2','3')"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SqlInSliceString(tt.opt, tt.field, tt.s); got != tt.want {
				t.Errorf("SqlInSliceInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
