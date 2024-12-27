package day13

import (
	"regexp"
	"strconv"

	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 13)
var machines = []*machine{}

type machine struct {
	A     [2]int
	B     [2]int
	Price [2]int
}

func init() {
	rgxX := regexp.MustCompile(`X[+=](\d+)`)
	rgxY := regexp.MustCompile(`Y[+=](\d+)`)
	for i := 0; i < len(input); i += 4 {
		ax, _ := strconv.Atoi(rgxX.FindStringSubmatch(input[i])[1])
		ay, _ := strconv.Atoi(rgxY.FindStringSubmatch(input[i])[1])
		bx, _ := strconv.Atoi(rgxX.FindStringSubmatch(input[i+1])[1])
		by, _ := strconv.Atoi(rgxY.FindStringSubmatch(input[i+1])[1])
		px, _ := strconv.Atoi(rgxX.FindStringSubmatch(input[i+2])[1])
		py, _ := strconv.Atoi(rgxY.FindStringSubmatch(input[i+2])[1])

		machines = append(machines, &machine{
			A:     [2]int{ax, ay},
			B:     [2]int{bx, by},
			Price: [2]int{px, py},
		})
	}
}

func Part1() int {
	res := 0

	for _, m := range machines {
		for a := range 100 {
			found := false
			for b := range 100 {
				if m.A[0]*a+m.B[0]*b == m.Price[0] && m.A[1]*a+m.B[1]*b == m.Price[1] {
					res += 3*a + b
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}

	return res
}

func Part2() int {
	res := 0

	for _, m := range machines {
		m.Price[0] += 10000000000000
		m.Price[1] += 10000000000000
		a := (m.Price[0]*m.B[1] - m.Price[1]*m.B[0]) / (m.A[0]*m.B[1] - m.A[1]*m.B[0])
		b := (m.Price[1]*m.A[0] - m.Price[0]*m.A[1]) / (m.A[0]*m.B[1] - m.A[1]*m.B[0])

		if m.A[0]*a+m.B[0]*b == m.Price[0] && m.A[1]*a+m.B[1]*b == m.Price[1] {
			res += 3*a + b
		}
	}

	return res
}
