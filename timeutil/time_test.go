package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getTimeFormat(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
	tf := getTimeFormat([]TimeFormat{YYYYMMDD})
	tf1 := getTimeFormat(nil)

	ass.Equal(YYYYMMDD, tf)
	ass.Equal(YYYYMMDDHHMMSS, tf1)
}

func TestParseTime(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	tmStr := FormatTime(tm.Local())

	ass.Equal(tm, ToCST(ParseTime(tmStr)))
}

func TestParseTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	ass.Equal(ToCST(tm), ToCST(ParseTimestamp(tm.Unix())))
}

func TestStartOfMinute(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	res := StartOfMinute(tm)

	ass.Equal(0, res.Second())

}

func TestEndOfMinute(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	res := EndOfMinute(tm)

	ass.Equal(59, res.Second())
}

func TestStartOfHour(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	res := StartOfHour(tm)

	ass.Equal(0, res.Second())
	ass.Equal(0, res.Minute())
}

func TestEndOfHour(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	res := EndOfHour(tm)

	ass.Equal(59, res.Second())
	ass.Equal(59, res.Minute())
}

func TestStartOfDay(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	res := StartOfDay(tm)

	ass.Equal(0, res.Second())
	ass.Equal(0, res.Minute())
	ass.Equal(0, res.Hour())
}

func TestEndOfDay(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	res := EndOfDay(tm)

	ass.Equal(59, res.Second())
	ass.Equal(59, res.Minute())
	ass.Equal(23, res.Hour())
}

func TestStartOfWeekMonday(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := StartOfWeekMonday(tm)

	ass.Equal(0, res.Second())
	ass.Equal(0, res.Minute())
	ass.Equal(0, res.Hour())
	ass.Equal(time.Monday, res.Weekday())
}

func TestEndOfWeekSunday(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	res := EndOfWeekSunday(tm)

	ass.Equal(59, res.Second())
	ass.Equal(59, res.Minute())
	ass.Equal(23, res.Hour())
	ass.Equal(time.Sunday, res.Weekday())
}

func TestStartOfWeekSunday(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	res := StartOfWeekSunday(tm)

	ass.Equal(0, res.Second())
	ass.Equal(0, res.Minute())
	ass.Equal(0, res.Hour())
	ass.Equal(time.Sunday, res.Weekday())
}

func TestEndOfWeekMonday(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	res := EndOfWeekMonday(tm)

	ass.Equal(59, res.Second())
	ass.Equal(59, res.Minute())
	ass.Equal(23, res.Hour())
	ass.Equal(time.Monday, res.Weekday())
}

func TestStartOfMonth(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	res := StartOfMonth(tm)

	ass.Equal(0, res.Second())
	ass.Equal(0, res.Minute())
	ass.Equal(0, res.Hour())
	ass.Equal(1, res.Day())
}

func TestEndOfMonth(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	res := EndOfMonth(tm)

	ass.Equal(59, res.Second())
	ass.Equal(59, res.Minute())
	ass.Equal(23, res.Hour())
	ass.Equal(31, res.Day())
}

func TestStartOfYear(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

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

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

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

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	ass.Equal(15, AddSecond(tm, 10).Second())
}

func TestAddMinute(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	ass.Equal(14, AddMinute(tm, 10).Minute())
}

func TestAddHour(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	ass.Equal(13, AddHour(tm, 10).Hour())
}

func TestAddDay(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	ass.Equal(12, AddDay(tm, 10).Day())
}

func TestAddWeek(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	ass.Equal(tm.Weekday(), AddWeek(tm, 7).Weekday())
}

func TestAddMonth(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	ass.Equal(time.November, AddMonth(tm, 10).Month())
}

func TestAddYear(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

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

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	ass.Equal(FormatTime(tm.Local()), FormatTimestamp(tm.Local().Unix()))
	ass.Equal(FormatTime(tm.Local(), YYYYMMDDHHMMSS2), FormatTimestamp(tm.Local().Unix(), YYYYMMDDHHMMSS2))
}

func TestFormatNow(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm := FormatTime(time.Now())
	tm2 := FormatTime(time.Now(), YYYYMMDDHHMMSS2)

	ass.Equal(tm, FormatNow())
	ass.Equal(tm2, FormatNow(YYYYMMDDHHMMSS2))
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

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	tm2 := time.Date(2022, 1, 2, 5, 55, 44, 0, CST())

	res := []time.Time{
		time.Date(2022, 1, 2, 3, 4, 5, 0, CST()),
		time.Date(2022, 1, 2, 4, 4, 5, 0, CST()),
		time.Date(2022, 1, 2, 5, 4, 5, 0, CST()),
	}

	ass.Equal(res, RangeHours(tm1, tm2))
}

func TestRangeDays(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	tm2 := time.Date(2022, 1, 4, 5, 55, 44, 0, CST())

	res := []time.Time{
		time.Date(2022, 1, 2, 3, 4, 5, 0, CST()),
		time.Date(2022, 1, 3, 3, 4, 5, 0, CST()),
		time.Date(2022, 1, 4, 3, 4, 5, 0, CST()),
	}

	ass.Equal(res, RangeDays(tm1, tm2))
}

func TestRangeMonths(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	tm2 := time.Date(2022, 3, 4, 5, 55, 44, 0, CST())

	res := []time.Time{
		time.Date(2022, 1, 2, 3, 4, 5, 0, CST()),
		time.Date(2022, 2, 2, 3, 4, 5, 0, CST()),
		time.Date(2022, 3, 2, 3, 4, 5, 0, CST()),
	}

	ass.Equal(res, RangeMonths(tm1, tm2))
}

func TestRangeYears(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	tm2 := time.Date(2024, 2, 4, 5, 55, 44, 0, CST())

	res := []time.Time{
		time.Date(2022, 1, 2, 3, 4, 5, 0, CST()),
		time.Date(2023, 1, 2, 3, 4, 5, 0, CST()),
		time.Date(2024, 1, 2, 3, 4, 5, 0, CST()),
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

	t1 := time.Date(2022, 1, 2, 22, 0, 0, 0, CST())
	t2 := time.Date(2023, 1, 2, 1, 59, 6, 0, CST())

	day, hour, min, sec := SubTime(t2, t1)

	ass.Equal(364, day)
	ass.Equal(3, hour)
	ass.Equal(59, min)
	ass.Equal(6, sec)
}

func TestSubTimestamp(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	t1 := time.Date(2022, 1, 2, 22, 0, 0, 0, CST())
	t2 := time.Date(2022, 1, 33, 1, 59, 6, 0, CST())

	day, hour, min, sec := SubTimestamp(t1.Unix(), t2.Unix())

	ass.Equal(30, day)
	ass.Equal(3, hour)
	ass.Equal(59, min)
	ass.Equal(6, sec)
}

func TestCST(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	local := CST()

	ass.Equal("Asia/Shanghai", local.String())
}

func TestToCST(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	t1 := time.Date(2022, 1, 2, 3, 4, 5, 0, time.UTC)

	t2 := time.Date(2022, 1, 2, 11, 4, 5, 0, CST())

	t3 := ToCST(t1)

	ass.Equal(t2, t3)

}
