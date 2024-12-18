package day1

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 1)

func Part1() int {
	var right, left []int

	for _, line := range input {
		parts := strings.Fields(line)

		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	var res int

	for i := 0; i < len(left); i++ {
		res += int(math.Abs(float64(left[i] - right[i])))
	}

	return res
}

func Part2() int {
	var res int
	left := make(map[int]int)
	right := make(map[int]int)

	for _, line := range input {
		parts := strings.Fields(line)

		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])

		left[l]++
		right[r]++
	}

	for id, times := range left {
		res += (id * right[id]) * times
	}

	return res
}
