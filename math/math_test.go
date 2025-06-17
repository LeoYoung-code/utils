package math

import (
	"testing"

	"github.com/LeoYoung-code/cast"
	"github.com/stretchr/testify/assert"
)

func TestDivInt64(t *testing.T) {
	type args struct {
		d  int64
		d2 int64
	}
	tests := []struct {
		name   string
		args   args
		want   float64
		isZero bool
	}{
		{name: "正常计算", args: args{d: 100, d2: 200}, want: 50},
		{name: "除数为零", args: args{d: 100, d2: 0}, want: 0, isZero: true},
		{name: "负数计算", args: args{d: -100, d2: 200}, want: -50},
		{name: "大数计算", args: args{d: 9223372036854775807, d2: 2}, want: 4611686018427387903.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DivPerInt64(tt.args.d, tt.args.d2)
			if tt.isZero {
				assert.Zero(t, got, "应该返回零")
			} else {
				assert.InDelta(t, tt.want, got, 0.001, "结果应该在预期范围内")
			}
		})
	}
}

func TestString2Float64(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want float64
	}{
		{name: "1", s: "34324.12", want: 34324.12},
		{"2", "1.47000", 1.47000},
		{"3", "3432412", float64(3432412)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := String2Float64(tt.s)
			t.Logf("%v", f)
			// assert.Equalf(t, tt.want, f, "String2Float64(%v)", tt.s)
		})
	}
}

func TestFen2YuanString(t *testing.T) {
	type args struct {
		fen string
		r   int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "fen", args: args{fen: "100", r: 2}, want: "1.00"},
		{name: "fen", args: args{fen: "300", r: 2}, want: "3.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Float2String2f(Fen2YuanString(tt.args.fen, tt.args.r)), "Fen2YuanString(%v, %v)", tt.args.fen, tt.args.r)
		})
	}
}

func TestDivInt641(t *testing.T) {
	type args struct {
		d  int64
		d2 int64
		r  int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{args: args{
			d:  9816259,
			d2: 55331,
			r:  4,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arpu := cast.ToString(DivInt64(tt.args.d, tt.args.d2, tt.args.r))
			t.Log(arpu)
		})
	}
}

func TestDivPerString(t *testing.T) {
	type args struct {
		d  string
		d2 string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{args: args{
			d:  "962.59129",
			d2: "331",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// assert.Equalf(t, tt.want, DivStringRound(tt.args.d, tt.args.d2,4), "DivPerString(%v, %v)", tt.args.d, tt.args.d2)
			t.Log(DivStringRound(tt.args.d, tt.args.d2, 8))
		})
	}
}

func TestFloat2String4f(t *testing.T) {
	type args struct {
		s float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-1",
			args: args{s: 0.00000000000},
			want: "0.0000",
		},
		{
			name: "test-2",
			args: args{s: 0.00000000001},
			want: "0.0000",
		},
		{
			name: "test-3",
			args: args{s: 0},
			want: "0.0000",
		},
		{
			name: "test-4",
			args: args{s: 0.00003000003},
			want: "0.0000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Float2String4f(tt.args.s), "Float2String4f(%v)", tt.args.s)
		})
	}
}

func TestFloat64Round(t *testing.T) {
	type args struct {
		f float64
		r int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test-1",
			args: args{f: 0.000500, r: 4},
			want: 0.0005,
		},
		{
			name: "test-2",
			args: args{f: 0.00000000001, r: 2},
			want: 0.00,
		},
		{
			name: "test-3",
			args: args{f: 0, r: 4},
			want: 0.0000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Float64Round(tt.args.f, tt.args.r), "Float64Round(%v, %v)", tt.args.f, tt.args.r)
		})
	}
}

func TestFen2YuanFloat64(t *testing.T) {
	type args struct {
		fen float64
		r   int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test-1",
			args: args{fen: 100, r: 2},
			want: 1,
		},
		{
			name: "test-2",
			args: args{fen: 0, r: 2},
			want: 0,
		},
		{
			name: "test-3",
			args: args{fen: 175, r: 4},
			want: 1.75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Fen2YuanFloat64(tt.args.fen, tt.args.r), "Fen2YuanFloat64(%v, %v)", tt.args.fen, tt.args.r)
		})
	}
}

func TestYuan2FenFloat64(t *testing.T) {
	type args struct {
		yuan float64
		r    int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test-1",
			args: args{yuan: 1, r: 2},
			want: 100,
		},
		{
			name: "test-2",
			args: args{yuan: 0, r: 2},
			want: 0,
		},
		{
			name: "test-3",
			args: args{yuan: 1.75, r: 4},
			want: 175,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Yuan2FenFloat64(tt.args.yuan, tt.args.r), "Yuan2FenFloat64(%v, %v)", tt.args.yuan, tt.args.r)
		})
	}
}

func TestPrecisionHandling(t *testing.T) {
	tests := []struct {
		name  string
		f     float64
		r     int32
		want  string
	}{
		{name: "两位小数", f: 3.14159, r: 2, want: "3.14"},
		{name: "四位小数", f: 2.71828, r: 4, want: "2.7183"},
		{name: "零精度", f: 123.456, r: 0, want: "123"},
		{name: "负数", f: -3.14159, r: 3, want: "-3.142"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Float2String4f(tt.f)
			assert.Equal(t, tt.want, result, "格式化结果应该匹配预期")

			rounded := Float64Round(tt.f, tt.r)
			assert.InDelta(t, cast.ToFloat64(tt.want), rounded, 0.01, "四舍五入应该准确")
		})
	}
}
