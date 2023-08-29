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
	HOUR_MIN_LAYOUT = "15:04:05"
)

var location, _ = time.LoadLocation("Asia/Shanghai")

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
	if len(dt) == 0 {
		return 0
	}
	t, err := time.ParseInLocation(DateLayout, dt, location)
	if err != nil {
		log.Error(err.Error())
		return 0
	}
	ymd, err2 := strconv.ParseInt(t.Format(YmdLayout), 10, 64)
	if err2 != nil {
		log.Error(err2.Error())
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
	t, err := time.ParseInLocation(YmdLayout, strconv.FormatInt(dt, 10), location)
	if err != nil {
		log.Error(err.Error())
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

func IntToTime(ts int64) time.Time {
	a := time.Unix(ts, 0)
	return a
}

func TimeStringToTime(ts string) time.Time {
	t, err := time.ParseInLocation(TimeLayout, ts, location)
	if err != nil {
		log.Error(err.Error())
		return time.Now()
	}
	return t
}

func TimeString2Time(ts string, layout string) time.Time {
	t, err := time.ParseInLocation(layout, ts, location)
	if err != nil {
		log.Error(err.Error())
		return time.Now()
	}
	return t
}

func DateStringToTime(ts string) time.Time {
	t, err := time.ParseInLocation(DateLayout, ts, location)
	if err != nil {
		log.Error(err.Error())
		return time.Now()
	}
	return t
}

// Int2Time 时间戳转 时间字符串 "Y-m-d H:i:s"
func Int2Time(ts int64) string {
	if ts == 0 {
		return ""
	}
	// return time.go.Unix(ts, 0).Format(TIME_LAYOUT)
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

func StringToUnixTime(ts string) int64 {
	t := TimeStringToTime(ts)
	return t.Unix()
}

func DateStringToUnixTime(ts string) int64 {
	if ts == "" {
		return 0
	}
	t := DateStringToTime(ts)
	return t.Unix()
}

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

func Time2Point(date string) int64 {
	ts := time.Now()
	point := ts.Hour()*12 + ts.Minute()/5
	return int64(point)
}

// DiffDateOfDay 计算两个日期（Ymd）间隔天数
func DiffDateOfDay(start, end string) float64 {
	t1, _ := time.Parse(DateLayout, start)
	t2, _ := time.Parse(DateLayout, end)
	return (t2.Sub(t1).Hours() / 24) + 1
}

// TodayLastTime 距离明天0点的时间（用于缓存自然天）
func TodayLastTime() time.Duration {
	loc := time.Now()
	tomorrow := time.Date(loc.Year(), loc.Month(), loc.Day()+1, 0, 0, 0, 0, location)
	return tomorrow.Sub(loc)
}

// FormatDuration 格式化时间
func FormatDuration(startTime, endTime string) (int64, int64) {
	if startTime == "" || endTime == "" {
		return 0, 0
	}
	s := now.New(DateStringToTime(startTime)).BeginningOfDay().Unix()
	e := now.New(DateStringToTime(endTime)).EndOfDay().Unix()
	return s, e
}

// FormatDuration2Time 格式化时间
func FormatDuration2Time(startTime, endTime string) (time.Time, time.Time) {
	if startTime == "" || endTime == "" {
		return time.Now(), time.Now()
	}
	s := now.New(DateStringToTime(startTime)).BeginningOfDay()
	e := now.New(DateStringToTime(endTime)).EndOfDay()
	return s, e
}
