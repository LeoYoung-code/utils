package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/rs/xid"
	"github.com/segmentio/ksuid"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	json "github.com/bytedance/sonic"
)

// GetStructName 获取当前函数所依赖名称 业务用于获取领域名称
func GetStructName(v any) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

// ByteToString 避免直接转换带了额外内存分配，使用断言转换string
func ByteToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

// ToString 序列化为json
func ToString(data any) string {
	b, _ := json.Marshal(data)
	return *(*string)(unsafe.Pointer(&b))
}

// IsInteger 判断字符串是否是整数
func IsInteger(s string) bool {
	_, err := strconv.ParseInt(s, 0, 64)
	return err == nil
}

// IsNumber 判断字符串是否是数字（包括整数和浮点数）
func IsNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsNum 判断字符串是否是数字 (已弃用，请使用 IsInteger 或 IsNumber)
func IsNum(s string) bool {
	return IsInteger(s)
}

// IsRepByLoop 判断切片中是否存在重复元素
func IsRepByLoop(origin []int64) error {
	seen := make(map[int64]struct{}, len(origin))
	for _, v := range origin {
		if _, exists := seen[v]; exists {
			return errors.New("切片存在重复元素")
		}
		seen[v] = struct{}{}
	}
	return nil
}

// IsStringIn 判断字符串切片中是否存在指定字符串
func IsStringIn(targetSlice []string, element string) bool {
	return ContainsString(targetSlice, element) != -1
}

// IsInt64In 判断数字是否在 指定数字切片中
func IsInt64In(element int64, targetSlice []int64) bool {
	for _, v := range targetSlice {
		if v == element {
			return true
		}
	}
	return false
}

// Contains 泛型函数，判断元素是否在切片中
func Contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func ContainsString(array []string, val string) (index int) {
	for i, item := range array {
		if item == val {
			return i
		}
	}
	return -1
}

// Ternary 三目运算的函数
func Ternary(a bool, b, c any) any {
	if a {
		return b
	}
	return c
}

// Explode 字符串转换字符串数组
func Explode(delimiter, text string) []string {
	if text == "" {
		return []string{}
	}
	return strings.Split(text, delimiter)
}

// MapStringVal 取map字段
func MapStringVal(sl map[string]string, k string) string {
	v, ok := sl[k]
	if ok {
		return v
	}
	return ""
}

func String2Int64(k string) int64 {
	i, err := strconv.ParseInt(k, 10, 64)
	if err != nil {
		return int64(0)
	}
	return i
}

func MapStringVal2Float64(sl map[string]string, k string) float64 {
	v, ok := sl[k]
	if ok {
		f, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return f
		}
	}
	return 0
}

func FilterEmpty(c []string) []string {
	return lo.Filter[string](c, func(v string, _ int) bool { return v != "" })
}

// GenerateUUID 生成uuid
func GenerateUUID() string {
	ksIdObj, err := ksuid.NewRandom()
	if err != nil {
		xIdObj := xid.New()
		return xIdObj.String()
	}
	return ksIdObj.String()
}

// UrlPath 获取url的path
func UrlPath(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}
	if u.RawQuery != "" {
		return u.Path + "?" + u.RawQuery
	}
	return u.Path
}

// IsOnline 判断是否在 上线状态 1上线   0下线
func IsOnline(upTime, downTime int64) int64 {
	var isOnline int64 = 0
	now := time.Now().Unix()
	if (upTime <= now && downTime > now) || (upTime < now && downTime >= now) {
		isOnline = 1
	}
	return isOnline
}

// Include 判断是否在切片中存在
func Include[T comparable](arr []T, check T) bool {
	for _, v := range arr {
		if v == check {
			return true
		}
	}
	return false
}

// IsUnion 判断是否存在交集
func IsUnion[T comparable](arr1 []T, arr2 []T) bool {
	m := make(map[T]struct{}, len(arr1))
	for _, v := range arr1 {
		m[v] = struct{}{}
	}
	for _, v := range arr2 {
		if _, ok := m[v]; ok {
			return true
		}
	}
	return false
}

// RandomElement 随机获取切片中的一个元素
func RandomElement(s []any) any {
	if len(s) == 0 {
		return nil
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(s))
	return s[index]
}

// ToAnySlice 将字符串切片转换为any切片
func ToAnySlice(s []string) []any {
	res := make([]any, len(s))
	for i := range s {
		res[i] = s[i]
	}
	return res
}

// Md5Parser md5解析
func Md5Parser(value string) int64 {
	// 计算md5
	hash := md5.Sum([]byte(value))
	md5Str := hex.EncodeToString(hash[:])

	// 取前5位转换成int
	prefix := md5Str[:5]
	i, err := strconv.ParseInt(prefix, 16, 32)
	if err != nil {
		return 0
	}
	i++

	// 取模
	return i % 40960
}
