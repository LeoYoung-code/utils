package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name   string
		args   args
		wantMd string
	}{
		{
			name: "测试正常时间返回",
			args: args{
				content: "123456",
			},
			wantMd: "e10adc3949ba59abbe56e057f20f883e",
		},
		{
			name: "测试空字符串",
			args: args{
				content: "",
			},
			wantMd: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantMd, Md5(tt.args.content), "Md5(%v)", tt.args.content)
		})
	}
}
