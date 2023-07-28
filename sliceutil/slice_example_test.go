package sliceutil

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func ExampleForEach() {

	list := []string{"a", "b", "c"}

	items := make([]string, 0, len(list))
	indexes := make([]int, 0, len(list))

	ForEach(list, func(i int, item string) {
		items = append(items, item)
		indexes = append(indexes, i)
	})

	fmt.Println(items)
	fmt.Println(indexes)
	// Output:
	// [a b c]
	// [0 1 2]
}

func ExampleForEachWithBreak() {

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

	fmt.Println(items)
	fmt.Println(indexes)

	// Output:
	// [a b]
	// [0 1]
}

func ExampleFilter() {

	list := []string{"a", "b", "c"}

	list2 := Filter(list, func(index int, item string) bool {
		return index != 1
	})
	fmt.Println(list2)

	// Output:
	// [a c]
}

func ExampleMap() {

	list := []string{"1", "2", "3"}

	list2 := Map(list, func(index int, item string) int {
		i, _ := strconv.Atoi(item)
		return i
	})

	fmt.Println(list2)

	// Output:
	// [1 2 3]
}

func ExampleContain() {

	list := []string{"1", "2", "3"}

	fmt.Println(Contain(list, "1"))
	fmt.Println(Contain(list, "4"))

	// Output:
	// true
	// false

}

func ExampleContainFilter() {

	list := []string{"1", "2", "3"}

	fmt.Println(ContainFilter(list, func(index int, item string) bool {
		return item == "2"
	}))

	fmt.Println(ContainFilter(list, func(index int, item string) bool {
		return item == "4"
	}))

	// Output:
	// true
	// false
}

func ExampleContainAll() {

	list := []string{"1", "2", "3"}

	fmt.Println(ContainAll(list, "1", "2", "3"))
	fmt.Println(ContainAll(list, "1", "2", "2"))
	fmt.Println(ContainAll(list, "1", "2"))
	fmt.Println(ContainAll(list, "1", "4"))

	// Output:
	// true
	// true
	// true
	// false

}

func ExampleFilterMap() {

	list := []string{"1", "2", "3"}

	list2 := FilterMap(list, func(index int, item string) (int, bool) {
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

func ExampleAppendIfNotDuplicate() {

	list := []string{"1", "2", "3"}

	fmt.Println(AppendIfNotDuplicate(list, "1"))
	fmt.Println(AppendIfNotDuplicate(list, "4"))

	// Output:
	// [1 2 3]
	// [1 2 3 4]
}

func ExampleAppendMultipleIfNotDuplicate() {

	list := []string{"1", "2", "3"}

	fmt.Println(AppendMultipleIfNotDuplicate(list, "1"))
	fmt.Println(AppendMultipleIfNotDuplicate(list, "3", "4"))
	fmt.Println(AppendMultipleIfNotDuplicate(list, "4", "5"))

	// Output:
	// [1 2 3]
	// [1 2 3 4]
	// [1 2 3 4 5]
}

func ExampleRemove() {

	list := []string{"1", "2", "3"}

	fmt.Println(Remove(list, "1"))
	fmt.Println(Remove(list, "1", "2"))

	// Output:
	// [2 3]
	// [3]
}

func ExampleRemoveFilter() {

	list := []string{"1", "2", "3"}

	fmt.Println(RemoveFilter(list, func(index int, item string) bool {
		if item == "1" || item == "2" {
			return true
		}

		return false
	}))

	// Output:
	// [3]
}

func ExampleUnique() {

	list := []string{"1", "2", "3"}

	fmt.Println(Unique(list))
	fmt.Println(Unique([]string{"1", "2", "3", "3", "2"}))

	// Output:
	// [1 2 3]
	// [1 2 3]
}

func ExampleUniqueBy() {

	list := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(UniqueBy(list, func(index int, item int) string {
		if item > 3 {
			return "3"
		}
		return strconv.Itoa(item)
	}))

	// Output:
	// [0 1 2 3]
}

func ExampleUniqueNonzero() {

	list := []int{0, 1, 0, 1, 2, 2, 2, 3}

	fmt.Println(UniqueNonzero(list))

	// Output:
	// [1 2 3]
}

func ExampleUniqueNonzeroBy() {

	list := []int{0, 1, 0, 1, 2, 2, 2, 3}

	fmt.Println(UniqueNonzeroBy(list, func(index int, item int) string {
		return strconv.Itoa(item)
	}))

	// Output:
	// [1 2 3]
}

func ExampleNonzero() {

	list := []int{0, 1, 1, 2, 2, 2, 3, 0}

	fmt.Println(Nonzero(list))

	// Output:
	// [1 1 2 2 2 3]
}

func ExampleReplace() {

	list := []int{0, 1, 1, 2, 2, 2, 3, 0}

	fmt.Println(Replace(list, 1, 6, 1))
	fmt.Println(Replace(list, 2, 6, -1))

	// Output:
	// [0 6 1 2 2 2 3 0]
	// [0 1 1 6 6 6 3 0]
}

func ExampleReplaceAll() {

	list := []int{0, 1, 1, 2, 2, 2, 3, 0}

	fmt.Println(ReplaceAll(list, 2, 6))

	// Output:
	// [0 1 1 6 6 6 3 0]
}

func ExampleDifference() {

	list1 := []int{0, 1, 1, 2, 2, 3, 3, 0}

	list2 := []int{0, 1, 1, 2, 2, 2, 6, 6}

	fmt.Println(Difference(list1, list2))

	// Output:
	// [3 3 6 6]
}

func ExampleDifferenceUnique() {

	list1 := []int{0, 1, 1, 2, 2, 3, 3, 0}

	list2 := []int{0, 1, 1, 2, 2, 2, 6, 6}

	fmt.Println(DifferenceUnique(list1, list2))

	// Output:
	// [3 6]
}

func ExampleIntersection() {

	list1 := []int{0, 1, 1, 2, 2, 3, 3, 0}

	list2 := []int{0, 1, 1, 2, 2, 2, 6, 6}

	fmt.Println(Intersection(list1, list2))

	list3 := []int{0, 1, 2, 3}
	list4 := []int{2, 3, 4, 5, 6}

	fmt.Println(Intersection(list3, list4))

	// Output:
	// [0 1 2]
	// [2 3]
}

func ExampleMutualDifference() {

	list1 := []int{0, 1, 1, 2, 2, 2, 3, 0}
	list2 := []int{0, 1, 1, 2, 2, 2, 6, 6}
	fmt.Println(MutualDifference(list1, list2))

	list3 := []int{0, 1, 2, 3}
	list4 := []int{1, 2, 7, 8}
	fmt.Println(MutualDifference(list3, list4))

	var list5 []int
	list6 := []int{1, 2, 3}
	fmt.Println(MutualDifference(list5, list6))

	list7 := []int{1, 2, 3}
	var list8 []int
	fmt.Println(MutualDifference(list7, list8))

	// Output:
	// [3] [6 6]
	// [0 3] [7 8]
	// [] [1 2 3]
	// [1 2 3] []
}

func ExampleToMapBy() {

	list := []int{1, 2, 3, 4, 5, 5, 5}

	m := ToMapBy(list, func(index int, item int) (string, string) {
		itemStr := strconv.Itoa(item)
		return itemStr + "_" + strconv.Itoa(index), itemStr
	})

	fmt.Println(m)

	// Output:
	// map[1_0:1 2_1:2 3_2:3 4_3:4 5_4:5 5_5:5 5_6:5]
}

func ExampleRepeat() {

	fmt.Println(Repeat(1, 5))
	fmt.Println(Repeat("1", 5))

	// Output:
	// [1 1 1 1 1]
	// [1 1 1 1 1]
}

func ExampleEqual() {

	list1 := []int{1, 2, 3, 3, 3}
	list2 := []int{1, 2, 3, 3, 3}
	list3 := []int{3, 2, 3, 3, 1}
	list4 := []int{1, 2}

	fmt.Println(Equal(list1, list2))
	fmt.Println(Equal(list1, list3))
	fmt.Println(Equal(list1, list4))

	// Output:
	// true
	// false
	// false
}

func ExampleEqualElement() {

	list1 := []int{1, 2, 3, 3, 3}
	list2 := []int{1, 2, 3, 3, 3}
	list3 := []int{3, 2, 3, 3, 1}
	list4 := []int{1, 2}
	list5 := []int{3, 2, 1}

	fmt.Println(EqualElement(list1, list2))
	fmt.Println(EqualElement(list1, list3))
	fmt.Println(EqualElement(list1, list4))
	fmt.Println(EqualElement(list1, list5))

	// Output:
	// true
	// true
	// false
	// false
}

func ExampleFindIndex() {

	list := []int{1, 2, 3, 3, 5, 5, 6}

	fmt.Println(FindIndex(list, 3))
	fmt.Println(FindIndex(list, 5))
	fmt.Println(FindIndex(list, 2))
	fmt.Println(FindIndex(list, 9))

	// Output:
	// 2
	// 4
	// 1
	// -1
}

func ExampleFindIndexFilter() {

	list := []int{1, 2, 3, 5, 3, 5, 6}

	n := 0

	fmt.Println(FindIndexFilter(list, func(index int, item int) bool {
		if item == 3 {
			n++
			if n > 1 {
				return true
			}
		}
		return false
	}))

	fmt.Println(FindIndexFilter(list, func(index int, item int) bool {
		return item > 3
	}))

	// Output:
	// 4
	// 3
}

func ExampleFindDuplicates() {

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}

	fmt.Println(FindDuplicates(list))

	// Output:
	// [3 3 5]

}

func ExampleFindDuplicatesBy() {

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}

	fmt.Println(FindDuplicatesBy(list, func(index int, item int) string {
		return strconv.Itoa(item)
	}))

	// Output:
	// [3 3 5]

}

func ExampleFindUniqueDuplicates() {

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}

	fmt.Println(FindUniqueDuplicates(list))

	// Output:
	// [3 5]
}

func ExampleFindUniqueDuplicatesBy() {

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}

	fmt.Println(FindUniqueDuplicatesBy(list, func(index int, item int) string {
		return strconv.Itoa(item)
	}))

	// Output:
	// [3 5]
}

func ExampleShuffle() {

	rand.Seed(time.Now().UnixNano())

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}

	list2 := Shuffle([]int{})

	fmt.Println(len(list) == len(Shuffle(list)))
	fmt.Println(list2)

	// Output:
	// true
	// []
}

func ExampleSample() {

	rand.Seed(time.Now().UnixNano())

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}
	var list2 []string

	fmt.Println(Contain(list, Sample(list)))

	fmt.Println(Sample(list2) == "")

	// Output:
	// true
	// true
}

func ExampleSamples() {

	rand.Seed(time.Now().UnixNano())

	list := []string{"1", "2", "3"}

	s2 := Samples(list, 3)

	fmt.Println(EqualElement(list, s2))

	s3 := Samples(list, 2)

	fmt.Println(ContainAll(list, s3...))

	fmt.Println(Samples([]string{}, 3))

	// Output:
	// true
	// true
	// []
}

func ExampleMin() {

	list := []string{"1", "2", "3"}
	fmt.Println(Min(list))

	list2 := []int{5, 2, 111, 3}
	fmt.Println(Min(list2))

	// Output:
	// 1
	// 2
}

func ExampleMinFilter() {

	list := []string{"ss12", "xs2", "xx2"}
	fmt.Println(MinFilter(list, func(index int, item string, max string) bool {
		if item == "xs2" {
			return item > max
		}

		return item < max
	}))

	list2 := []int{5, 2, 111, 3}

	fmt.Println(MinFilter(list2, func(index int, item int, max int) bool {
		return item < max
	}))

	// Output:
	// xs2
	// 2
}

func ExampleMax() {

	list := []string{"1", "2", "3"}
	fmt.Println(Max(list))

	list2 := []int{5, 2, 111, 3}
	fmt.Println(Max(list2))

	// Output:
	// 3
	// 111

}

func ExampleMaxFilter() {

	list := []string{"ss12", "xs2", "xx2"}

	fmt.Println(MaxFilter(list, func(index int, item string, max string) bool {
		if item == "xs2" {
			return item < max
		}

		return item > max
	}))

	list2 := []int{5, 2, 111, 3}

	fmt.Println(MaxFilter(list2, func(index int, item int, max int) bool {
		return item > max
	}))

	// Output:
	// xx2
	// 111

}

func ExampleDrop() {

	s := []string{"1", "2", "3", "4"}

	fmt.Println(Drop(s, 2))
	fmt.Println(Drop(s, 1111))

	// Output:
	// [3 4]
	// []
}

func ExampleDropLast() {

	s := []string{"1", "2", "3", "4"}

	fmt.Println(DropLast(s, 2))
	fmt.Println(DropLast(s, 1111))

	// Output:
	// [1 2]
	// []
}

func ExampleSlice() {

	s := []string{"1", "2", "3", "4"}

	fmt.Println(Slice(s, 1, 3))
	fmt.Println(Slice(s, 1, 4))
	fmt.Println(Slice(s, 5, 10))
	fmt.Println(Slice(s, 10, 5))
	fmt.Println(Slice(s, 1, 6))

	// Output:
	// [2 3]
	// [2 3 4]
	// []
	// []
	// [2 3 4]
}

func ExampleIsSorted() {

	fmt.Println(IsSorted([]int{0, 1, 2, 3}))
	fmt.Println(IsSorted([]string{"a", "b", "c", "d"}))
	fmt.Println(IsSorted([]int{0, 1, 4, 3}))
	fmt.Println(IsSorted([]string{"a", "e", "d", "c"}))

	// Output:
	// true
	// true
	// false
	// false

}

func ExampleIsSortedBy() {

	fmt.Println(IsSortedBy([]string{"a", "bb", "ccc"}, func(s string) int {
		return len(s)
	}))

	fmt.Println(IsSortedBy([]string{"1", "2", "3", "11"}, func(s string) int {
		ret, _ := strconv.Atoi(s)
		return ret
	}))

	fmt.Println(IsSortedBy([]string{"aa", "b", "ccc"}, func(s string) int {
		return len(s)
	}))

	// Output:
	// true
	// true
	// false
}

func ExampleReverse() {

	list := []int{1, 2, 3, 4, 5}
	Reverse(list)

	fmt.Println(list)

	// Output:
	// [5 4 3 2 1]
}

func ExampleFlatten() {
	list := [][]int{
		{1, 2, 3},
		{6, 7, 8},
	}
	list1 := Flatten(list)

	fmt.Println(list1)

	// Output:
	// [1 2 3 6 7 8]
}

func ExampleInsertAt() {
	list := []int{1, 2, 3, 4, 5}

	list1 := InsertAt(list, 2, 666, 777, 888)

	fmt.Println(list1)

	// Output:
	// [1 2 666 777 888 3 4 5]
}
