package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ass.Equal(
		time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local),
		ParseTime("2022-01-02 03:04:05", YYYYMMDDHHMMSS))

	ass.Equal(
		time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local),
		ParseTime("2022/01/02 03:04:05", YYYYMMDDHHMMSS2))

	ass.Equal(
		time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
		ParseTime("2022-01", YYYYMM))

	ass.Equal(
		time.Date(2022, 1, 2, 0, 0, 0, 0, time.Local),
		ParseTime("2022-01-02", YYYYMMDD))

	ass.Equal(
		time.Date(0, 1, 1, 3, 4, 5, 0, time.Local),
		ParseTime("03:04:05", HHMMSS))

	ass.Equal(
		time.Date(0, 1, 1, 3, 4, 0, 0, time.Local),
		ParseTime("03:04", HHMM))
}

func TestParseTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	ass.Equal(tm, ParseTimestamp(1641063845))

}

func TestStartOfMinute(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := StartOfMinute(tm)

	ass.Equal(0, res.Second())

}

func TestEndOfMinute(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := EndOfMinute(tm)

	ass.Equal(59, res.Second())
}

func TestStartOfHour(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := StartOfHour(tm)

	ass.Equal(0, res.Second())
	ass.Equal(0, res.Minute())
}

func TestEndOfHour(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := EndOfHour(tm)

	ass.Equal(59, res.Second())
	ass.Equal(59, res.Minute())
}

func TestStartOfDay(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := StartOfDay(tm)

	ass.Equal(0, res.Second())
	ass.Equal(0, res.Minute())
	ass.Equal(0, res.Hour())
}

func TestEndOfDay(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := EndOfDay(tm)

	ass.Equal(59, res.Second())
	ass.Equal(59, res.Minute())
	ass.Equal(23, res.Hour())
}

func TestStartOfWeekMonday(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := StartOfWeekMonday(tm)

	ass.Equal(0, res.Second())
	ass.Equal(0, res.Minute())
	ass.Equal(0, res.Hour())
	ass.Equal(time.Monday, res.Weekday())
}

func TestEndOfWeekSunday(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := EndOfWeekSunday(tm)

	ass.Equal(59, res.Second())
	ass.Equal(59, res.Minute())
	ass.Equal(23, res.Hour())
	ass.Equal(time.Sunday, res.Weekday())
}

func TestStartOfWeekSunday(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := StartOfWeekSunday(tm)

	ass.Equal(0, res.Second())
	ass.Equal(0, res.Minute())
	ass.Equal(0, res.Hour())
	ass.Equal(time.Sunday, res.Weekday())
}

func TestEndOfWeekMonday(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := EndOfWeekMonday(tm)

	ass.Equal(59, res.Second())
	ass.Equal(59, res.Minute())
	ass.Equal(23, res.Hour())
	ass.Equal(time.Monday, res.Weekday())
}

func TestStartOfMonth(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := StartOfMonth(tm)

	ass.Equal(0, res.Second())
	ass.Equal(0, res.Minute())
	ass.Equal(0, res.Hour())
	ass.Equal(1, res.Day())
}

func TestEndOfMonth(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := EndOfMonth(tm)

	ass.Equal(59, res.Second())
	ass.Equal(59, res.Minute())
	ass.Equal(23, res.Hour())
	ass.Equal(31, res.Day())
}

func TestStartOfYear(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := StartOfYear(tm)

	ass.Equal(0, res.Second())
	ass.Equal(0, res.Minute())
	ass.Equal(0, res.Hour())
	ass.Equal(1, res.Day())
	ass.Equal(time.January, res.Month())
}

func TestEndOfYear(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := EndOfYear(tm)

	ass.Equal(59, res.Second())
	ass.Equal(59, res.Minute())
	ass.Equal(23, res.Hour())
	ass.Equal(31, res.Day())
	ass.Equal(time.December, res.Month())
}

func TestAddSecond(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	ass.Equal(15, AddSecond(tm, 10).Second())
}

func TestAddMinute(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	ass.Equal(14, AddMinute(tm, 10).Minute())
}

func TestAddHour(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	ass.Equal(13, AddHour(tm, 10).Hour())
}

func TestAddDay(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	ass.Equal(12, AddDay(tm, 10).Day())
}

func TestAddWeek(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	ass.Equal(tm.Weekday(), AddWeek(tm, 7).Weekday())
}

func TestAddMonth(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	ass.Equal(time.November, AddMonth(tm, 10).Month())
}

func TestAddYear(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	ass.Equal(2032, AddYear(tm, 10).Year())
}

func TestFormatTime(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	ass.Equal("2022-01-02 03:04:05", FormatTime(tm))
	ass.Equal("2022/01/02 03:04:05", FormatTime(tm, YYYYMMDDHHMMSS2))
}

func TestFormatTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ass.Equal("2022-01-02 03:04:05", FormatTimestamp(1641063845))
	ass.Equal("2022/01/02 03:04:05", FormatTimestamp(1641063845, YYYYMMDDHHMMSS2))

}

func TestFormatNow(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := FormatTime(time.Now())
	tm2 := FormatTime(time.Now(), YYYYMMDDHHMMSS2)

	ass.Equal(tm, FormatNow())
	ass.Equal(tm2, FormatNow(YYYYMMDDHHMMSS2))
}

func TestFormatStartOfMinute(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfMinute(tm)

	f2 := FormatStartOfMinute(tm, YYYYMMDDHHMMSS2)

	ass.Equal("2022-01-02 03:04:00", f1)
	ass.Equal("2022/01/02 03:04:00", f2)
}

func TestFormatEndOfMinute(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfMinute(tm)

	f2 := FormatEndOfMinute(tm, YYYYMMDDHHMMSS2)

	ass.Equal("2022-01-02 03:04:59", f1)
	ass.Equal("2022/01/02 03:04:59", f2)
}

func TestFormatStartOfHour(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfHour(tm)

	f2 := FormatStartOfHour(tm, YYYYMMDDHHMMSS2)

	ass.Equal("2022-01-02 03:00:00", f1)
	ass.Equal("2022/01/02 03:00:00", f2)
}

func TestFormatEndOfHour(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfHour(tm)

	f2 := FormatEndOfHour(tm, YYYYMMDDHHMMSS2)

	ass.Equal("2022-01-02 03:59:59", f1)
	ass.Equal("2022/01/02 03:59:59", f2)
}

func TestFormatStartOfDay(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfDay(tm)

	f2 := FormatStartOfDay(tm, YYYYMMDDHHMMSS2)

	ass.Equal("2022-01-02 00:00:00", f1)
	ass.Equal("2022/01/02 00:00:00", f2)
}

func TestFormatEndOfDay(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfDay(tm)

	f2 := FormatEndOfDay(tm, YYYYMMDDHHMMSS2)

	ass.Equal("2022-01-02 23:59:59", f1)
	ass.Equal("2022/01/02 23:59:59", f2)
}

func TestFormatStartOfWeekMonday(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfWeekMonday(tm)
	f1Tm := ParseTime(f1)

	ass.Equal(time.Monday, f1Tm.Weekday())

	f2 := FormatStartOfWeekMonday(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	ass.Equal(time.Monday, f2Tm.Weekday())

}

func TestFormatEndOfWeekSunday(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfWeekSunday(tm)
	f1Tm := ParseTime(f1)

	ass.Equal(time.Sunday, f1Tm.Weekday())

	f2 := FormatEndOfWeekSunday(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	ass.Equal(time.Sunday, f2Tm.Weekday())
}

func TestFormatStartOfWeekSunday(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfWeekSunday(tm)
	f1Tm := ParseTime(f1)

	ass.Equal(time.Sunday, f1Tm.Weekday())

	f2 := FormatStartOfWeekSunday(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	ass.Equal(time.Sunday, f2Tm.Weekday())

}

func TestFormatEndOfWeekMonday(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfWeekMonday(tm)
	f1Tm := ParseTime(f1)

	ass.Equal(time.Monday, f1Tm.Weekday())

	f2 := FormatEndOfWeekMonday(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	ass.Equal(time.Monday, f2Tm.Weekday())
}

func TestFormatStartOfMonth(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfMonth(tm)
	f1Tm := ParseTime(f1)

	ass.Equal(1, f1Tm.Day())

	f2 := FormatStartOfMonth(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	ass.Equal(1, f2Tm.Day())

}

func TestFormatEndOfMonth(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfMonth(tm)
	f1Tm := ParseTime(f1)

	ass.Equal(31, f1Tm.Day())

	f2 := FormatEndOfMonth(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	ass.Equal(31, f2Tm.Day())
}

func TestFormatStartOfYear(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfYear(tm)
	f1Tm := ParseTime(f1)

	ass.Equal(time.January, f1Tm.Month())

	f2 := FormatStartOfYear(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	ass.Equal(time.January, f2Tm.Month())
}

func TestFormatEndOfYear(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfYear(tm)
	f1Tm := ParseTime(f1)

	ass.Equal(time.December, f1Tm.Month())

	f2 := FormatEndOfYear(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	ass.Equal(time.December, f2Tm.Month())
}

func TestStartOfMinuteTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := StartOfMinuteTimestamp(tm)

	ass.Equal(StartOfMinute(tm).Unix(), t1)
}

func TestEndOfMinuteTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := EndOfMinuteTimestamp(tm)

	ass.Equal(EndOfMinute(tm).Unix(), t1)
}

func TestStartOfHourTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := StartOfHourTimestamp(tm)

	ass.Equal(StartOfHour(tm).Unix(), t1)
}

func TestEndOfHourTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := EndOfHourTimestamp(tm)

	ass.Equal(EndOfHour(tm).Unix(), t1)
}

func TestStartOfDayTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := StartOfDayTimestamp(tm)

	ass.Equal(StartOfDay(tm).Unix(), t1)
}

func TestEndOfDayTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := EndOfDayTimestamp(tm)

	ass.Equal(EndOfDay(tm).Unix(), t1)
}

func TestStartOfWeekMondayTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := StartOfWeekMondayTimestamp(tm)

	ass.Equal(StartOfWeekMonday(tm).Unix(), t1)
}

func TestEndOfWeekSundayTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := EndOfWeekSundayTimestamp(tm)

	ass.Equal(EndOfWeekSunday(tm).Unix(), t1)
}

func TestStartOfWeekSundayTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := StartOfWeekSundayTimestamp(tm)

	ass.Equal(StartOfWeekSunday(tm).Unix(), t1)
}

func TestEndOfWeekMondayTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := EndOfWeekMondayTimestamp(tm)

	ass.Equal(EndOfWeekMonday(tm).Unix(), t1)
}

func TestStartOfMonthTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := StartOfMonthTimestamp(tm)

	ass.Equal(StartOfMonth(tm).Unix(), t1)
}

func TestEndOfMonthTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := EndOfMonthTimestamp(tm)

	ass.Equal(EndOfMonth(tm).Unix(), t1)
}

func TestStartOfYearTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := StartOfYearTimestamp(tm)

	ass.Equal(StartOfYear(tm).Unix(), t1)
}

func TestEndOfYearTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	t1 := EndOfYearTimestamp(tm)

	ass.Equal(EndOfYear(tm).Unix(), t1)
}

func TestIsLeapYear(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	y1 := IsLeapYear(2012)
	y2 := IsLeapYear(2013)

	ass.True(y1)
	ass.False(y2)

	// IsLeapYear()
}

func TestRangeHours(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	tm2 := time.Date(2022, 1, 2, 5, 55, 44, 0, time.Local)

	res := []time.Time{
		time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local),
		time.Date(2022, 1, 2, 4, 4, 5, 0, time.Local),
		time.Date(2022, 1, 2, 5, 4, 5, 0, time.Local),
	}

	ass.Equal(res, RangeHours(tm1, tm2))
}

func TestRangeDays(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	tm2 := time.Date(2022, 1, 4, 5, 55, 44, 0, time.Local)

	res := []time.Time{
		time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local),
		time.Date(2022, 1, 3, 3, 4, 5, 0, time.Local),
		time.Date(2022, 1, 4, 3, 4, 5, 0, time.Local),
	}

	ass.Equal(res, RangeDays(tm1, tm2))
}

func TestRangeMonths(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	tm2 := time.Date(2022, 3, 4, 5, 55, 44, 0, time.Local)

	res := []time.Time{
		time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local),
		time.Date(2022, 2, 2, 3, 4, 5, 0, time.Local),
		time.Date(2022, 3, 2, 3, 4, 5, 0, time.Local),
	}

	ass.Equal(res, RangeMonths(tm1, tm2))
}

func TestRangeYears(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	tm2 := time.Date(2024, 2, 4, 5, 55, 44, 0, time.Local)

	res := []time.Time{
		time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local),
		time.Date(2023, 1, 2, 3, 4, 5, 0, time.Local),
		time.Date(2024, 1, 2, 3, 4, 5, 0, time.Local),
	}

	ass.Equal(res, RangeYears(tm1, tm2))
}

func TestNewTime(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ass.Equal(time.Now().Unix(), NewTime().Unix())
	ass.Equal(time.Date(2022, 1, 2, 0, 0, 0, 0, time.Local), NewTime(2022, 1, 2))
	ass.Equal(time.Date(2022, 0, 0, 0, 0, 0, 0, time.Local), NewTime(2022))
}

func TestSubTime(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	t1 := time.Date(2022, 1, 2, 22, 0, 0, 0, time.Local)
	t2 := time.Date(2023, 1, 2, 1, 59, 6, 0, time.Local)

	day, hour, min, sec := SubTime(t2, t1)

	ass.Equal(364, day)
	ass.Equal(3, hour)
	ass.Equal(59, min)
	ass.Equal(6, sec)
}

func TestSubTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	t1 := time.Date(2022, 1, 2, 22, 0, 0, 0, time.Local)
	t2 := time.Date(2022, 1, 33, 1, 59, 6, 0, time.Local)

	day, hour, min, sec := SubTimestamp(t1.Unix(), t2.Unix())

	ass.Equal(30, day)
	ass.Equal(3, hour)
	ass.Equal(59, min)
	ass.Equal(6, sec)
}
