# Slicejez

------

提供了一些对切片的操作，包括遍历、映射、过滤、去重、求差集、求交集等。

------

## 用法

```go
import "github.com/dengrandpa/jez/slicejez"
```

------

## 目录

-   [ForEach](#forEach)
-   [ForEachWithBreak](#forEachWithBreak)
-   [Filter](#filter)
-   [Map](#map)
-   [Contain](#contain)
-   [ContainAll](#containAll)
-   [FilterMap](#filterMap)
-   [AppendIfNotDuplicate](#appendIfNotDuplicate)
-   [AppendMultipleIfNotDuplicate](#appendMultipleIfNotDuplicate)
-   [Remove](#remove)
-   [RemoveFilter](#removeFilter)
-   [Unique](#unique)
-   [UniqueBy](#uniqueBy)
-   [UniqueNonzero](#uniqueNonzero)
-   [UniqueNonzeroBy](#uniqueNonzeroBy)
-   [Nonzero](#nonzero)
-   [Replace](#replace)
-   [ReplaceAll](#replaceAll)
-   [Difference](#difference)
-   [DifferenceUnique](#differenceUnique)
-   [Intersection](#intersection)
-   [MutualDifference](#mutualDifference)
-   [ToMapBy](#toMapBy)
-   [Repeat](#repeat)
-   [Equal](#equal)
-   [EqualElement](#equalElement)
-   [FindIndex](#findIndex)
-   [FindIndexFilter](#findIndexFilter)
-   [FindDuplicates](#findDuplicates)
-   [FindUniqueDuplicates](#findUniqueDuplicates)
-   [Min](#min)
-   [Max](#max)
-   [Drop](#drop)
-   [DropLast](#dropLast)
-   [Slice](#slice)
-   [IsSorted](#isSorted)
-   [IsSortedBy](#isSortedBy)
-   [Reverse](#reverse)
-   [Flatten](#flatten)
-   [InsertAt](#insertAt)
-   [NewSafeSlice](#newSafeSlice)
-   [SafeSlice_ForEach](#safeSliceForEach)
-   [SafeSlice_ForEachWithBreak](#safeSliceForEachWithBreak)
-   [SafeSlice_Filter](#safeSliceFilter)
-   [SafeSlice_Append](#safeSliceAppend)
-   [SafeSlice_AppendIfNotDuplicate](#safeSliceAppendIfNotDuplicate)
-   [SafeSlice_AppendMultipleIfNotDuplicate](#safeSliceAppendMultipleIfNotDuplicate)
-   [SafeSlice_Load](#safeSliceLoad)
-   [SafeSlice_LoadByIndex](#safeSliceLoadByIndex)
-   [SafeSlice_Index](#safeSliceIndex)
-   [SafeSlice_Insert](#safeSliceInsert)
-   [SafeSlice_Len](#safeSliceLen)
-   [SafeSlice_Remove](#safeSliceRemove)
-   [SafeSlice_RemoveByIndex](#safeSliceRemoveByIndex)
-   [SafeSlice_Replace](#safeSliceReplace)
-   [SafeSlice_ReplaceByIndex](#safeSliceReplaceByIndex)
-   [SafeSlice_Slice](#safeSliceSlice)

------

### ForEach
遍历切片并为每个元素调用 iteratee 函数。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"a", "b", "c"}

	items := make([]string, 0, len(list))
	indexes := make([]int, 0, len(list))

    slicejez.ForEach(list, func(i int, item string) {
		items = append(items, item)
		indexes = append(indexes, i)
	})

	fmt.Println(items)
	fmt.Println(indexes)
	// Output:
	// [a b c]
	// [0 1 2]
}
```

### ForEachWithBreak
遍历切片并为每个元素调用 iteratee 函数，如果返回 false，则停止遍历。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"a", "b", "c"}

	items := make([]string, 0, len(list))
	indexes := make([]int, 0, len(list))

    slicejez.ForEachWithBreak(list, func(index int, item string) bool {
		if index >= 2 {
			return false
		}

		items = append(items, item)
		indexes = append(indexes, index)
		return true
	})

	fmt.Println(items)
	fmt.Println(indexes)

	// Output:
	// [a b]
	// [0 1]
}
```

### Filter
遍历切片并为每个元素调用 iteratee 函数，只返回调用结果为true的元素。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"a", "b", "c"}

	list2 := slicejez.Filter(list, func(index int, item string) bool {
		return index != 1
	})
	fmt.Println(list2)

	// Output:
	// [a c]
}
```

### Map
遍历切片并为每个元素调用 iteratee 函数，并返回调用后结果。

```go
package main

import (
	"fmt"
	"strconv"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"1", "2", "3"}

	list2 := slicejez.Map(list, func(index int, item string) int {
		i, _ := strconv.Atoi(item)
		return i
	})

	fmt.Println(list2)

	// Output:
	// [1 2 3]
}
```

### Contain
效验切片是否包含目标元素。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"1", "2", "3"}

	fmt.Println(slicejez.Contain(list, "1"))
	fmt.Println(slicejez.Contain(list, "4"))

	// Output:
	// true
	// false

}
```

### ContainAll
效验切片是否包含所有的目标元素。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"1", "2", "3"}

	fmt.Println(slicejez.ContainAll(list, "1", "2", "3"))
	fmt.Println(slicejez.ContainAll(list, "1", "2", "2"))
	fmt.Println(slicejez.ContainAll(list, "1", "2"))
	fmt.Println(slicejez.ContainAll(list, "1", "4"))

	// Output:
	// true
	// true
	// true
	// false

}
```

### FilterMap
遍历切片并为每个元素调用 iteratee 函数，如果调用结果为true，则返回调用后元素。

```go
package main

import (
	"fmt"
	"strconv"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"1", "2", "3"}

	list2 := slicejez.FilterMap(list, func(index int, item string) (int, bool) {
		i, _ := strconv.Atoi(item)
		if i == 2 {
			return i, true
		}
		return 0, false
	})

	fmt.Println(list2)

	// Output:
	// [2]
}
```

### AppendIfNotDuplicate
添加元素到切片，如果元素已经存在，则不添加。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"1", "2", "3"}

	fmt.Println(slicejez.AppendIfNotDuplicate(list, "1"))
	fmt.Println(slicejez.AppendIfNotDuplicate(list, "4"))

	// Output:
	// [1 2 3]
	// [1 2 3 4]
}
```

### AppendMultipleIfNotDuplicate
添加多个元素到切片，如果元素已经存在，则不添加。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"1", "2", "3"}

	fmt.Println(slicejez.AppendMultipleIfNotDuplicate(list, "1"))
	fmt.Println(slicejez.AppendMultipleIfNotDuplicate(list, "3", "4"))
	fmt.Println(slicejez.AppendMultipleIfNotDuplicate(list, "4", "5"))

	// Output:
	// [1 2 3]
	// [1 2 3 4]
	// [1 2 3 4 5]
}
```

### Remove
从切片中删除元素。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"1", "2", "3"}

	fmt.Println(slicejez.Remove(list, "1"))
	fmt.Println(slicejez.Remove(list, "1", "2"))

	// Output:
	// [2 3]
	// [3]
}
```

### RemoveFilter
遍历切片并为每个元素调用 iteratee 函数，如果调用结果为true，则删除该元素。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"1", "2", "3"}

	fmt.Println(slicejez.RemoveFilter(list, func(index int, item string) bool {
		if item == "1" || item == "2" {
			return true
		}

		return false
	}))

	// Output:
	// [3]
}
```

### Unique
去重。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"1", "2", "3"}

	fmt.Println(slicejez.Unique(list))
	fmt.Println(slicejez.Unique([]string{"1", "2", "3", "3", "2"}))

	// Output:
	// [1 2 3]
	// [1 2 3]
}
```

### UniqueBy
遍历切片并为每个元素调用 iteratee 函数，返回唯一的元素。

```go
package main

import (
	"fmt"
	"strconv"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(slicejez.UniqueBy(list, func(index int, item int) string {
		if item > 3 {
			return "3"
		}
		return strconv.Itoa(item)
	}))

	// Output:
	// [0 1 2 3]
}
```

### UniqueNonzero
删除重复元素及零值。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []int{0, 1, 0, 1, 2, 2, 2, 3}

	fmt.Println(slicejez.UniqueNonzero(list))

	// Output:
	// [1 2 3]
}
```

### UniqueNonzeroBy
遍历切片并为每个元素调用 iteratee 函数，返回唯一的、非零值的元素。

```go
package main

import (
	"fmt"
	"strconv"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []int{0, 1, 0, 1, 2, 2, 2, 3}

	fmt.Println(slicejez.UniqueNonzeroBy(list, func(index int, item int) string {
		return strconv.Itoa(item)
	}))

	// Output:
	// [1 2 3]
}
```

### Nonzero
删除零值。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []int{0, 1, 1, 2, 2, 2, 3, 0}

	fmt.Println(slicejez.Nonzero(list))

	// Output:
	// [1 1 2 2 2 3]
}
```

### Replace
将切片中的元素 old 替换为 new ，最多替换 n 次，如果 n 为-1，则替换所有的 old 元素。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []int{0, 1, 1, 2, 2, 2, 3, 0}

	fmt.Println(slicejez.Replace(list, 1, 6, 1))
	fmt.Println(slicejez.Replace(list, 2, 6, -1))

	// Output:
	// [0 6 1 2 2 2 3 0]
	// [0 1 1 6 6 6 3 0]
}
```

### ReplaceAll
将切片中的元素 old 替换为 new ，替换所有的 old 元素。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []int{0, 1, 1, 2, 2, 2, 3, 0}

	fmt.Println(slicejez.ReplaceAll(list, 2, 6))

	// Output:
	// [0 1 1 6 6 6 3 0]
}
```

### Difference
差集，结果不去重。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list1 := []int{0, 1, 1, 2, 2, 3, 3, 0}

	list2 := []int{0, 1, 1, 2, 2, 2, 6, 6}

	fmt.Println(slicejez.Difference(list1, list2))

	// Output:
	// [3 3 6 6]
}
```

### DifferenceUnique
差集，结果去重。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list1 := []int{0, 1, 1, 2, 2, 3, 3, 0}

	list2 := []int{0, 1, 1, 2, 2, 2, 6, 6}

	fmt.Println(slicejez.DifferenceUnique(list1, list2))

	// Output:
	// [3 6]
}
```

### Intersection
交集，结果元素唯一。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list1 := []int{0, 1, 1, 2, 2, 3, 3, 0}

	list2 := []int{0, 1, 1, 2, 2, 2, 6, 6}

	fmt.Println(slicejez.Intersection(list1, list2))

	list3 := []int{0, 1, 2, 3}
	list4 := []int{2, 3, 4, 5, 6}

	fmt.Println(slicejez.Intersection(list3, list4))

	// Output:
	// [0 1 2]
	// [2 3]
}
```

### MutualDifference
差异，结果不去重。
- 返回值：
  - list1 中存在， list2 中不存在。
  - list2 中存在， list1 中不存在。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list1 := []int{0, 1, 1, 2, 2, 2, 3, 0}
	list2 := []int{0, 1, 1, 2, 2, 2, 6, 6}
	fmt.Println(slicejez.MutualDifference(list1, list2))

	list3 := []int{0, 1, 2, 3}
	list4 := []int{1, 2, 7, 8}
	fmt.Println(slicejez.MutualDifference(list3, list4))

	var list5 []int
	list6 := []int{1, 2, 3}
	fmt.Println(slicejez.MutualDifference(list5, list6))

	list7 := []int{1, 2, 3}
	var list8 []int
	fmt.Println(slicejez.MutualDifference(list7, list8))

	// Output:
	// [3] [6 6]
	// [0 3] [7 8]
	// [] [1 2 3]
	// [1 2 3] []
}
```

### ToMapBy
遍历切片，将切片中的元素转换为map的key和value。

```go
package main

import (
	"fmt"
	"strconv"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []int{1, 2, 3, 4, 5, 5, 5}

	m := slicejez.ToMapBy(list, func(index int, item int) (string, string) {
		itemStr := strconv.Itoa(item)
		return itemStr + "_" + strconv.Itoa(index), itemStr
	})

	fmt.Println(m)

	// Output:
	// map[1_0:1 2_1:2 3_2:3 4_3:4 5_4:5 5_5:5 5_6:5]
}
```

### Repeat
返回包含 n 个 item 的切片。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	fmt.Println(slicejez.Repeat(1, 5))
	fmt.Println(slicejez.Repeat("1", 5))

	// Output:
	// [1 1 1 1 1]
	// [1 1 1 1 1]
}
```

### Equal
长度、顺序、值都相等时返回 true 。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list1 := []int{1, 2, 3, 3, 3}
	list2 := []int{1, 2, 3, 3, 3}
	list3 := []int{3, 2, 3, 3, 1}
	list4 := []int{1, 2}

	fmt.Println(slicejez.Equal(list1, list2))
	fmt.Println(slicejez.Equal(list1, list3))
	fmt.Println(slicejez.Equal(list1, list4))

	// Output:
	// true
	// false
	// false
}
```

### EqualElement
长度、值相等时返回 true ，不考虑顺序。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list1 := []int{1, 2, 3, 3, 3}
	list2 := []int{1, 2, 3, 3, 3}
	list3 := []int{3, 2, 3, 3, 1}
	list4 := []int{1, 2}
	list5 := []int{3, 2, 1}

	fmt.Println(slicejez.EqualElement(list1, list2))
	fmt.Println(slicejez.EqualElement(list1, list3))
	fmt.Println(slicejez.EqualElement(list1, list4))
	fmt.Println(slicejez.EqualElement(list1, list5))

	// Output:
	// true
	// true
	// false
	// false
}
```

### FindIndex
返回第一个匹配的元素的索引，不存在则返回 -1 。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []int{1, 2, 3, 3, 5, 5, 6}

	fmt.Println(slicejez.FindIndex(list, 3))
	fmt.Println(slicejez.FindIndex(list, 5))
	fmt.Println(slicejez.FindIndex(list, 2))
	fmt.Println(slicejez.FindIndex(list, 9))

	// Output:
	// 2
	// 4
	// 1
	// -1
}
```

### FindIndexFilter
返回调用 iteratee 函数返回 true 的第一个元素的索引，不存在则返回 -1 。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []int{1, 2, 3, 5, 3, 5, 6}

	n := 0

	fmt.Println(slicejez.FindIndexFilter(list, func(index int, item int) bool {
		if item == 3 {
			n++
			if n > 1 {
				return true
			}
		}
		return false
	}))

	fmt.Println(slicejez.FindIndexFilter(list, func(index int, item int) bool {
		return item > 3
	}))

	// Output:
	// 4
	// 3
}
```

### FindDuplicates
返回切片中所有重复的元素，结果不去重。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}

	fmt.Println(slicejez.FindDuplicates(list))

	// Output:
	// [3 3 5]

}
```

### FindUniqueDuplicates
返回切片中所有重复的元素，结果去重。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}

	fmt.Println(slicejez.FindUniqueDuplicates(list))

	// Output:
	// [3 5]
}
```

### Min
返回最小值

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"1", "2", "3"}
	fmt.Println(slicejez.Min(list))

	list2 := []int{5, 2, 111, 3}
	fmt.Println(slicejez.Min(list2))

	// Output:
	// 1
	// 2
}
```

### Max
返回最大值

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []string{"1", "2", "3"}
	fmt.Println(slicejez.Max(list))

	list2 := []int{5, 2, 111, 3}
	fmt.Println(slicejez.Max(list2))

	// Output:
	// 3
	// 111

}
```

### Drop
返回从开头删除n个元素的切片，如果 n 大于切片的长度，则返回空切片。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	s := []string{"1", "2", "3", "4"}

	fmt.Println(slicejez.Drop(s, 2))
	fmt.Println(slicejez.Drop(s, 1111))

	// Output:
	// [3 4]
	// []
}
```

### DropLast
返回从末尾删除n个元素的切片，如果 n 大于切片的长度，则返回空切片。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	s := []string{"1", "2", "3", "4"}

	fmt.Println(slicejez.DropLast(s, 2))
	fmt.Println(slicejez.DropLast(s, 1111))

	// Output:
	// [1 2]
	// []
}
```

### Slice
返回索引从 n 到 m 的切片，但不包括 m，等同于 slice[n:m]，即[min,max)，但不会在溢出时panic。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	s := []string{"1", "2", "3", "4"}

	fmt.Println(slicejez.Slice(s, 1, 3))
	fmt.Println(slicejez.Slice(s, 1, 4))
	fmt.Println(slicejez.Slice(s, 5, 10))
	fmt.Println(slicejez.Slice(s, 10, 5))
	fmt.Println(slicejez.Slice(s, 1, 6))

	// Output:
	// [2 3]
	// [2 3 4]
	// []
	// []
	// [2 3 4]
}
```

### IsSorted
判断切片是否已排序。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	fmt.Println(slicejez.IsSorted([]int{0, 1, 2, 3}))
	fmt.Println(slicejez.IsSorted([]string{"a", "b", "c", "d"}))
	fmt.Println(slicejez.IsSorted([]int{0, 1, 4, 3}))
	fmt.Println(slicejez.IsSorted([]string{"a", "e", "d", "c"}))

	// Output:
	// true
	// true
	// false
	// false

}
```

### IsSortedBy
遍历切片并为每个元素调用 iteratee 函数，以确定它是否已排序。

```go
package main

import (
	"fmt"
	"strconv"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	fmt.Println(slicejez.IsSortedBy([]string{"a", "bb", "ccc"}, func(s string) int {
		return len(s)
	}))

	fmt.Println(slicejez.IsSortedBy([]string{"1", "2", "3", "11"}, func(s string) int {
		ret, _ := strconv.Atoi(s)
		return ret
	}))

	fmt.Println(slicejez.IsSortedBy([]string{"aa", "b", "ccc"}, func(s string) int {
		return len(s)
	}))

	// Output:
	// true
	// true
	// false
}
```

### Reverse
将切片中的元素顺序反转。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {

	list := []int{1, 2, 3, 4, 5}
    slicejez.Reverse(list)

	fmt.Println(list)

	// Output:
	// [5 4 3 2 1]
}
```

### Flatten
将二维切片转换为一维切片。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
	list := [][]int{
		{1, 2, 3},
		{6, 7, 8},
	}
	list1 := slicejez.Flatten(list)

	fmt.Println(list1)

	// Output:
	// [1 2 3 6 7 8]
}
```

### InsertAt
在切片的指定索引处插入值，如果索引大于切片的长度或小于 0，则将值附加到切片的末尾。

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
	list := []int{1, 2, 3, 4, 5}

	list1 := slicejez.InsertAt(list, 2, 666, 777, 888)

	fmt.Println(list1)

	// Output:
	// [1 2 666 777 888 3 4 5]
}
```

### NewSafeSlice
创建一个并发安全的切片。

```go
package main

import (
    "fmt"
  
    "github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{1, 2})
    fmt.Println(ss.Len())
  
    // Output:
    // 2
}
```

### SafeSlice_ForEach
遍历切片并为每个元素调用 iteratee 函数。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }
  
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        ss.ForEach(func(index int, item int) {
          if item == 100 {
            fmt.Println(item)
          }
        })
      }(i)
    }
  
    wg.Wait()
  
    // Output:
    // 100
    // 100
    // 100
    // 100
    // 100
}
```

### SafeSlice_ForEachWithBreak
遍历切片并为每个元素调用 iteratee 函数，如果返回 false，则停止遍历。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }
  
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        ss.ForEachWithBreak(func(index int, item int) bool {
          ok := item < 500
          if ok {
            if item == 100 || item == 600 {
              fmt.Println(item)
            }
          }
          return ok
        })
      }(i)
    }
  
    // Output:
    // 100
    // 100
    // 100
    // 100
    // 100
}

```

### SafeSlice_Filter
遍历切片并为每个元素调用 iteratee 函数，只返回调用结果为true的元素。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }
  
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        res := ss.Filter(func(index int, item int) bool {
          return item == 100
        })
        fmt.Println(res)
      }(i)
    }
  
    // Output:
    // 100
    // 100
    // 100
    // 100
    // 100
}

```

### SafeSlice_Append
添加元素到切片。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        ss.Append(v)
      }(i)
    }
    wg.Wait()
	
    fmt.Println(ss.Len())
  
    // Output:
    // 1000
}

```

### SafeSlice_AppendIfNotDuplicate
添加元素到切片，如果元素已经存在，则不添加。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        ss.AppendIfNotDuplicate(v)
      }(i)
    }
  
    wg.Wait()
  
    actual := slicejez.ToMapBy(ss.Load(), func(index int, item int) (int, struct{}) {
      return item, struct{}{}
    })
	
    fmt.Println(ss.Len())
    fmt.Println(len(actual))
	
    // Output:
    // 1000
    // 1000
}

```

### SafeSlice_AppendMultipleIfNotDuplicate
添加多个元素到切片，如果元素已经存在，则不添加。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        ss.AppendMultipleIfNotDuplicate(v, v+1)
      }(i)
    }
  
    wg.Wait()
  
    actual := slicejez.ToMapBy(ss.Load(), func(index int, item int) (int, struct{}) {
      return item, struct{}{}
    })
	
    fmt.Println(ss.Len())
    fmt.Println(len(actual))
	
    // Output:
    // 1001
    // 1001
}

```

### SafeSlice_Load
返回切片的副本。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        ss.AppendMultipleIfNotDuplicate(v, v+1)
      }(i)
    }
  
    wg.Wait()
  
    actual := slicejez.ToMapBy(ss.Load(), func(index int, item int) (int, struct{}) {
      return item, struct{}{}
    })
	
    fmt.Println(ss.Len())
    fmt.Println(len(actual))
	
    // Output:
    // 1001
    // 1001
}

```

### SafeSlice_LoadByIndex
返回指定索引位置的元素，-1 则返回最后一个元素，如果索引超出范围，panic。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        fmt.Println(ss.LoadByIndex(0))
      }(i)
    }
	
    // Output:
    // 0
    // 0
}

```

### SafeSlice_Index
返回指定元素在切片中的索引位置。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        fmt.Println(ss.Index(1))
      }(i)
    }
	
    // Output:
    // 1
    // 1
}

```

### SafeSlice_Insert
在指定索引位置插入元素。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
  
    for i := 0; i < 1000; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        ss.Insert(1, v)
      }(i)
    }
  
    wg.Wait()
  
    fmt.Println(ss.Load())
    fmt.Println(ss.LoadByIndex(0))
    fmt.Println(ss.LoadByIndex(-1))
	
    // Output:
    // 1002
    // 0
    // 1
}

```

### SafeSlice_Len
返回切片的长度。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        fmt.Println(ss.Len())
      }(i)
    }
	
    // Output:
    // 1000
    // 1000
    // 1000
    // 1000
    // 1000
}

```

### SafeSlice_Remove
从切片中移除元素。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
    for i := 0; i < 500; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        ss.Remove(v, v+500)
      }(i)
    }
  
    wg.Wait()
  
	fmt.Println(ss.Len())
	
    // Output:
    // 0
}

```

### SafeSlice_RemoveByIndex
从切片中移除指定索引位置的元素。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
	
    for i := 0; i < 1000; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        ss.RemoveByIndex(0)
      }(i)
    }

    wg.Wait()
  
    fmt.Println(ss.Len())
	
    // Output:
    // 0
}

```

### SafeSlice_Replace
将切片中的元素 old 替换为 new ，最多替换 n 次，如果 n 为-1，则替换所有的 old 元素。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        ss.Replace(v, 1000+v, -1)
      }(i)
    }
    wg.Wait()

    for i := 0; i < 5; i++ {
        fmt.Println(ss.LoadByIndex(i))
    }
	
    // Output:
    // 1000
    // 1001
    // 1002
    // 1003
    // 1004
}

```

### SafeSlice_ReplaceByIndex
将指定索引位置的元素替换为 new 。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        ss.ReplaceByIndex(v, 999-v)
      }(i)
    }
  
    wg.Wait()
	
	for i := 0; i < 5; i++ {
        fmt.Println(ss.LoadByIndex(i))
    }
	
    // Output:
    // 999
    // 998
    // 997
    // 996
    // 995
}

```

### SafeSlice_Slice
返回索引从 n 到 m 的切片，但不包括 m，等同于 slice[n:m]，即[min,max)，但不会在溢出时panic。

```go
package main

import (
	"fmt"
	"sync"
	
	"github.com/dengrandpa/jez/slicejez"
)

func main() {
    ss := slicejez.NewSafeSlice([]int{})
    for i := 0; i < 1000; i++ {
      ss.Append(i)
    }

    var wg sync.WaitGroup
  
    for i := 0; i < 5; i++ {
      wg.Add(1)
      go func(v int) {
        defer wg.Done()
        actual := ss.Slice(500, 501)
        fmt.Println(actual)
      }(i)
    }
	
    // Output:
    // 500
    // 500
    // 500
    // 500
    // 500
}

```
