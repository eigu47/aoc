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

	// fmt.Printf("problems %+v\n", problems)

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

func Day06_2(input []string) int {
	// input = test_06
	ans := 0

	operations := input[len(input)-1]
	for start, operation := range operations {
		if operation == ' ' {
			continue
		}

		colAns := 0
		for i := start; i < len(operations) && !(i+1 < len(operations) && operations[i+1] != ' '); i++ {
			num := 0
			for _, row := range input[:len(input)-1] {
				if row[i] != ' ' {
					num = num*10 + int(row[i]-'0')
				}
			}

			switch operation {
			case '+':
				colAns += num
			case '*':
				if colAns == 0 {
					colAns = 1
				}
				colAns *= num
			}
		}

		ans += colAns
	}

	return ans
}

var test_06 = strings.Split(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `, "\n")
