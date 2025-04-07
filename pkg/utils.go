// Package pkg 提供了所有工具包的集中引用点
package pkg

import (
	"utils/pkg/common"
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
	// 通用工具函数
	GetStructName  = common.GetStructName
	ByteToString   = common.ByteToString
	ToString       = common.ToString
	IsNum          = common.IsNum
	IsStringIn     = common.IsStringIn
	IsInt64In      = common.IsInt64In
	ContainsString = common.ContainsString
	Ternary        = common.Ternary
	Explode        = common.Explode
	FilterEmpty    = common.FilterEmpty
	GenerateUUID   = common.GenerateUUID
	UrlPath        = common.UrlPath
	IsOnline       = common.IsOnline

	// 对于泛型函数，我们提供一些常用类型的实例化版本
	IncludeString = func(arr []string, check string) bool {
		return common.Include(arr, check)
	}
	IncludeInt = func(arr []int, check int) bool {
		return common.Include(arr, check)
	}
	IncludeInt64 = func(arr []int64, check int64) bool {
		return common.Include(arr, check)
	}

	IsUnionString = func(arr1, arr2 []string) bool {
		return common.IsUnion(arr1, arr2)
	}
	IsUnionInt = func(arr1, arr2 []int) bool {
		return common.IsUnion(arr1, arr2)
	}
	IsUnionInt64 = func(arr1, arr2 []int64) bool {
		return common.IsUnion(arr1, arr2)
	}

	RandomElement = common.RandomElement
	ToAnySlice    = common.ToAnySlice
	Md5Parser     = common.Md5Parser
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
