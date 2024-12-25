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

func run(pos [2]int, height int, peaks map[[2]int]bool) int {
	if !util.IsInBounds(pos, input) || grid[pos[0]][pos[1]] != height {
		return 0
	}

	if grid[pos[0]][pos[1]] == 9 {
		peaks[pos] = true
		return 0
	}

	for _, dir := range directions {
		run([2]int{pos[0] + dir[0], pos[1] + dir[1]}, height+1, peaks)
	}

	return len(peaks)
}

func Part1() int {
	res := 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == 0 {
				res += run([2]int{i, j}, 0, map[[2]int]bool{})
			}
		}
	}

	return res
}

func run2(pos [2]int, height int) int {
	if !util.IsInBounds(pos, input) || grid[pos[0]][pos[1]] != height {
		return 0
	}

	if grid[pos[0]][pos[1]] == 9 {
		return 1
	}

	trails := 0
	for _, dir := range directions {
		trails += run2([2]int{pos[0] + dir[0], pos[1] + dir[1]}, height+1)
	}

	return trails
}

func Part2() int {
	res := 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == 0 {
				res += run2([2]int{i, j}, 0)
			}
		}
	}

	return res
}
