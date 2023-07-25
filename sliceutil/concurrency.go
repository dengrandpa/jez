package sliceutil

import (
	"sync"
)

// ParallelForEach Executes parallel iterator functions on each element in the slice.
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

// ConcurrentForEach Executes concurrent iterator functions on each element in the slice.
func ConcurrentForEach[T any](list []T, iteratee func(index int, item T)) {
	for i, item := range list {
		go iteratee(i, item)
	}
}

// ParallelMap Returns a slice containing the results of parallel calls to the provided function on each element in the called slice.
// Note: The order of the results is not guaranteed.
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

// ConcurrentMap Returns a slice containing the results of concurrent calls to the provided function on each element in the called slice.
// Note: The order of the results is not guaranteed.
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
