// Package sliceutil slice相关函数
package sliceutil

import (
	"golang.org/x/exp/constraints"
)

// ForEach 遍历切片并为每个元素调用 iteratee 函数。
func ForEach[T any](list []T, iteratee func(index int, item T)) {
	for i, item := range list {
		iteratee(i, item)
	}
}

// ForEachWithBreak 遍历切片并为每个元素调用 iteratee 函数，如果返回 false，则停止遍历。
func ForEachWithBreak[T any](list []T, iteratee func(index int, item T) bool) {
	for i, v := range list {
		if !iteratee(i, v) {
			break
		}
	}
}

// Filter 遍历切片并为每个元素调用 iteratee 函数，只返回调用结果为true的元素。
func Filter[T any](list []T, iteratee func(index int, item T) bool) []T {
	result := make([]T, 0, len(list))

	for i, item := range list {
		if iteratee(i, item) {
			result = append(result, item)
		}
	}

	return result
}

// Map 遍历切片并为每个元素调用 iteratee 函数，并返回调用后结果。
func Map[T, U any](list []T, iteratee func(index int, item T) U) []U {
	result := make([]U, len(list))

	for i, item := range list {
		result[i] = iteratee(i, item)
	}

	return result
}

// Contain 效验切片是否包含目标元素。
func Contain[T comparable](list []T, target T) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}

	return false
}

// ContainAll 效验切片是否包含所有的目标元素。
func ContainAll[T comparable](list []T, targets ...T) bool {
	if len(targets) == 0 {
		return true
	}

	if len(list) == 0 {
		return false
	}

	exist := make(map[T]struct{}, len(list))
	for _, v := range list {
		exist[v] = struct{}{}
	}

	for _, v := range targets {
		if _, ok := exist[v]; !ok {
			return false
		}
	}

	return true
}

// FilterMap 遍历切片并为每个元素调用 iteratee 函数，如果调用结果为true，则返回调用后元素。
func FilterMap[T, U any](list []T, iteratee func(index int, item T) (U, bool)) []U {
	result := make([]U, 0, len(list))

	for i, item := range list {
		if r, ok := iteratee(i, item); ok {
			result = append(result, r)
		}
	}

	return result
}

// AppendIfNotDuplicate 添加元素到切片，如果元素已经存在，则不添加。
func AppendIfNotDuplicate[T comparable](list []T, item T) []T {
	if !Contain(list, item) {
		list = append(list, item)
	}

	return list
}

// AppendMultipleIfNotDuplicate 添加多个元素到切片，如果元素已经存在，则不添加。
func AppendMultipleIfNotDuplicate[T comparable](list []T, items ...T) []T {
	if len(items) == 0 {
		return list
	}

	exist := make(map[T]struct{}, len(list))
	for _, item := range list {
		exist[item] = struct{}{}
	}

	for _, item := range items {
		if _, ok := exist[item]; !ok {
			list = append(list, item)
			exist[item] = struct{}{}
		}
	}

	return list
}

// Remove 从切片中删除元素。
func Remove[T comparable](list []T, items ...T) []T {
	if len(list) == 0 || len(items) == 0 {
		return list
	}

	toRemove := make(map[T]struct{}, len(items))
	for _, item := range items {
		toRemove[item] = struct{}{}
	}

	result := make([]T, 0, len(list))
	for _, item := range list {
		if _, ok := toRemove[item]; !ok {
			result = append(result, item)
		}
	}

	return result
}

// RemoveFilter 遍历切片并为每个元素调用 iteratee 函数，如果调用结果为true，则删除该元素。
func RemoveFilter[T comparable](list []T, iteratee func(index int, item T) bool) []T {
	if len(list) == 0 {
		return list
	}

	result := make([]T, 0, len(list))

	for i, item := range list {
		if !iteratee(i, item) {
			result = append(result, item)
		}
	}

	return result
}

// Unique 去重。
func Unique[T comparable](list []T) []T {
	result := make([]T, 0, len(list))
	exist := make(map[T]struct{}, len(list))

	for _, item := range list {
		if _, ok := exist[item]; !ok {
			result = append(result, item)
			exist[item] = struct{}{}
		}
	}

	return result
}

// UniqueBy 遍历切片并为每个元素调用 iteratee 函数，返回唯一的元素。
func UniqueBy[T, U comparable](list []T, iteratee func(index int, item T) U) []U {
	result := make([]U, 0, len(list))
	seen := make(map[U]struct{}, len(list))

	for i, item := range list {
		key := iteratee(i, item)

		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, key)
	}

	return result
}

// UniqueNonzero 删除重复元素及零值。
func UniqueNonzero[T comparable](list []T) []T {
	result := make([]T, 0, len(list))
	exist := make(map[T]struct{}, len(list))

	var zero T

	for _, item := range list {
		if item == zero {
			continue
		}
		if _, ok := exist[item]; !ok {
			result = append(result, item)
			exist[item] = struct{}{}
		}
	}

	return result
}

// UniqueNonzeroBy 遍历切片并为每个元素调用 iteratee 函数，返回唯一的、非零值的元素。
func UniqueNonzeroBy[T, U comparable](list []T, iteratee func(index int, item T) U) []U {
	result := make([]U, 0, len(list))
	seen := make(map[U]struct{}, len(list))

	var zero T

	for i, item := range list {
		if item == zero {
			continue
		}

		key := iteratee(i, item)

		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, key)
	}

	return result
}

// Nonzero 删除零值。
func Nonzero[T comparable](list []T) []T {
	var zero T

	result := make([]T, 0, len(list))

	for _, item := range list {
		if item != zero {
			result = append(result, item)
		}
	}

	return result
}

// Replace 将切片中的元素 old 替换为 new ，最多替换 n 次，如果 n 为-1，则替换所有的 old 元素。
func Replace[T comparable](list []T, old T, new T, n int) []T {
	if n == 0 {
		return list
	}

	result := make([]T, len(list))
	copy(result, list)

	for i := range result {
		if result[i] == old && n != 0 {
			result[i] = new
			n--
		}
	}

	return result
}

// ReplaceAll 将切片中的元素 old 替换为 new ，替换所有的 old 元素。
func ReplaceAll[T comparable](list []T, old T, new T) []T {
	return Replace(list, old, new, -1)
}

// Difference 差集，结果不去重。
func Difference[T comparable](list1, list2 []T) []T {

	if len(list1) == 0 {
		return list2
	}

	if len(list2) == 0 {
		return list1
	}

	result := make([]T, 0, len(list1)+len(list2))

	l1 := make(map[T]struct{}, len(list1))
	l2 := make(map[T]struct{}, len(list2))

	for _, v := range list1 {
		l1[v] = struct{}{}
	}

	for _, v := range list2 {
		l2[v] = struct{}{}
	}

	for _, v := range list1 {
		if _, ok := l2[v]; !ok {
			result = append(result, v)
		}
	}

	for _, v := range list2 {
		if _, ok := l1[v]; !ok {
			result = append(result, v)
		}
	}

	return result
}

// DifferenceUnique 差集，结果去重。
func DifferenceUnique[T comparable](list1, list2 []T) []T {

	if len(list1) == 0 {
		return Unique(list2)
	}

	if len(list2) == 0 {
		return Unique(list1)
	}

	result := make([]T, 0, len(list1)+len(list2))

	l1 := make(map[T]struct{}, len(list1))
	l2 := make(map[T]struct{}, len(list2))

	for _, v := range list1 {
		l1[v] = struct{}{}
	}

	for _, v := range list2 {
		l2[v] = struct{}{}
	}

	for _, v := range list1 {
		if _, ok := l2[v]; !ok {
			result = append(result, v)
			l2[v] = struct{}{}
		}
	}

	for _, v := range list2 {
		if _, ok := l1[v]; !ok {
			result = append(result, v)
			l1[v] = struct{}{}
		}
	}

	return result
}

// Intersection 交集，结果元素唯一。
func Intersection[T comparable](list1, list2 []T) []T {
	if len(list1) == 0 || len(list2) == 0 {
		return []T{}
	}

	result := make([]T, 0, len(list1))
	exist := make(map[T]struct{}, len(list1)+len(list2))

	for _, v := range list2 {
		exist[v] = struct{}{}
	}

	for _, v := range list1 {
		if _, ok := exist[v]; ok {
			result = append(result, v)
			delete(exist, v)
		}
	}

	return result
}

// MutualDifference 差异，结果不去重。
//
// 返回值：
//   - list1 中存在， list2 中不存在。
//   - list2 中存在， list1 中不存在。
func MutualDifference[T comparable](list1, list2 []T) ([]T, []T) {
	if len(list1) == 0 {
		return []T{}, list2
	}

	if len(list2) == 0 {
		return list1, []T{}
	}

	left := make([]T, 0, len(list1)+len(list2))
	right := make([]T, 0, len(list1)+len(list2))

	seenLeft := make(map[T]struct{}, len(list1))
	seenRight := make(map[T]struct{}, len(list2))

	for _, v := range list1 {
		seenLeft[v] = struct{}{}
	}

	for _, v := range list2 {
		seenRight[v] = struct{}{}
	}

	for _, v := range list1 {
		if _, ok := seenRight[v]; !ok {
			left = append(left, v)
		}
	}

	for _, v := range list2 {
		if _, ok := seenLeft[v]; !ok {
			right = append(right, v)
		}
	}

	return left, right
}

// ToMapBy 遍历切片，将切片中的元素转换为map的key和value。
func ToMapBy[T any, K comparable, V any](list []T, iteratee func(index int, item T) (K, V)) map[K]V {
	result := make(map[K]V, len(list))

	for index, t := range list {
		k, v := iteratee(index, t)
		result[k] = v
	}

	return result
}

// Repeat 返回包含 n 个 item 的切片。
func Repeat[T any](item T, n int) []T {
	if n <= 0 {
		return []T{}
	}

	result := make([]T, n)

	for i := range result {
		result[i] = item
	}

	return result
}

// Equal 长度、顺序、值都相等时返回 true 。
func Equal[T comparable](list1, list2 []T) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i, item := range list1 {
		if item != list2[i] {
			return false
		}
	}

	return true
}

// EqualElement 长度、值相等时返回 true ，不考虑顺序。
func EqualElement[T comparable](list1 []T, list2 []T) bool {
	if len(list1) != len(list2) {
		return false
	}

	exist1 := make(map[T]int, len(list1))
	exist2 := make(map[T]int, len(list2))

	for _, item := range list1 {
		exist1[item]++
	}

	for _, item := range list2 {
		exist2[item]++
	}

	for item, count1 := range exist1 {
		count2, ok := exist2[item]
		if !ok || count1 != count2 {
			return false
		}
	}

	return true
}

// FindIndex 返回第一个匹配的元素的索引，不存在则返回 -1 。
func FindIndex[T comparable](list []T, target T) int {
	for i, item := range list {
		if item == target {
			return i
		}
	}

	return -1
}

// FindIndexFilter 返回调用 iteratee 函数返回 true 的第一个元素的索引，不存在则返回 -1 。
func FindIndexFilter[T any](list []T, iteratee func(index int, item T) bool) int {
	for i, item := range list {
		if iteratee(i, item) {
			return i
		}
	}

	return -1
}

// FindDuplicates 返回切片中所有重复的元素，结果不去重。
func FindDuplicates[T comparable](list []T) []T {
	exist := make(map[T]struct{}, len(list))
	result := make([]T, 0, len(list))

	for _, item := range list {
		if _, ok := exist[item]; ok {
			result = append(result, item)
		} else {
			exist[item] = struct{}{}
		}
	}

	return result
}

// FindUniqueDuplicates 返回切片中所有重复的元素，结果去重。
func FindUniqueDuplicates[T comparable](list []T) []T {
	exist := make(map[T]int, len(list))
	result := make([]T, 0, len(list))

	for _, item := range list {
		c, ok := exist[item]

		if ok && c == 1 {
			result = append(result, item)
		}
		exist[item]++
	}

	return result
}

// Min 返回最小值
func Min[T constraints.Ordered](list []T) T {
	var min T

	if len(list) == 0 {
		return min
	}

	min = list[0]

	for i := 1; i < len(list); i++ {
		if list[i] < min {
			min = list[i]
		}
	}

	return min
}

// Max 返回最大值
func Max[T constraints.Ordered](list []T) T {
	var max T

	if len(list) == 0 {
		return max
	}

	max = list[0]

	for i := 1; i < len(list); i++ {

		if list[i] > max {
			max = list[i]
		}
	}

	return max
}

// Drop 返回从开头删除n个元素的切片，如果 n 大于切片的长度，则返回空切片。
func Drop[T any](list []T, n int) []T {
	if n <= 0 {
		return list
	}

	if len(list) <= n {
		return []T{}
	}

	return list[n:]
}

// DropLast 返回从末尾删除n个元素的切片，如果 n 大于切片的长度，则返回空切片。
func DropLast[T any](list []T, n int) []T {
	if n <= 0 {
		return list
	}

	if len(list) <= n {
		return []T{}
	}

	return list[:len(list)-n]
}

// Slice 返回索引从 n 到 m 的切片，但不包括 m，等同于 slice[n:m]，即[min,max)，但不会在溢出时panic。
func Slice[T any](list []T, n, m int) []T {

	if n >= m {
		return []T{}
	}

	if n >= len(list) {
		return []T{}
	}

	if m >= len(list) {
		return list[n:]
	}

	return list[n:m]
}

// IsSorted 判断切片是否已排序。
func IsSorted[T constraints.Ordered](list []T) bool {
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			return false
		}
	}

	return true
}

// IsSortedBy 遍历切片并为每个元素调用 iteratee 函数，以确定它是否已排序。
func IsSortedBy[T any, K constraints.Ordered](list []T, iteratee func(item T) K) bool {
	size := len(list)

	for i := 0; i < size-1; i++ {
		if iteratee(list[i]) > iteratee(list[i+1]) {
			return false
		}
	}

	return true
}

// Reverse 将切片中的元素顺序反转。
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Flatten 将二维切片转换为一维切片。
func Flatten[T any](collection [][]T) []T {
	var result []T
	for i := range collection {
		result = append(result, collection[i]...)
	}

	return result
}

// InsertAt 在切片的指定索引处插入值，如果索引大于切片的长度或小于 0，则将值附加到切片的末尾。
func InsertAt[T any](slice []T, index int, value ...T) []T {
	if len(value) == 0 {
		return slice
	}

	if index < 0 || index > len(slice) {
		slice = append(slice, value...)
		return slice
	}

	slice = append(slice[:index], append(value, slice[index:]...)...)
	return slice
}
