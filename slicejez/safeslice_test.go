package slicejez

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildTestSafeSlice(start, end int) *SafeSlice[int] {
	ss := NewSafeSlice([]int{})
	for i := start; i < end; i++ {
		ss.Append(i)
	}

	return ss
}

func TestSafeSlice_ForEach(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
	ss := buildTestSafeSlice(0, 1000)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ss.ForEach(func(index int, item int) {
				ass.Equal(index, item)
			})
		}(i)
	}

	wg.Wait()

}

func TestSafeSlice_ForEachWithBreak(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
	ss := buildTestSafeSlice(0, 1000)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()

			stop := 500
			expected := buildTestSafeSlice(0, 0)
			ss.ForEachWithBreak(func(index int, item int) bool {
				ok := index < stop
				if ok {
					expected.Append(item)
				}
				return ok
			})

			actual := buildTestSafeSlice(0, stop).Load()

			ass.Equal(expected.Load(), actual)
		}(i)
	}
}

func TestSafeSlice_Filter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
	ss := buildTestSafeSlice(0, 1000)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			stop := 500

			actual := ss.Filter(func(index int, item int) bool {
				return item < stop
			})

			expected := buildTestSafeSlice(0, stop).Load()

			ass.Equal(expected, actual)
		}(i)

	}
	wg.Wait()
}

func TestSafeSlice_Append(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := NewSafeSlice([]int{})

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ss.Append(v)
		}(i)
	}
	wg.Wait()

	ass.Len(ss.Load(), 1000)
}

func TestSafeSlice_AppendIfNotDuplicate(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := buildTestSafeSlice(0, 0)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ss.AppendIfNotDuplicate(v)
		}(i)
	}

	wg.Wait()

	actual := ToMapBy(ss.Load(), func(index int, item int) (int, struct{}) {
		return item, struct{}{}
	})

	ass.Len(ss.Load(), 1000)
	ass.Len(ss.Load(), len(actual))
}

func TestSafeSlice_AppendMultipleIfNotDuplicate(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := buildTestSafeSlice(0, 0)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ss.AppendMultipleIfNotDuplicate(v, v+1)
		}(i)
	}

	wg.Wait()

	actual := ToMapBy(ss.Load(), func(index int, item int) (int, struct{}) {
		return item, struct{}{}
	})

	ass.Len(ss.Load(), 1001)
	ass.Len(ss.Load(), len(actual))
}

func TestSafeSlice_Load(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := buildTestSafeSlice(0, 1000)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ass.Len(ss.Load(), 1000)
		}(i)
	}
}

func TestSafeSlice_LoadByIndex(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := buildTestSafeSlice(0, 1000)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			actual := ss.LoadByIndex(0)
			ass.Equal(0, actual)
		}(i)
	}
}

func TestSafeSlice_Index(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := buildTestSafeSlice(0, 1000)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			actual := ss.Index(1)
			ass.Equal(1, actual)
		}(i)
	}
}

func TestSafeSlice_Insert(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := buildTestSafeSlice(0, 2)

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ss.Insert(1, v)
		}(i)
	}

	wg.Wait()

	ass.Len(ss.Load(), 1002)
	ass.Equal(0, ss.LoadByIndex(0))
	ass.Equal(1, ss.LoadByIndex(-1))
}

func TestSafeSlice_Len(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := buildTestSafeSlice(0, 1000)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ass.Equal(1000, ss.Len())
		}(i)
	}
}

func TestSafeSlice_Remove(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := buildTestSafeSlice(0, 1000)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ss.Remove(v)
		}(i)
	}

	wg.Wait()

	ass.Len(ss.Load(), 0)
}

func TestSafeSlice_RemoveByIndex(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := buildTestSafeSlice(0, 1000)

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ss.RemoveByIndex(0)
		}(i)
	}

	wg.Wait()

	ass.Len(ss.Load(), 0)
}

func TestSafeSlice_Replace(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := buildTestSafeSlice(0, 1000)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ss.Replace(v, 1000+v, -1)
		}(i)
	}
	wg.Wait()

	expected := buildTestSafeSlice(1000, 2000).Load()

	ass.Equal(expected, ss.Load())
}

func TestSafeSlice_ReplaceByIndex(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := buildTestSafeSlice(0, 1000)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ss.ReplaceByIndex(v, 999-v)
		}(i)
	}

	wg.Wait()

	expected := buildTestSafeSlice(0, 1000).Load()
	Reverse(expected)
	ass.Equal(expected, ss.Load())
}

func TestSafeSlice_Slice(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ss := buildTestSafeSlice(0, 1000)

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			actual := ss.Slice(500, 501)
			ass.Equal([]int{500}, actual)
		}(i)
	}
}
