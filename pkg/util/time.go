package util

import (
	"fmt"
	"strconv"
	"time"
	"woods/pkg/setting"
)

// TimeNowUnix int32
func TimeNowUnix() int32 {
	return int32(time.Now().Unix())
}

// Time2Unix int32
func Time2Unix(t time.Time) int32 {
	return int32(t.Unix())
}

// Time2DayNumber 给定一个时间，把当日时间格式的字符串转换为int类型
func Time2DayNumber(t time.Time) int32 {
	tInt, _ := strconv.Atoi(t.Format("20060102"))
	return int32(tInt)
}

// Unix2Time timestamp to time
func Unix2Time(u int64) time.Time {
	return time.Unix(u, 0)
}

// FormatWeekday 格式化星期
func FormatWeekday(t time.Time) string {
	switch t.Weekday() {
	case time.Sunday:
		return "星期日"
	case time.Monday:
		return "星期一"
	case time.Tuesday:
		return "星期二"
	case time.Wednesday:
		return "星期三"
	case time.Thursday:
		return "星期四"
	case time.Friday:
		return "星期五"
	case time.Saturday:
		return "星期六"
	default:
		return "星期一"
	}
}

// GetZeroTime 给定一个时间，获取当天的零点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func Now() time.Time {
	return time.Now().UTC().In(GetLocation())
}

func NowMilli() int64 {
	return Now().UnixMilli()
}

// 指定时间转为日期
func Time2Date(t time.Time) string {
	return t.Format("2006-01-02")
}

// 指定日期转换为时间
func Date2Time(date string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", date, GetLocation())
}

func GetDateTimestamp(time time.Time) int64 {
	dateTime := Time2Date(time)
	fmt.Println("转换为时间：", dateTime)
	date, _ := Date2Time(dateTime)
	return date.In(GetLocation()).Unix()
}

// 获取指定天数的日期
func GetDayTime(day int) time.Time {
	currentTime := Now()
	return currentTime.AddDate(0, 0, day)
}

// 获取指定月数的日期
func GetMonthTime(month int) time.Time {
	currentTime := Now()
	return currentTime.AddDate(0, month, 0)
}

// 获取指定月数的日期
func GetYearTime(year int) time.Time {
	currentTime := Now()
	return currentTime.AddDate(0, year, 0)
}

func GetLocation() *time.Location {
	sh, _ := time.LoadLocation(setting.AppSetting.TimeZone)
	return sh
}
