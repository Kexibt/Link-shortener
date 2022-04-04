package url

import (
	"math"
)

const (
	st = 63
)

func ConvertToShort(num int) string {
	array := []int{}

	if num == 0 {
		return "0"
	}
	for i := num; i != 0; i /= st {
		array = append(array, (i%st)+48)
	}

	shortLink := ""
	for i := 0; i < len(array); i++ {
		numb := array[i]
		if numb > 57 {
			numb += 7
		}
		if numb > 90 {
			numb += 4
		}
		if numb > 95 {
			numb += 1
		}

		shortLink = string(rune(numb)) + shortLink
	}

	for i := len(shortLink); i <= 10; i++ {
		shortLink = "0" + shortLink
	}

	return shortLink
}

func ConvertToID(shortLink string) int {
	res := 0
	size := len(shortLink)

	for i, r := range shortLink {
		num := int(rune(r))
		if num > 95 {
			num -= 1
		}
		if num > 90 {
			num -= 4
		}
		if num > 57 {
			num -= 7
		}

		num -= 48

		res += num * int(math.Pow(float64(st), float64(size-i-1)))
	}

	return res
}
