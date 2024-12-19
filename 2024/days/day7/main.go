package day7

import (
	"strconv"
	"strings"

	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 7)
var results = []int{}
var numbers = [][]int{}

func init() {
	for _, line := range input {
		vals := strings.Split(line, ":")
		if len(vals) < 2 {
			break
		}
		num, _ := strconv.Atoi(vals[0])
		results = append(results, num)
		lft := []int{}
		for _, val := range strings.Fields(vals[1]) {
			num, _ := strconv.Atoi(val)
			lft = append(lft, num)
		}
		numbers = append(numbers, lft)
	}
}

func test(nums *[]int, sum int, idx int, res int) bool {
	if idx == len(*nums) {
		return sum == res
	}

	return test(nums, sum+(*nums)[idx], idx+1, res) || test(nums, sum*(*nums)[idx], idx+1, res)
}

func Part1() int {
	totalRes := 0

	for i, res := range results {
		if test(&numbers[i], 0, 0, res) {
			totalRes += res
		}
	}

	return totalRes
}
