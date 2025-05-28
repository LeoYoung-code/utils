package version

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSysVersion(t *testing.T) {
	type args struct {
		versionStr string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{
				versionStr: "10.2",
			},
			want: 10,
		},
		{
			name: "test2",
			args: args{
				versionStr: "10.10.5",
			},
			want: 10,
		},
		{
			name: "empty_string",
			args: args{
				versionStr: "",
			},
			want: 0,
		},
		{
			name: "non_numeric_prefix",
			args: args{
				versionStr: "v1.2.3",
			},
			want: 0,
		},
		{
			name: "invalid_number",
			args: args{
				versionStr: "abc.def",
			},
			want: 0,
		},
		{
			name: "single_number",
			args: args{
				versionStr: "5",
			},
			want: 5,
		},
		{
			name: "multiple_dots",
			args: args{
				versionStr: "10.2.3.4",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(GetSysVersion(tt.args.versionStr))
			assert.Equalf(t, tt.want, GetSysVersion(tt.args.versionStr), "GetSysVersion(%v)", tt.args.versionStr)
		})
	}
}
