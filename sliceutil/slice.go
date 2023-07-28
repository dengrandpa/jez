// Package sliceutil provides some useful methods for the slice.
package sliceutil

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// ForEach iterates over a slice and calls the iteratee for each item.
func ForEach[T any](list []T, iteratee func(index int, item T)) {
	for i, item := range list {
		iteratee(i, item)
	}
}

// ForEachWithBreak iterates over a slice and calls the iteratee for each item until iteratee return false.
func ForEachWithBreak[T any](list []T, iteratee func(index int, item T) bool) {
	for i, v := range list {
		if !iteratee(i, v) {
			break
		}
	}
}

// Filter iterates over elements of a slice, returning an array of all elements predicate returns truthy for.
func Filter[V any](list []V, iteratee func(index int, item V) bool) []V {
	result := make([]V, 0, len(list))

	for i, item := range list {
		if iteratee(i, item) {
			result = append(result, item)
		}
	}

	return result
}

// Map return a slice with the results of calling a provided function on every element in the calling slice.
func Map[T, U any](list []T, iteratee func(index int, item T) U) []U {
	result := make([]U, len(list))

	for i, item := range list {
		result[i] = iteratee(i, item)
	}

	return result
}

// Contain checks if a slice contains a value.
func Contain[T comparable](list []T, target T) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}

	return false
}

// ContainFilter returns true if iteratee function return true.
func ContainFilter[T any](list []T, iteratee func(index int, item T) bool) bool {
	for i, item := range list {
		if iteratee(i, item) {
			return true
		}
	}

	return false
}

// ContainAll checks if a slice contains all values.
func ContainAll[T comparable](list []T, sub ...T) bool {
	if len(sub) == 0 {
		return true
	}

	if len(list) == 0 {
		return false
	}

	exist := make(map[T]struct{}, len(list))
	for _, v := range list {
		exist[v] = struct{}{}
	}

	for _, v := range sub {
		if _, ok := exist[v]; !ok {
			return false
		}
	}

	return true
}

// FilterMap return a slice with the results of calling a provided function on every element in the calling slice.
func FilterMap[T, U any](list []T, iteratee func(index int, item T) (U, bool)) []U {
	result := make([]U, 0, len(list))

	for i, item := range list {
		if r, ok := iteratee(i, item); ok {
			result = append(result, r)
		}
	}

	return result
}

// AppendIfNotDuplicate appends an item to a slice if it does not already exist.
func AppendIfNotDuplicate[T comparable](list []T, item T) []T {
	if !Contain(list, item) {
		list = append(list, item)
	}

	return list
}

// AppendMultipleIfNotDuplicate appends multiple items to a slice if it does not already exist.
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

// Remove removes the first occurrence of the target from the slice.
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

// RemoveFilter removes the first occurrence from the result of the iteratee function.
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

// Unique returns a slice with unique values.
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

// UniqueBy returns a slice with unique elements from the result of the iteratee function.
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

// UniqueNonzero returns a slice with unique non-zero elements from the original slice.
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

// UniqueNonzeroBy returns a slice with unique non-zero elements from the result of the iteratee function.
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

// Nonzero returns a slice with non-zero elements from the original slice.
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

// Replace replaces the first n occurrences of old with new.
func Replace[T comparable](list []T, old T, new T, n int) []T {
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

// ReplaceAll replaces all occurrences of old with new.
func ReplaceAll[T comparable](list []T, old T, new T) []T {
	return Replace(list, old, new, -1)
}

// Difference returns the difference between two slices.
// Note: This method does not remove duplicate elements from the result.
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

// DifferenceUnique returns the difference between two slices.
// Note: This method removes duplicate elements from the result.
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

// Intersection returns the intersection between two slices.
// Note: This method removes duplicate elements from the result.
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

// MutualDifference returns the mutual difference between two slices.
// The first value is the collection of elements absent of list2.
// The second value is the collection of elements absent of list1.
// Note: This method does not remove duplicate elements from the result.
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

// ToMapBy converts a slice to a map by the given key function.
func ToMapBy[T any, K comparable, V any](list []T, iteratee func(index int, item T) (K, V)) map[K]V {
	result := make(map[K]V, len(list))

	for index, t := range list {
		k, v := iteratee(index, t)
		result[k] = v
	}

	return result
}

// Repeat returns a slice consisting of n copies of item.
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

// Equal returns true if the slices have the same length, element order and values are the same.
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

// EqualElement returns true if the slices have the same length and values are the same(regardless of order).
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

// FindIndex returns the index the first occurrence of the target in the slice, or -1 if there is no match.
func FindIndex[T comparable](list []T, target T) int {
	for i, item := range list {
		if item == target {
			return i
		}
	}

	return -1
}

// FindIndexFilter returns the index of the first element in the slice that the iteratee function returns true, or -1 if there is no match.
func FindIndexFilter[T any](list []T, iteratee func(index int, item T) bool) int {
	for i, item := range list {
		if iteratee(i, item) {
			return i
		}
	}

	return -1
}

// FindDuplicates returns a slice with all the duplicate elements of the list.
// Note: This method does not remove duplicate elements from the result.
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

// FindDuplicatesBy returns a slice with all the duplicate elements from the result of the iteratee function.
func FindDuplicatesBy[T, U comparable](list []T, iteratee func(index int, item T) U) []U {
	exist := make(map[U]struct{}, len(list))
	result := make([]U, 0, len(list))

	for i, item := range list {
		key := iteratee(i, item)

		if _, ok := exist[key]; ok {
			result = append(result, key)
		} else {
			exist[key] = struct{}{}
		}
	}

	return result
}

// FindUniqueDuplicates returns a slice with all the unique duplicate elements of the slice.
// Note: This method removes duplicate elements from the result.
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

// FindUniqueDuplicatesBy returns a slice with all the unique duplicate elements from the result of the iteratee function.
func FindUniqueDuplicatesBy[T, U comparable](list []T, iteratee func(index int, item T) U) []U {
	exist := make(map[U]int, len(list))
	result := make([]U, 0, len(list))

	for i, item := range list {
		key := iteratee(i, item)

		c, ok := exist[key]

		if ok && c == 1 {
			result = append(result, key)
		}
		exist[key]++
	}

	return result
}

// Shuffle returns randomly shuffle the elements in the given slice.
// Note: Before using the Shuffle function, make sure to call rand.Seed in your program's initialization code.
func Shuffle[T any](list []T) []T {
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	return list
}

// Sample returns a random item from slice.
// Note: Before using the Sample function, make sure to call rand.Seed in your program's initialization code.
func Sample[T any](list []T) T {
	var zero T

	if len(list) == 0 {
		return zero
	}

	return list[rand.Intn(len(list))]
}

// Samples returns N random items from slice.
// Note1: If n is greater than the length of the slice, the entire slice after randomization is returned.
// Note2: If there are duplicates in the slice, the returned one may also have duplicates.
func Samples[T any](list []T, n int) []T {
	if n <= 0 {
		return list
	}

	newList := make([]T, len(list))
	copy(newList, list)

	newList = Shuffle(newList)

	if n > len(newList) {
		return newList
	}

	return newList[:n]
}

// Min returns the minimum value of a slice.
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

// MinFilter returns the minimum value of a slice using the given iteratee function.
func MinFilter[T any](list []T, iteratee func(index int, item T, max T) bool) T {
	var min T

	if len(list) == 0 {
		return min
	}

	min = list[0]

	for i := 1; i < len(list); i++ {

		if iteratee(i, list[i], min) {
			min = list[i]
		}
	}

	return min
}

// Max searches the maximum value of a slice.
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

// MaxFilter search the maximum value of a slice using the given iteratee function.
func MaxFilter[T any](list []T, iteratee func(index int, item T, max T) bool) T {
	var max T

	if len(list) == 0 {
		return max
	}

	max = list[0]

	for i := 1; i < len(list); i++ {

		if iteratee(i, list[i], max) {
			max = list[i]
		}
	}

	return max
}

// Drop returns a slice with n elements dropped from the start.
// Note: if "n" is greater than the length of the slice, an empty slice is returned.
func Drop[T any](list []T, n int) []T {
	if n <= 0 {
		return list
	}

	if len(list) <= n {
		return []T{}
	}

	return list[n:]
}

// DropLast returns a slice with n elements dropped from the end.
// Note: if "n" is greater than the length of the slice, an empty slice is returned.
func DropLast[T any](list []T, n int) []T {
	if n <= 0 {
		return list
	}

	if len(list) <= n {
		return []T{}
	}

	return list[:len(list)-n]
}

// Slice Returns a slice from "n" to "m" but not including "m". Like "slice[n:m]", but doesn't panic on overflow
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

// IsSorted checks if a slice is sorted.
func IsSorted[T constraints.Ordered](list []T) bool {
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			return false
		}
	}

	return true
}

// IsSortedBy checks if a slice is sorted by iteratee.
func IsSortedBy[T any, K constraints.Ordered](list []T, iteratee func(item T) K) bool {
	size := len(list)

	for i := 0; i < size-1; i++ {
		if iteratee(list[i]) > iteratee(list[i+1]) {
			return false
		}
	}

	return true
}

// Reverse reverses a slice.
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Flatten flattens a slice of slices.
func Flatten[T any](collection [][]T) []T {
	var result []T
	for i := range collection {
		result = append(result, collection[i]...)
	}

	return result
}

// InsertAt inserts values into a slice at the specified index.
// Note: if the index is greater than the length of the slice or less than 0,
// the values are appended to the end of the slice.
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
