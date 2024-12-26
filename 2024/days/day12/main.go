package day12

import (
	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 12)
var grid = map[[2]int]rune{}

// type plant struct {
// 	Char   rune
// 	Area   int
// 	Perim  int
// 	Border map[[2]int]bool
// }

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

// func Part2() int {
// 	res := 0

// 	visited := map[[2]int]bool{}

// 	for i, line := range input {
// 		for j, char := range line {
// 			pos := [2]int{i, j}
// 			if visited[pos] {
// 				continue
// 			}

// 			region := visit(pos, &plant{
// 				Char:   char,
// 				Area:   map[[2]int]bool{},
// 				Perim:  0,
// 				Border: map[[2]int]bool{},
// 			}, visited)
// 			// res += len(region.Area) * region.Perim

// 		}
// 	}

// 	return res
// }
