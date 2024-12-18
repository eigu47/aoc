package day2

import (
	"strconv"
	"strings"

	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 2)

func Part1() int {
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

			d := diff(lastReport, report)
			if !(inc || dec) || (d < 1 || d > 3) {
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

func Part2() int {
	res := 0

	for _, line := range input {
		level := strings.Fields(line)

		isSafe := true
		lastReport, _ := strconv.Atoi(level[0])
		inc := true
		dec := true
		bad := 0

		for i := 1; i < len(level); i++ {
			report, _ := strconv.Atoi(level[i])
			if report <= lastReport {
				inc = false
			} else if report >= lastReport {
				dec = false
			}

			d := diff(lastReport, report)
			if !(inc || dec) || (d < 1 || d > 3) {
				if bad > 0 {
					isSafe = false
					break
				}
				bad++
				inc = true
				dec = true
			} else {
				lastReport = report
			}
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
