package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetStruct_Add(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		s    SetStruct
		args args
	}{
		{"test1", SetStruct{
			"a": {},
			"b": {},
			"c": {},
		}, args{"d(new)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.key)
		})
		fmt.Println(tt.s)
	}
}

func TestSetStruct_Delete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		s    SetStruct
		args args
	}{
		{"test1", SetStruct{
			"a": {},
			"b": {},
			"c": {},
		}, args{"a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Delete(tt.args.key)
		})
		fmt.Println(tt.s)
	}
}

func TestSetStruct_Has(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		s    SetStruct
		args args
		want bool
	}{
		{"test1", SetStruct{
			"a": {},
			"b": {},
			"c": {},
		}, args{"a"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Has(tt.args.key), "Has(%v)", tt.args.key)
		})
	}
}
