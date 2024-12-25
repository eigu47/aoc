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
