package day12

import (
	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 12)
var grid = map[[2]int]rune{}

func init() {
	for i, line := range input {
		for j, char := range line {
			grid[[2]int{i, j}] = char
		}
	}
}

func visit(pos [2]int, char rune, visited map[[2]int]bool) (int, int) {
	if _, ok := grid[pos]; !ok || char != grid[pos] {
		return 0, 1
	}

	if visited[pos] {
		return 0, 0
	}

	area := 1
	perim := 0
	visited[pos] = true

	for _, dir := range util.Directions {
		newArea, newPerim := visit([2]int{pos[0] + dir[0], pos[1] + dir[1]}, char, visited)
		area += newArea
		perim += newPerim
	}

	return area, perim
}

func Part1() int {
	res := 0
	visited := map[[2]int]bool{}

	for i, line := range input {
		for j, char := range line {
			pos := [2]int{i, j}
			if visited[pos] {
				continue
			}

			area, perim := visit(pos, char, visited)
			res += area * perim
		}
	}

	return res
}

func visit2(pos [2]int, char rune, visited map[[2]int]bool) (int, int) {
	if _, ok := grid[pos]; !ok || char != grid[pos] {
		return 0, 0
	}

	if visited[pos] {
		return 0, 0
	}

	area := 1
	sides := 0
	visited[pos] = true

	for _, dir := range util.Diagonals {
		diag := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if grid[[2]int{diag[0], pos[1]}] != char &&
			grid[[2]int{pos[0], diag[1]}] != char {
			sides++
		}

		if grid[[2]int{diag[0], pos[1]}] == char &&
			grid[[2]int{pos[0], diag[1]}] == char &&
			grid[diag] != char {
			sides++
		}
	}

	for _, dir := range util.Directions {
		newArea, newSides := visit2([2]int{pos[0] + dir[0], pos[1] + dir[1]}, char, visited)
		area += newArea
		sides += newSides
	}

	return area, sides
}

func Part2() int {
	res := 0
	visited := map[[2]int]bool{}

	for i, line := range input {
		for j, char := range line {
			pos := [2]int{i, j}
			if visited[pos] {
				continue
			}

			area, sides := visit2(pos, char, visited)
			res += area * sides
		}
	}

	return res
}
