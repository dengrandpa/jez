package maputil

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

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

	ass.Equal([]string{"a", "b", "c", "d", "e", "f"}, keys)
	ass.Equal([]int{1, 1, 2, 3, 4, 5}, values)
}

func TestFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}

	m2 := map[string]int{
		"b": 2,
		"d": 4,
	}

	ass.Equal(m2, Filter(m, func(_ string, value int) bool {
		return value%2 == 0
	}))

}

func TestKeys(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
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

	ass.Equal([]string{"a", "b", "c", "d", "e", "f"}, k)
}

func TestKeysBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
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

	ass.Equal([]string{"a1", "b1", "c1", "d1", "e1", "f1"}, kb)
}

func TestValues(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
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

	ass.Equal([]int{1, 1, 2, 3, 4, 5}, v)
}

func TestValuesBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
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

	ass.Equal([]int{2, 2, 3, 4, 5, 6}, v)
}

func TestValuesUnique(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
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

	ass.Equal([]int{1, 2, 3, 4, 5}, v)
}

func TestKeysAndValues(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
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

	ass.Equal([]string{"a", "b", "c", "d", "e", "f"}, keys)
	ass.Equal([]int{1, 1, 2, 3, 4, 5}, values)
}

func TestKeysAndValuesFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
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

	ass.Equal([]string{"b", "d"}, keys)
	ass.Equal([]int{2, 4}, values)
}

func TestMapToSliceBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
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

	ass.Equal([]string{"a 1", "b 2", "c 3", "d 4", "e 5", "f 1"}, list)
}

func TestMapToSliceFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
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

	ass.Equal([]string{"b 2", "d 4"}, list)
}

func TestDeletes(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}
	Deletes(m, "a", "b")
	ass.Equal(map[string]int{
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}, m)
}

func TestDeleteByValues(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 1,
	}
	DeleteByValues(m, 1, 2)
	ass.Equal(map[string]int{
		"c": 3,
		"d": 4,
		"e": 5,
	}, m)
}

func TestDeleteFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
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
	ass.Equal(map[string]int{
		"a": 1,
		"c": 3,
		"e": 5,
		"f": 1,
	}, m)
}

func TestReplaceValue(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
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
	ass.Equal(map[string]int{
		"a": 222,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 222,
		"g": 222,
	}, m)
}
