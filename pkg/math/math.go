package math

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

// FloatStringAdd 将两个字符串表示的浮点数相加，返回 d + d2
func FloatStringAdd(d, d2 string) float64 {
	dd, err := decimal.NewFromString(d)
	if err != nil {
		return 0
	}
	dd2, err := decimal.NewFromString(d2)
	if err != nil {
		return 0
	}
	res, _ := dd.Add(dd2).Float64()
	return res
}

// Float64Add 将两个float64相加，返回 d + d2
func Float64Add(d, d2 float64) float64 {
	dd := decimal.NewFromFloat(d)
	dd2 := decimal.NewFromFloat(d2)
	res, _ := dd.Add(dd2).Float64()
	return res
}

// SubString 将两个字符串表示的浮点数相减，返回 d - d2
func SubString(d, d2 string) float64 {
	dd, err := decimal.NewFromString(d)
	if err != nil {
		return 0
	}
	dd2, err := decimal.NewFromString(d2)
	if err != nil {
		return 0
	}
	res, _ := dd.Sub(dd2).Float64()
	return res
}

// StringDiv100 将字符串表示的数除以100
func StringDiv100(d string) string {
	dd, err := decimal.NewFromString(d)
	if err != nil {
		return "0"
	}
	return dd.Div(decimal.NewFromInt(100)).String()
}

// GapString2String 计算涨跌幅：(现价/昨收价-1)*100%，返回保留2位小数的字符串
func GapString2String(d, d2 string) string {
	if d == "0" || d2 == "0" {
		return "0"
	}
	dd, err := decimal.NewFromString(d)
	if err != nil {
		return "0"
	}
	dd2, err := decimal.NewFromString(d2)
	if err != nil {
		return "0"
	}
	if dd2.IsZero() {
		return "0"
	}
	f, _ := dd.Div(dd2).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100)).Float64()
	return fmt.Sprintf("%.2f", f)
}

// GapInt642String 使用int64类型计算涨跌幅：(现价/昨收价-1)*100%，返回保留2位小数的字符串
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

// Float64Mul 将两个float64相乘，返回保留r位小数的结果
func Float64Mul(d, d2 float64, r int32) float64 {
	dd := decimal.NewFromFloat(d)
	dd2 := decimal.NewFromFloat(d2)
	res, _ := dd.Mul(dd2).Round(r).Float64()
	return res
}

// DivPerInt64 计算百分比：(d/d2)*100，使用int64类型输入
func DivPerInt64(d, d2 int64) float64 {
	if d2 == 0 {
		return 0
	}
	dd := decimal.NewFromInt(d).Mul(decimal.NewFromInt(100))
	dd2 := decimal.NewFromInt(d2)
	res, _ := dd.Div(dd2).Float64()
	return res
}

// DivPerInt64Multi000 计算千分比：(d/d2)*1000，使用int64类型输入
func DivPerInt64Multi000(d, d2 int64) float64 {
	if d2 == 0 {
		return 0
	}
	dd := decimal.NewFromInt(d).Mul(decimal.NewFromInt(1000))
	dd2 := decimal.NewFromInt(d2)
	res, _ := dd.Div(dd2).Float64()
	return res
}

// DivPerString 计算百分比：(d/d2)*100，使用字符串类型输入
func DivPerString(d, d2 string) float64 {
	dd, err := decimal.NewFromString(d)
	if err != nil {
		return 0
	}
	dd2, err := decimal.NewFromString(d2)
	if err != nil {
		return 0
	}
	if dd2.IsZero() {
		return 0
	}
	res, _ := dd.Mul(decimal.NewFromInt(100)).Div(dd2).Float64()
	return res
}

// DivPerInt 计算百分比：(d/d2)*100，使用int类型输入
func DivPerInt(d, d2 int) float64 {
	if d2 == 0 {
		return 0
	}
	dd := decimal.NewFromInt(int64(d))
	dd2 := decimal.NewFromInt(int64(d2))
	res, _ := dd.Mul(decimal.NewFromInt(100)).Div(dd2).Float64()
	return res
}

// DivPerFloat64 计算百分比：(d/d2)*100，使用float64类型输入
func DivPerFloat64(d, d2 float64) float64 {
	if d2 == 0 {
		return 0
	}
	dd := decimal.NewFromFloat(d)
	dd2 := decimal.NewFromFloat(d2)
	res, _ := dd.Mul(decimal.NewFromInt(100)).Div(dd2).Float64()
	return res
}

// DivInt64 返回d/d2的商，保留r位小数，使用int64类型输入
func DivInt64(d, d2 int64, r int32) float64 {
	if d2 == 0 {
		return 0
	}
	dd := decimal.NewFromInt(d)
	dd2 := decimal.NewFromInt(d2)
	res, _ := dd.Div(dd2).Round(r).Float64()
	return res
}

// DivStringRound 返回d/d2的商，保留r位小数，使用字符串类型输入
func DivStringRound(d, d2 string, r int32) float64 {
	dd, err := decimal.NewFromString(d)
	if err != nil {
		return 0
	}
	dd2, err := decimal.NewFromString(d2)
	if err != nil {
		return 0
	}
	if dd2.IsZero() {
		return 0
	}
	res, _ := dd.Div(dd2).Round(r).Float64()
	return res
}

// DivFloat64 返回d/d2的商，保留r位小数，使用float64类型输入
func DivFloat64(d, d2 float64, r int32) float64 {
	if d2 == 0 {
		return 0
	}
	dd := decimal.NewFromFloat(d)
	dd2 := decimal.NewFromFloat(d2)
	res, _ := dd.Div(dd2).Round(r).Float64()
	return res
}

// Yuan2FenFloat64 将元转换为分（乘以100），保留r位小数
func Yuan2FenFloat64(yuan float64, r int32) float64 {
	return Float64Mul(yuan, 100, r)
}

// Fen2YuanFloat64 将分转换为元（除以100），保留r位小数
func Fen2YuanFloat64(fen float64, r int32) float64 {
	return DivFloat64(fen, 100, r)
}

// Fen2YuanString 将字符串表示的分转换为元（除以100），保留r位小数
func Fen2YuanString(fen string, r int32) float64 {
	f, err := decimal.NewFromString(fen)
	if err != nil {
		return 0
	}
	d := decimal.NewFromInt(100)
	res, _ := f.Div(d).Round(r).Float64()
	return res
}

// DivString 返回d/d2的商，保留r位小数，使用字符串类型输入
func DivString(d, d2 string, r int32) float64 {
	dd, err := decimal.NewFromString(d)
	if err != nil {
		return 0
	}
	dd2, err := decimal.NewFromString(d2)
	if err != nil {
		return 0
	}
	if dd2.IsZero() {
		return 0
	}
	res, _ := dd.Div(dd2).Round(r).Float64()
	return res
}

// DivInt 返回d/d2的商，保留r位小数，使用int类型输入
func DivInt(d, d2 int, r int32) float64 {
	if d2 == 0 {
		return 0
	}
	dd := decimal.NewFromInt(int64(d))
	dd2 := decimal.NewFromInt(int64(d2))
	res, _ := dd.Div(dd2).Round(r).Float64()
	return res
}

// SubFloat 返回d-d2的差，使用float64类型输入
func SubFloat(d, d2 float64) float64 {
	dd := decimal.NewFromFloat(d)
	dd2 := decimal.NewFromFloat(d2)
	res, _ := dd.Sub(dd2).Float64()
	return res
}

// SubInt64 返回d-d2的差，使用float64类型输入（函数名与实现不符，保留原函数名以兼容旧代码）
func SubInt64(d, d2 float64) float64 {
	dd := decimal.NewFromFloat(d)
	dd2 := decimal.NewFromFloat(d2)
	res, _ := dd.Sub(dd2).Float64()
	return res
}

// Float2String2f 将float64转为保留2位小数的字符串
func Float2String2f(s float64) string {
	return fmt.Sprintf("%.2f", s)
}

// String2Float64 将字符串转为float64，错误时返回0
func String2Float64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

// Float2String4f 将float64转为保留4位小数的字符串
func Float2String4f(s float64) string {
	return fmt.Sprintf("%.4f", s)
}

// Float64Round 将float64四舍五入到r位小数
func Float64Round(f float64, r int32) float64 {
	res, _ := decimal.NewFromFloat(f).Round(r).Float64()
	return res
}
