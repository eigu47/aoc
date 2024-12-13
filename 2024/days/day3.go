package days

import (
	"regexp"
	"strconv"

	"github.com/eigu47/aoc2023/util"
)

func Day3_1() int {
	input := util.GetInput(2024, 3)

	res := 0

	for _, line := range input {
		for _, mult := range regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatch(line, -1) {
			x, _ := strconv.Atoi(mult[1])
			y, _ := strconv.Atoi(mult[2])
			res += x * y
		}
	}

	return res
}
