// Package pkg 提供了所有工具包的集中引用点
package pkg

import (
	commonpkg "utils/pkg/common"
	utilspkg "utils/pkg/common" // util.go 的包名为 utils
	atomicpkg "utils/pkg/concurrent/atomic"
	poolpkg "utils/pkg/concurrent/pool"
	"utils/pkg/database"
	stringpkg "utils/pkg/string"
	timepkg "utils/pkg/time"
)

// 版本信息
const Version = "2.0.0"

// 导出常用的函数和类型，方便用户直接使用

// Common 导出通用工具函数
var (
	// 通用工具函数 (来自 pkg/common/util.go，包名为 utils)
	GetStructName  = utilspkg.GetStructName
	ByteToString   = utilspkg.ByteToString
	ToString       = utilspkg.ToString
	IsNum          = utilspkg.IsNum
	IsStringIn     = utilspkg.IsStringIn
	IsInt64In      = utilspkg.IsInt64In
	ContainsString = utilspkg.ContainsString
	Ternary        = utilspkg.Ternary
	Explode        = utilspkg.Explode
	FilterEmpty    = utilspkg.FilterEmpty
	GenerateUUID   = utilspkg.GenerateUUID
	UrlPath        = utilspkg.UrlPath
	IsOnline       = utilspkg.IsOnline

	// 对于泛型函数，我们提供一些常用类型的实例化版本
	IncludeString = func(arr []string, check string) bool {
		return utilspkg.Include(arr, check)
	}
	IncludeInt = func(arr []int, check int) bool {
		return utilspkg.Include(arr, check)
	}
	IncludeInt64 = func(arr []int64, check int64) bool {
		return utilspkg.Include(arr, check)
	}

	IsUnionString = func(arr1, arr2 []string) bool {
		return utilspkg.IsUnion(arr1, arr2)
	}
	IsUnionInt = func(arr1, arr2 []int) bool {
		return utilspkg.IsUnion(arr1, arr2)
	}
	IsUnionInt64 = func(arr1, arr2 []int64) bool {
		return utilspkg.IsUnion(arr1, arr2)
	}

	RandomElement = utilspkg.RandomElement
	ToAnySlice    = utilspkg.ToAnySlice
	Md5Parser     = utilspkg.Md5Parser

	// 来自 pkg/common/common.go，包名为 common
	Num2Version  = commonpkg.Num2Version
	GetSum32     = commonpkg.GetSum32
	UntilSuccess = commonpkg.UntilSuccess
	IsCancelled  = commonpkg.IsCancelled
)

// 导出其他常用功能

// Concurrent 导出并发相关工具
var (
	// 并发工具
	NewGoPool   = poolpkg.NewGoPool
	NewAtomicID = atomicpkg.NewID
)

// String 导出字符串相关工具
var (
	// 字符串工具
	StrToBytes       = stringpkg.StrToBytes
	TrimSpace        = stringpkg.TrimSpace
	Split            = stringpkg.Split
	Join             = stringpkg.Join
	ToUpper          = stringpkg.ToUpper
	ToLower          = stringpkg.ToLower
	Contains         = stringpkg.Contains
	CamelToSnake     = stringpkg.CamelToSnake
	SnakeToCamel     = stringpkg.SnakeToCamel
	FirstCharToUpper = stringpkg.FirstCharToUpper
	// 新增导出的优化函数
	UcFirst              = stringpkg.UcFirst              // 首字母大写
	LcFirst              = stringpkg.LcFirst              // 首字母小写
	GoSanitized          = stringpkg.GoSanitized          // 转换为有效的Go标识符
	BytesToString        = stringpkg.BytesToString        // 字节切片转字符串（零拷贝）
	StringToBytes        = stringpkg.StringToBytes        // 字符串转字节切片（零拷贝）
	GenerateRandomString = stringpkg.GenerateRandomString // 生成随机字符串
)

// Time 导出时间相关工具
var (
	// 时间格式常量
	DateFormat       = timepkg.DateFormat
	TimeFormat       = timepkg.TimeFormat
	DateTimeFormat   = timepkg.DateTimeFormat
	DateTimeFormatT  = timepkg.DateTimeFormatT
	DateTimeFormatZ  = timepkg.DateTimeFormatZ
	DateTimeFormatTZ = timepkg.DateTimeFormatTZ

	// 时间工具函数
	FormatTime         = timepkg.FormatTime
	ParseTime          = timepkg.ParseTime
	GetCurrentTime     = timepkg.GetCurrentTime
	GetCurrentDate     = timepkg.GetCurrentDate
	GetCurrentMonth    = timepkg.GetCurrentMonth
	GetFirstDayOfMonth = timepkg.GetFirstDayOfMonth
	GetLastDayOfMonth  = timepkg.GetLastDayOfMonth
	WeekStart          = timepkg.WeekStart
	WeekEnd            = timepkg.WeekEnd
	BeginningOfDay     = timepkg.BeginningOfDay
	EndOfDay           = timepkg.EndOfDay
	DiffDays           = timepkg.DiffDays
	IsSameDay          = timepkg.IsSameDay
)

// Database 导出数据库相关工具
var (
	// 数据库工具
	Transaction = database.Transaction
)
