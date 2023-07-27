// Package timeutil provide some useful time functions.
package timeutil

import (
	"time"
)

// get time format
func getTimeFormat(format []TimeFormat) TimeFormat {
	t := YYYYMMDDHHMMSS
	if len(format) > 0 {
		t = format[0]
	}
	return t
}

// ParseTime convert string to time.
func ParseTime(tm string, format ...TimeFormat) time.Time {
	t, _ := time.ParseInLocation(getTimeFormat(format).string(), tm, time.Local)
	return t
}

// ParseTimestamp parse unix timestamp to time.
func ParseTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// StartOfMinute return the start time of current minute.
func StartOfMinute(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMinute return the end time of current minute.
func EndOfMinute(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), t.Minute(), 59, maxNesc, t.Location())
}

// StartOfHour return the start time of current hour.
func StartOfHour(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour return the end time of current hour.
func EndOfHour(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), 59, 59, maxNesc, t.Location())
}

// StartOfDay return the start time of today.
func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// EndOfDay return the end time of today.
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, maxNesc, t.Location())
}

// StartOfWeekMonday return the start time of the current week, starting on Monday.
func StartOfWeekMonday(t time.Time) time.Time {
	year, month, day := t.Date()

	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	return time.Date(year, month, day-weekday+1, 0, 0, 0, 0, t.Location())
}

// EndOfWeekSunday return the end time of current week, ending on Sunday.
func EndOfWeekSunday(t time.Time) time.Time {
	year, month, day := t.Date()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return time.Date(year, month, day+(7-weekday), 23, 59, 59, maxNesc, t.Location())
}

// StartOfWeekSunday return the start time of the current week, starting on Sunday.
func StartOfWeekSunday(t time.Time) time.Time {
	year, month, day := t.Date()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return time.Date(year, month, day-weekday, 0, 0, 0, 0, t.Location())
}

// EndOfWeekMonday return the end time of current week, ending on Monday.
func EndOfWeekMonday(t time.Time) time.Time {
	year, month, day := t.Date()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return time.Date(year, month, day+(8-weekday), 23, 59, 59, maxNesc, t.Location())
}

// StartOfMonth return the start time of current month.
func StartOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth return the end time of current month.
func EndOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month+1, 1, 0, 0, 0, 0, t.Location()).Add(-time.Nanosecond)
}

// StartOfYear return the start time of current year.
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear return the end time of current year.
func EndOfYear(t time.Time) time.Time {
	return time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location()).Add(-time.Nanosecond)
}

// AddSecond add or sub second to the time.
func AddSecond(t time.Time, second int) time.Time {
	return t.Add(time.Second * time.Duration(second))
}

// AddMinute add or sub minute to the time.
func AddMinute(t time.Time, minute int) time.Time {
	return t.Add(time.Minute * time.Duration(minute))
}

// AddHour add or sub hour to the time.
func AddHour(t time.Time, hour int) time.Time {
	return t.Add(time.Hour * time.Duration(hour))
}

// AddDay add or sub day to the time.
func AddDay(t time.Time, day int) time.Time {
	return t.AddDate(0, 0, day)
}

// AddWeek add or sub week to the time.
func AddWeek(t time.Time, week int) time.Time {
	return t.AddDate(0, 0, week*7)
}

// AddMonth add or sub month to the time.
func AddMonth(t time.Time, month int) time.Time {
	return t.AddDate(0, month, 0)
}

// AddYear add or sub year to the time.
func AddYear(t time.Time, year int) time.Time {
	return t.AddDate(year, 0, 0)
}

// FormatTime convert time to string.
func FormatTime(t time.Time, format ...TimeFormat) string {
	return t.Format(getTimeFormat(format).string())
}

// FormatTimestamp convert unix timestamp to string.
func FormatTimestamp(timestamp int64, format ...TimeFormat) string {
	return FormatTime(ParseTimestamp(timestamp), format...)
}

// FormatNow return current time, default format: yyyy-mm-dd hh:mm:ss.
func FormatNow(format ...TimeFormat) string {
	return FormatTime(time.Now(), format...)
}

// FormatStartOfMinute return the start time of current minute, default format: yyyy-mm-dd hh:mm:00.
func FormatStartOfMinute(t time.Time, format ...TimeFormat) string {
	return FormatTime(StartOfMinute(t), format...)
}

// FormatEndOfMinute return the end time of current minute, default format: yyyy-mm-dd hh:mm:59.
func FormatEndOfMinute(t time.Time, format ...TimeFormat) string {
	return FormatTime(EndOfMinute(t), format...)
}

// FormatStartOfHour return the start time of current hour, default format: yyyy-mm-dd hh:00:00.
func FormatStartOfHour(t time.Time, format ...TimeFormat) string {
	return FormatTime(StartOfHour(t), format...)
}

// FormatEndOfHour return the end time of current hour, default format: yyyy-mm-dd hh:59:59.
func FormatEndOfHour(t time.Time, format ...TimeFormat) string {
	return FormatTime(EndOfHour(t), format...)
}

// FormatStartOfDay return the start time of today, default format: yyyy-mm-dd 00:00:00.
func FormatStartOfDay(t time.Time, format ...TimeFormat) string {
	return FormatTime(StartOfDay(t), format...)
}

// FormatEndOfDay return the end time of today, default format: yyyy-mm-dd 23:59:59.
func FormatEndOfDay(t time.Time, format ...TimeFormat) string {
	return FormatTime(EndOfDay(t), format...)
}

// FormatStartOfWeekMonday return the start time of the current week, starting on Monday, default format: yyyy-mm-dd 00:00:00.
func FormatStartOfWeekMonday(t time.Time, format ...TimeFormat) string {
	return FormatTime(StartOfWeekMonday(t), format...)
}

// FormatEndOfWeekSunday return the end time of current week, ending on Sunday, default format: yyyy-mm-dd 23:59:59.
func FormatEndOfWeekSunday(t time.Time, format ...TimeFormat) string {
	return FormatTime(EndOfWeekSunday(t), format...)
}

// FormatStartOfWeekSunday return the start time of the current week, starting on Sunday, default format: yyyy-mm-dd 00:00:00.
func FormatStartOfWeekSunday(t time.Time, format ...TimeFormat) string {
	return FormatTime(StartOfWeekSunday(t), format...)
}

// FormatEndOfWeekMonday return the end time of current week, ending on Monday, default format: yyyy-mm-dd 23:59:59.
func FormatEndOfWeekMonday(t time.Time, format ...TimeFormat) string {
	return FormatTime(EndOfWeekMonday(t), format...)
}

// FormatStartOfMonth return the start time of current month, default format: yyyy-mm-01 00:00:00.
func FormatStartOfMonth(t time.Time, format ...TimeFormat) string {
	return FormatTime(StartOfMonth(t), format...)
}

// FormatEndOfMonth return the end time of current month, default format: yyyy-mm-31 23:59:59.
func FormatEndOfMonth(t time.Time, format ...TimeFormat) string {
	return FormatTime(EndOfMonth(t), format...)
}

// FormatStartOfYear return the start time of current year, default format: yyyy-01-01 00:00:00.
func FormatStartOfYear(t time.Time, format ...TimeFormat) string {
	return FormatTime(StartOfYear(t), format...)
}

// FormatEndOfYear return the end time of current year, default format: yyyy-12-31 23:59:59.
func FormatEndOfYear(t time.Time, format ...TimeFormat) string {
	return FormatTime(EndOfYear(t), format...)
}

// StartOfMinuteTimestamp return the start timestamp of current minute.
func StartOfMinuteTimestamp(t time.Time) int64 {
	return StartOfMinute(t).Unix()
}

// EndOfMinuteTimestamp return the end timestamp of current minute.
func EndOfMinuteTimestamp(t time.Time) int64 {
	return EndOfMinute(t).Unix()
}

// StartOfHourTimestamp return the start timestamp of current hour.
func StartOfHourTimestamp(t time.Time) int64 {
	return StartOfHour(t).Unix()
}

// EndOfHourTimestamp return the end timestamp of current hour.
func EndOfHourTimestamp(t time.Time) int64 {
	return EndOfHour(t).Unix()
}

// StartOfDayTimestamp return the start timestamp of today.
func StartOfDayTimestamp(t time.Time) int64 {
	return StartOfDay(t).Unix()
}

// EndOfDayTimestamp return the end timestamp of today.
func EndOfDayTimestamp(t time.Time) int64 {
	return EndOfDay(t).Unix()
}

// StartOfWeekMondayTimestamp return the start timestamp of the current week, starting on Monday.
func StartOfWeekMondayTimestamp(t time.Time) int64 {
	return StartOfWeekMonday(t).Unix()
}

// EndOfWeekSundayTimestamp return the end timestamp of current week, ending on Sunday.
func EndOfWeekSundayTimestamp(t time.Time) int64 {
	return EndOfWeekSunday(t).Unix()
}

// StartOfWeekSundayTimestamp return the start timestamp of the current week, starting on Sunday.
func StartOfWeekSundayTimestamp(t time.Time) int64 {
	return StartOfWeekSunday(t).Unix()
}

// EndOfWeekMondayTimestamp return the end timestamp of current week, ending on Monday.
func EndOfWeekMondayTimestamp(t time.Time) int64 {
	return EndOfWeekMonday(t).Unix()
}

// StartOfMonthTimestamp return the start timestamp of current month.
func StartOfMonthTimestamp(t time.Time) int64 {
	return StartOfMonth(t).Unix()
}

// EndOfMonthTimestamp return the end timestamp of current month.
func EndOfMonthTimestamp(t time.Time) int64 {
	return EndOfMonth(t).Unix()
}

// StartOfYearTimestamp return the start timestamp of current year.
func StartOfYearTimestamp(t time.Time) int64 {
	return StartOfYear(t).Unix()
}

// EndOfYearTimestamp return the end timestamp of current year.
func EndOfYearTimestamp(t time.Time) int64 {
	return EndOfYear(t).Unix()
}

// IsLeapYear check if param year is leap year or not.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// RangeHours return a list of hours between start and end, containing start and end hours.
func RangeHours(start, end time.Time) []time.Time {
	// Calculate the number of hours between start and end, including both hours.
	hours := int(end.Sub(start).Hours())

	dates := make([]time.Time, 0, hours+1)

	dates = append(dates, start)

	for i := 1; i <= hours; i++ {
		dates = append(dates, start.Add(time.Hour*time.Duration(i)))
	}
	return dates
}

// RangeDays return a list of dates between start and end, containing start and end days.
func RangeDays(start, end time.Time) []time.Time {
	// Calculate the number of days between start and end, including both days.
	days := int(end.Sub(start).Hours() / 24)

	dates := make([]time.Time, 0, days+1)

	dates = append(dates, start)

	for i := 1; i <= days; i++ {
		dates = append(dates, start.AddDate(0, 0, i))
	}

	return dates
}

// RangeMonths return a list of months between start and end, containing start and end months.
func RangeMonths(start, end time.Time) []time.Time {
	// Calculate the number of months between start and end, including both months.
	months := (end.Year()-start.Year())*12 + int(end.Month()) - int(start.Month())

	dates := make([]time.Time, 0, months+1)

	dates = append(dates, start)

	for i := 1; i <= months; i++ {
		dates = append(dates, start.AddDate(0, i, 0))
	}

	return dates
}

// RangeYears return a list of years between start and end, containing start and end years.
func RangeYears(start, end time.Time) []time.Time {
	// Calculate the number of years between start and end, including both years.
	years := end.Year() - start.Year()

	dates := make([]time.Time, 0, years+1)

	dates = append(dates, start)

	for i := 1; i <= years; i++ {
		dates = append(dates, start.AddDate(i, 0, 0))
	}

	return dates
}

// NewTime It is equivalent to time.Date, if no parameter is passed, it is equivalent to time.Now.
func NewTime(t ...int) time.Time {
	size := len(t)
	if size == 0 {
		return time.Now()
	}

	param := make([]int, 7)
	for i := 0; i < size; i++ {
		param[i] = t[i]
	}

	return time.Date(param[0], time.Month(param[1]), param[2], param[3], param[4], param[5], param[6], time.Local)
}

// SubTime Returns the day, hour, minute, and second of the difference between t1 and t2
func SubTime(t1, t2 time.Time) (days, hours, minutes, seconds int) {
	return SubTimestamp(t1.Unix(), t2.Unix())
}

// SubTimestamp Returns the day, hour, minute, and second of the difference between t1 and t2
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
