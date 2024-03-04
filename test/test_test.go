package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_testBase(t *testing.T) {
	type args struct {
		i    int64
		base int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "testBase",
			args: args{
				i:    123456789,
				base: 10,
			},
			want: 123456789,
		},
		{
			name: "testBase",
			args: args{
				i:    123456789,
				base: 0,
			},
			want: 123456789,
		},
		{
			name: "testBase",
			args: args{
				i:    6,
				base: 2,
			},
			want: 110,
		},
		{
			name: "testBase",
			args: args{
				i:    886543,
				base: 8,
			},
			want: 3303417,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, testBase(tt.args.i, tt.args.base), "testBase(%v, %v)", tt.args.i, tt.args.base)
		})
	}
}
