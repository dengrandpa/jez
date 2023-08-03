# Randomjez

------

提供了一些随机数生成的函数。

------

## 用法

```go
import "github.com/dengrandpa/jez/randomjez"
```

------

## 目录

-   [Random](#random)
-   [RandomLower](#randomLower)
-   [RandomUpper](#randomUpper)
-   [RandomNumeral](#randomNumeral)
-   [RandomCaseLetters](#randomCaseLetters)
-   [RandomLowerNumeral](#randomLowerNumeral)
-   [RandomUpperNumeral](#randomUpperNumeral)
-   [RandomCharset](#randomCharset)
-   [RandomInt](#randomInt)
-   [RandomIntSlice](#randomIntSlice)
-   [RandomIntSliceUnique](#randomIntSliceUnique)
-   [RandomBytes](#randomBytes)
-   [UUIDv4](#uUIDv4)
-   [Shuffle](#shuffle)
-   [Sample](#sample)
-   [Samples](#samples)

------

### Random
随机生成字符串

- 注意事项：
  - 当 n < 1 时，会 panic
  - 即使 unique 为 true 时，但如果传入的字符串有重复元素，那么生成的字符串可能也会重复
  - 当 unique 为 true 时，如果 n 大于 s 的长度，那么会 panic

```go
package main

import (
	"fmt"
	"regexp"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {
	res := randomjez.Random(randomjez.Numeral, 11)

	reg := regexp.MustCompile(`^[0-9]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { randomjez.Random(randomjez.Numeral, 0) })
	isPanic(func() { randomjez.Random(randomjez.Numeral, 11, true) })

	// Output:
	// true
	// 11
	// true
	// true
}
`````````

### RandomLower
随机生成小写字母字符串

- 注意事项：
  - 当 n < 1 时，会 panic
  - 当 unique 为 true 时，如果 n 大于 26（26个字母），那么会 panic

```go
package main

import (
	"fmt"
	"regexp"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {
	res := randomjez.RandomLower(27)

	reg := regexp.MustCompile(`^[a-z]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { randomjez.RandomLower(0) })
	isPanic(func() { randomjez.RandomLower(27, true) })

	// Output:
	// true
	// 27
	// true
	// true
}
```

### RandomUpper
随机生成大写字母字符串

- 注意事项：
  - 当 n < 1 时，会 panic
  - 当 unique 为 true 时，如果 n 大于 26（26个字母），那么会 panic

```go
package main

import (
	"fmt"
	"regexp"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {
	res := randomjez.RandomUpper(27)

	reg := regexp.MustCompile(`^[A-Z]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { randomjez.RandomUpper(0) })
	isPanic(func() { randomjez.RandomUpper(27, true) })

	// Output:
	// true
	// 27
	// true
	// true
}
```

### RandomNumeral
随机生成数字字符串

- 注意事项：
  - 当 n < 1 时，会 panic
  - 当 unique 为 true 时，如果 n 大于 10（10个数字），那么会 panic

```go
package main

import (
	"fmt"
	"regexp"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {
	res := randomjez.RandomNumeral(11)

	reg := regexp.MustCompile(`^[0-9]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { randomjez.RandomNumeral(0) })
	isPanic(func() { randomjez.RandomNumeral(11, true) })

	// Output:
	// true
	// 11
	// true
	// true
}
```

### RandomCaseLetters
随机生成大小写字母字符串。

- 注意事项：
  - 当 n < 1 时，会 panic
  - 当 unique 为 true 时，如果 n 大于 52（26个大写字母+26个小写字母），那么会 panic

```go
package main

import (
	"fmt"
	"regexp"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {
	res := randomjez.RandomCaseLetters(53)

	reg := regexp.MustCompile(`^[A-Za-z]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { randomjez.RandomCaseLetters(0) })
	isPanic(func() { randomjez.RandomCaseLetters(53, true) })

	// Output:
	// true
	// 53
	// true
	// true
}
```

### RandomLowerNumeral
随机生成小写字母和数字字符串

- 注意事项：
  - 当 n < 1 时，会 panic
  - 当 unique 为 true 时，如果 n 大于 36（10个数字+26个字母），那么会 panic

```go
package main

import (
	"fmt"
	"regexp"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {
	res := randomjez.RandomLowerNumeral(37)

	reg := regexp.MustCompile(`^[a-z0-9]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { randomjez.RandomLowerNumeral(0) })
	isPanic(func() { randomjez.RandomLowerNumeral(37, true) })

	// Output:
	// true
	// 37
	// true
	// true
}
```

### RandomUpperNumeral
随机生成大写字母和数字字符串

- 注意事项：
  - 当 n < 1 时，会 panic
  - 当 unique 为 true 时，如果 n 大于 36（10个数字+26个字母），那么会 panic

```go
package main

import (
	"fmt"
	"regexp"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {
	res := randomjez.RandomUpperNumeral(37)

	reg := regexp.MustCompile(`^[A-Z0-9]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { randomjez.RandomUpperNumeral(0) })
	isPanic(func() { randomjez.RandomUpperNumeral(37, true) })

	// Output:
	// true
	// 37
	// true
	// true
}
```

### RandomCharset
随机生成字符串，包含数字、大小写字母

- 注意事项：
  - 当 n < 1 时，会 panic
  - 当 unique 为 true 时，如果 n 大于 62（10个数字+26个大写字母+26个小写字母），那么会 panic

```go
package main

import (
	"fmt"
	"regexp"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {
	res := randomjez.RandomCharset(63)

	reg := regexp.MustCompile(`^[A-Za-z0-9]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { randomjez.RandomUpperNumeral(0) })
	isPanic(func() { randomjez.RandomUpperNumeral(63, true) })

	// Output:
	// true
	// 63
	// true
	// true
}
```

### RandomInt
随机生成整数，包含 min，不包含 max，即 [min,max)。

- 注意事项：
  - 当 min > max 时，会交换 min 和 max 的值

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {
	min := 10
	max := 15

	res := randomjez.RandomInt(min, max)

	if res < min || res >= max {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}

	// Output:
	// true
}
```

### RandomIntSlice
随机生成整数切片

- 注意事项：
  - 当 min > max 时，会交换 min 和 max 的值。
  - 当 n < 1 时，会 panic

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {
	min := 10
	max := 15

	res := randomjez.RandomIntSlice(min, max, 10)

	ok := true

	for _, r := range res {
		if r < min || r >= max {
			ok = false
			break
		}
	}

	fmt.Println(ok)
	fmt.Println(len(res))
	isPanic(func() { randomjez.RandomIntSlice(min, max, 0) })

	// Output:
	// true
	// 10
	// true
}
```

### RandomIntSliceUnique
随机生成不重复的整数切片

- 注意事项：
  - 当 min > max 时，会交换 min 和 max 的值
  - 当 n < 1 时，会 panic
  - 当 n 大于 max-min+1 时，会 panic

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {
	min := 10
	max := 15

	res := randomjez.RandomIntSliceUnique(min, max, 5)

	var (
		ok    = true
		exist = make(map[int]struct{}, 5)
	)

	for _, r := range res {
		if r < min || r >= max {
			ok = false
			break
		}
		if _, ok1 := exist[r]; ok1 {
			ok = false
			break
		}
		exist[r] = struct{}{}
	}

	fmt.Println(ok)
	fmt.Println(len(res))

	isPanic(func() { randomjez.RandomIntSliceUnique(min, max, 0) })
	isPanic(func() { randomjez.RandomIntSliceUnique(min, max, 10) })

	// Output:
	// true
	// 5
	// true
	// true
}
```

### RandomBytes
随机生成字节切片

- 注意事项：
  - 当 n < 1 时，会 panic

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {

	res := randomjez.RandomBytes(10)

	fmt.Println(len(res))
	isPanic(func() { randomjez.RandomBytes(0) })

	// Output:
	// 10
	// true
}
```

### UUIDv4
根据 RFC4122 生成 UUID v4版本

- 注意事项：
  - 这种方法生成的 UUID v4 可能不是完全符合标准，但在大多数情况下应该是足够的。

```go
package main

import (
	"fmt"
	"regexp"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {
	uuid := randomjez.UUIDv4()

	reg := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)

	fmt.Println(len(uuid))
	fmt.Println(reg.MatchString(uuid))

	// Output:
	// 36
	// true
}
```

### Shuffle
打乱切片中元素的顺序。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}

	list2 := randomjez.Shuffle([]int{})

	fmt.Println(len(list) == len(randomjez.Shuffle(list)))
	fmt.Println(list2)

	// Output:
	// true
	// []
}
```

### Sample
从切片中随机返回一个元素。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {

  list := []int{1, 2, 3, 3, 5, 3, 5, 6}
  var list2 []string

  s := randomjez.Sample(list)

  ok := false
  for _, v := range list {
    if v == s {
      ok = true
      break
    }
  }

  fmt.Println(ok)
  fmt.Println(randomjez.Sample(list2) == "")

  // Output:
  // true
  // true
}
```

### Samples
从切片中随机返回n个元素，结果不去重。
- 注意事项：
  - 如果 n 大于切片的长度，则返回打乱顺序后的切片，如果 n 小于等于0，则返回空切片。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/randomjez"
)

func main() {

  list := []int{1, 2, 3}

  s2 := randomjez.Samples(list, 3)

  ok := true
  for _, v := range s2 {
    if v > 3 || v < 1 {
      ok = false
      break
    }
  }

  fmt.Println(ok)

  fmt.Println(randomjez.Samples([]string{}, 3))

  // Output:
  // true
  // []
}
```
