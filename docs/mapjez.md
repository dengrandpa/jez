# Mapjez

------

提供了一些常用的 map 操作函数。

------

## 用法

```go
import "github.com/dengrandpa/jez/mapjez"
```

------

## 目录

-   [ForEach](#forEach)
-   [Filter](#filter)
-   [Keys](#keys)
-   [KeysBy](#keysBy)
-   [Values](#values)
-   [ValuesBy](#valuesBy)
-   [ValuesUnique](#valuesUnique)
-   [KeysAndValues](#keysAndValues)
-   [KeysAndValuesFilter](#keysAndValuesFilter)
-   [Deletes](#deletes)
-   [DeleteByValues](#deleteByValues)
-   [DeleteFilter](#deleteFilter)
-   [ReplaceValue](#replaceValue)
-   [MapToSliceBy](#mapToSliceBy)
-   [MapToSliceFilter](#mapToSliceFilter)

------

### ForEach
遍历map，对每个元素调用 iteratee 函数。

```go
package main

import (
	"fmt"
	"sort"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	keys := make([]string, 0, len(m))
	values := make([]int, 0, len(m))

	mapjez.ForEach(m, func(key string, value int) {
		keys = append(keys, key)
		values = append(values, value)
	})

	sort.Strings(keys)
	sort.Ints(values)

	fmt.Println(keys)
	fmt.Println(values)

	// Output:
	// [a b c d e f]
	// [1 1 2 3 4 5]
}

`````````

### Filter
遍历map，对每个元素调用 iteratee 函数，如果 iteratee 返回 true，则将该元素添加到结果map中。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	fmt.Println(mapjez.Filter(m, func(_ string, value int) bool {
		return value%2 == 0
	}))

	// Output:
	// map[b:2 d:4]

}

```

### Keys
遍历map，将每个key添加到结果slice中。

```go
package main

import (
	"fmt"
	"sort"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	k := mapjez.Keys(m)
	sort.Strings(k)

	fmt.Println(k)

	// Output:
	// [a b c d e f]
}

```

### KeysBy
遍历map，对每个元素调用 iteratee 函数，并返回调用后结果。

```go
package main

import (
	"fmt"
	"sort"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	kb := mapjez.KeysBy(m, func(key string) string {
		return key + "1"
	})
	sort.Strings(kb)
	fmt.Println(kb)

	// Output:
	// [a1 b1 c1 d1 e1 f1]

}

```

### Values
返回map中所有的value。

```go
package main

import (
	"fmt"
	"sort"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}
	v := mapjez.Values(m)
	sort.Ints(v)

	fmt.Println(v)

	// Output:
	// [1 1 2 3 4 5]
}

```

### ValuesBy
遍历map，对每个元素调用 iteratee 函数，并返回调用后结果。

```go
package main

import (
	"fmt"
	"sort"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	v := mapjez.ValuesBy(m, func(value int) int {
		return value + 1
	})
	sort.Ints(v)

	fmt.Println(v)

	// Output:
	// [2 2 3 4 5 6]
}

```

### ValuesUnique
返回map中所有的value，结果去重。

```go
package main

import (
	"fmt"
	"sort"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}
	v := mapjez.ValuesUnique(m)
	sort.Ints(v)
	fmt.Println(v)

	// Output:
	// [1 2 3 4 5]
}

```

### KeysAndValues
返回map中所有的key和value。

```go
package main

import (
	"fmt"
	"sort"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	keys, values := mapjez.KeysAndValues(m)
	sort.Strings(keys)
	sort.Ints(values)

	fmt.Println(keys)
	fmt.Println(values)

	// Output:
	// [a b c d e f]
	// [1 1 2 3 4 5]
}

```

### KeysAndValuesFilter
遍历map，对每个元素调用 iteratee 函数，如果 iteratee 返回true，则将该元素添加到结果slice中。

```go
package main

import (
	"fmt"
	"sort"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	keys, values := mapjez.KeysAndValuesFilter(m, func(key string, value int) bool {
		return value%2 == 0
	})
	sort.Strings(keys)
	sort.Ints(values)

	fmt.Println(keys)
	fmt.Println(values)

	// Output:
	// [b d]
	// [2 4]
}

```

### Deletes
通过key删除多个元素。

```go
package main

import (
	"fmt"
	"sort"
	"strconv"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	list := mapjez.MapToSliceBy(m, func(key string, value int) string {
		return key + " " + strconv.Itoa(value)
	})
	sort.Strings(list)

	fmt.Println(list)

	// Output:
	// [a 1 b 2 c 3 d 4 e 5 f 1]
}

```

### DeleteByValues
通过value删除多个元素。

```go
package main

import (
	"fmt"
	"sort"
	"strconv"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	list := mapjez.MapToSliceFilter(m, func(key string, value int) (string, bool) {
		return key + " " + strconv.Itoa(value), value%2 == 0
	})
	sort.Strings(list)

	fmt.Println(list)

	// Output:
	// [b 2 d 4]
}

```

### DeleteFilter
遍历map，对每个元素调用 iteratee 函数，如果 iteratee 返回true，则删除该元素。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}
	mapjez.Deletes(m, "a", "b")

	fmt.Println(m)

	// Output:
	// map[c:3 d:4 e:5 f:1]
}

```

### ReplaceValue
替换所有value等于 old 的元素。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}
	mapjez.DeleteByValues(m, 1, 2)

	fmt.Println(m)

	// Output:
	// map[c:3 d:4 e:5]
}

```

### MapToSliceBy
map转切片，遍历map，对每个元素调用 iteratee 函数，并返回调用后结果切片。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}
	mapjez.DeleteFilter(m, func(key string, value int) bool {
		return value%2 == 0
	})

	fmt.Println(m)

	// Output:
	// map[a:1 c:3 e:5 f:1]
}

```

### MapToSliceFilter
map转切片，遍历map，对每个元素调用 iteratee 函数，如果 iteratee 返回true，则将该元素添加到结果切片中。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/mapjez"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
		"g": 1,
	}
	mapjez.ReplaceValue(m, 1, 222)

	fmt.Println(m)

	// Output:
	// map[a:222 b:2 c:3 d:4 e:5 f:222 g:222]
}

```
