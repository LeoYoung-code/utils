package _map

import (
	"reflect"
	"testing"
)

func TestSequenceTraversalMap(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		want []string
	}{
		{
			name: "empty map",
			m:    map[string]int{},
			// want: []string{},
		},
		{
			name: "single element",
			m:    map[string]int{"hello": 1},
			// want: []string{"hello"},
		},
		{
			name: "multiple elements",
			m:    map[string]int{"hello": 1, "go": 2, "land": 3},
			// want: []string{"go", "hello", "land"},
		},
		{
			name: "special characters in keys",
			m:    map[string]int{"!@#$": 1, "%^&*": 2, "()*": 3},
			// want: []string{"!@#$", "%^&*", "()*"},
		},
		{
			name: "numeric keys",
			m:    map[string]int{"1": 1, "2": 2, "3": 3},
			// want: []string{"1", "2", "3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]string, 0, len(tt.m))
			SequenceTraversalMap(tt.m)

			// change SequenceTraversalMap to return the keys and values in a format that can be compared
			// for this example, we're just comparing keys

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SequenceTraversalMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
