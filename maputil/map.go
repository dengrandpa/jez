// Package maputil provides some useful methods for the map.
package maputil

// ForEach iterates over the map and calls the iteratee for each key/value pair.
func ForEach[K comparable, V any](m map[K]V, iteratee func(key K, value V)) {
	for k, v := range m {
		iteratee(k, v)
	}
}

// Filter iterates over the map and calls the iteratee for each key/value pair.
func Filter[K comparable, V any](m map[K]V, iteratee func(key K, value V) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if iteratee(k, v) {
			result[k] = v
		}
	}
	return result
}

// Keys returns all keys of the map.
func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))

	for k := range m {
		result = append(result, k)
	}

	return result
}

// KeysBy returns all keys of the map after applying the iteratee to each key.
func KeysBy[K comparable, V any](m map[K]V, iteratee func(item K) K) []K {
	result := make([]K, 0, len(m))

	for k := range m {
		result = append(result, iteratee(k))
	}

	return result
}

// Values returns all values of the map.
func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

// ValuesBy returns all values of the map after applying the iteratee to each value.
func ValuesBy[K comparable, V any](m map[K]V, iteratee func(item V) V) []V {
	result := make([]V, 0, len(m))

	for _, v := range m {
		result = append(result, iteratee(v))
	}

	return result
}

// ValuesUnique returns all unique values of the map.
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

// KeysAndValues returns all keys and values of the map.
func KeysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))

	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}

// KeysAndValuesFilter returns all keys and values of the map after applying the iteratee to each key/value pair.
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

// DeleteByKeys deletes all keys from the map.
func DeleteByKeys[K comparable, V any](m map[K]V, keys ...K) {
	for _, k := range keys {
		delete(m, k)
	}
}

// DeleteByValues deletes all values from the map.
// Note: This method will delete all values that satisfy the condition.
func DeleteByValues[K, V comparable](m map[K]V, key ...V) {
	for k, v := range m {
		for _, item := range key {
			if v == item {
				delete(m, k)
				break
			}
		}
	}
}

// DeleteByValue deletes the first value from the map.
// Note: This method will delete all values that satisfy the condition.
func DeleteByValue[K, V comparable](m map[K]V, item V) {
	for k, v := range m {
		if v == item {
			delete(m, k)
		}
	}
}

// DeleteFilter deletes all key/value pairs from the map after applying the iteratee to each key/value pair.
func DeleteFilter[K comparable, V any](m map[K]V, iteratee func(key K, value V) bool) {
	for k, v := range m {
		if iteratee(k, v) {
			delete(m, k)
		}
	}
}

// ReplaceValue replaces all values in the map.
func ReplaceValue[K, V comparable](m map[K]V, old, new V) {
	for k, v := range m {
		if v == old {
			m[k] = new
		}
	}
}

// MapToSliceBy returns a slice of the results of applying the iteratee to each key/value pair.
func MapToSliceBy[K comparable, V any, R any](m map[K]V, iteratee func(key K, value V) R) []R {
	result := make([]R, 0, len(m))

	for k, v := range m {
		result = append(result, iteratee(k, v))
	}

	return result
}

// MapToSliceFilter returns a slice of the results of applying the iteratee to each key/value pair.
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
