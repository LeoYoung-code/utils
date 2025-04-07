package encrypt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "测试正常时间返回",
			args: args{
				s: "123456",
			},
			want: []byte{0x5, 0x31, 0x32, 0x33, 0x34, 0x35, 0x1, 0x36},
		},
		{
			name: "测试空字符串",
			args: args{
				s: "",
			},
			want: []byte(""),
		},
		{
			name: "测试异常字符串",
			args: args{
				s: "💵💵💵💵",
			},
			want: []byte{0x5, 0xf0, 0x9f, 0x92, 0xb5, 0xf0, 0x5, 0x9f, 0x92, 0xb5, 0xf0, 0x9f, 0x5, 0x92, 0xb5, 0xf0, 0x9f, 0x92, 0x1, 0xb5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Compress(tt.args.s), "Compress(%v)", tt.args.s)
		})
	}
}

func TestCompressBytes(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "测试正常时间返回",
			args: args{
				s: []byte("123456"),
			},
			want: []byte{0x5, 0x31, 0x32, 0x33, 0x34, 0x35, 0x1, 0x36},
		},
		{
			name: "测试空字符串",
			args: args{
				s: []byte(""),
			},
			want: []byte(""),
		},
		{
			name: "测试异常字符串",
			args: args{
				s: []byte("💵💵💵💵"),
			},
			want: []byte{0x5, 0xf0, 0x9f, 0x92, 0xb5, 0xf0, 0x5, 0x9f, 0x92, 0xb5, 0xf0, 0x9f, 0x5, 0x92, 0xb5, 0xf0, 0x9f, 0x92, 0x1, 0xb5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CompressBytes(tt.args.s), "CompressBytes(%v)", tt.args.s)
		})
	}
}

func TestDecompress(t *testing.T) {
	type args struct {
		c []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "测试正常时间返回",
			args: args{
				c: []byte{0x5, 0x31, 0x32, 0x33, 0x34, 0x35, 0x1, 0x36},
			},
			want:    "123456",
			wantErr: assert.NoError,
		},
		{
			name: "测试空字符串",
			args: args{
				c: []byte(""),
			},
			want:    "",
			wantErr: assert.NoError,
		},
		{
			name: "测试异常字符串",
			args: args{
				c: []byte{0x5, 0xf0, 0x9f, 0x92, 0xb5, 0xf0, 0x5, 0x9f, 0x92, 0xb5, 0xf0, 0x9f, 0x5, 0x92, 0xb5, 0xf0, 0x9f, 0x92, 0x1, 0xb5},
			},
			want:    "💵💵💵💵",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decompress(tt.args.c)
			if !tt.wantErr(t, err, fmt.Sprintf("Decompress(%v)", tt.args.c)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Decompress(%v)", tt.args.c)
		})
	}
}

func TestDecompressToBytes(t *testing.T) {
	type args struct {
		c []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecompressToBytes(tt.args.c)
			if !tt.wantErr(t, err, fmt.Sprintf("DecompressToBytes(%v)", tt.args.c)) {
				return
			}
			assert.Equalf(t, tt.want, got, "DecompressToBytes(%v)", tt.args.c)
		})
	}
}
