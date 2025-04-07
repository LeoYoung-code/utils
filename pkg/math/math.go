package math

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

// FloatStringAdd Add returns d + d2.
func FloatStringAdd(d, d2 string) float64 {
	dd, _ := decimal.NewFromString(d)
	dd2, _ := decimal.NewFromString(d2)
	res, _ := dd.Add(dd2).Float64()
	return res
}

// Float64Add Div returns d + d2
func Float64Add(d, d2 float64) float64 {
	dd := decimal.NewFromFloat(d)
	dd2 := decimal.NewFromFloat(d2)
	res, _ := dd.Add(dd2).Float64()
	return res
}

// SubString Sub returns d - d2.
func SubString(d, d2 string) float64 {
	dd, _ := decimal.NewFromString(d)
	dd2, _ := decimal.NewFromString(d2)
	res, _ := dd.Sub(dd2).Float64()
	return res
}

// StringDiv100 除以100
func StringDiv100(d string) string {
	dd, _ := decimal.NewFromString(d)
	return dd.Div(decimal.NewFromInt(100)).String()
}

// GapString2String 计算涨跌幅  就是用现价除以昨收价 ，减去1
func GapString2String(d, d2 string) string {
	if d == "0" || d2 == "0" {
		return "0"
	}
	dd, _ := decimal.NewFromString(d)
	dd2, _ := decimal.NewFromString(d2)
	if dd2.IsZero() {
		return "0"
	}
	f, _ := dd.Div(dd2).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100)).Float64()
	return fmt.Sprintf("%.2f", f)
}

// GapInt642String 计算涨跌幅  就是用现价除以昨收价 ，减去1
func GapInt642String(d, d2 int64) string {
	if d == 0 || d2 == 0 {
		return "0"
	}
	dd := decimal.NewFromInt(d)
	dd2 := decimal.NewFromInt(d2)
	if dd2.IsZero() {
		return "0"
	}
	f, _ := dd.Div(dd2).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100)).Float64()
	return fmt.Sprintf("%.2f", f)
}

// Float64Mul  returns d * d2.
func Float64Mul(d, d2 float64, r int32) float64 {
	dd := decimal.NewFromFloat(d)
	dd2 := decimal.NewFromFloat(d2)
	res, _ := dd.Mul(dd2).Round(r).Float64()
	return res
}

// DivPerInt64 DivPerInt Div returns d / d2 * 100. If it doesn't divide exactly, the result will have
func DivPerInt64(d, d2 int64) float64 {
	if d2 == 0 {
		return float64(0)
	}
	dd := decimal.NewFromInt(d).Mul(decimal.NewFromInt(100))
	dd2 := decimal.NewFromInt(d2)
	res, _ := dd.Div(dd2).Float64()
	return res
}

// DivPerInt64Multi000  Div returns d / d2 * 100. If it doesn't divide exactly, the result will have
func DivPerInt64Multi000(d, d2 int64) float64 {
	if d2 == 0 {
		return float64(0)
	}
	dd := decimal.NewFromInt(d).Mul(decimal.NewFromInt(1000))
	dd2 := decimal.NewFromInt(d2)
	res, _ := dd.Div(dd2).Float64()
	return res
}

// DivPerString DivPerInt Div returns d / d2 * 100. If it doesn't divide exactly, the result will have
func DivPerString(d, d2 string) float64 {
	dd, _ := decimal.NewFromString(d)
	dd2, _ := decimal.NewFromString(d2)
	if dd2.IsZero() {
		return float64(0)
	}
	res, _ := dd.Mul(decimal.NewFromInt(100)).Div(dd2).Float64()
	return res
}

// DivPerInt Div returns d / d2 * 100. If it doesn't divide exactly, the result will have
func DivPerInt(d, d2 int) float64 {
	if d2 == 0 {
		return float64(0)
	}
	dd := decimal.NewFromInt(int64(d))
	dd2 := decimal.NewFromInt(int64(d2))
	res, _ := dd.Mul(decimal.NewFromInt(100)).Div(dd2).Float64()
	return res
}

// DivPerFloat64 Div returns d / d2 * 100. If it doesn't divide exactly, the result will have
func DivPerFloat64(d, d2 float64) float64 {
	if d2 == 0 {
		return float64(0)
	}
	dd := decimal.NewFromFloat(d)
	dd2 := decimal.NewFromFloat(d2)
	res, _ := dd.Mul(decimal.NewFromInt(100)).Div(dd2).Float64()
	return res
}

// DivInt64 DivPerInt Div returns d / d2.
func DivInt64(d, d2 int64, r int32) float64 {
	if d2 == 0 {
		return float64(0)
	}
	dd := decimal.NewFromInt(d)
	dd2 := decimal.NewFromInt(d2)
	res, _ := dd.Div(dd2).Round(r).Float64()
	return res
}

// DivStringRound DivPerInt Div returns d / d2.
func DivStringRound(d, d2 string, r int32) float64 {
	dd, _ := decimal.NewFromString(d)
	if dd.IsZero() {
		return float64(0)
	}
	dd2, _ := decimal.NewFromString(d2)
	res, _ := dd.Div(dd2).Round(r).Float64()
	return res
}

// DivFloat64 Div returns d / d2. If it doesn't divide exactly, the result will have
func DivFloat64(d, d2 float64, r int32) float64 {
	if d2 == 0 {
		return float64(0)
	}
	dd := decimal.NewFromFloat(d)
	dd2 := decimal.NewFromFloat(d2)
	res, _ := dd.Div(dd2).Round(r).Float64()
	return res
}

// Yuan2FenFloat64 元 -> 分
func Yuan2FenFloat64(yuan float64, r int32) float64 {
	return Float64Mul(yuan, float64(100), r)
}

// Fen2YuanFloat64 分 -> 元
func Fen2YuanFloat64(fen float64, r int32) float64 {
	return DivFloat64(fen, float64(100), r)
}

// Yuan2FenString 元 -> 分
func Fen2YuanString(fen string, r int32) float64 {
	f, _ := decimal.NewFromString(fen)
	d := decimal.NewFromInt(100)
	res, _ := f.Div(d).Round(r).Float64()
	return res
}

// DivString DivPerInt Div returns d / d2. If it doesn't divide exactly, the result will have
func DivString(d, d2 string, r int32) float64 {
	dd, _ := decimal.NewFromString(d)
	dd2, _ := decimal.NewFromString(d2)
	if dd2.IsZero() {
		return float64(0)
	}
	res, _ := dd.Div(dd2).Round(r).Float64()
	return res
}

// DivInt Div returns d / d2. If it doesn't divide exactly, the result will have
func DivInt(d, d2 int, r int32) float64 {
	if d2 == 0 {
		return float64(0)
	}
	dd := decimal.NewFromInt(int64(d))
	dd2 := decimal.NewFromInt(int64(d2))
	res, _ := dd.Div(dd2).Round(r).Float64()
	return res
}

// SubFloat Sub returns d - d2.
func SubFloat(d, d2 float64) float64 {
	dd := decimal.NewFromFloat(d)
	dd2 := decimal.NewFromFloat(d2)
	res, _ := dd.Sub(dd2).Float64()
	return res
}

// SubInt64 Sub returns d - d2.
func SubInt64(d, d2 float64) float64 {
	dd := decimal.NewFromFloat(d)
	dd2 := decimal.NewFromFloat(d2)
	res, _ := dd.Sub(dd2).Float64()
	return res
}

// SubString Sub returns d - d2.
func Float2String2f(s float64) string {
	return fmt.Sprintf("%.2f", s)
}

// Float2String2f 保留2位小数
func String2Float64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return float64(0)
	}
	return f
}

// Float2String4f 保留4位小数
func Float2String4f(s float64) string {
	return fmt.Sprintf("%.4f", s)
}

// Float64Round 保留几位小数
func Float64Round(f float64, r int32) float64 {
	res, _ := decimal.NewFromFloat(f).Round(r).Float64()
	return res
}
