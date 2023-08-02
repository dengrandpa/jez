package randomutil

import (
	"fmt"
	"regexp"
)

func isPanic(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(true)
		}
	}()
	f()
	fmt.Println(false)
}

func ExampleRandom() {
	res := Random(Numeral, 11)

	reg := regexp.MustCompile(`^[0-9]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { Random(Numeral, 0) })
	isPanic(func() { Random(Numeral, 11, true) })

	// Output:
	// true
	// 11
	// true
	// true
}

func ExampleRandomLower() {
	res := RandomLower(27)

	reg := regexp.MustCompile(`^[a-z]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { RandomLower(0) })
	isPanic(func() { RandomLower(27, true) })

	// Output:
	// true
	// 27
	// true
	// true
}

func ExampleRandomUpper() {
	res := RandomUpper(27)

	reg := regexp.MustCompile(`^[A-Z]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { RandomUpper(0) })
	isPanic(func() { RandomUpper(27, true) })

	// Output:
	// true
	// 27
	// true
	// true
}

func ExampleRandomNumeral() {
	res := RandomNumeral(11)

	reg := regexp.MustCompile(`^[0-9]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { RandomNumeral(0) })
	isPanic(func() { RandomNumeral(11, true) })

	// Output:
	// true
	// 11
	// true
	// true
}

func ExampleRandomCaseLetters() {
	res := RandomCaseLetters(53)

	reg := regexp.MustCompile(`^[A-Za-z]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { RandomCaseLetters(0) })
	isPanic(func() { RandomCaseLetters(53, true) })

	// Output:
	// true
	// 53
	// true
	// true
}

func ExampleRandomLowerNumeral() {
	res := RandomLowerNumeral(37)

	reg := regexp.MustCompile(`^[a-z0-9]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { RandomLowerNumeral(0) })
	isPanic(func() { RandomLowerNumeral(37, true) })

	// Output:
	// true
	// 37
	// true
	// true
}

func ExampleRandomUpperNumeral() {
	res := RandomUpperNumeral(37)

	reg := regexp.MustCompile(`^[A-Z0-9]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { RandomUpperNumeral(0) })
	isPanic(func() { RandomUpperNumeral(37, true) })

	// Output:
	// true
	// 37
	// true
	// true
}

func ExampleRandomCharset() {
	res := RandomCharset(63)

	reg := regexp.MustCompile(`^[A-Za-z0-9]+$`)

	fmt.Println(reg.MatchString(res))
	fmt.Println(len(res))
	isPanic(func() { RandomUpperNumeral(0) })
	isPanic(func() { RandomUpperNumeral(63, true) })

	// Output:
	// true
	// 63
	// true
	// true
}

func ExampleRandomInt() {
	min := 10
	max := 15

	res := RandomInt(min, max)

	if res < min || res >= max {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}

	// Output:
	// true
}

func ExampleRandomIntSlice() {
	min := 10
	max := 15

	res := RandomIntSlice(min, max, 10)

	ok := true

	for _, r := range res {
		if r < min || r >= max {
			ok = false
			break
		}
	}

	fmt.Println(ok)
	fmt.Println(len(res))
	isPanic(func() { RandomIntSlice(min, max, 0) })

	// Output:
	// true
	// 10
	// true
}

func ExampleRandomIntSliceUnique() {
	min := 10
	max := 15

	res := RandomIntSliceUnique(min, max, 5)

	var (
		ok    = true
		exist = make(map[int]struct{}, 5)
	)

	for _, r := range res {
		if r < min || r >= max {
			ok = false
			break
		}
		if _, ok1 := exist[r]; ok1 {
			ok = false
			break
		}
		exist[r] = struct{}{}
	}

	fmt.Println(ok)
	fmt.Println(len(res))

	isPanic(func() { RandomIntSliceUnique(min, max, 0) })
	isPanic(func() { RandomIntSliceUnique(min, max, 10) })

	// Output:
	// true
	// 5
	// true
	// true
}

func ExampleRandomBytes() {

	res := RandomBytes(10)

	fmt.Println(len(res))
	isPanic(func() { RandomBytes(0) })

	// Output:
	// 10
	// true
}

func ExampleUUIDv4() {
	uuid := UUIDv4()

	reg := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)

	fmt.Println(len(uuid))
	fmt.Println(reg.MatchString(uuid))

	// Output:
	// 36
	// true
}

func ExampleShuffle() {

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}

	list2 := Shuffle([]int{})

	fmt.Println(len(list) == len(Shuffle(list)))
	fmt.Println(list2)

	// Output:
	// true
	// []
}

func ExampleSample() {

	list := []int{1, 2, 3, 3, 5, 3, 5, 6}
	var list2 []string

	s := Sample(list)

	ok := false
	for _, v := range list {
		if v == s {
			ok = true
			break
		}
	}

	fmt.Println(ok)
	fmt.Println(Sample(list2) == "")

	// Output:
	// true
	// true
}

func ExampleSamples() {

	list := []int{1, 2, 3}

	s2 := Samples(list, 3)

	ok := true
	for _, v := range s2 {
		if v > 3 || v < 1 {
			ok = false
			break
		}
	}

	fmt.Println(ok)

	fmt.Println(Samples([]string{}, 3))

	// Output:
	// true
	// []
}
