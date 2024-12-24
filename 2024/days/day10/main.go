package day10

import (
	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 10)
var grid = make([][]int, len(input))

var directions = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func init() {
	for i, line := range input {
		grid[i] = make([]int, len(line))
		for j, cell := range line {
			grid[i][j] = int(cell - '0')
		}
	}
}

func run(pos [2]int, peaks map[[2]int]bool) {
	if grid[pos[0]][pos[1]] == 9 {
		peaks[pos] = true
		return
	}

	for _, dir := range directions {
		next := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if util.IsInBounds(next, input) && grid[pos[0]][pos[1]]+1 == grid[next[0]][next[1]] {
			run(next, peaks)
		}
	}
}

func Part1() int {
	res := 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == 0 {
				peaks := map[[2]int]bool{}
				run([2]int{i, j}, peaks)
				res += len(peaks)
			}
		}
	}

	return res
}
