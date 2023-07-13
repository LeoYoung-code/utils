package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObj2String(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试正常时间返回",
			args: args{
				obj: 123456,
			},
			want: "123456",
		},
		{
			name: "测试空字符串",
			args: args{
				obj: "",
			},
			want: "\"\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Obj2String(tt.args.obj), "Obj2String(%v)", tt.args.obj)
		})
	}
}
