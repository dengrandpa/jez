# Mapjez

------

Provides some common map operation functions.

------

## Usage

```go
import "github.com/dengrandpa/jez/mapjez"
```

------

## Index

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
Traverse the map and call the iteratee function for each element.

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
Traverse the map and call the iteratee function for each element. If iteratee returns true, the element is added to the result map.

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
Traverse the map and add each key to the result slice.

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
Traverse the map, call the iteratee function for each element, and return the result after the call.

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
Return all values in the map.

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
Traverse the map, call the iteratee function for each element, and return the result after the call.

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
Return all the values in the map, and the result will be degened.

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
Returns all keys and values in the map.

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
Traverse the map, call the iteratee function for each element, and if iteratee returns true, the element is added to the result slice.

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
Delete multiple elements by key.

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

### DeleteByValues
Delete multiple elements by value.

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

### DeleteFilter
Traverse the map, call the iteratee function for each element, and delete the element if iteratee returns true.

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

### ReplaceValue
Replace all elements whose value is equal to old.

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

### MapToSliceBy
The map is sliced, the map is traversed, the iteratee function is called for each element, and the result slice is returned after the call.

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

### MapToSliceFilter
The map is sliced, the map is traversed, and the iteratee function is called on each element. If iteratee returns true, the element is added to the result slice.

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
