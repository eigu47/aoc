package day11

import (
	"math"
	"strconv"
	"strings"

	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 11)
var runes = []int{}

func init() {
	str := strings.Fields(input[0])
	runes = make([]int, len(str))

	for i, rune := range str {
		num, _ := strconv.Atoi(rune)
		runes[i] = num
	}
}

func blink(runes []int, cur, total int) int {
	if cur == total {
		return len(runes)
	}

	tmp := []int{}
	for _, rune := range runes {
		if rune == 0 {
			tmp = append(tmp, 1)
			continue
		}

		if digits := int(math.Log10(float64(rune))) + 1; digits%2 == 0 {
			half := int(math.Pow(10, float64(digits/2)))
			tmp = append(tmp, rune/half, rune%half)
			continue
		}

		tmp = append(tmp, rune*2024)
	}

	return blink(tmp, cur+1, total)
}

func Part1() int {
	return blink(runes, 0, 25)
}

func blink2(rune int, cur, total int, cache map[[2]int]int) int {
	pos := [2]int{rune, cur}
	if val, ok := cache[pos]; ok {
		return val
	}

	if cur == total {
		cache[pos] = 1
		return 1
	}

	if rune == 0 {
		cache[pos] = blink2(1, cur+1, total, cache)
		return cache[pos]
	}

	if digits := int(math.Log10(float64(rune))) + 1; digits%2 == 0 {
		half := int(math.Pow(10, float64(digits/2)))
		cache[pos] = blink2(rune/half, cur+1, total, cache) + blink2(rune%half, cur+1, total, cache)
		return cache[pos]
	}

	cache[pos] = blink2(rune*2024, cur+1, total, cache)
	return cache[pos]
}

func Part2() int {
	res := 0
	cache := map[[2]int]int{}

	for _, rune := range runes {
		res += blink2(rune, 0, 75, cache)
	}

	return res
}
