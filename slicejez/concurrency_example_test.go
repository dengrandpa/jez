package slicejez

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"
)

func ExampleParallelForEach() {

	list := []int32{1, 2, 3, 4, 5}

	var i atomic.Int32

	ParallelForEach(list, func(index int, item int32) {
		i.Add(item)
	})

	fmt.Println(i.Load())

	// Output:
	// 15
}

func ExampleConcurrentForEach() {

	list := []int32{1, 2, 3, 4, 5}

	var i atomic.Int32

	ConcurrentForEach(list, func(index int, item int32) {
		i.Add(item)
	})

	time.Sleep(1 * time.Millisecond)

	fmt.Println(i.Load())

	// Output:
	// 15
}

func ExampleParallelMap() {

	list := []int64{1, 2, 3, 4, 5}

	result := ParallelMap(list, func(index int, item int64) string {
		return strconv.FormatInt(item, 10)
	})

	fmt.Println(result)

	// Output:
	// [1 2 3 4 5]
}

func ExampleConcurrentMap() {

	list := []int64{1, 2, 3, 4, 5}

	result := ConcurrentMap(list, func(index int, item int64) string {
		return strconv.FormatInt(item, 10)
	})

	time.Sleep(1 * time.Millisecond)

	fmt.Println(result)

	// Output:
	// [1 2 3 4 5]
}
