package utils

import (
	"testing"

	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
)

func TestDivInt64(t *testing.T) {
	type args struct {
		d  int64
		d2 int64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "test", args: args{d: 100, d2: 200}, want: 50},
		{name: "test", args: args{d: 100, d2: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DivPerInt64(tt.args.d, tt.args.d2), "DivPerInt64(%v, %v)", tt.args.d, tt.args.d2)
			// assert.Equalf(t, tt.want, DivInt64(tt.args.d, tt.args.d2), "DivInt64(%v, %v)", tt.args.d, tt.args.d2)
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
