package days

import (
	"github.com/eigu47/aoc2023/util"
)

const CHAR = "XMAS"

var DIR = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func Day4_1() int {
	input := util.GetInput(2024, 4)
	res := 0

	var checkXmas func(x, y, dir, idx int) bool
	checkXmas = func(x, y, dir, idx int) bool {
		if idx >= len(CHAR) {
			return true
		}

		if x < 0 || x >= len(input) || y < 0 || y >= len(input[x]) || input[x][y] != CHAR[idx] {
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
