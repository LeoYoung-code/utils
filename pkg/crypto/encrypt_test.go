package encrypt

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
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
		{
			name: "测试异常字符串",
			args: args{
				content: "💵💵💵💵",
			},
			wantMd: "b64b7443b785bd34290925f33c8afb82",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantMd, Md5(tt.args.content), "Md5(%v)", tt.args.content)
		})
	}
}

// aesDecrypt 测试用，如需使用，请做好全面测试
func aesDecrypt(decodeStr string) (string, error) {
	secretSalt := "242ccb8230d709e1"
	decodeBytes, _ := base64.StdEncoding.DecodeString(decodeStr)
	iv := decodeBytes[0:16]
	got, err := AesDecrypt(decodeStr, []byte(secretSalt), string(iv))

	return string(got), err
}

func TestEncrypt(t *testing.T) {
	en1 := Encrypt("aa")
	de1, _ := aesDecrypt(en1)
	en2 := Encrypt("Encrypt")
	de2, _ := aesDecrypt(en2)
	en3 := Encrypt("700")
	de3, _ := aesDecrypt(en3)
	en4 := Encrypt("0")
	de4, _ := aesDecrypt(en4)
	assert.Equal(t, de1, "aa")
	assert.Equal(t, de2, "Encrypt")
	assert.Equal(t, de3, "700")
	assert.Equal(t, de4, "0")
}
