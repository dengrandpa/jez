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
