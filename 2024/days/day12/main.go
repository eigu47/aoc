package day12

import (
	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 12)
var grid = map[[2]int]rune{}

type plant struct {
	Area  map[[2]int]bool
	Perim int
}

func init() {
	for i, line := range input {
		for j, char := range line {
			grid[[2]int{i, j}] = char
		}
	}
}

func visit(pos [2]int, char rune, plant *plant, visited map[[2]int]bool) *plant {
	if _, ok := grid[pos]; !ok || char != grid[pos] {
		plant.Perim++
		return nil
	}

	if visited[pos] {
		return nil
	}

	plant.Area[pos] = true
	visited[pos] = true

	for _, dir := range util.Directions {
		visit([2]int{pos[0] + dir[0], pos[1] + dir[1]}, char, plant, visited)
	}

	return plant
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

			region := visit(pos, char, &plant{map[[2]int]bool{}, 0}, visited)
			res += len(region.Area) * region.Perim
		}
	}

	return res
}
