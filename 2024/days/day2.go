package days

import (
	"fmt"
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
		
		for i := 1; i < len(level); i++ {
			inc := true

			

			report, _ := strconv.Atoi(level[i])
			d := diff(lastReport, report)

			if d < 1 || d > 4 {
				isSafe = false
				fmt.Printf("%d %d", lastReport, report)
				continue
			}
		}

		if isSafe {
			res++
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
