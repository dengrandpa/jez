package mapjez

import (
	"fmt"
	"sort"
	"strconv"
)

func ExampleForEach() {

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

	ForEach(m, func(key string, value int) {
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

func ExampleFilter() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	fmt.Println(Filter(m, func(_ string, value int) bool {
		return value%2 == 0
	}))

	// Output:
	// map[b:2 d:4]

}

func ExampleKeys() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	k := Keys(m)
	sort.Strings(k)

	fmt.Println(k)

	// Output:
	// [a b c d e f]
}

func ExampleKeysBy() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	kb := KeysBy(m, func(key string) string {
		return key + "1"
	})
	sort.Strings(kb)
	fmt.Println(kb)

	// Output:
	// [a1 b1 c1 d1 e1 f1]

}

func ExampleValues() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}
	v := Values(m)
	sort.Ints(v)

	fmt.Println(v)

	// Output:
	// [1 1 2 3 4 5]
}

func ExampleValuesBy() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	v := ValuesBy(m, func(value int) int {
		return value + 1
	})
	sort.Ints(v)

	fmt.Println(v)

	// Output:
	// [2 2 3 4 5 6]
}

func ExampleValuesUnique() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}
	v := ValuesUnique(m)
	sort.Ints(v)
	fmt.Println(v)

	// Output:
	// [1 2 3 4 5]
}

func ExampleKeysAndValues() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	keys, values := KeysAndValues(m)
	sort.Strings(keys)
	sort.Ints(values)

	fmt.Println(keys)
	fmt.Println(values)

	// Output:
	// [a b c d e f]
	// [1 1 2 3 4 5]
}

func ExampleKeysAndValuesFilter() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	keys, values := KeysAndValuesFilter(m, func(key string, value int) bool {
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

func ExampleMapToSliceBy() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	list := MapToSliceBy(m, func(key string, value int) string {
		return key + " " + strconv.Itoa(value)
	})
	sort.Strings(list)

	fmt.Println(list)

	// Output:
	// [a 1 b 2 c 3 d 4 e 5 f 1]
}

func ExampleMapToSliceFilter() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	list := MapToSliceFilter(m, func(key string, value int) (string, bool) {
		return key + " " + strconv.Itoa(value), value%2 == 0
	})
	sort.Strings(list)

	fmt.Println(list)

	// Output:
	// [b 2 d 4]
}

func ExampleDeletes() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}
	Deletes(m, "a", "b")

	fmt.Println(m)

	// Output:
	// map[c:3 d:4 e:5 f:1]
}

func ExampleDeleteByValues() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}
	DeleteByValues(m, 1, 2)

	fmt.Println(m)

	// Output:
	// map[c:3 d:4 e:5]
}

func ExampleDeleteFilter() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}
	DeleteFilter(m, func(key string, value int) bool {
		return value%2 == 0
	})

	fmt.Println(m)

	// Output:
	// map[a:1 c:3 e:5 f:1]
}

func ExampleReplaceValue() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
		"g": 1,
	}
	ReplaceValue(m, 1, 222)

	fmt.Println(m)

	// Output:
	// map[a:222 b:2 c:3 d:4 e:5 f:222 g:222]
}
