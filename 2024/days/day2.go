package days

import (
	"strconv"
	"strings"

	"github.com/eigu47/aoc2023/util"
)

func Day2_1() int {
	input := util.GetInput(2024, 2)

	res := 0

	for _, line := range input {
		level := strings.Fields(line)

		isSafe := true
		lastReport, _ := strconv.Atoi(level[0])
		inc := true
		dec := true

		for i := 1; i < len(level); i++ {
			report, _ := strconv.Atoi(level[i])
			if report <= lastReport {
				inc = false
			} else if report >= lastReport {
				dec = false
			}

			if !(inc || dec) {
				isSafe = false
				break
			}

			d := diff(lastReport, report)
			if d < 1 || d > 3 {
				isSafe = false
				// fmt.Printf("%d %d", lastReport, report)
				break
			}

			lastReport = report
		}

		if isSafe {
			res++
			// fmt.Println(line)
		}
	}

	return res
}

func diff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
