package math

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// parseDecimal 将字符串转换为decimal.Decimal
// 如果转换失败返回nil
func parseDecimal(s string) *decimal.Decimal {
	dd, err := decimal.NewFromString(s)
	if err != nil {
		return nil
	}
	return &dd
}

// safeDiv 执行安全除法运算
// 如果除数为零返回nil
func safeDiv(dividend, divisor *decimal.Decimal) *decimal.Decimal {
	if divisor == nil || divisor.IsZero() {
		return nil
	}
	result := dividend.Div(*divisor)
	return &result
}

// FloatStringAdd Add returns d + d2.
func FloatStringAdd(d, d2 string) float64 {
	dd := parseDecimal(d)
	dd2 := parseDecimal(d2)
	if dd == nil || dd2 == nil {
		return 0
	}
	res, _ := dd.Add(*dd2).Float64()
	return res
}

// SubString Sub returns d - d2.
func SubString(d, d2 string) float64 {
	dd := parseDecimal(d)
	dd2 := parseDecimal(d2)
	if dd == nil || dd2 == nil {
		return 0
	}
	res, _ := dd.Sub(*dd2).Float64()
	return res
}

// StringDiv100 除以100
func StringDiv100(d string) string {
	dd := parseDecimal(d)
	if dd == nil {
		return "0"
	}
	return dd.Div(decimal.NewFromInt(100)).String()
}

// GapString2String 计算涨跌幅  就是用现价除以昨收价 ，减去1
func GapString2String(d, d2 string) string {
	dd := parseDecimal(d)
	dd2 := parseDecimal(d2)
	if dd == nil || dd2 == nil || dd2.IsZero() {
		return "0"
	}
	f := dd.Div(*dd2).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100))
	result, _ := f.Float64()
	return fmt.Sprintf("%.2f", result)
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
	f := dd.Div(dd2).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100))
	result, _ := f.Float64()
	return fmt.Sprintf("%.2f", result)
}

// DivPerString DivPerInt Div returns d / d2 * 100. If it doesn't divide exactly, the result will have
func DivPerString(d, d2 string) float64 {
	dd := parseDecimal(d)
	dd2 := parseDecimal(d2)
	if dd == nil || dd2 == nil || dd2.IsZero() {
		return 0
	}
	result := dd.Mul(decimal.NewFromInt(100)).Div(*dd2)
	f, _ := result.Float64()
	return f
}

// DivString DivPerInt Div returns d / d2. If it doesn't divide exactly, the result will have
func DivString(d, d2 string, r int32) float64 {
	dd := parseDecimal(d)
	dd2 := parseDecimal(d2)
	if dd == nil || dd2 == nil || dd2.IsZero() {
		return 0
	}
	result := dd.Div(*dd2).Round(r)
	f, _ := result.Float64()
	return f
}

// DivInt Div returns d / d2. If it doesn't divide exactly, the result will have
func DivInt(d, d2 int, r int32) float64 {
	if d2 == 0 {
		return 0
	}
	dd := decimal.NewFromInt(int64(d))
	dd2 := decimal.NewFromInt(int64(d2))
	result := dd.Div(dd2).Round(r)
	f, _ := result.Float64()
	return f
}

// Fen2YuanString 元 -> 分
func Fen2YuanString(fen string, r int32) float64 {
	f := parseDecimal(fen)
	if f == nil {
		return 0
	}
	result := f.Div(decimal.NewFromInt(100)).Round(r)
	fRes, _ := result.Float64()
	return fRes
}

// Float2String2f 保留2位小数
func Float2String2f(s float64) string {
	return fmt.Sprintf("%.2f", s)
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
