package time

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeToString(t *testing.T) {
	type args struct {
		t *time.Time
	}
	now := time.Now()
	year := time.Date(1970, 1, 1, 8, 0, 0, 0, time.Local)
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "测试正常时间返回", args: args{t: &now}, want: now.Format(TimeLayout)},
		{name: "测试正常时间返回", args: args{t: nil}, want: now.Format(TimeLayout)},
		{name: "测试正常时间返回", args: args{t: &year}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeToString(tt.args.t); got != tt.want {
				t.Log(got, tt.args.t)
				t.Errorf("TimeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseTimeWithLayout(t *testing.T) {
	tests := []struct {
		name      string
		timeStr   string
		layout    string
		wantValid bool
	}{
		{name: "有效时间字符串", timeStr: "2022-01-01", layout: DateLayout, wantValid: true},
		{name: "无效时间字符串", timeStr: "2022-13-01", layout: DateLayout, wantValid: false},
		{name: "空时间字符串", timeStr: "", layout: DateLayout, wantValid: false},
		{name: "格式不匹配", timeStr: "2022/01/01", layout: DateLayout, wantValid: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, valid := parseTimeWithLayout(tt.timeStr, tt.layout)
			assert.Equal(t, tt.wantValid, valid, "parseTimeWithLayout valid result")
			if tt.wantValid {
				assert.False(t, result.IsZero(), "parsed time should not be zero")
			}
		})
	}
}

func TestStringToTime(t *testing.T) {
	tests := []struct {
		name      string
		timeStr   string
		layout    string
		wantValid bool
	}{
		{name: "有效日期字符串", timeStr: "2022-01-01", layout: DateLayout, wantValid: true},
		{name: "有效时间字符串", timeStr: "2022-01-01 12:30:45", layout: TimeLayout, wantValid: true},
		{name: "无效字符串", timeStr: "invalid", layout: DateLayout, wantValid: false},
		{name: "空字符串", timeStr: "", layout: DateLayout, wantValid: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, valid := StringToTime(tt.timeStr, tt.layout)
			assert.Equal(t, tt.wantValid, valid, "StringToTime valid result")
		})
	}
}

func TestDateStringToTime(t *testing.T) {
	location, err := time.LoadLocation("Asia/Shanghai")
	assert.NoError(t, err)

	tt, err := time.ParseInLocation(DateLayout, "2022-01-01", location)
	assert.NoError(t, err)
	unix := tt.Unix()
	format := time.Unix(unix, 0).Format(TimeLayout)
	assert.Equal(t, "2022-01-01 00:00:00", format)
	// "2022-01-01" time.Parse 之后的时间戳是  "2022-01-01 08:00:00"

	// 测试新的实现
	result := DateStringToTime("2022-01-01")
	assert.Equal(t, tt.Year(), result.Year())
	assert.Equal(t, tt.Month(), result.Month())
	assert.Equal(t, tt.Day(), result.Day())

	// 测试无效日期
	invalidResult := DateStringToTime("invalid-date")
	assert.NotEqual(t, 1970, invalidResult.Year(), "应该返回当前时间而非零时间")
}

func TestDateStringToYmdInt(t *testing.T) {
	tests := []struct {
		name string
		dt   string
		want int64
	}{
		{name: "1", dt: "2022-11-01", want: 20221101},
		{name: "2", dt: "2022-01-01", want: 20220101},
		{name: "3", dt: "2022-12-1", want: 0},
		{name: "4", dt: "2022-2-1", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymd := DateStringToYmdInt(tt.dt)
			t.Log(tt.name, ymd, tt.want)
			assert.Equalf(t, tt.want, ymd, "DateStringToYmdInt(%v)", tt.dt)
		})
	}
}

func TestYmdIntToDateString(t *testing.T) {
	tests := []struct {
		name string
		dt   int64
		want string
	}{
		{"1", 20221101, "2022-11-01"},
		{"2", 2022111, "-"},
		{"3", 20221134, "-"},
		{"4", 20221130, "2022-11-30"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, YmdIntToDateString(tt.dt), "YmdIntToDateString(%v)", tt.dt)
		})
	}
}

func TestHour2TimeString(t *testing.T) {
	tests := []struct {
		name string
		hour int
		want string
	}{
		{"0", 0, "00:00"},
		{"1", 1, "01:00"},
		{"12", 12, "12:00"},
		{"22", 22, "22:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hour := Hour2DateTime("", tt.hour)
			t.Log(hour)
			assert.Equalf(t, tt.want, hour, "Hour2TimeString(%v)", tt.hour)
		})
	}
}

func TestPoint2DateTime(t *testing.T) {
	type args struct {
		date  string
		point int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"2022-11-01", 0}, "2022-11-01 00:00"},
		{"2", args{"2022-11-01", 1}, "2022-11-01 00:05"},
		{"3", args{"2022-11-01", 20}, "2022-11-01 01:40"},
		{"4", args{"2022-11-01", 61}, "2022-11-01 05:05"},
		{"5", args{"2022-11-01", 287}, "2022-11-01 23:55"},
		{"11", args{"", 0}, "00:00"},
		{"12", args{"", 1}, "00:05"},
		{"13", args{"", 20}, "01:40"},
		{"14", args{"", 61}, "05:05"},
		{"15", args{"", 287}, "23:55"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Point2DateTime(tt.args.date, tt.args.point), "Point2DateTime(%v, %v)", tt.args.date, tt.args.point)
		})
	}
}

func TestStringToUnixTime(t *testing.T) {
	tests := []struct {
		name string
		ts   string
		want int64
	}{
		{"有效时间", "2022-01-01 12:00:00", 1641013200},
		{"无效时间", "invalid", 0},
		{"空字符串", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringToUnixTime(tt.ts)
			if tt.want > 0 {
				assert.Greater(t, result, int64(0), "有效时间应该返回正时间戳")
			} else {
				assert.Equal(t, tt.want, result, "无效时间应该返回0")
			}
		})
	}
}

func TestDateStringToUnixTime(t *testing.T) {
	tests := []struct {
		name string
		ts   string
		want int64
	}{
		{"有效日期", "2022-01-01", 1640966400},
		{"无效日期", "invalid", 0},
		{"空字符串", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DateStringToUnixTime(tt.ts)
			if tt.want > 0 {
				assert.Greater(t, result, int64(0), "有效日期应该返回正时间戳")
			} else {
				assert.Equal(t, tt.want, result, "无效日期应该返回0")
			}
		})
	}
}

func TestDiffDateOfDay(t *testing.T) {
	type args struct {
		d1 string
		d2 string
	}
	tests := []struct {
		name string
		d1   string
		d2   string
		want float64
	}{
		{"1", "2022-11-01", "2022-11-03", float64(3)},
		{"无效日期", "invalid", "2022-11-03", float64(0)},
		{"无效日期2", "2022-11-01", "invalid", float64(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DiffDateOfDay(tt.d1, tt.d2), "DiffDateOfDay(%v, %v)", tt.d1, tt.d2)
		})
	}
}

func TestFormatDuration2Time(t *testing.T) {
	type args struct {
		startTime string
		endTime   string
	}
	tests := []struct {
		name  string
		args  args
		want  time.Time
		want1 time.Time
	}{
		{"1", args{"2022-11-01", "2022-11-15"}, time.Date(2022, 11, 1, 0, 0, 0, 0, time.Local), time.Date(2022, 11, 2, 0, 0, 0, 0, time.Local)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FormatDuration2Time(tt.args.startTime, tt.args.endTime)
			fmt.Println(got)
			fmt.Println(got1)
		})
	}
}

func TestIsNowTimeIn(t *testing.T) {
	type args struct {
		startTime int64
		endTime   int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "测试正常时间返回", args: args{startTime: time.Now().Unix() - 300, endTime: time.Now().Unix() + 300}, want: true},
		{name: "测试异常时间返回", args: args{startTime: time.Now().Unix() + 300, endTime: time.Now().Unix() + 600}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsNowTimeIn(tt.args.startTime, tt.args.endTime), "IsNowTimeIn(%v, %v)", tt.args.startTime, tt.args.endTime)
		})
	}
}

func TestGetZeroTime(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		{"1", time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetZeroTime(), "GetZeroTime()")
			println(GetZeroTime().Unix())
		})
	}
}
