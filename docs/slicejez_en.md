# Slicejez

------

Provides some operations on slices, including traversal, mapping, filtering, deduplication, difference set, intersection set, etc.

------

## Usage

```go
import "github.com/dengrandpa/jez/slicejez"
```

------

## Index

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
Traverse the slice and call the iteratee function for each element.

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
`````````

### ForEachWithBreak
Traverse the slice and call the iteratee function for each element, and if it returns false, stop the traversal.

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
Traverse the slice and call the iteratee function for each element, returning only the element with the call result of true.

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
Traverse the slice and call the iteratee function for each element, and return the result after the call.

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
Check whether the slice contains the target element.

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
Verify whether the slice contains all the target elements.

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
The iteratee function is called for each element by traversing the slice. If the call result is true, the called element is returned.

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
Add elements to slices, and if the elements already exist, they will not be added.

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
Add multiple elements to the slice. If the element already exists, it will not be added.

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
Delete elements from slices.

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
The iteratee function is called for each element by traversing the slice. If the call result is true, the element is deleted.

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
Remove duplicates.

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
Traverse the slice and call the iteratee function for each element, returning the unique element.

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
Delete duplicate elements and zero values.

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
Traverse the slice and call the iteratee function for each element, returning a unique, non-zero element.

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
Delete the zero value.

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
Replace the element old in the slice with new, up to n times, and if n is -1, replace all old elements.

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
Replace the element old in the slice with new, and replace all the old elements.

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
The difference, the result is not heavy.

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
The difference is set, and the result is to be repeated.

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
The intersection, the result element is unique.

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
The difference, the result is not serious.
- Return:
  - It exists in list1 and does not exist in list2.
  - It exists in list2 and does not exist in list1.

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
Traverse the slice and convert the elements in the slice to the key and value of the map.

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
Returns a slice containing n items.

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
Return true when the length, order and value are equal.

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
Return true when the length and value are equal, regardless of the order.

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
Returns the index of the first matched element, and -1 if it does not exist.

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
Returns the index of the first element of true by calling the iteratee function, and -1 if it does not exist.

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
Return all the duplicate elements in the slice, and the result will not be repeated.

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
Return all duplicate elements in the slice, and the result will be deduplicate.

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
Return the minimum value

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
Return the maximum value

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
Returns a slice that deletes n elements from the beginning, and returns an empty slice if n is greater than the length of the slice.

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
Returns a slice that deletes n elements from the end, and returns an empty slice if n is greater than the length of the slice.

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
Returns the slice of the index from n to m, but does not include m, which is equivalent to slice[n:m], that is, [min,max), but not panic when overflowing.

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
Determine whether the slices have been sorted.

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
Traverse the slice and call the iteratee function for each element to determine whether it has been sorted.

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
Invert the order of the elements in the slice.

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
Convert two-dimensional slices to one-dimensional slices.

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
Insert a value at the specified index of the slice. If the index is greater than or less than 0 in length of the slice, the value is attached to the end of the slice.

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
Traverse the slice and call the iterate function for each element.

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
Traverse the slice and call the iterate function for each element. If false is returned, stop traversing.

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
Traverse the slice and call the iterate function for each element, returning only the elements with a call result of true.

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
Add elements to the slice.

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
Add an element to the slice, if the element already exists, do not add it.

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
Add multiple elements to the slice, if the element already exists, do not add it.

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
Returns a copy of the slice.

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
Returns the element at the specified index position, -1 returns the last element, and if the index is out of range, panic.

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
Returns the index position of the specified element in the slice.

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
Insert an element at the specified index position.

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
Returns the length of the slice.

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
Remove elements from slices.

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
Remove the element at the specified index position from the slice.

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
Replace the old element in the slice with new, up to n times. If n is -1, replace all old elements.

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
Replace the element at the specified index position with new.

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
Returns a slice with an index from n to m, but does not include m, which is equivalent to slice [n: m], i.e. [min, max], but does not panic during overflow.

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
