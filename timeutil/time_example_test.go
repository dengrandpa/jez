package timeutil

import (
	"fmt"
	"time"
)

func ExampleParseTime() {

	fmt.Println(FormatTime(time.Date(2022, 1, 2, 3, 4, 5, 0, CST())))
	fmt.Println(FormatTime(time.Date(2022, 1, 2, 3, 4, 5, 0, CST()), YYYYMMDDHHMMSS2))
	fmt.Println(FormatTime(time.Date(2022, 1, 2, 0, 0, 0, 0, CST()), YYYYMM))
	fmt.Println(FormatTime(time.Date(2022, 1, 2, 0, 0, 0, 0, CST()), YYYYMMDD))
	fmt.Println(FormatTime(time.Date(0, 1, 1, 3, 4, 5, 0, CST()), HHMMSS))
	fmt.Println(FormatTime(time.Date(0, 1, 1, 3, 4, 0, 0, CST()), HHMM))

	// Output:
	// 2022-01-02 03:04:05
	// 2022/01/02 03:04:05
	// 2022-01
	// 2022-01-02
	// 03:04:05
	// 03:04
}

func ExampleParseTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	localTm := tm.Local()

	fmt.Println(ToCST(ParseTimestamp(localTm.Unix())))

	// Output:
	// 2022-01-02 03:04:05 +0800 CST

}

func ExampleStartOfMinute() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	fmt.Println(StartOfMinute(tm))

	// Output:
	// 2022-01-02 03:04:00 +0800 CST
}

func ExampleEndOfMinute() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := EndOfMinute(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 03:04:59.999999999 +0800 CST
}

func ExampleStartOfHour() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := StartOfHour(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 03:00:00 +0800 CST
}

func ExampleEndOfHour() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := EndOfHour(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 03:59:59.999999999 +0800 CST
}

func ExampleStartOfDay() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := StartOfDay(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 00:00:00 +0800 CST
}

func ExampleEndOfDay() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := EndOfDay(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 23:59:59.999999999 +0800 CST
}

func ExampleStartOfWeekMonday() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := StartOfWeekMonday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday())

	// Output:
	// 2021-12-27 00:00:00 +0800 CST
	// Monday
}

func ExampleEndOfWeekSunday() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := EndOfWeekSunday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday().String())

	// Output:
	// 2022-01-02 23:59:59.999999999 +0800 CST
	// Sunday
}

func ExampleStartOfWeekSunday() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := StartOfWeekSunday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday().String())

	// Output:
	// 2021-12-26 00:00:00 +0800 CST
	// Sunday
}

func ExampleEndOfWeekMonday() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := EndOfWeekMonday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday())

	// Output:
	// 2022-01-03 23:59:59.999999999 +0800 CST
	// Monday
}

func ExampleStartOfMonth() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := StartOfMonth(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-01 00:00:00 +0800 CST
}

func ExampleEndOfMonth() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := EndOfMonth(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-31 23:59:59.999999999 +0800 CST
}

func ExampleStartOfYear() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := StartOfYear(tm)

	fmt.Println(res)
	fmt.Println(res.Month())

	// Output:
	// 2022-01-01 00:00:00 +0800 CST
	// January
}

func ExampleEndOfYear() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	res := EndOfYear(tm)

	fmt.Println(res)
	fmt.Println(res.Month())

	// Output:
	// 2022-12-31 23:59:59.999999999 +0800 CST
	// December
}

func ExampleAddSecond() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	fmt.Println(AddSecond(tm, 10))

	// Output:
	// 2022-01-02 03:04:15 +0800 CST
}

func ExampleAddMinute() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	fmt.Println(AddMinute(tm, 10))

	// Output:
	// 2022-01-02 03:14:05 +0800 CST
}

func ExampleAddHour() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	fmt.Println(AddHour(tm, 10))

	// Output:
	// 2022-01-02 13:04:05 +0800 CST
}

func ExampleAddDay() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	fmt.Println(AddDay(tm, 10))

	// Output:
	// 2022-01-12 03:04:05 +0800 CST
}

func ExampleAddWeek() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	res := AddWeek(tm, 10).Weekday()

	fmt.Println(res == tm.Weekday())

	// Output:
	// true
}

func ExampleAddMonth() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	fmt.Println(AddMonth(tm, 10).Month().String())

	// Output:
	// November
}

func ExampleAddYear() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	fmt.Println(AddYear(tm, 10))

	// Output:
	// 2032-01-02 03:04:05 +0800 CST
}

func ExampleFormatTime() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	fmt.Println(FormatTime(tm))
	fmt.Println(FormatTime(tm, YYYYMMDDHHMMSS2))

	// Output:
	// 2022-01-02 03:04:05
	// 2022/01/02 03:04:05

}

func ExampleFormatTimestamp() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	fmt.Println(FormatTimestamp(tm.Local().Unix()) == FormatTime(tm.Local()))

	// Output:
	// true
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

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	tm2 := time.Date(2022, 1, 2, 5, 55, 44, 0, CST())

	fmt.Println(RangeHours(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2022-01-02 04:04:05 +0800 CST 2022-01-02 05:04:05 +0800 CST]

}

func ExampleRangeDays() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	tm2 := time.Date(2022, 1, 4, 5, 55, 44, 0, CST())

	fmt.Println(RangeDays(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2022-01-03 03:04:05 +0800 CST 2022-01-04 03:04:05 +0800 CST]
}

func ExampleRangeMonths() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	tm2 := time.Date(2022, 3, 4, 5, 55, 44, 0, CST())

	fmt.Println(RangeMonths(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2022-02-02 03:04:05 +0800 CST 2022-03-02 03:04:05 +0800 CST]
}

func ExampleRangeYears() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())
	tm2 := time.Date(2024, 2, 4, 5, 55, 44, 0, CST())

	fmt.Println(RangeYears(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2023-01-02 03:04:05 +0800 CST 2024-01-02 03:04:05 +0800 CST]
}

func ExampleNewTime() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, CST())

	localTm := tm.In(time.Local)

	tm1 := NewTime(localTm.Year(), int(localTm.Month()), localTm.Day(), localTm.Hour(), localTm.Minute(), localTm.Second())

	fmt.Println(NewTime().Unix() == time.Now().Unix())
	fmt.Println(ToCST(tm1))

	// Output:
	// true
	// 2022-01-02 03:04:05 +0800 CST

}

func ExampleSubTime() {

	t1 := time.Date(2022, 1, 2, 22, 0, 0, 0, CST())
	t2 := time.Date(2023, 1, 2, 1, 59, 6, 0, CST())

	fmt.Println(SubTime(t2, t1))

	// Output:
	// 364 3 59 6
}

func ExampleSubTimestamp() {

	t1 := time.Date(2022, 1, 2, 22, 0, 0, 0, CST())
	t2 := time.Date(2022, 1, 33, 1, 59, 6, 0, CST())

	fmt.Println(SubTimestamp(t1.Unix(), t2.Unix()))

	// Output:
	// 30 3 59 6
}

func ExampleCST() {
	local := CST()

	fmt.Println(local)

	// Output:
	// Asia/Shanghai
}

func ExampleToCST() {

	t1 := time.Date(2022, 1, 2, 3, 4, 5, 0, time.UTC)

	t2 := time.Date(2022, 1, 2, 11, 4, 5, 0, CST())

	t3 := ToCST(t1)

	fmt.Println(t3)

	fmt.Println(t2.Unix() == t3.Unix())

	// Output:
	// 2022-01-02 11:04:05 +0800 CST
	// true
}
