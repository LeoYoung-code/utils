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

func TestByteToString(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ByteToString(tt.args.data), "ByteToString(%v)", tt.args.data)
		})
	}
}

func TestContainsString(t *testing.T) {
	type args struct {
		array []string
		val   string
	}
	tests := []struct {
		name      string
		args      args
		wantIndex int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantIndex, ContainsString(tt.args.array, tt.args.val), "ContainsString(%v, %v)", tt.args.array, tt.args.val)
		})
	}
}

func TestExplode(t *testing.T) {
	type args struct {
		delimiter string
		text      string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Explode(tt.args.delimiter, tt.args.text), "Explode(%v, %v)", tt.args.delimiter, tt.args.text)
		})
	}
}

func TestFilterEmpty(t *testing.T) {
	type args struct {
		c []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FilterEmpty(tt.args.c), "FilterEmpty(%v)", tt.args.c)
		})
	}
}

func TestGenerateUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GenerateUUID(), "GenerateUUID()")
		})
	}
}

func TestInclude(t *testing.T) {
	type args[T comparable] struct {
		arr   []T
		check T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Include(tt.args.arr, tt.args.check), "Include(%v, %v)", tt.args.arr, tt.args.check)
		})
	}
}

func TestIsInt64In(t *testing.T) {
	type args struct {
		element     int64
		targetSlice []int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsInt64In(tt.args.element, tt.args.targetSlice), "IsInt64In(%v, %v)", tt.args.element, tt.args.targetSlice)
		})
	}
}

func TestIsNum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsNum(tt.args.s), "IsNum(%v)", tt.args.s)
		})
	}
}

func TestIsOnline(t *testing.T) {
	type args struct {
		upTime   int64
		downTime int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsOnline(tt.args.upTime, tt.args.downTime), "IsOnline(%v, %v)", tt.args.upTime, tt.args.downTime)
		})
	}
}

func TestIsRepByLoop1(t *testing.T) {
	type args struct {
		origin []int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, IsRepByLoop(tt.args.origin), fmt.Sprintf("IsRepByLoop(%v)", tt.args.origin))
		})
	}
}

func TestIsStringIn(t *testing.T) {
	type args struct {
		targetSlice []string
		element     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsStringIn(tt.args.targetSlice, tt.args.element), "IsStringIn(%v, %v)", tt.args.targetSlice, tt.args.element)
		})
	}
}

func TestIsUnion(t *testing.T) {
	type args[T comparable] struct {
		arr1 []T
		arr2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsUnion(tt.args.arr1, tt.args.arr2), "IsUnion(%v, %v)", tt.args.arr1, tt.args.arr2)
		})
	}
}

func TestJoin(t *testing.T) {
	type args[T comparable] struct {
		t1  []T
		sep string
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want string
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Join(tt.args.t1, tt.args.sep), "Join(%v, %v)", tt.args.t1, tt.args.sep)
		})
	}
}

func TestMapStringVal(t *testing.T) {
	type args struct {
		sl map[string]string
		k  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MapStringVal(tt.args.sl, tt.args.k), "MapStringVal(%v, %v)", tt.args.sl, tt.args.k)
		})
	}
}

func TestMapStringVal2Float64(t *testing.T) {
	type args struct {
		sl map[string]string
		k  string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MapStringVal2Float64(tt.args.sl, tt.args.k), "MapStringVal2Float64(%v, %v)", tt.args.sl, tt.args.k)
		})
	}
}

func TestMd5Parser1(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Md5Parser(tt.args.value), "Md5Parser(%v)", tt.args.value)
		})
	}
}

func TestRandomElement1(t *testing.T) {
	type args struct {
		s []any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RandomElement(tt.args.s), "RandomElement(%v)", tt.args.s)
		})
	}
}

func TestString2Int64(t *testing.T) {
	type args struct {
		k string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, String2Int64(tt.args.k), "String2Int64(%v)", tt.args.k)
		})
	}
}

func TestTernary(t *testing.T) {
	type args struct {
		a bool
		b any
		c any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Ternary(tt.args.a, tt.args.b, tt.args.c), "Ternary(%v, %v, %v)", tt.args.a, tt.args.b, tt.args.c)
		})
	}
}

func TestToAnySlice(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ToAnySlice(tt.args.s), "ToAnySlice(%v)", tt.args.s)
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		data any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ToString(tt.args.data), "ToString(%v)", tt.args.data)
		})
	}
}

func TestUrlPath1(t *testing.T) {
	type args struct {
		rawURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, UrlPath(tt.args.rawURL), "UrlPath(%v)", tt.args.rawURL)
		})
	}
}
