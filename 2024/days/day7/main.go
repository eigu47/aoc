package day7

import (
	"fmt"
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

func test1(nums *[]int, sum int, idx int, res int) bool {
	if idx == len(*nums) {
		return sum == res
	}

	return test1(nums, sum+(*nums)[idx], idx+1, res) || test1(nums, sum*(*nums)[idx], idx+1, res)
}

func Part1() int {
	totalRes := 0

	for i, res := range results {
		if test1(&numbers[i], 0, 0, res) {
			totalRes += res
		}
	}

	return totalRes
}

func test2(nums *[]int, sum int, idx int, res int) bool {
	if idx == len(*nums) {
		return sum == res
	}

	or, _ := strconv.Atoi(fmt.Sprintf("%d%d", sum, (*nums)[idx]))
	return test2(nums, sum+(*nums)[idx], idx+1, res) || test2(nums, sum*(*nums)[idx], idx+1, res) || test2(nums, or, idx+1, res)
}

func Part2() int {
	totalRes := 0

	for i, res := range results {
		if test2(&numbers[i], 0, 0, res) {
			totalRes += res
		}
	}

	return totalRes
}
