package sliceutil

import (
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParallelForEach(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int32{1, 2, 3, 4, 5}

	var i atomic.Int32

	ParallelForEach(list, func(index int, item int32) {
		i.Add(item)
	})

	ass.Equal(int32(15), i.Load())

	ParallelForEach([]int32{}, func(index int, item int32) {
		i.Add(item)
	})
	ass.Equal(int32(15), i.Load())

}

func TestConcurrentForEach(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int32{1, 2, 3, 4, 5}

	var i atomic.Int32

	ConcurrentForEach(list, func(index int, item int32) {
		i.Add(item)
	})

	time.Sleep(1 * time.Millisecond)

	ass.Equal(int32(15), i.Load())
}

func TestParallelMap(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{1, 2, 3, 4, 5}

	result := ParallelMap(list, func(index int, item int) string {
		return strconv.Itoa(item)
	})

	ass.Equal([]string{"1", "2", "3", "4", "5"}, result)

	ass.Empty(ParallelMap([]int{}, func(index int, item int) string {
		return strconv.Itoa(item)
	}))
}

func TestConcurrentMap(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	list := []int{1, 2, 3, 4, 5}

	result := ConcurrentMap(list, func(index int, item int) string {
		return strconv.Itoa(item)
	})

	time.Sleep(1 * time.Millisecond)

	ass.Equal([]string{"1", "2", "3", "4", "5"}, result)

	ass.Empty(ConcurrentMap([]int{}, func(index int, item int) string {
		return strconv.Itoa(item)
	}))
}
