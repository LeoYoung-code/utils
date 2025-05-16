package math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatStringAdd(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		expect float64
	}{
		{"1.23", "4.56", 5.79},
		{"0", "0", 0},
		{"-1.23", "1.23", 0},
		{"invalid", "1.23", 0},
		{"1.23", "invalid", 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s+%s=%f", test.a, test.b, test.expect), func(t *testing.T) {
			result := FloatStringAdd(test.a, test.b)
			assert.InDelta(t, test.expect, result, 0.0001)
		})
	}
}

func TestFloat64Add(t *testing.T) {
	tests := []struct {
		a      float64
		b      float64
		expect float64
	}{
		{1.23, 4.56, 5.79},
		{0, 0, 0},
		{-1.23, 1.23, 0},
		{1e+10, 1, 1e+10 + 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%f+%f=%f", test.a, test.b, test.expect), func(t *testing.T) {
			result := Float64Add(test.a, test.b)
			assert.InDelta(t, test.expect, result, 0.0001)
		})
	}
}

func TestSubString(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		expect float64
	}{
		{"10.5", "4.3", 6.2},
		{"0", "0", 0},
		{"-1.23", "1.23", -2.46},
		{"invalid", "1.23", 0},
		{"1.23", "invalid", 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s-%s=%f", test.a, test.b, test.expect), func(t *testing.T) {
			result := SubString(test.a, test.b)
			assert.InDelta(t, test.expect, result, 0.0001)
		})
	}
}

func TestStringDiv100(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"100", "1"},
		{"123.45", "1.2345"},
		{"0", "0"},
		{"invalid", "0"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s/100=%s", test.input, test.expect), func(t *testing.T) {
			result := StringDiv100(test.input)
			assert.Equal(t, test.expect, result)
		})
	}
}

func TestGapString2String(t *testing.T) {
	tests := []struct {
		current string
		prev    string
		expect  string
	}{
		{"110", "100", "10.00"},
		{"90", "100", "-10.00"},
		{"100", "100", "0.00"},
		{"0", "100", "0"},
		{"100", "0", "0"},
		{"invalid", "100", "0"},
		{"100", "invalid", "0"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("GapString(%s,%s)=%s", test.current, test.prev, test.expect), func(t *testing.T) {
			result := GapString2String(test.current, test.prev)
			assert.Equal(t, test.expect, result)
		})
	}
}

func TestGapInt642String(t *testing.T) {
	tests := []struct {
		current int64
		prev    int64
		expect  string
	}{
		{110, 100, "10.00"},
		{90, 100, "-10.00"},
		{100, 100, "0.00"},
		{0, 100, "0"},
		{100, 0, "0"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("GapInt64(%d,%d)=%s", test.current, test.prev, test.expect), func(t *testing.T) {
			result := GapInt642String(test.current, test.prev)
			assert.Equal(t, test.expect, result)
		})
	}
}

func TestFloat64Mul(t *testing.T) {
	tests := []struct {
		a      float64
		b      float64
		r      int32
		expect float64
	}{
		{1.23, 4.56, 2, 5.61},
		{0, 4.56, 2, 0},
		{1.23, 0, 2, 0},
		{1.2345, 1.2345, 4, 1.5240},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%f*%f(r=%d)=%f", test.a, test.b, test.r, test.expect), func(t *testing.T) {
			result := Float64Mul(test.a, test.b, test.r)
			assert.InDelta(t, test.expect, result, 0.0001)
		})
	}
}

func TestDivPerInt64(t *testing.T) {
	tests := []struct {
		a      int64
		b      int64
		expect float64
	}{
		{50, 100, 50},
		{200, 100, 200},
		{0, 100, 0},
		{100, 0, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("DivPerInt64(%d,%d)=%f", test.a, test.b, test.expect), func(t *testing.T) {
			result := DivPerInt64(test.a, test.b)
			assert.InDelta(t, test.expect, result, 0.0001)
		})
	}
}

func TestDivPerInt64Multi000(t *testing.T) {
	tests := []struct {
		a      int64
		b      int64
		expect float64
	}{
		{50, 100, 500},
		{200, 100, 2000},
		{0, 100, 0},
		{100, 0, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("DivPerInt64Multi000(%d,%d)=%f", test.a, test.b, test.expect), func(t *testing.T) {
			result := DivPerInt64Multi000(test.a, test.b)
			assert.InDelta(t, test.expect, result, 0.0001)
		})
	}
}

func TestDivPerString(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		expect float64
	}{
		{"50", "100", 50},
		{"200", "100", 200},
		{"0", "100", 0},
		{"100", "0", 0},
		{"invalid", "100", 0},
		{"100", "invalid", 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("DivPerString(%s,%s)=%f", test.a, test.b, test.expect), func(t *testing.T) {
			result := DivPerString(test.a, test.b)
			assert.InDelta(t, test.expect, result, 0.0001)
		})
	}
}

func TestDivPerInt(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		expect float64
	}{
		{50, 100, 50},
		{200, 100, 200},
		{0, 100, 0},
		{100, 0, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("DivPerInt(%d,%d)=%f", test.a, test.b, test.expect), func(t *testing.T) {
			result := DivPerInt(test.a, test.b)
			assert.InDelta(t, test.expect, result, 0.0001)
		})
	}
}

func TestDivPerFloat64(t *testing.T) {
	tests := []struct {
		a      float64
		b      float64
		expect float64
	}{
		{50.5, 100.5, 50.2488},
		{200.5, 100.5, 199.5025},
		{0, 100.5, 0},
		{100.5, 0, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("DivPerFloat64(%f,%f)=%f", test.a, test.b, test.expect), func(t *testing.T) {
			result := DivPerFloat64(test.a, test.b)
			assert.InDelta(t, test.expect, result, 0.01)
		})
	}
}

func TestDivInt64(t *testing.T) {
	tests := []struct {
		a      int64
		b      int64
		r      int32
		expect float64
	}{
		{10, 3, 2, 3.33},
		{0, 3, 2, 0},
		{10, 0, 2, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("DivInt64(%d,%d,r=%d)=%f", test.a, test.b, test.r, test.expect), func(t *testing.T) {
			result := DivInt64(test.a, test.b, test.r)
			assert.InDelta(t, test.expect, result, 0.01)
		})
	}
}

func TestDivStringRound(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		r      int32
		expect float64
	}{
		{"10", "3", 2, 3.33},
		{"0", "3", 2, 0},
		{"10", "0", 2, 0},
		{"invalid", "3", 2, 0},
		{"10", "invalid", 2, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("DivStringRound(%s,%s,r=%d)=%f", test.a, test.b, test.r, test.expect), func(t *testing.T) {
			result := DivStringRound(test.a, test.b, test.r)
			assert.InDelta(t, test.expect, result, 0.01)
		})
	}
}

func TestDivFloat64(t *testing.T) {
	tests := []struct {
		a      float64
		b      float64
		r      int32
		expect float64
	}{
		{10.5, 3.5, 2, 3},
		{0, 3.5, 2, 0},
		{10.5, 0, 2, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("DivFloat64(%f,%f,r=%d)=%f", test.a, test.b, test.r, test.expect), func(t *testing.T) {
			result := DivFloat64(test.a, test.b, test.r)
			assert.InDelta(t, test.expect, result, 0.01)
		})
	}
}

func TestYuan2FenFloat64(t *testing.T) {
	tests := []struct {
		yuan   float64
		r      int32
		expect float64
	}{
		{1.23, 2, 123},
		{0, 2, 0},
		{0.01, 2, 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Yuan2FenFloat64(%f,r=%d)=%f", test.yuan, test.r, test.expect), func(t *testing.T) {
			result := Yuan2FenFloat64(test.yuan, test.r)
			assert.InDelta(t, test.expect, result, 0.01)
		})
	}
}

func TestFen2YuanFloat64(t *testing.T) {
	tests := []struct {
		fen    float64
		r      int32
		expect float64
	}{
		{123, 2, 1.23},
		{0, 2, 0},
		{1, 2, 0.01},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Fen2YuanFloat64(%f,r=%d)=%f", test.fen, test.r, test.expect), func(t *testing.T) {
			result := Fen2YuanFloat64(test.fen, test.r)
			assert.InDelta(t, test.expect, result, 0.01)
		})
	}
}

func TestFen2YuanString(t *testing.T) {
	tests := []struct {
		fen    string
		r      int32
		expect float64
	}{
		{"123", 2, 1.23},
		{"0", 2, 0},
		{"1", 2, 0.01},
		{"invalid", 2, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Fen2YuanString(%s,r=%d)=%f", test.fen, test.r, test.expect), func(t *testing.T) {
			result := Fen2YuanString(test.fen, test.r)
			assert.InDelta(t, test.expect, result, 0.01)
		})
	}
}

func TestDivString(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		r      int32
		expect float64
	}{
		{"10", "2", 2, 5},
		{"0", "2", 2, 0},
		{"10", "0", 2, 0},
		{"invalid", "2", 2, 0},
		{"10", "invalid", 2, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("DivString(%s,%s,r=%d)=%f", test.a, test.b, test.r, test.expect), func(t *testing.T) {
			result := DivString(test.a, test.b, test.r)
			assert.InDelta(t, test.expect, result, 0.01)
		})
	}
}

func TestDivInt(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		r      int32
		expect float64
	}{
		{10, 2, 2, 5},
		{0, 2, 2, 0},
		{10, 0, 2, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("DivInt(%d,%d,r=%d)=%f", test.a, test.b, test.r, test.expect), func(t *testing.T) {
			result := DivInt(test.a, test.b, test.r)
			assert.InDelta(t, test.expect, result, 0.01)
		})
	}
}

func TestSubFloat(t *testing.T) {
	tests := []struct {
		a      float64
		b      float64
		expect float64
	}{
		{10.5, 2.5, 8},
		{0, 2.5, -2.5},
		{10.5, 0, 10.5},
		{2.5, 10.5, -8},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("SubFloat(%f,%f)=%f", test.a, test.b, test.expect), func(t *testing.T) {
			result := SubFloat(test.a, test.b)
			assert.InDelta(t, test.expect, result, 0.01)
		})
	}
}

func TestSubInt64(t *testing.T) {
	tests := []struct {
		a      float64
		b      float64
		expect float64
	}{
		{10.5, 2.5, 8},
		{0, 2.5, -2.5},
		{10.5, 0, 10.5},
		{2.5, 10.5, -8},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("SubInt64(%f,%f)=%f", test.a, test.b, test.expect), func(t *testing.T) {
			result := SubInt64(test.a, test.b)
			assert.InDelta(t, test.expect, result, 0.01)
		})
	}
}

func TestFloat2String2f(t *testing.T) {
	tests := []struct {
		input  float64
		expect string
	}{
		{1.234, "1.23"},
		{1.235, "1.24"},
		{0, "0.00"},
		{-1.234, "-1.23"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Float2String2f(%f)=%s", test.input, test.expect), func(t *testing.T) {
			result := Float2String2f(test.input)
			assert.Equal(t, test.expect, result)
		})
	}
}

func TestString2Float64(t *testing.T) {
	tests := []struct {
		input  string
		expect float64
	}{
		{"1.234", 1.234},
		{"0", 0},
		{"-1.234", -1.234},
		{"invalid", 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("String2Float64(%s)=%f", test.input, test.expect), func(t *testing.T) {
			result := String2Float64(test.input)
			assert.InDelta(t, test.expect, result, 0.0001)
		})
	}
}

func TestFloat2String4f(t *testing.T) {
	tests := []struct {
		input  float64
		expect string
	}{
		{1.23456, "1.2346"},
		{1.23451, "1.2345"},
		{0, "0.0000"},
		{-1.23456, "-1.2346"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Float2String4f(%f)=%s", test.input, test.expect), func(t *testing.T) {
			result := Float2String4f(test.input)
			assert.Equal(t, test.expect, result)
		})
	}
}

func TestFloat64Round(t *testing.T) {
	tests := []struct {
		input  float64
		r      int32
		expect float64
	}{
		{1.23456, 2, 1.23},
		{1.23556, 2, 1.24},
		{0, 2, 0},
		{-1.23456, 2, -1.23},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Float64Round(%f,r=%d)=%f", test.input, test.r, test.expect), func(t *testing.T) {
			result := Float64Round(test.input, test.r)
			assert.InDelta(t, test.expect, result, 0.0001)
		})
	}
}

// 基准测试
func BenchmarkFloatStringAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FloatStringAdd("123.456", "789.012")
	}
}

func BenchmarkFloat64Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64Add(123.456, 789.012)
	}
}

func BenchmarkGapString2String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GapString2String("110", "100")
	}
}

func BenchmarkGapInt642String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GapInt642String(110, 100)
	}
}

func BenchmarkFloat64Round(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64Round(123.4567, 2)
	}
}
