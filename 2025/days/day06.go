package days

import (
	"strconv"
	"strings"
)

func Day06_1(input []string) int {
	// input = test_06
	ans := 0

	problems := make([][]int, len(strings.Fields(input[0])))

	for _, row := range input[:len(input)-1] {
		for idx, num := range strings.Fields(row) {
			n, _ := strconv.Atoi(num)
			problems[idx] = append(problems[idx], n)
		}
	}

	for idx, operation := range strings.Fields(input[len(input)-1]) {
		rowAns := problems[idx][0]
		for _, num := range problems[idx][1:] {
			switch operation {
			case "+":
				rowAns += num
			case "*":
				rowAns *= num
			}
		}

		ans += rowAns
	}

	return ans
}

var test_06 = strings.Split(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `, "\n")
