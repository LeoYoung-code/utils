package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_signature(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "测试正常返回",
			want: "85d2cf664be821400f5fc5963cb952f324e80679",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, signature(), "signature()")
		})
	}
}
