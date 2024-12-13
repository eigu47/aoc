package days

import (
	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 4)

func Day4_1() int {
	STR := "XMAS"
	DIR := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	res := 0

	var checkXmas func(x, y, dir, idx int) bool
	checkXmas = func(x, y, dir, idx int) bool {
		if idx >= len(STR) {
			return true
		}

		if x < 0 || x >= len(input) || y < 0 || y >= len(input[x]) || input[x][y] != STR[idx] {
			return false
		}

		return checkXmas(x+DIR[dir][0], y+DIR[dir][1], dir, idx+1)
	}

	for x, row := range input {
		for y := range row {
			for dir := range DIR {
				if checkXmas(x, y, dir, 0) {
					res++
				}
			}
		}
	}

	return res
}

func Day4_2() int {
	DIR := [2][2][2]int{
		{{-1, -1}, {1, 1}},
		{{-1, 1}, {1, -1}},
	}

	res := 0

	checkXmas := func(x, y int, dir [2][2]int) bool {
		if x-1 < 0 || x+1 >= len(input) || y-1 < 0 || y+1 >= len(input[x]) {
			return false
		}

		a, b := rune(input[x+dir[0][0]][y+dir[0][1]]), rune(input[x+dir[1][0]][y+dir[1][1]])
		if (a == 'M' && b == 'S') || (a == 'S' && b == 'M') {
			return true
		}

		return false
	}

	for x, row := range input {
		for y, char := range row {
			if char == 'A' && checkXmas(x, y, DIR[0]) && checkXmas(x, y, DIR[1]) {
				res++
			}
		}
	}

	return res
}
