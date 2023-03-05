package utils

import (
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

func TestDateStringToTime(t *testing.T) {
	location, err := time.LoadLocation("Asia/Shanghai")
	assert.NoError(t, err)

	tt, err := time.ParseInLocation(DateLayout, "2022-01-01", location)
	assert.NoError(t, err)
	unix := tt.Unix()
	format := time.Unix(unix, 0).Format(TimeLayout)
	assert.Equal(t, "2022-01-01 00:00:00", format)
	// "2022-01-01" time.Parse 之后的时间戳是  "2022-01-01 08:00:00"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DiffDateOfDay(tt.d1, tt.d2), "DiffDateOfDay(%v, %v)", tt.d1, tt.d2)
		})
	}
}
