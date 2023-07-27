package timeutil

import (
	"fmt"
	"time"
)

func ExampleParseTime() {

	fmt.Println(ParseTime("2022-01-02 03:04:05", YYYYMMDDHHMMSS))
	fmt.Println(ParseTime("2022-01-02 03:04:05", YYYYMMDDHHMMSS))
	fmt.Println(ParseTime("2022/01/02 03:04:05", YYYYMMDDHHMMSS2))
	fmt.Println(ParseTime("2022-01", YYYYMM))
	fmt.Println(ParseTime("2022-01-02", YYYYMMDD))
	fmt.Println(ParseTime("03:04:05", HHMMSS))
	fmt.Println(ParseTime("03:04", HHMM))

	// Output:
	// 2022-01-02 03:04:05 +0800 CST
	// 2022-01-02 03:04:05 +0800 CST
	// 2022-01-02 03:04:05 +0800 CST
	// 2022-01-01 00:00:00 +0800 CST
	// 2022-01-02 00:00:00 +0800 CST
	// 0000-01-01 03:04:05 +0805 LMT
	// 0000-01-01 03:04:00 +0805 LMT
}

func ExampleParseTimestamp() {

	fmt.Println(ParseTimestamp(1641063845))

	// Output:
	// 2022-01-02 03:04:05 +0800 CST

}

func ExampleStartOfMinute() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	fmt.Println(StartOfMinute(tm))

	// Output:
	// 2022-01-02 03:04:00 +0800 CST
}

func ExampleEndOfMinute() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := EndOfMinute(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 03:04:59.999999999 +0800 CST
}

func ExampleStartOfHour() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := StartOfHour(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 03:00:00 +0800 CST
}

func ExampleEndOfHour() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := EndOfHour(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 03:59:59.999999999 +0800 CST
}

func ExampleStartOfDay() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := StartOfDay(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 00:00:00 +0800 CST
}

func ExampleEndOfDay() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := EndOfDay(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 23:59:59.999999999 +0800 CST
}

func ExampleStartOfWeekMonday() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := StartOfWeekMonday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday())

	// Output:
	// 2021-12-27 00:00:00 +0800 CST
	// Monday
}

func ExampleEndOfWeekSunday() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := EndOfWeekSunday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday().String())

	// Output:
	// 2022-01-02 23:59:59.999999999 +0800 CST
	// Sunday
}

func ExampleStartOfWeekSunday() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := StartOfWeekSunday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday().String())

	// Output:
	// 2021-12-26 00:00:00 +0800 CST
	// Sunday
}

func ExampleEndOfWeekMonday() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := EndOfWeekMonday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday())

	// Output:
	// 2022-01-03 23:59:59.999999999 +0800 CST
	// Monday
}

func ExampleStartOfMonth() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := StartOfMonth(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-01 00:00:00 +0800 CST
}

func ExampleEndOfMonth() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := EndOfMonth(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-31 23:59:59.999999999 +0800 CST
}

func ExampleStartOfYear() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := StartOfYear(tm)

	fmt.Println(res)
	fmt.Println(res.Month())

	// Output:
	// 2022-01-01 00:00:00 +0800 CST
	// January
}

func ExampleEndOfYear() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	res := EndOfYear(tm)

	fmt.Println(res)
	fmt.Println(res.Month())

	// Output:
	// 2022-12-31 23:59:59.999999999 +0800 CST
	// December
}

func ExampleAddSecond() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	fmt.Println(AddSecond(tm, 10))

	// Output:
	// 2022-01-02 03:04:15 +0800 CST
}

func ExampleAddMinute() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	fmt.Println(AddMinute(tm, 10))

	// Output:
	// 2022-01-02 03:14:05 +0800 CST
}

func ExampleAddHour() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	fmt.Println(AddHour(tm, 10))

	// Output:
	// 2022-01-02 13:04:05 +0800 CST
}

func ExampleAddDay() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	fmt.Println(AddDay(tm, 10))

	// Output:
	// 2022-01-12 03:04:05 +0800 CST
}

func ExampleAddWeek() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	res := AddWeek(tm, 10).Weekday()

	fmt.Println(res == tm.Weekday())

	// Output:
	// true
}

func ExampleAddMonth() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	fmt.Println(AddMonth(tm, 10).Month().String())

	// Output:
	// November
}

func ExampleAddYear() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	fmt.Println(AddYear(tm, 10))

	// Output:
	// 2032-01-02 03:04:05 +0800 CST
}

func ExampleFormatTime() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	fmt.Println(FormatTime(tm))
	fmt.Println(FormatTime(tm, YYYYMMDDHHMMSS2))

	// Output:
	// 2022-01-02 03:04:05
	// 2022/01/02 03:04:05

}

func ExampleFormatTimestamp() {

	tm := int64(1641063845)

	fmt.Println(FormatTimestamp(tm))
	fmt.Println(FormatTimestamp(tm, YYYYMMDDHHMMSS2))

	// Output:
	// 2022-01-02 03:04:05
	// 2022/01/02 03:04:05
}

func ExampleFormatNow() {

	now := FormatTime(time.Now())
	res := FormatNow()

	now1 := FormatTime(time.Now(), YYYYMMDDHHMMSS2)
	res1 := FormatNow(YYYYMMDDHHMMSS2)

	fmt.Println(now == res)
	fmt.Println(now1 == res1)

	// Output:
	// true
	// true
}

func ExampleFormatStartOfMinute() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfMinute(tm)
	f2 := FormatStartOfMinute(tm, YYYYMMDDHHMMSS2)

	fmt.Println(f1)
	fmt.Println(f2)

	// Output:
	// 2022-01-02 03:04:00
	// 2022/01/02 03:04:00
}

func ExampleFormatEndOfMinute() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfMinute(tm)
	f2 := FormatEndOfMinute(tm, YYYYMMDDHHMMSS2)

	fmt.Println(f1)
	fmt.Println(f2)

	// Output:
	// 2022-01-02 03:04:59
	// 2022/01/02 03:04:59
}

func ExampleFormatStartOfHour() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfHour(tm)
	f2 := FormatStartOfHour(tm, YYYYMMDDHHMMSS2)

	fmt.Println(f1)
	fmt.Println(f2)

	// Output:
	// 2022-01-02 03:00:00
	// 2022/01/02 03:00:00
}

func ExampleFormatEndOfHour() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfHour(tm)
	f2 := FormatEndOfHour(tm, YYYYMMDDHHMMSS2)

	fmt.Println(f1)
	fmt.Println(f2)

	// Output:
	// 2022-01-02 03:59:59
	// 2022/01/02 03:59:59
}

func ExampleFormatStartOfDay() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfDay(tm)
	f2 := FormatStartOfDay(tm, YYYYMMDDHHMMSS2)

	fmt.Println(f1)
	fmt.Println(f2)

	// Output:
	// 2022-01-02 00:00:00
	// 2022/01/02 00:00:00
}

func ExampleFormatEndOfDay() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfDay(tm)
	f2 := FormatEndOfDay(tm, YYYYMMDDHHMMSS2)

	fmt.Println(f1)
	fmt.Println(f2)

	// Output:
	// 2022-01-02 23:59:59
	// 2022/01/02 23:59:59
}

func ExampleFormatStartOfWeekMonday() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfWeekMonday(tm)
	f1Tm := ParseTime(f1)

	fmt.Println(f1Tm.Weekday().String())

	f2 := FormatStartOfWeekMonday(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	fmt.Println(f2Tm.Weekday())

	// Output:
	// Monday
	// Monday
}

func ExampleFormatEndOfWeekSunday() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfWeekSunday(tm)
	f1Tm := ParseTime(f1)

	fmt.Println(f1Tm.Weekday().String())

	f2 := FormatEndOfWeekSunday(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	fmt.Println(f2Tm.Weekday().String())

	// Output:
	// Sunday
	// Sunday
}

func ExampleFormatStartOfWeekSunday() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfWeekSunday(tm)
	f1Tm := ParseTime(f1)

	fmt.Println(f1Tm.Weekday().String())

	f2 := FormatStartOfWeekSunday(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	fmt.Println(f2Tm.Weekday().String())

	// Output:
	// Sunday
	// Sunday

}

func ExampleFormatEndOfWeekMonday() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfWeekMonday(tm)
	f1Tm := ParseTime(f1)

	fmt.Println(f1Tm.Weekday().String())

	f2 := FormatEndOfWeekMonday(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	fmt.Println(f2Tm.Weekday().String())

	// Output:
	// Monday
	// Monday
}

func ExampleFormatStartOfMonth() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfMonth(tm)
	f1Tm := ParseTime(f1)

	fmt.Println(f1Tm.Day())

	f2 := FormatStartOfMonth(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	fmt.Println(f2Tm.Day())

	// Output:
	// 1
	// 1

}

func ExampleFormatEndOfMonth() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfMonth(tm)
	f1Tm := ParseTime(f1)

	fmt.Println(f1Tm.Day())

	f2 := FormatEndOfMonth(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	fmt.Println(f2Tm.Day())

	// Output:
	// 31
	// 31
}

func ExampleFormatStartOfYear() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatStartOfYear(tm)
	f1Tm := ParseTime(f1)

	fmt.Println(f1Tm.Month().String())

	f2 := FormatStartOfYear(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)

	fmt.Println(f2Tm.Month().String())

	// Output:
	// January
	// January
}

func ExampleFormatEndOfYear() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)

	f1 := FormatEndOfYear(tm)
	f1Tm := ParseTime(f1)
	fmt.Println(f1Tm.Month().String())

	f2 := FormatEndOfYear(tm, YYYYMMDDHHMMSS2)
	f2Tm := ParseTime(f2, YYYYMMDDHHMMSS2)
	fmt.Println(f2Tm.Month().String())

	// Output:
	// December
	// December
}

func ExampleStartOfMinuteTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := StartOfMinuteTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1641063840
}

func ExampleEndOfMinuteTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := EndOfMinuteTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1641063899
}

func ExampleStartOfHourTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := StartOfHourTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1641063600
}

func ExampleEndOfHourTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := EndOfHourTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1641067199
}

func ExampleStartOfDayTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := StartOfDayTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1641052800
}

func ExampleEndOfDayTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := EndOfDayTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1641139199
}

func ExampleStartOfWeekMondayTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := StartOfWeekMondayTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1640534400
}

func ExampleEndOfWeekSundayTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := EndOfWeekSundayTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1641139199
}

func ExampleStartOfWeekSundayTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := StartOfWeekSundayTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1640448000
}

func ExampleEndOfWeekMondayTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := EndOfWeekMondayTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1641225599
}

func ExampleStartOfMonthTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := StartOfMonthTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1640966400
}

func ExampleEndOfMonthTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := EndOfMonthTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1643644799
}

func ExampleStartOfYearTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := StartOfYearTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1640966400
}

func ExampleEndOfYearTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	t1 := EndOfYearTimestamp(tm)

	fmt.Println(t1)

	// Output:
	// 1672502399
}

func ExampleIsLeapYear() {

	fmt.Println(IsLeapYear(2012))
	fmt.Println(IsLeapYear(2013))

	// Output:
	// true
	// false

}

func ExampleRangeHours() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	tm2 := time.Date(2022, 1, 2, 5, 55, 44, 0, time.Local)

	fmt.Println(RangeHours(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2022-01-02 04:04:05 +0800 CST 2022-01-02 05:04:05 +0800 CST]

}

func ExampleRangeDays() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	tm2 := time.Date(2022, 1, 4, 5, 55, 44, 0, time.Local)

	fmt.Println(RangeDays(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2022-01-03 03:04:05 +0800 CST 2022-01-04 03:04:05 +0800 CST]
}

func ExampleRangeMonths() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	tm2 := time.Date(2022, 3, 4, 5, 55, 44, 0, time.Local)

	fmt.Println(RangeMonths(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2022-02-02 03:04:05 +0800 CST 2022-03-02 03:04:05 +0800 CST]
}

func ExampleRangeYears() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, time.Local)
	tm2 := time.Date(2024, 2, 4, 5, 55, 44, 0, time.Local)

	fmt.Println(RangeYears(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2023-01-02 03:04:05 +0800 CST 2024-01-02 03:04:05 +0800 CST]
}

func ExampleNewTime() {

	timestamp := NewTime().Unix()
	tm1 := NewTime(2022, 1, 2)
	tm2 := NewTime(2022, 1, 2, 3, 4, 5)

	fmt.Println(timestamp == time.Now().Unix())
	fmt.Println(tm1)
	fmt.Println(tm2)

	// Output:
	// true
	// 2022-01-02 00:00:00 +0800 CST
	// 2022-01-02 03:04:05 +0800 CST

}

func ExampleSubTime() {

	t1 := time.Date(2022, 1, 2, 22, 0, 0, 0, time.Local)
	t2 := time.Date(2023, 1, 2, 1, 59, 6, 0, time.Local)

	fmt.Println(SubTime(t2, t1))

	// Output:
	// 364 3 59 6
}

func ExampleSubTimestamp() {

	t1 := time.Date(2022, 1, 2, 22, 0, 0, 0, time.Local)
	t2 := time.Date(2022, 1, 33, 1, 59, 6, 0, time.Local)

	fmt.Println(SubTimestamp(t1.Unix(), t2.Unix()))

	// Output:
	// 30 3 59 6
}
