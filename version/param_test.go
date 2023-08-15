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
			name: "test3",
			args: args{
				versionStr: "",
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				versionStr: "https://gopherize.me/goph",
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
