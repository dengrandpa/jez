package randomutil

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	res := Random(Numeral, 11)

	reg := regexp.MustCompile(`^[0-9]+$`)

	ass.True(reg.MatchString(res))
	ass.Len(res, 11)
	ass.Panics(func() { Random(Numeral, 0) })
	ass.Panics(func() { Random(Numeral, 11, true) })
	ass.True(reg.MatchString(Random(Numeral, 5, true)))

}

func TestRandomLower(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	res := RandomLower(27)

	reg := regexp.MustCompile(`^[a-z]+$`)

	ass.True(reg.MatchString(res))
	ass.Len(res, 27)
	ass.Panics(func() { RandomLower(0) })
	ass.Panics(func() { RandomLower(27, true) })
}

func TestRandomUpper(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	res := RandomUpper(27)

	reg := regexp.MustCompile(`^[A-Z]+$`)

	ass.True(reg.MatchString(res))
	ass.Len(res, 27)
	ass.Panics(func() { RandomUpper(0) })
	ass.Panics(func() { RandomUpper(27, true) })
}

func TestRandomNumeral(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	res := RandomNumeral(11)

	reg := regexp.MustCompile(`^[0-9]+$`)

	ass.True(reg.MatchString(res))
	ass.Len(res, 11)
	ass.Panics(func() { RandomNumeral(0) })
	ass.Panics(func() { RandomNumeral(11, true) })
}

func TestRandomCaseLetters(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	res := RandomCaseLetters(53)

	reg := regexp.MustCompile(`^[A-Za-z]+$`)

	ass.True(reg.MatchString(res))
	ass.Len(res, 53)
	ass.Panics(func() { RandomCaseLetters(0) })
	ass.Panics(func() { RandomCaseLetters(53, true) })
}

func TestRandomLowerNumeral(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	res := RandomLowerNumeral(37)

	reg := regexp.MustCompile(`^[a-z0-9]+$`)

	ass.True(reg.MatchString(res))
	ass.Len(res, 37)
	ass.Panics(func() { RandomLowerNumeral(0) })
	ass.Panics(func() { RandomLowerNumeral(37, true) })
}

func TestRandomUpperNumeral(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	res := RandomUpperNumeral(37)

	reg := regexp.MustCompile(`^[A-Z0-9]+$`)

	ass.True(reg.MatchString(res))
	ass.Len(res, 37)
	ass.Panics(func() { RandomUpperNumeral(0) })
	ass.Panics(func() { RandomUpperNumeral(37, true) })
}

func TestRandomCharset(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	res := RandomCharset(63)

	reg := regexp.MustCompile(`^[A-Za-z0-9]+$`)

	ass.True(reg.MatchString(res))
	ass.Len(res, 63)
	ass.Panics(func() { RandomUpperNumeral(0) })
	ass.Panics(func() { RandomUpperNumeral(63, true) })
}

func TestRandomInt(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	min := 10
	max := 15

	res := RandomInt(min, max)

	ass.Less(res, max)
	ass.GreaterOrEqual(res, min)

	res1 := RandomInt(max, min)
	ass.Less(res1, max)
	ass.GreaterOrEqual(res1, min)

	ass.Equal(RandomInt(min, min), 10)
}

func testRandomIntSlice(ass *assert.Assertions, min, max int) {
	res := RandomIntSlice(min, max, 10)

	if min > max {
		max, min = min, max
	}

	ok := true
	for _, r := range res {
		if r < min || r >= max {
			ok = false
			break
		}
	}

	ass.True(ok)
	ass.Len(res, 10)
}

func TestRandomIntSlice(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	min := 10
	max := 15

	ass.Panics(func() { RandomIntSlice(min, max, 0) })

	testRandomIntSlice(ass, max, min)

	testRandomIntSlice(ass, min, max)

	ass.Equal(RandomIntSlice(min, min, 2), []int{min, min})
}

func testRandomIntSliceUnique(ass *assert.Assertions, min, max int) {
	res := RandomIntSliceUnique(min, max, 5)

	if min > max {
		max, min = min, max
	}

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

	ass.True(ok)
	ass.Len(res, 5)
}

func TestRandomIntSliceUnique(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	min := 10
	max := 15

	ass.Panics(func() { RandomIntSliceUnique(min, max, 0) })
	ass.Panics(func() { RandomIntSliceUnique(min, max, 10) })

	testRandomIntSliceUnique(ass, max, min)
	testRandomIntSliceUnique(ass, min, max)

	ass.Equal(RandomIntSliceUnique(min, min, 1), []int{10})
}

func TestRandomBytes(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ass.Panics(func() { RandomBytes(0) })

	res := RandomBytes(10)

	ass.Len(res, 10)
}

func TestUUIDv4(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	uuid := UUIDv4()

	reg := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)

	ass.True(reg.MatchString(uuid))
	ass.Len(uuid, 36)
}
