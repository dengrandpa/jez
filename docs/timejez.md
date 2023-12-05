# Timejez

------

提供了时间相关的一些操作，包括时间格式化、时间计算、时间区间计算、时间转换等

------

## 用法

```go
import "github.com/dengrandpa/jez/timejez"
```

------

## 目录

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
将字符串转换为时间，默认格式为 YYYYMMDDHHMMSS

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	tmStr := timejez.FormatTime(tm.Local())

	fmt.Println(timejez.ToCST(timejez.ParseTime(tmStr)))

	// Output:
	// 2022-01-02 03:04:05 +0800 CST
}
```

### ParseTimestamp
将时间戳转换为时间

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	localTm := tm.Local()

	fmt.Println(timejez.ToCST(timejez.ParseTimestamp(localTm.Unix())))

	// Output:
	// 2022-01-02 03:04:05 +0800 CST

}
```

### StartOfMinute
返回时间 t 所在分钟的开始时间 yyyy-mm-dd hh:mm:00

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	fmt.Println(timejez.StartOfMinute(tm))

	// Output:
	// 2022-01-02 03:04:00 +0800 CST
}
```

### EndOfMinute
返回时间 t 所在分钟的结束时间 yyyy-mm-dd hh:mm:59

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.EndOfMinute(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 03:04:59.999999999 +0800 CST
}
```

### StartOfHour
返回时间 t 所在小时的开始时间 yyyy-mm-dd hh:00:00

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.StartOfHour(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 03:00:00 +0800 CST
}
```

### EndOfHour
返回时间 t 所在小时的结束时间 yyyy-mm-dd hh:59:59

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.EndOfHour(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 03:59:59.999999999 +0800 CST
}
```

### StartOfDay
返回时间 t 所在天的开始时间 yyyy-mm-dd 00:00:00

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.StartOfDay(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 00:00:00 +0800 CST
}
```

### EndOfDay
返回时间 t 所在天的结束时间 yyyy-mm-dd 23:59:59

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.EndOfDay(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-02 23:59:59.999999999 +0800 CST
}
```

### StartOfWeekMonday
返回时间 t 所在周的开始时间，周一为第一天 yyyy-mm-dd 00:00:00

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.StartOfWeekMonday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday())

	// Output:
	// 2021-12-27 00:00:00 +0800 CST
	// Monday
}
```

### EndOfWeekSunday
返回时间 t 所在周的结束时间，周日为最后一天 yyyy-mm-dd 23:59:59

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.EndOfWeekSunday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday().String())

	// Output:
	// 2022-01-02 23:59:59.999999999 +0800 CST
	// Sunday
}
```

### StartOfWeekSunday
返回时间 t 所在周的开始时间，周日为第一天 yyyy-mm-dd 00:00:00

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.StartOfWeekSunday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday().String())

	// Output:
	// 2021-12-26 00:00:00 +0800 CST
	// Sunday
}
```

### EndOfWeekMonday
返回时间 t 所在周的结束时间，周一为最后一天 yyyy-mm-dd 23:59:59

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.EndOfWeekMonday(tm)

	fmt.Println(res)
	fmt.Println(res.Weekday())

	// Output:
	// 2022-01-03 23:59:59.999999999 +0800 CST
	// Monday
}
```

### StartOfMonth
返回时间 t 所在月的开始时间 yyyy-mm-01 00:00:00

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.StartOfMonth(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-01 00:00:00 +0800 CST
}
```

### EndOfMonth
返回时间 t 所在月的结束时间 yyyy-mm-dd 23:59:59

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.EndOfMonth(tm)

	fmt.Println(res)

	// Output:
	// 2022-01-31 23:59:59.999999999 +0800 CST
}
```

### StartOfYear
返回时间 t 所在年的开始时间 yyyy-01-01 00:00:00

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.StartOfYear(tm)

	fmt.Println(res)
	fmt.Println(res.Month())

	// Output:
	// 2022-01-01 00:00:00 +0800 CST
	// January
}
```

### EndOfYear
返回时间 t 所在年的结束时间 yyyy-12-31 23:59:59

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	res := timejez.EndOfYear(tm)

	fmt.Println(res)
	fmt.Println(res.Month())

	// Output:
	// 2022-12-31 23:59:59.999999999 +0800 CST
	// December
}
```

### AddSecond
添加或删除秒数

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	fmt.Println(timejez.AddSecond(tm, 10))

	// Output:
	// 2022-01-02 03:04:15 +0800 CST
}
```

### AddMinute
添加或删除分钟数

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	fmt.Println(timejez.AddMinute(tm, 10))

	// Output:
	// 2022-01-02 03:14:05 +0800 CST
}
```

### AddHour
添加或删除小时数

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	fmt.Println(timejez.AddHour(tm, 10))

	// Output:
	// 2022-01-02 13:04:05 +0800 CST
}
```

### AddDay
添加或删除天数

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	fmt.Println(timejez.AddDay(tm, 10))

	// Output:
	// 2022-01-12 03:04:05 +0800 CST
}
```

### AddWeek
添加或删除周数

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	res := timejez.AddWeek(tm, 10).Weekday()

	fmt.Println(res == tm.Weekday())

	// Output:
	// true
}
```

### AddMonth
添加或删除月数

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	fmt.Println(timejez.AddMonth(tm, 10).Month().String())

	// Output:
	// November
}
```

### AddYear
添加或删除年数

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	fmt.Println(timejez.AddYear(tm, 10))

	// Output:
	// 2032-01-02 03:04:05 +0800 CST
}
```

### FormatTime
将时间格式化为字符串，默认格式为 yyyy-mm-dd hh:mm:ss

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	fmt.Println(timejez.FormatTime(tm))
	fmt.Println(timejez.FormatTime(tm, timejez.YYYYMMDDHHMMSS2))

	// Output:
	// 2022-01-02 03:04:05
	// 2022/01/02 03:04:05

}
```

### FormatTimestamp
将时间戳格式化为字符串，默认格式为 yyyy-mm-dd hh:mm:ss

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	fmt.Println(timejez.FormatTimestamp(tm.Local().Unix()) == timejez.FormatTime(tm.Local()))

	// Output:
	// true
}
```

### FormatNow
返回当前时间的字符串格式，默认格式为 yyyy-mm-dd hh:mm:ss

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	now := timejez.FormatTime(time.Now())
	res := timejez.FormatNow()

	now1 := timejez.FormatTime(time.Now(), timejez.YYYYMMDDHHMMSS2)
	res1 := timejez.FormatNow(timejez.YYYYMMDDHHMMSS2)

	fmt.Println(now == res)
	fmt.Println(now1 == res1)

	// Output:
	// true
	// true
}
```

### IsLeapYear
判断年份 year 是否为闰年

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	fmt.Println(timejez.IsLeapYear(2012))
	fmt.Println(timejez.IsLeapYear(2013))

	// Output:
	// true
	// false

}
```

### RangeHours
返回两个时间之间的所有小时的切片，包含 start 和 end，即[start,end]，如果start和end结果一样，则只返回1个

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	tm2 := time.Date(2022, 1, 2, 5, 55, 44, 0, timejez.CST())

	fmt.Println(timejez.RangeHours(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2022-01-02 04:04:05 +0800 CST 2022-01-02 05:04:05 +0800 CST]

}
```

### RangeDays
返回两个时间之间的所有天的切片，包含 start 和 end，即[start,end]，如果start和end结果一样，则只返回1个

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	tm2 := time.Date(2022, 1, 4, 5, 55, 44, 0, timejez.CST())

	fmt.Println(timejez.RangeDays(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2022-01-03 03:04:05 +0800 CST 2022-01-04 03:04:05 +0800 CST]
}
```

### RangeMonths
返回两个时间之间的所有月的切片，包含 start 和 end，即[start,end]，如果start和end结果一样，则只返回1个

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	tm2 := time.Date(2022, 3, 4, 5, 55, 44, 0, timejez.CST())

	fmt.Println(timejez.RangeMonths(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2022-02-02 03:04:05 +0800 CST 2022-03-02 03:04:05 +0800 CST]
}
```

### RangeYears
返回两个时间之间的所有年的切片，包含 start 和 end，即[start,end]，如果start和end结果一样，则只返回1个

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm1 := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())
	tm2 := time.Date(2024, 2, 4, 5, 55, 44, 0, timejez.CST())

	fmt.Println(timejez.RangeYears(tm1, tm2))

	// Output:
	// [2022-01-02 03:04:05 +0800 CST 2023-01-02 03:04:05 +0800 CST 2024-01-02 03:04:05 +0800 CST]
}
```

### NewTime
相当于 time.Date，如果不传参数则相当于 time.Now。

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	tm := time.Date(2022, 1, 2, 3, 4, 5, 0, timejez.CST())

	localTm := tm.In(time.Local)

	tm1 := timejez.NewTime(localTm.Year(), int(localTm.Month()), localTm.Day(), localTm.Hour(), localTm.Minute(), localTm.Second())

	fmt.Println(timejez.NewTime().Unix() == time.Now().Unix())
	fmt.Println(timejez.ToCST(tm1))

	// Output:
	// true
	// 2022-01-02 03:04:05 +0800 CST

}
```

### SubTime
返回 t1 和 t2 之间的差值的日、小时、分钟和秒

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	t1 := time.Date(2022, 1, 2, 22, 0, 0, 0, timejez.CST())
	t2 := time.Date(2023, 1, 2, 1, 59, 6, 0, timejez.CST())

	fmt.Println(timejez.SubTime(t2, t1))

	// Output:
	// 364 3 59 6
}
```

### SubTimestamp
返回 t1 和 t2 之间的差值的日、小时、分钟和秒

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	t1 := time.Date(2022, 1, 2, 22, 0, 0, 0, timejez.CST())
	t2 := time.Date(2022, 1, 33, 1, 59, 6, 0, timejez.CST())

	fmt.Println(timejez.SubTimestamp(t1.Unix(), t2.Unix()))

	// Output:
	// 30 3 59 6
}
```

### CST
返回中国时区

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {
	local := timejez.CST()

	fmt.Println(local)

	// Output:
	// Asia/Shanghai
}
```

### ToCST
将时间转换为中国时区

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/dengrandpa/jez/timejez"
)

func main() {

	t1 := time.Date(2022, 1, 2, 3, 4, 5, 0, time.UTC)

	t2 := time.Date(2022, 1, 2, 11, 4, 5, 0, timejez.CST())

	t3 := timejez.ToCST(t1)

	fmt.Println(t3)

	fmt.Println(t2.Unix() == t3.Unix())

	// Output:
	// 2022-01-02 11:04:05 +0800 CST
	// true
}
```
