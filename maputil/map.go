// Package maputil map相关函数
package maputil

// ForEach 遍历map，对每个元素调用 iteratee 函数。
func ForEach[K comparable, V any](m map[K]V, iteratee func(key K, value V)) {
	for k, v := range m {
		iteratee(k, v)
	}
}

// Filter 遍历map，对每个元素调用 iteratee 函数，如果 iteratee 返回 true，则将该元素添加到结果map中。
func Filter[K comparable, V any](m map[K]V, iteratee func(key K, value V) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if iteratee(k, v) {
			result[k] = v
		}
	}
	return result
}

// Keys 遍历map，将每个key添加到结果slice中。
func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))

	for k := range m {
		result = append(result, k)
	}

	return result
}

// KeysBy 遍历map，对每个元素调用 iteratee 函数，并返回调用后结果。
func KeysBy[K comparable, V any](m map[K]V, iteratee func(item K) K) []K {
	result := make([]K, 0, len(m))

	for k := range m {
		result = append(result, iteratee(k))
	}

	return result
}

// Values 返回map中所有的value。
func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

// ValuesBy 遍历map，对每个元素调用 iteratee 函数，并返回调用后结果。
func ValuesBy[K comparable, V any](m map[K]V, iteratee func(item V) V) []V {
	result := make([]V, 0, len(m))

	for _, v := range m {
		result = append(result, iteratee(v))
	}

	return result
}

// ValuesUnique 返回map中所有的value，结果去重。
func ValuesUnique[K, V comparable](m map[K]V) []V {
	result := make([]V, 0, len(m))
	exist := make(map[V]struct{}, len(m))

	for _, v := range m {
		if _, ok := exist[v]; !ok {
			exist[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

// KeysAndValues 返回map中所有的key和value。
func KeysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))

	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}

// KeysAndValuesFilter 遍历map，对每个元素调用 iteratee 函数，如果 iteratee 返回true，则将该元素添加到结果slice中。
func KeysAndValuesFilter[K comparable, V any](m map[K]V, iteratee func(key K, value V) bool) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))

	for k, v := range m {
		if iteratee(k, v) {
			keys = append(keys, k)
			values = append(values, v)
		}
	}

	return keys, values
}

// Deletes 通过key删除多个元素。
func Deletes[K comparable, V any](m map[K]V, keys ...K) {
	for _, k := range keys {
		delete(m, k)
	}
}

// DeleteByValues 通过value删除多个元素。
func DeleteByValues[K, V comparable](m map[K]V, values ...V) {
	for k, v := range m {
		for _, item := range values {
			if v == item {
				delete(m, k)
				break
			}
		}
	}
}

// DeleteFilter 遍历map，对每个元素调用 iteratee 函数，如果 iteratee 返回true，则删除该元素。
func DeleteFilter[K comparable, V any](m map[K]V, iteratee func(key K, value V) bool) {
	for k, v := range m {
		if iteratee(k, v) {
			delete(m, k)
		}
	}
}

// ReplaceValue 替换所有value等于 old 的元素。
func ReplaceValue[K, V comparable](m map[K]V, old, new V) {
	for k, v := range m {
		if v == old {
			m[k] = new
		}
	}
}

// MapToSliceBy map转切片，遍历map，对每个元素调用 iteratee 函数，并返回调用后结果切片。
func MapToSliceBy[K comparable, V any, R any](m map[K]V, iteratee func(key K, value V) R) []R {
	result := make([]R, 0, len(m))

	for k, v := range m {
		result = append(result, iteratee(k, v))
	}

	return result
}

// MapToSliceFilter map转切片，遍历map，对每个元素调用 iteratee 函数，如果 iteratee 返回true，则将该元素添加到结果切片中。
func MapToSliceFilter[K comparable, V any, R any](m map[K]V, iteratee func(key K, value V) (R, bool)) []R {
	result := make([]R, 0, len(m))

	for k, v := range m {
		r, ok := iteratee(k, v)
		if ok {
			result = append(result, r)
		}
	}

	return result
}
