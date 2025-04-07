package test

import (
	"reflect"
	"testing"
)

func Test_calculateRealSize(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{
			name: "Example struct",
			args: args{
				value: Example{
					Name:    "Alice",
					Age:     30,
					Friends: []string{"Bob", "Charlie", "Daisy"},
				},
			},
			want: uintptr(116),
		},
		{
			name: "Slice of strings",
			args: args{
				value: []string{"Alice", "Bob", "Charlie"},
			},
			want: uintptr(87),
		},
		{
			name: "Map of strings",
			args: args{
				value: map[string]string{
					"Alice":   "Bob",
					"Charlie": "Daisy",
				},
			},
			want: uintptr(92),
		},
		{
			name: "Pointer to int",
			args: args{
				value: new(int),
			},
			want: uintptr(8), // 假设每个int占用8字节
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateRealSize(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateRealSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
