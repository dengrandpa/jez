package slicejez

import (
	"sync"
)

// ParallelForEach 遍历切片并对每个元素并行调用 iteratee 函数。
func ParallelForEach[T any](list []T, iteratee func(index int, item T)) {
	if len(list) == 0 {
		return
	}

	var wg sync.WaitGroup

	wg.Add(len(list))

	for i, item := range list {
		go func(index int, item T) {
			defer wg.Done()
			iteratee(index, item)
		}(i, item)
	}

	wg.Wait()
}

// ConcurrentForEach 遍历切片并对每个元素并发调用 iteratee 函数。
func ConcurrentForEach[T any](list []T, iteratee func(index int, item T)) {
	for i, item := range list {
		go iteratee(i, item)
	}
}

// ParallelMap 遍历切片并对每个元素并行调用 iteratee 函数，返回一个包含每次调用结果的切片。
func ParallelMap[T, U any](list []T, iteratee func(index int, item T) U) []U {
	if len(list) == 0 {
		return []U{}
	}

	result := make([]U, len(list))

	var wg sync.WaitGroup
	wg.Add(len(list))

	for i, item := range list {
		go func(index int, item T) {
			defer wg.Done()
			result[index] = iteratee(index, item)
		}(i, item)
	}

	wg.Wait()

	return result
}

// ConcurrentMap 遍历切片并对每个元素并发调用 iteratee 函数，返回一个包含每次调用结果的切片。
func ConcurrentMap[T, U any](list []T, iteratee func(index int, item T) U) []U {
	if len(list) == 0 {
		return []U{}
	}

	result := make([]U, len(list))

	for i, item := range list {

		go func(index int, item T) {
			result[index] = iteratee(index, item)
		}(i, item)
	}

	return result
}
