package time

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/now"
)

const (
	TimeLayout      = "2006-01-02 15:04:05"
	DateLayout      = "2006-01-02"
	YmdLayout       = "20060102"
	YearMonthLayout = "200601"
	HourLayout      = "15:04"
	MinLayout       = "15:04"
	HourMinLayout   = "15:04:05"
)

var location, _ = time.LoadLocation("Asia/Shanghai")

func parseTimeWithLayout(timeStr, layout string) (time.Time, bool) {
	if timeStr == "" {
		return time.Time{}, false
	}
	t, err := time.ParseInLocation(layout, timeStr, location)
	if err != nil {
		log.Error(err.Error())
		return time.Time{}, false
	}
	return t, true
}

func TimeToString(t *time.Time) string {
	// 获取指定时间
	if t != nil {
		if t.Year() == 1970 {
			return ""
		}
		return t.Format(TimeLayout)
	}
	// 获取当前时间
	return time.Now().Format(TimeLayout)
}

// TimeToDateString 2006-01-02
func TimeToDateString(t *time.Time) string {
	// 获取指定时间
	if t != nil {
		if t.Year() == 1970 {
			return ""
		}
		return t.Format(DateLayout)
	}
	// 获取当前时间
	return time.Now().Format(DateLayout)
}

// DateStringToYmdInt 2006-01-02(string) -> 20060102(int64)
func DateStringToYmdInt(dt string) int64 {
	t, ok := parseTimeWithLayout(dt, DateLayout)
	if !ok {
		return 0
	}
	ymd, err := strconv.ParseInt(t.Format(YmdLayout), 10, 64)
	if err != nil {
		log.Error(err.Error())
		return 0
	}
	return ymd
}

// YesterdayDateString 昨天Ymd
func YesterdayDateString(layout string) string {
	loc := time.Now()
	return time.Date(loc.Year(), loc.Month(), loc.Day()-1, 0, 0, 0, 0, location).Format(layout)
}

// YmdIntToDateString 20060102(int64) -> 2006-01-02(string)
func YmdIntToDateString(dt int64) string {
	if dt == 0 {
		return "-"
	}
	t, ok := parseTimeWithLayout(strconv.FormatInt(dt, 10), YmdLayout)
	if !ok {
		return "-"
	}
	return t.Format(DateLayout)
}

// CurrentMonth 当月起始日期
func CurrentMonth() (string, string) {
	loc := time.Now()
	firstOfMonth := time.Date(loc.Year(), loc.Month(), 1, 0, 0, 0, 0, location)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	return firstOfMonth.Format(YmdLayout), lastOfMonth.Format(YmdLayout)
}

// CurrentYearMonth 200601
func CurrentYearMonth() string {
	return time.Now().Format(YearMonthLayout)
}

// IntToTime 时间戳转时间对象
func IntToTime(ts int64) time.Time {
	return time.Unix(ts, 0)
}

// StringToTime 通用的字符串转时间函数
func StringToTime(ts string, layout string) (time.Time, bool) {
	return parseTimeWithLayout(ts, layout)
}

// TimeStringToTime 标准时间字符串转Time
func TimeStringToTime(ts string) time.Time {
	t, ok := parseTimeWithLayout(ts, TimeLayout)
	if !ok {
		return time.Now()
	}
	return t
}

// DateStringToTime 日期字符串转Time
func DateStringToTime(ts string) time.Time {
	t, ok := parseTimeWithLayout(ts, DateLayout)
	if !ok {
		return time.Now()
	}
	return t
}

// Int2Time 时间戳转 时间字符串 "Y-m-d H:i:s"
func Int2Time(ts int64) string {
	if ts == 0 {
		return ""
	}
	t := IntToTime(ts)
	return TimeToString(&t)
}

// Int2Date 时间字符串 "Y-m-d"
func Int2Date(ts int64) string {
	if ts == 0 {
		return ""
	}
	t := IntToTime(ts)
	return TimeToDateString(&t)
}

// StringToUnixTime 时间字符串转时间戳
func StringToUnixTime(ts string) int64 {
	t, ok := parseTimeWithLayout(ts, TimeLayout)
	if !ok {
		return 0
	}
	return t.Unix()
}

// DateStringToUnixTime 日期字符串转时间戳
func DateStringToUnixTime(ts string) int64 {
	t, ok := parseTimeWithLayout(ts, DateLayout)
	if !ok {
		return 0
	}
	return t.Unix()
}

// FormatDurationTxt 格式化时间段为可读文本
func FormatDurationTxt(startTime, endTime int64) string {
	if startTime == 0 || endTime == 0 {
		return "不限"
	}
	txt := fmt.Sprintf("%s - %s", IntToTime(startTime).Format(DateLayout), IntToTime(endTime).Format(DateLayout))
	if strings.TrimSpace(txt) == "" {
		return "不限"
	}
	return txt
}

// Hour2DateTime hour转时间点
// 当date为""时：  1=> 01:00、 2=> 02:00、 ... 、 23=>23:00
// 当date为"2022-11-11"时:   1=> 2022-11-11 01:00、 2=> 2022-11-11 02:00、 ... 、 23=>2022-11-11 23:00
func Hour2DateTime(date string, hour int) string {
	hourTime := time.Date(0, 0, 0, hour, 0, 0, 0, location).Format(HourLayout)
	if len(date) > 0 {
		return date + " " + hourTime
	}
	return hourTime
}

// Point2DateTime point转时间点（1个point代表5分钟）
// 当date为""时：   0 => 00:00 、 1 => 00:05、 ... 、287 => 23:55
// 当date为"2022-11-11"时：  0 => 2022-11-11 00:00 、 1 => 2022-11-11 00:05、 ... 、287 => 2022-11-11 23:55
func Point2DateTime(date string, point int) string {
	min := point * 5
	hourMin := time.Date(0, 0, 0, 0, min, 0, 0, location).Format(MinLayout)
	if len(date) > 0 {
		return date + " " + hourMin
	}
	return hourMin
}

// Time2Point 当前时间转换为point值
func Time2Point(date string) int64 {
	ts := time.Now()
	point := ts.Hour()*12 + ts.Minute()/5
	return int64(point)
}

// DiffDateOfDay 计算两个日期（Ymd）间隔天数
func DiffDateOfDay(start, end string) float64 {
	t1, ok1 := parseTimeWithLayout(start, DateLayout)
	t2, ok2 := parseTimeWithLayout(end, DateLayout)
	if !ok1 || !ok2 {
		return 0
	}
	return (t2.Sub(t1).Hours() / 24) + 1
}

// TodayLastTime 距离明天0点的时间（用于缓存自然天）
func TodayLastTime() time.Duration {
	loc := time.Now()
	tomorrow := time.Date(loc.Year(), loc.Month(), loc.Day()+1, 0, 0, 0, 0, location)
	return tomorrow.Sub(loc)
}

// FormatDuration 格式化时间范围为时间戳
func FormatDuration(startTime, endTime string) (int64, int64) {
	if startTime == "" || endTime == "" {
		return 0, 0
	}
	s := now.New(DateStringToTime(startTime)).BeginningOfDay().Unix()
	e := now.New(DateStringToTime(endTime)).EndOfDay().Unix()
	return s, e
}

// FormatDuration2Time 格式化时间范围为时间对象
func FormatDuration2Time(startTime, endTime string) (time.Time, time.Time) {
	if startTime == "" || endTime == "" {
		return time.Now(), time.Now()
	}
	s := now.New(DateStringToTime(startTime)).BeginningOfDay()
	e := now.New(DateStringToTime(endTime)).EndOfDay()
	return s, e
}

// IsNowTimeIn 现在的时候是否在某个时间段内
func IsNowTimeIn(startTime, endTime int64) bool {
	if endTime == 0 {
		return true
	}
	now := time.Now().Unix()
	if endTime > 0 && now > endTime {
		return false
	}

	if startTime > 0 && now < startTime {
		return false
	}
	return true
}

// GetZeroTime 获取零时间
func GetZeroTime() time.Time {
	return time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
}
