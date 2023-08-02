package sliceutil

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"a", "b", "c"}

	items := make([]string, 0, len(list))
	indexes := make([]int, 0, len(list))

	ForEach(list, func(i int, item string) {
		items = append(items, item)
		indexes = append(indexes, i)
	})

	ass.Equal(list, items)
	ass.Equal([]int{0, 1, 2}, indexes)
}

func TestForEachWithBreak(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"a", "b", "c"}

	items := make([]string, 0, len(list))
	indexes := make([]int, 0, len(list))

	ForEachWithBreak(list, func(index int, item string) bool {
		if index >= 2 {
			return false
		}

		items = append(items, item)
		indexes = append(indexes, index)
		return true
	})

	ass.Equal([]string{"a", "b"}, items)
	ass.Equal([]int{0, 1}, indexes)

}

func TestFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"a", "b", "c"}

	list2 := Filter(list, func(index int, item string) bool {
		return index != 1
	})

	ass.Equal([]string{"a", "c"}, list2)
}

func TestMap(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"1", "2", "3"}

	list2 := Map(list, func(index int, item string) int {
		i, _ := strconv.Atoi(item)
		return i
	})

	ass.Equal([]int{1, 2, 3}, list2)
}

func TestContain(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"1", "2", "3"}

	ass.Equal(true, Contain(list, "1"))
	ass.Equal(false, Contain(list, "4"))

}

func TestContainAll(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"1", "2", "3"}

	ass.Equal(true, ContainAll(list, "1", "2", "3"))
	ass.Equal(true, ContainAll(list, "1", "2", "2"))
	ass.Equal(true, ContainAll(list, "1", "2"))
	ass.Equal(false, ContainAll(list, "1", "4"))
	ass.True(ContainAll(list))
	ass.False(ContainAll([]string{}, "1"))
}

func TestFilterMap(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"1", "2", "3"}

	list2 := FilterMap(list, func(index int, item string) (int, bool) {
		i, _ := strconv.Atoi(item)
		if i == 2 {
			return i, true
		}
		return 0, false
	})

	ass.Equal([]int{2}, list2)
}

func TestAppendIfNotDuplicate(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"1", "2", "3"}

	ass.Equal(list, AppendIfNotDuplicate(list, "1"))
	ass.Equal([]string{"1", "2", "3", "4"}, AppendIfNotDuplicate(list, "4"))
}

func TestAppendMultipleIfNotDuplicate(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"1", "2", "3"}

	ass.Equal(list, AppendMultipleIfNotDuplicate(list, "1"))
	ass.Equal([]string{"1", "2", "3", "4"}, AppendMultipleIfNotDuplicate(list, "3", "4"))
	ass.Equal([]string{"1", "2", "3", "4", "5"}, AppendMultipleIfNotDuplicate(list, "4", "5"))
	ass.Equal(list, AppendMultipleIfNotDuplicate(list))

}

func TestRemove(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"1", "2", "3"}

	ass.Equal([]string{"2", "3"}, Remove(list, "1"))
	ass.Equal([]string{"3"}, Remove(list, "1", "2"))
	ass.Equal(list, Remove(list))
	ass.Empty(Remove([]string{}, "1"))

}

func TestRemoveFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"1", "2", "3"}

	ass.Equal([]string{"3"}, RemoveFilter(list, func(index int, item string) bool {
		if item == "1" || item == "2" {
			return true
		}

		return false
	}))

	ass.Empty(RemoveFilter([]string{}, func(index int, item string) bool {
		return true
	}))
}

func TestUnique(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	list := []string{"1", "2", "3"}

	ass.Equal(list, Unique(list))

	ass.Equal(list, Unique([]string{"1", "2", "3", "3", "2"}))
}

func TestUniqueBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{0, 1, 2, 3, 4, 5, 6}

	ass.Equal([]string{"0", "1", "2", "3"}, UniqueBy(list, func(index int, item int) string {
		if item > 3 {
			return "3"
		}
		return strconv.Itoa(item)
	}))
}

func TestUniqueNonzero(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{0, 1, 0, 1, 2, 2, 2, 3}

	ass.Equal([]int{1, 2, 3}, UniqueNonzero(list))

}

func TestUniqueNonzeroBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{0, 1, 0, 1, 2, 2, 2, 3}

	ass.Equal([]string{"1", "2", "3"}, UniqueNonzeroBy(list, func(index int, item int) string {
		return strconv.Itoa(item)
	}))
}

func TestNonzero(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{0, 1, 1, 2, 2, 2, 3, 0}

	list2 := Nonzero(list)

	ass.Equal(Slice(list, 1, 7), list2)
}

func TestReplace(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{0, 1, 1, 2, 2, 2, 3, 0}

	ass.Equal([]int{0, 6, 1, 2, 2, 2, 3, 0}, Replace(list, 1, 6, 1))

	ass.Equal([]int{0, 1, 1, 6, 6, 6, 3, 0}, Replace(list, 2, 6, -1))

	ass.Empty(Replace([]string{}, "1", "2", -1))
	ass.Equal(list, Replace(list, 1, 2, 0))

}

func TestReplaceAll(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{0, 1, 1, 2, 2, 2, 3, 0}

	ass.Equal([]int{0, 1, 1, 6, 6, 6, 3, 0}, ReplaceAll(list, 2, 6))
}

func TestDifference(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list1 := []int{0, 1, 1, 2, 2, 3, 3, 0}

	list2 := []int{0, 1, 1, 2, 2, 2, 6, 6}

	ass.Equal([]int{3, 3, 6, 6}, Difference(list1, list2))

	ass.Equal(list1, Difference(list1, []int{}))
	ass.Equal(list1, Difference([]int{}, list1))
}

func TestDifferenceUnique(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list1 := []int{0, 1, 1, 2, 2, 3, 3, 0}

	list2 := []int{0, 1, 1, 2, 2, 2, 6, 6}

	list3 := []int{0, 1, 2}

	ass.Equal([]int{3, 6}, DifferenceUnique(list1, list2))

	ass.Equal(list3, DifferenceUnique(list3, []int{}))
	ass.Equal(list3, DifferenceUnique([]int{}, list3))
}

func TestIntersection(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list1 := []int{0, 1, 1, 2, 2, 3, 3, 0}

	list2 := []int{0, 1, 1, 2, 2, 2, 6, 6}

	ass.Equal([]int{0, 1, 2}, Intersection(list1, list2))

	list3 := []int{0, 1, 2, 3}
	list4 := []int{2, 3, 4, 5, 6}

	ass.Equal([]int{2, 3}, Intersection(list3, list4))

	ass.Empty(Intersection(list1, []int{}))
	ass.Empty(Intersection([]int{}, list1))
}

func TestMutualDifference(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
	list1 := []int{0, 1, 1, 2, 2, 2, 3, 0}

	list2 := []int{0, 1, 1, 2, 2, 2, 6, 6}

	add, del := MutualDifference(list1, list2)

	ass.Equal([]int{3}, add)
	ass.Equal([]int{6, 6}, del)

	list3 := []int{0, 1, 2, 3}
	list4 := []int{1, 2, 7, 8}

	add1, del1 := MutualDifference(list3, list4)

	ass.Equal([]int{0, 3}, add1)
	ass.Equal([]int{7, 8}, del1)

	var list5 []int
	list6 := []int{1, 2, 3}
	add2, del2 := MutualDifference(list5, list6)

	ass.Equal([]int{}, add2)
	ass.Equal(list6, del2)

	list7 := []int{1, 2, 3}
	var list8 []int
	add3, del3 := MutualDifference(list7, list8)
	ass.Equal(list6, add3)
	ass.Equal([]int{}, del3)
}

func TestToMapBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{1, 2, 3, 4, 5, 5, 5}

	m := ToMapBy(list, func(index int, item int) (string, string) {
		itemStr := strconv.Itoa(item)
		return itemStr + "_" + strconv.Itoa(index), itemStr
	})

	ass.Equal(map[string]string{
		"1_0": "1",
		"2_1": "2",
		"3_2": "3",
		"4_3": "4",
		"5_4": "5",
		"5_5": "5",
		"5_6": "5",
	}, m)
}

func TestRepeat(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ass.Equal([]int{1, 1, 1, 1, 1}, Repeat(1, 5))
	ass.Equal([]string{"1", "1", "1", "1", "1"}, Repeat("1", 5))
	ass.Empty(Repeat(1, 0))
}

func TestEqual(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list1 := []int{1, 2, 3, 3, 3}
	list2 := []int{1, 2, 3, 3, 3}
	list3 := []int{3, 2, 3, 3, 1}
	list4 := []int{1, 2}

	ass.Equal(true, Equal(list1, list2))
	ass.Equal(false, Equal(list1, list3))
	ass.Equal(false, Equal(list1, list4))
}

func TestEqualElement(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list1 := []int{1, 2, 3, 3, 3}
	list2 := []int{1, 2, 3, 3, 3}
	list3 := []int{3, 2, 3, 3, 1}
	list4 := []int{1, 2}
	list5 := []int{3, 2, 1}
	list6 := []int{3, 1, 1}

	ass.True(EqualElement(list1, list2))
	ass.True(EqualElement(list1, list3))
	ass.False(EqualElement(list1, list4))
	ass.False(EqualElement(list1, list5))
	ass.False(EqualElement(list5, list6))
}

func TestFindIndex(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{1, 2, 3, 3, 5, 5, 6}

	ass.Equal(2, FindIndex(list, 3))
	ass.Equal(4, FindIndex(list, 5))
	ass.Equal(1, FindIndex(list, 2))
	ass.Equal(-1, FindIndex(list, 9))

}

func TestFindIndexFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{1, 2, 3, 5, 3, 5, 6}

	n := 0
	ass.Equal(4, FindIndexFilter(list, func(index int, item int) bool {
		if item == 3 {
			n++
			if n > 1 {
				return true
			}
		}
		return false
	}))

	ass.Equal(3, FindIndexFilter(list, func(index int, item int) bool {
		return item > 3
	}))

	ass.Equal(-1, FindIndexFilter(list, func(index int, item int) bool {
		return item > 10
	}))
}

func TestFindDuplicates(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}

	ass.Equal([]int{3, 3, 5}, FindDuplicates(list))
}

func TestFindUniqueDuplicates(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}

	ass.Equal([]int{3, 5}, FindUniqueDuplicates(list))

}

func TestMin(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"1", "2", "3"}
	ass.Equal("1", Min(list))

	list2 := []int{5, 2, 111, 3}
	ass.Equal(2, Min(list2))

	ass.Equal(0, Min([]int{}))
}

func TestMax(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []string{"1", "2", "3"}
	ass.Equal("3", Max(list))

	list2 := []int{5, 2, 111, 3}
	ass.Equal(111, Max(list2))

	ass.Equal(0, Max([]int{}))
}

func TestDrop(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	s := []string{"1", "2", "3", "4"}

	ass.Equal([]string{"3", "4"}, Drop(s, 2))
	ass.Equal([]string{}, Drop(s, 1111))
	ass.Equal(s, Drop(s, 0))
}

func TestDropLast(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	s := []string{"1", "2", "3", "4"}

	ass.Equal([]string{"1", "2"}, DropLast(s, 2))
	ass.Equal([]string{}, DropLast(s, 1111))

	ass.Equal(s, DropLast(s, 0))
}

func TestSlice(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	s := []string{"1", "2", "3", "4"}

	ass.Equal([]string{"2", "3"}, Slice(s, 1, 3))
	ass.Equal([]string{"2", "3", "4"}, Slice(s, 1, 4))
	ass.Equal([]string{}, Slice(s, 5, 10))
	ass.Equal([]string{}, Slice(s, 10, 5))
	ass.Equal([]string{"2", "3", "4"}, Slice(s, 1, 6))
}

func TestIsSorted(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
	ass.True(IsSorted([]int{0, 1, 2, 3}))
	ass.True(IsSorted([]string{"a", "b", "c", "d"}))

	ass.False(IsSorted([]int{0, 1, 4, 3}))
	ass.False(IsSorted([]string{"a", "e", "d", "c"}))
}

func TestIsSortedByKey(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ass.True(IsSortedBy([]string{"a", "bb", "ccc"}, func(s string) int {
		return len(s)
	}))

	ass.True(IsSortedBy([]string{"1", "2", "3", "11"}, func(s string) int {
		ret, _ := strconv.Atoi(s)
		return ret
	}))

	ass.False(IsSortedBy([]string{"aa", "b", "ccc"}, func(s string) int {
		return len(s)
	}))

}

func TestReverse(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
	list := []int{1, 2, 3, 4, 5}
	Reverse(list)

	ass.Equal([]int{5, 4, 3, 2, 1}, list)
}

func TestFlatten(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
	list := [][]int{
		{1, 2, 3},
		{6, 7, 8},
	}
	list1 := Flatten(list)

	ass.Equal([]int{1, 2, 3, 6, 7, 8}, list1)

}

func TestInsertAt(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{1, 2, 3, 4, 5}

	list1 := InsertAt(list, 2, 666, 777, 888)

	ass.Equal([]int{1, 2, 666, 777, 888, 3, 4, 5}, list1)

	ass.Equal(list, InsertAt(list, 0))
	ass.Equal([]int{1, 2, 3, 4, 5, 6}, InsertAt(list, -1, 6))
	ass.Equal([]int{1, 2, 3, 4, 5, 6}, InsertAt(list, 10, 6))
}
