# Timeutil

------

Provides some time-related operations, including time formatting, time calculation, time interval calculation, time conversion, etc.

------

## Usage

```go
import "github.com/dengrandpa/jez/timeutil"
```

------

## Index

-   [ParseTime](#parseTime)
-   [ParseTimestamp](#parseTimestamp)
-   [StartOfMinute](#startOfMinute)
-   [EndOfMinute](#endOfMinute)
-   [StartOfHour](#startOfHour)
-   [EndOfHour](#endOfHour)
-   [StartOfDay](#startOfDay)
-   [EndOfDay](#endOfDay)
-   [StartOfWeekMonday](#startOfWeekMonday)
-   [EndOfWeekSunday](#endOfWeekSunday)
-   [StartOfWeekSunday](#startOfWeekSunday)
-   [EndOfWeekMonday](#endOfWeekMonday)
-   [StartOfMonth](#startOfMonth)
-   [EndOfMonth](#endOfMonth)
-   [StartOfYear](#startOfYear)
-   [EndOfYear](#endOfYear)
-   [AddSecond](#addSecond)
-   [AddMinute](#addMinute)
-   [AddHour](#addHour)
-   [AddDay](#addDay)
-   [AddWeek](#addWeek)
-   [AddMonth](#addMonth)
-   [AddYear](#addYear)
-   [FormatTime](#formatTime)
-   [FormatTimestamp](#formatTimestamp)
-   [FormatNow](#formatNow)
-   [IsLeapYear](#isLeapYear)
-   [RangeHours](#rangeHours)
-   [RangeDays](#rangeDays)
-   [RangeMonths](#rangeMonths)
-   [RangeYears](#rangeYears)
-   [NewTime](#newTime)
-   [SubTime](#subTime)
-   [SubTimestamp](#subTimestamp)
-   [CST](#cST)
-   [ToCST](#toCST)
------

### ParseTime
Convert the string to time, the default format is YYYYMMDDHHMMSS.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	tmStr := timeutil.FormatTime(tm.Local())

	fmt.Println(timeutil.ToCST(timeutil.ParseTime(tmStr)))

	// Output:
	// 2022-01-02 03:04:05 +0800 CST
}
`````````

### ParseTimestamp
Convert timestamps to time.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	localTm := tm.Local()

	fmt.Println(timeutil.ToCST(timeutil.ParseTimestamp(localTm.Unix())))

	// Output:
	// 2022-01-02 03:04:05 +0800 CST

}
```

### StartOfMinute
Return time t Start time of the minute yyyy-mm-dd hh:mm:00.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	fmt.Println(timeutil.StartOfMinute(tm))

	// Output:
	// 2022-01-02 03:04:00 +0800 CST
}
```

### EndOfMinute
Return time t The end time of the minute yyyy-mm-dd hh:mm:59.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.EndOfMinute(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 03:04:59.999999999 +0800 CST
}
```

### StartOfHour
Return time t The start time of the hour yyyy-mm-dd hh:00:00.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.StartOfHour(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 03:00:00 +0800 CST
}
```

### EndOfHour
Return time t The end time of the hour yyyy-mm-dd hh:59:59.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.EndOfHour(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 03:59:59.999999999 +0800 CST
}
```

### StartOfDay
Return time t Start time of the day yyyy-mm-dd 00:00:00.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.StartOfDay(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 00:00:00 +0800 CST
}
```

### EndOfDay
Return time t The end time of the day yyyy-mm-dd 23:59:59.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.EndOfDay(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 23:59:59.999999999 +0800 CST
}
```

### StartOfWeekMonday
Return time t The start time of the week, Monday is the first day yyyy-mm-dd 00:00:00.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.StartOfWeekMonday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday())

	// Output:
	// 2021-12-27 00:00:00 +0800 CST
	// Monday
}
```

### EndOfWeekSunday
Return time t The end time of the week, Sunday is the last day. yyyy-mm-dd 23:59:59.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.EndOfWeekSunday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday().String())

	// Output:
	// 2022-01-02 23:59:59.999999999 +0800 CST
	// Sunday
}
```

### StartOfWeekSunday
Return time t The start time of the week, Sunday is the first day. yyyy-mm-dd 00:00:00.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.StartOfWeekSunday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday().String())

	// Output:
	// 2021-12-26 00:00:00 +0800 CST
	// Sunday
}
```

### EndOfWeekMonday
Return time t The end time of the week where the week is located, Monday is the last day. yyyy-mm-dd 23:59:59.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.EndOfWeekMonday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday())

	// Output:
	// 2022-01-03 23:59:59.999999999 +0800 CST
	// Monday
}
```

### StartOfMonth
Return time t The start time of the month yyyy-mm-01 00:00:00.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.StartOfMonth(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-01 00:00:00 +0800 CST
}
```

### EndOfMonth
Return time t The end time of the month yyyy-mm-dd 23:59:59.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.EndOfMonth(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-31 23:59:59.999999999 +0800 CST
}
```

### StartOfYear
Return time t The start time of the year yyyy-01-01 00:00:00.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.StartOfYear(tm)

	fmt.Println(res)
	fmt.Println(res.Month())

	// Output:
	// 2022-01-01 00:00:00 +0800 CST
	// January
}
```

### EndOfYear
Return time t The end time of the year yyyy-12-31 23:59:59.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	res := timeutil.EndOfYear(tm)

	fmt.Println(res)
	fmt.Println(res.Month())

	// Output:
	// 2022-12-31 23:59:59.999999999 +0800 CST
	// December
}
```

### AddSecond
Add or delete seconds.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	fmt.Println(timeutil.AddSecond(tm, 10))

	// Output:
	// 2022-01-02 03:04:15 +0800 CST
}
```

### AddMinute
Add or delete minutes.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	fmt.Println(timeutil.AddMinute(tm, 10))

	// Output:
	// 2022-01-02 03:14:05 +0800 CST
}
```

### AddHour
Add or delete hours.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	fmt.Println(timeutil.AddHour(tm, 10))

	// Output:
	// 2022-01-02 13:04:05 +0800 CST
}
```

### AddDay
Add or delete days.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	fmt.Println(timeutil.AddDay(tm, 10))

	// Output:
	// 2022-01-12 03:04:05 +0800 CST
}
```

### AddWeek
Add or delete weeks.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	res := timeutil.AddWeek(tm, 10).Weekday()

	fmt.Println(res == tm.Weekday())

	// Output:
	// true
}
```

### AddMonth
Add or delete months.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	fmt.Println(timeutil.AddMonth(tm, 10).Month().String())

	// Output:
	// November
}
```

### AddYear
Add or delete years.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	fmt.Println(timeutil.AddYear(tm, 10))

	// Output:
	// 2032-01-02 03:04:05 +0800 CST
}
```

### FormatTime
Format the time as a string, the default format is yyyy-mm-dd hh:mm:ss.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	fmt.Println(timeutil.FormatTime(tm))
	fmt.Println(timeutil.FormatTime(tm, timeutil.YYYYMMDDHHMMSS2))

	// Output:
	// 2022-01-02 03:04:05
	// 2022/01/02 03:04:05

}
```

### FormatTimestamp
Format the timestamp as a string, the default format is yyyy-mm-dd hh:mm:ss.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	fmt.Println(timeutil.FormatTimestamp(tm.Local().Unix()) == timeutil.FormatTime(tm.Local()))

	// Output:
	// true
}
```

### FormatNow
Returns the string format of the current time, the default format is yyyy-mm-dd hh:mm:ss.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	now := timeutil.FormatTime(time.Now())
	res := timeutil.FormatNow()

	now1 := timeutil.FormatTime(time.Now(), timeutil.YYYYMMDDHHMMSS2)
	res1 := timeutil.FormatNow(timeutil.YYYYMMDDHHMMSS2)

	fmt.Println(now == res)
	fmt.Println(now1 == res1)

	// Output:
	// true
	// true
}
```

### IsLeapYear
Determine whether the year is a leap year.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	fmt.Println(timeutil.IsLeapYear(2012))
	fmt.Println(timeutil.IsLeapYear(2013))

	// Output:
	// true
	// false

}
```

### RangeHours
Returns a slice of all hours between two times, including start and end, that is, [start,end]. If start and end results are the same, only 1.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	tm2 := time.Date(2022, 1, 2, 5, 55, 44, 0, timeutil.CST())

	fmt.Println(timeutil.RangeHours(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2022-01-02 04:04:05 +0800 CST 2022-01-02 05:04:05 +0800 CST]

}
```

### RangeDays
Returns the slice of all days between two times, including start and end, that is, [start, end]. If the result of start and end is the same, only 1 is returned..

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	tm2 := time.Date(2022, 1, 4, 5, 55, 44, 0, timeutil.CST())

	fmt.Println(timeutil.RangeDays(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2022-01-03 03:04:05 +0800 CST 2022-01-04 03:04:05 +0800 CST]
}
```

### RangeMonths
Returns slices of all months between two times, including start and end, that is, [start,end]. If start and end results are the same, only 1.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	tm2 := time.Date(2022, 3, 4, 5, 55, 44, 0, timeutil.CST())

	fmt.Println(timeutil.RangeMonths(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2022-02-02 03:04:05 +0800 CST 2022-03-02 03:04:05 +0800 CST]
}
```

### RangeYears
Returns slices of all years between two times, including start and end, that is, [start, end]. If start and end results are the same, only 1.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())
	tm2 := time.Date(2024, 2, 4, 5, 55, 44, 0, timeutil.CST())

	fmt.Println(timeutil.RangeYears(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2023-01-02 03:04:05 +0800 CST 2024-01-02 03:04:05 +0800 CST]
}
```

### NewTime
It is equivalent to time.Date, and if the parameters are not passed, it is equivalent to time.Now..

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timeutil.CST())

	localTm := tm.In(time.Local)

	tm1 := timeutil.NewTime(localTm.Year(), int(localTm.Month()), localTm.Day(), localTm.Hour(), localTm.Minute(), localTm.Second())

	fmt.Println(timeutil.NewTime().Unix() == time.Now().Unix())
	fmt.Println(timeutil.ToCST(tm1))

	// Output:
	// true
	// 2022-01-02 03:04:05 +0800 CST

}
```

### SubTime
Returns days, hours, minutes and seconds of the difference between t1 and t2.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	t1 := time.Date(2022, 1, 2, 22, 0, 0, 0, timeutil.CST())
	t2 := time.Date(2023, 1, 2, 1, 59, 6, 0, timeutil.CST())

	fmt.Println(timeutil.SubTime(t2, t1))

	// Output:
	// 364 3 59 6
}
```

### SubTimestamp
Returns days, hours, minutes and seconds of the difference between t1 and t2.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	t1 := time.Date(2022, 1, 2, 22, 0, 0, 0, timeutil.CST())
	t2 := time.Date(2022, 1, 33, 1, 59, 6, 0, timeutil.CST())

	fmt.Println(timeutil.SubTimestamp(t1.Unix(), t2.Unix()))

	// Output:
	// 30 3 59 6
}
```

### CST
Return to China time zone.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {
	local := timeutil.CST()

	fmt.Println(local)

	// Output:
	// Asia/Shanghai
}
```

### ToCST
Convert time to Chinese time zone.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timeutil"
)

func main() {

	t1 := time.Date(2022, 1, 2, 3, 4, 5, 0, time.UTC)

	t2 := time.Date(2022, 1, 2, 11, 4, 5, 0, timeutil.CST())

	t3 := timeutil.ToCST(t1)

	fmt.Println(t3)

	fmt.Println(t2.Unix() == t3.Unix())

	// Output:
	// 2022-01-02 11:04:05 +0800 CST
	// true
}
```
