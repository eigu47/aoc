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
	{1, -1},
}

func Day4_1() int {
	input := util.GetInput(2024, 4)
	res := 0

	var checkXmas func(x, y, dir, char int) bool
	checkXmas = func(x, y, dir, char int) bool {
		if char >= len(CHAR) {
			return true
		}

		if x < 0 || x >= len(input) || y < 0 || y >= len(input[x]) {
			return false
		}

		if input[x][y] == CHAR[char] {
			return checkXmas(x+DIR[dir][0], y+DIR[dir][1], dir, char+1)
		}

		return false
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
