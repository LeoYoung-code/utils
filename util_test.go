package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRepByLoop(t *testing.T) {
	type args struct {
		origin []int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "测试正常时间返回",
			args: args{
				origin: []int64{1, 3, 5, 6, 6, 6, 7},
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, IsRepByLoop(tt.args.origin), fmt.Sprintf("IsRepByLoop(%v)", tt.args.origin))
		})
	}
}

func TestUrlPath(t *testing.T) {
	type args struct {
		rawURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"path+query", args{rawURL: "https://c.test.com/a.png?t=1112233"}, "/a.png?t=1112233"},
		{"path", args{rawURL: "https://c.test.com/a.png"}, "/a.png"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, UrlPath(tt.args.rawURL), "UrlPath(%v)", tt.args.rawURL)
		})
	}
}

func TestRandomElement(t *testing.T) {
	type args struct {
		s []any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			"nil",
			args{s: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := RandomElement(tt.args.s)
			fmt.Printf("%+v", res)
			// assert.Equalf(t, tt.want, res, "RandomElement(%v)", tt.args.s)
		})
	}
}

func TestMd5Parser(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{value: "#21CCC916-6811-43B6-9D88-8176F1CBC820"},
			want: 10663,
		},
		{
			name: "test2",
			args: args{value: "#1166F695-2507-4D25-BF15-1186B8FEF99D"},
			want: 10758,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Md5Parser(tt.args.value), "Md5Parser(%v)", tt.args.value)
		})
	}
}
