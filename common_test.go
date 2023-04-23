package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSum32(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "测试正常时间返回",
			args: args{
				data: "123456",
			},
			want: 0x9995b6aa,
		},
		{
			name: "测试空字符串",
			args: args{
				data: "",
			},
			want: 0x811c9dc5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetSum32(tt.args.data), "GetSum32(%v)", tt.args.data)
		})
	}
}

func TestNum2Version(t *testing.T) {
	type args struct {
		innerVersion string
	}
	tests := []struct {
		name           string
		args           args
		wantOutVersion string
	}{
		{
			name: "测试正常时间返回",
			args: args{
				innerVersion: "1.2.3",
			},
			wantOutVersion: "1.2.3",
		},
		{
			name: "测试空字符串",
			args: args{
				innerVersion: "",
			},
			wantOutVersion: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantOutVersion, Num2Version(tt.args.innerVersion), "Num2Version(%v)", tt.args.innerVersion)
		})
	}
}
