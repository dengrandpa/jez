// Package timejez 时间相关函数
package timejez

import (
	"time"
)

// 返回时间格式化的格式，如果没有传入，则默认为 YYYYMMDDHHMMSS
func getTimeFormat(format []TimeFormat) TimeFormat {
	t := YYYYMMDDHHMMSS
	if len(format) > 0 {
		t = format[0]
	}
	return t
}

// ParseTime 将字符串转换为时间，默认格式为 YYYYMMDDHHMMSS
func ParseTime(tm string, format ...TimeFormat) time.Time {
	t, _ := time.ParseInLocation(getTimeFormat(format).string(), tm, time.Local)
	return t
}

// ParseTimestamp 将时间戳转换为时间
func ParseTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0).Local()
}

// StartOfMinute 返回时间 t 所在分钟的开始时间 yyyy-mm-dd hh:mm:00
func StartOfMinute(t time.Time) time.Time {
	year, month, day := t.Date()
	return NewTime(t.Location(), year, int(month), day, t.Hour(), t.Minute())
}

// EndOfMinute 返回时间 t 所在分钟的结束时间 yyyy-mm-dd hh:mm:59
func EndOfMinute(t time.Time) time.Time {
	year, month, day := t.Date()
	return NewTime(t.Location(), year, int(month), day, t.Hour(), t.Minute()+1).Add(-time.Nanosecond)
}

// StartOfHour 返回时间 t 所在小时的开始时间 yyyy-mm-dd hh:00:00
func StartOfHour(t time.Time) time.Time {
	year, month, day := t.Date()
	return NewTime(t.Location(), year, int(month), day, t.Hour())
}

// EndOfHour 返回时间 t 所在小时的结束时间 yyyy-mm-dd hh:59:59
func EndOfHour(t time.Time) time.Time {
	year, month, day := t.Date()
	return NewTime(t.Location(), year, int(month), day, t.Hour()+1).Add(-time.Nanosecond)
}

// StartOfDay 返回时间 t 所在天的开始时间 yyyy-mm-dd 00:00:00
func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return NewTime(t.Location(), year, int(month), day)
}

// EndOfDay 返回时间 t 所在天的结束时间 yyyy-mm-dd 23:59:59
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return NewTime(t.Location(), year, int(month), day+1).Add(-time.Nanosecond)
}

// StartOfWeekMonday 返回时间 t 所在周的开始时间，周一为第一天 yyyy-mm-dd 00:00:00
func StartOfWeekMonday(t time.Time) time.Time {
	year, month, day := t.Date()

	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	return NewTime(t.Location(), year, int(month), day-weekday+1)
}

// EndOfWeekSunday 返回时间 t 所在周的结束时间，周日为最后一天 yyyy-mm-dd 23:59:59
func EndOfWeekSunday(t time.Time) time.Time {
	year, month, day := t.Date()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return NewTime(t.Location(), year, int(month), day+(7-weekday)+1).Add(-time.Nanosecond)
}

// StartOfWeekSunday 返回时间 t 所在周的开始时间，周日为第一天 yyyy-mm-dd 00:00:00
func StartOfWeekSunday(t time.Time) time.Time {
	year, month, day := t.Date()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return NewTime(t.Location(), year, int(month), day-weekday)
}

// EndOfWeekMonday 返回时间 t 所在周的结束时间，周一为最后一天 yyyy-mm-dd 23:59:59
func EndOfWeekMonday(t time.Time) time.Time {
	year, month, day := t.Date()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return NewTime(t.Location(), year, int(month), day+(8-weekday)+1).Add(-time.Nanosecond)
}

// StartOfMonth 返回时间 t 所在月的开始时间 yyyy-mm-01 00:00:00
func StartOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return NewTime(t.Location(), year, int(month), 1)
}

// EndOfMonth 返回时间 t 所在月的结束时间 yyyy-mm-dd 23:59:59
func EndOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return NewTime(t.Location(), year, int(month)+1, 1).Add(-time.Nanosecond)
}

// StartOfYear 返回时间 t 所在年的开始时间 yyyy-01-01 00:00:00
func StartOfYear(t time.Time) time.Time {
	return NewTime(t.Location(), t.Year(), 1, 1)
}

// EndOfYear 返回时间 t 所在年的结束时间 yyyy-12-31 23:59:59
func EndOfYear(t time.Time) time.Time {
	return NewTime(t.Location(), t.Year()+1, 1, 1).Add(-time.Nanosecond)
}

// AddSecond 添加或删除秒数
func AddSecond(t time.Time, second int) time.Time {
	return t.Add(time.Second * time.Duration(second))
}

// AddMinute 添加或删除分钟数
func AddMinute(t time.Time, minute int) time.Time {
	return t.Add(time.Minute * time.Duration(minute))
}

// AddHour 添加或删除小时数
func AddHour(t time.Time, hour int) time.Time {
	return t.Add(time.Hour * time.Duration(hour))
}

// AddDay 添加或删除天数
func AddDay(t time.Time, day int) time.Time {
	return t.AddDate(0, 0, day)
}

// AddWeek 添加或删除周数
func AddWeek(t time.Time, week int) time.Time {
	return t.AddDate(0, 0, week*7)
}

// AddMonth 添加或删除月数
func AddMonth(t time.Time, month int) time.Time {
	return t.AddDate(0, month, 0)
}

// AddYear 添加或删除年数
func AddYear(t time.Time, year int) time.Time {
	return t.AddDate(year, 0, 0)
}

// FormatTime 将时间格式化为字符串，默认格式为 yyyy-mm-dd hh:mm:ss
func FormatTime(t time.Time, format ...TimeFormat) string {
	return t.Format(getTimeFormat(format).string())
}

// FormatTimestamp 将时间戳格式化为字符串，默认格式为 yyyy-mm-dd hh:mm:ss
func FormatTimestamp(timestamp int64, format ...TimeFormat) string {
	return FormatTime(ParseTimestamp(timestamp), format...)
}

// FormatNow 返回当前时间的字符串格式，默认格式为 yyyy-mm-dd hh:mm:ss
func FormatNow(format ...TimeFormat) string {
	return FormatTime(time.Now(), format...)
}

// IsLeapYear 判断年份 year 是否为闰年
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// RangeHours 返回两个时间之间的所有小时的切片，包含 start 和 end，即[start,end]，如果start和end结果一样，则只返回1个
func RangeHours(start, end time.Time) []time.Time {
	hours := int(end.Sub(start).Hours())

	dates := make([]time.Time, 0, hours+1)

	dates = append(dates, start)

	for i := 1; i <= hours; i++ {
		dates = append(dates, start.Add(time.Hour*time.Duration(i)))
	}
	return dates
}

// RangeDays 返回两个时间之间的所有天的切片，包含 start 和 end，即[start,end]，如果start和end结果一样，则只返回1个
func RangeDays(start, end time.Time) []time.Time {
	days := int(end.Sub(start).Hours() / 24)

	dates := make([]time.Time, 0, days+1)

	dates = append(dates, start)

	for i := 1; i <= days; i++ {
		dates = append(dates, start.AddDate(0, 0, i))
	}

	return dates
}

// RangeMonths 返回两个时间之间的所有月的切片，包含 start 和 end，即[start,end]，如果start和end结果一样，则只返回1个
func RangeMonths(start, end time.Time) []time.Time {
	months := (end.Year()-start.Year())*12 + int(end.Month()) - int(start.Month())

	dates := make([]time.Time, 0, months+1)

	dates = append(dates, start)

	for i := 1; i <= months; i++ {
		dates = append(dates, start.AddDate(0, i, 0))
	}

	return dates
}

// RangeYears 返回两个时间之间的所有年的切片，包含 start 和 end，即[start,end]，如果start和end结果一样，则只返回1个
func RangeYears(start, end time.Time) []time.Time {
	years := end.Year() - start.Year()

	dates := make([]time.Time, 0, years+1)

	dates = append(dates, start)

	for i := 1; i <= years; i++ {
		dates = append(dates, start.AddDate(i, 0, 0))
	}

	return dates
}

// NewTime 相当于 time.Date，如果不传参数则相当于 time.Now。
func NewTime(loc *time.Location, t ...int) time.Time {
	if loc == nil {
		loc = time.Local
	}
	size := len(t)
	if size == 0 {
		return time.Now()
	}

	param := make([]int, 7)
	for i := 0; i < size; i++ {
		param[i] = t[i]
	}

	return time.Date(param[0], time.Month(param[1]), param[2], param[3], param[4], param[5], param[6], loc)
}

// SubTime 返回 t1 和 t2 之间的差值的日、小时、分钟和秒
func SubTime(t1, t2 time.Time) (days, hours, minutes, seconds int) {
	return SubTimestamp(t1.Unix(), t2.Unix())
}

// SubTimestamp 返回 t1 和 t2 之间的差值的日、小时、分钟和秒
func SubTimestamp(t1, t2 int64) (days, hours, minutes, seconds int) {
	sub := int(t1 - t2)

	if sub < 0 {
		sub = -sub
	}

	days = sub / secondsInDay
	remainingSeconds := sub % secondsInDay

	hours = remainingSeconds / secondsInHour
	remainingSeconds = remainingSeconds % secondsInHour

	minutes = remainingSeconds / secondsInMinute

	seconds = remainingSeconds % secondsInMinute

	return days, hours, minutes, seconds
}

// CST 返回中国时区
func CST() *time.Location {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return loc
}

// ToCST 将时间转换为中国时区
func ToCST(t time.Time) time.Time {
	return t.In(CST())
}
