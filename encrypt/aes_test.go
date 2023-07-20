package encrypt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAesDecrypt(t *testing.T) {
	type args struct {
		decodeStr string
		key       []byte
		iv        string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test1",
			args: args{
				decodeStr: "abc",
				key:       []byte{0x61, 0x62, 0x63},
				iv:        "bcd",
			},
			want: []byte{0x61, 0x62, 0x63},
		},
		{
			name: "test2",
			args: args{
				decodeStr: "",
				key:       []byte{0x61, 0x62, 0x63},
				iv:        "bcd",
			},
			want: []byte{0x61, 0x62, 0x63},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesDecrypt(tt.args.decodeStr, tt.args.key, tt.args.iv)
			if !tt.wantErr(t, err, fmt.Sprintf("AesDecrypt(%v, %v, %v)", tt.args.decodeStr, tt.args.key, tt.args.iv)) {
				return
			}
			assert.Equalf(t, tt.want, got, "AesDecrypt(%v, %v, %v)", tt.args.decodeStr, tt.args.key, tt.args.iv)
		})
	}
}

func TestAesEncrypt(t *testing.T) {
	type args struct {
		encodeStr string
		key       string
		iv        string
		process   func(crypt []byte) []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test1",
			args: args{
				"abc",
				"abc",
				"kkk",
				nil,
			},
			want:    "",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesEncrypt(tt.args.encodeStr, tt.args.key, tt.args.iv, tt.args.process)
			if !tt.wantErr(t, err, fmt.Sprintf("AesEncrypt(%v, %v, %v)", tt.args.encodeStr, tt.args.key, tt.args.iv)) {
				return
			}
			assert.Equalf(t, tt.want, got, "AesEncrypt(%v, %v, %v, %v)", tt.args.encodeStr, tt.args.key, tt.args.iv, tt.args.process)
		})
	}
}

func TestPKCS5Padding(t *testing.T) {
	type args struct {
		ciphertext []byte
		blockSize  int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PKCS5Padding(tt.args.ciphertext, tt.args.blockSize), "PKCS5Padding(%v, %v)", tt.args.ciphertext, tt.args.blockSize)
		})
	}
}

func TestPKCS5UnPadding(t *testing.T) {
	type args struct {
		origData []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PKCS5UnPadding(tt.args.origData), "PKCS5UnPadding(%v)", tt.args.origData)
		})
	}
}
