package randomutil

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Numeral          = "0123456789"
	LowerCaseLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperCaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random 随机生成字符串
//
// 参数：
//   - s: 字符集合，用于生成随机字符串
//   - n: 生成的随机字符串的长度
//   - unique: 可选参数，指定是否生成不重复的字符串，默认为 false
//
// 返回值：
//   - 生成的随机字符串
//
// 注意事项：
//   - 当 n < 1 时，会 panic
//   - 即使 unique 为 true 时，但如果传入的字符串有重复元素，那么生成的字符串可能也会重复
//   - 当 unique 为 true 时，如果 n 大于 s 的长度，那么会 panic
func Random(s string, n int, unique ...bool) string {
	if n < 1 {
		panic("n must be greater than 0")
	}

	b := make([]byte, n)

	size := len(s)

	if len(unique) > 0 && unique[0] {
		if n > size {
			panic("n must be less than or equal to the length of s")
		}

		indices := rand.Perm(size)

		if len(indices)-n > 0 {
			indices = indices[:n]
		}

		for i, idx := range indices {
			b[i] = s[idx]
		}
	} else {
		for i := range b {
			b[i] = s[rand.Intn(size)]
		}
	}

	return string(b)
}

// RandomLower 随机生成小写字母字符串
//
// 注意事项：
//   - 当 n < 1 时，会 panic
//   - 当 unique 为 true 时，如果 n 大于 26（26个字母），那么会 panic
func RandomLower(n int, unique ...bool) string {
	return Random(LowerCaseLetters, n, unique...)
}

// RandomUpper 随机生成大写字母字符串
//
// 注意事项：
//   - 当 n < 1 时，会 panic
//   - 当 unique 为 true 时，如果 n 大于 26（26个字母），那么会 panic
func RandomUpper(n int, unique ...bool) string {
	return Random(UpperCaseLetters, n, unique...)
}

// RandomNumeral 随机生成数字字符串
//
// 注意事项：
//   - 当 n < 1 时，会 panic
//   - 当 unique 为 true 时，如果 n 大于 10（10个数字），那么会 panic
func RandomNumeral(n int, unique ...bool) string {
	return Random(Numeral, n, unique...)
}

// RandomCaseLetters 随机生成大小写字母字符串
//
// 注意事项：
//   - 当 n < 1 时，会 panic
//   - 当 unique 为 true 时，如果 n 大于 52（26个大写字母+26个小写字母），那么会 panic
func RandomCaseLetters(n int, unique ...bool) string {
	return Random(LowerCaseLetters+UpperCaseLetters, n, unique...)
}

// RandomLowerNumeral 随机生成小写字母和数字字符串
//
// 注意事项：
//   - 当 n < 1 时，会 panic
//   - 当 unique 为 true 时，如果 n 大于 36（10个数字+26个字母），那么会 panic
func RandomLowerNumeral(n int, unique ...bool) string {
	return Random(LowerCaseLetters+Numeral, n, unique...)
}

// RandomUpperNumeral 随机生成大写字母和数字字符串
//
// 注意事项：
//   - 当 n < 1 时，会 panic
//   - 当 unique 为 true 时，如果 n 大于 36（10个数字+26个字母），那么会 panic
func RandomUpperNumeral(n int, unique ...bool) string {
	return Random(UpperCaseLetters+Numeral, n, unique...)
}

// RandomCharset 随机生成字符串，包含数字、大小写字母
//
// 注意事项：
//   - 当 n < 1 时，会 panic
//   - 当 unique 为 true 时，如果 n 大于 62（10个数字+26个大写字母+26个小写字母），那么会 panic
func RandomCharset(n int, unique ...bool) string {
	return Random(Numeral+LowerCaseLetters+UpperCaseLetters, n, unique...)
}

// RandomInt 随机生成整数，包含 min，不包含 max，即 [min,max)
//
// 注意事项：
//   - 当 min > max 时，会交换 min 和 max 的值
func RandomInt(min, max int) int {
	if min == max {
		return min
	}

	if max < min {
		min, max = max, min
	}

	return rand.Intn(max-min) + min
}

// RandomIntSlice 随机生成整数切片
//
// 注意事项：
//   - 当 min > max 时，会交换 min 和 max 的值
//   - 当 n < 1 时，会 panic
func RandomIntSlice(min, max, n int) []int {
	if n < 1 {
		panic("n must be greater than 0")
	}

	s := make([]int, n)

	sub := max - min

	if sub == 0 {
		for i := range s {
			s[i] = min
		}

		return s
	}

	if sub < 0 {
		min = max
		sub = -sub
	}

	for i := range s {
		s[i] = rand.Intn(sub) + min
	}

	return s
}

// RandomIntSliceUnique 随机生成不重复的整数切片
//
// 注意事项：
//   - 当 min > max 时，会交换 min 和 max 的值
//   - 当 n < 1 时，会 panic
//   - 当 n 大于 max-min+1 时，会 panic
func RandomIntSliceUnique(min, max, n int) []int {
	if n < 1 {
		panic("n must be greater than 0")
	}

	if min > max {
		min, max = max, min
	}

	sub := max - min + 1

	if n > sub {
		panic(fmt.Sprintf("n must be less than or equal to %d", sub))
	}

	if sub == 1 {
		return []int{min}
	}

	s := make([]int, n)

	for i := 0; i < n; i++ {
		s[i] = min + i
	}

	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// RandomBytes 随机生成字节切片
//
// 注意事项：
//   - 当 n < 1 时，会 panic
func RandomBytes(n int) []byte {
	if n < 1 {
		panic("n must be greater than 0")
	}

	b := make([]byte, n)

	_, _ = rand.Read(b)

	return b
}

// UUIDv4 根据 RFC4122 生成 UUID v4版本
//
// 注意事项：
//   - 这种方法生成的 UUID v4 可能不是完全符合标准，但在大多数情况下应该是足够的。
func UUIDv4() string {
	uuid := make([]byte, 16)
	_, _ = rand.Read(uuid)

	// 设置 version 为 4
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	// 设置 variant 为 RFC4122
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

// Shuffle 打乱切片中元素的顺序。
func Shuffle[T any](list []T) []T {
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	return list
}

// Sample 从切片中随机返回一个元素。
func Sample[T any](list []T) T {
	var zero T

	if len(list) == 0 {
		return zero
	}

	return list[rand.Intn(len(list))]
}

// Samples 从切片中随机返回n个元素，结果不去重。
//
// 注意事项：
//   - 如果 n 大于切片的长度，则返回打乱顺序后的切片，如果 n 小于等于0，则返回空切片。
func Samples[T any](list []T, n int) []T {
	if n <= 0 {
		return []T{}
	}

	newList := make([]T, len(list))
	copy(newList, list)

	newList = Shuffle(newList)

	if n > len(newList) {
		return newList
	}

	return newList[:n]
}
