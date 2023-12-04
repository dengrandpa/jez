package slicejez

import (
	"sync"
)

// SafeSlice 并发安全的切片。
type SafeSlice[T comparable] struct {
	list []T
	lock *sync.RWMutex
}

// NewSafeSlice 创建一个并发安全的切片。
func NewSafeSlice[T comparable](list []T) *SafeSlice[T] {
	return &SafeSlice[T]{
		list: list,
		lock: new(sync.RWMutex),
	}
}

// ForEach 遍历切片并为每个元素调用 iteratee 函数。
func (s *SafeSlice[T]) ForEach(iteratee func(index int, item T)) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	ForEach(s.list, iteratee)
}

// ForEachWithBreak 遍历切片并为每个元素调用 iteratee 函数，如果返回 false，则停止遍历。
func (s *SafeSlice[T]) ForEachWithBreak(iteratee func(index int, item T) bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	ForEachWithBreak(s.list, iteratee)
}

// Filter 遍历切片并为每个元素调用 iteratee 函数，只返回调用结果为true的元素。
func (s *SafeSlice[T]) Filter(iteratee func(index int, item T) bool) []T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return Filter(s.list, iteratee)
}

// Append 添加元素到切片。
func (s *SafeSlice[T]) Append(items ...T) {
	s.lock.Lock()
	s.list = append(s.list, items...)
	s.lock.Unlock()
}

// AppendIfNotDuplicate 添加元素到切片，如果元素已经存在，则不添加。
func (s *SafeSlice[T]) AppendIfNotDuplicate(item T) {
	s.lock.Lock()
	s.list = AppendIfNotDuplicate(s.list, item)
	s.lock.Unlock()
}

// AppendMultipleIfNotDuplicate 添加多个元素到切片，如果元素已经存在，则不添加。
func (s *SafeSlice[T]) AppendMultipleIfNotDuplicate(items ...T) {
	s.lock.Lock()
	s.list = AppendMultipleIfNotDuplicate(s.list, items...)
	s.lock.Unlock()
}

// Load 返回切片的副本。
func (s *SafeSlice[T]) Load() []T {
	s.lock.RLock()
	defer s.lock.RUnlock()

	dst := make([]T, len(s.list))
	copy(dst, s.list)
	return dst
}

// LoadByIndex 返回指定索引位置的元素，-1 则返回最后一个元素，如果索引超出范围，panic。
func (s *SafeSlice[T]) LoadByIndex(index int) T {
	if index == -1 {
		index = len(s.list) - 1
	}
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.list[index]
}

// Index 返回指定元素在切片中的索引位置。
func (s *SafeSlice[T]) Index(item T) int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return FindIndex(s.list, item)
}

// Insert 在指定索引位置插入元素。
func (s *SafeSlice[T]) Insert(index int, items ...T) {
	s.lock.Lock()
	s.list = InsertAt(s.list, index, items...)
	s.lock.Unlock()
}

// Len 返回切片的长度。
func (s *SafeSlice[T]) Len() int {
	return len(s.list)
}

// Remove 从切片中移除元素。
func (s *SafeSlice[T]) Remove(items ...T) {
	s.lock.Lock()
	s.list = Remove(s.list, items...)
	s.lock.Unlock()
}

// RemoveByIndex 从切片中移除指定索引位置的元素。
func (s *SafeSlice[T]) RemoveByIndex(index int) {
	s.lock.Lock()
	s.list = append(s.list[:index], s.list[index+1:]...)
	s.lock.Unlock()
}

// Replace 将切片中的元素 old 替换为 new ，最多替换 n 次，如果 n 为-1，则替换所有的 old 元素。
func (s *SafeSlice[T]) Replace(old T, new T, n int) {
	s.lock.Lock()
	s.list = Replace(s.list, old, new, n)
	s.lock.Unlock()
}

// ReplaceByIndex 将指定索引位置的元素替换为 new 。
func (s *SafeSlice[T]) ReplaceByIndex(index int, new T) {
	s.lock.Lock()
	s.list[index] = new
	s.lock.Unlock()
}

// Slice 返回索引从 n 到 m 的切片，但不包括 m，等同于 slice[n:m]，即[min,max)，但不会在溢出时panic。
func (s *SafeSlice[T]) Slice(n, m int) []T {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return Slice(s.list, n, m)
}
