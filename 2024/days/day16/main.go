package day16

import (
	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 16)

const (
	wall  = '#'
	start = 'S'
	end   = 'E'
	empty = '.'
)

var grid = map[[2]int]rune{}
var reindeer = [2]int{}
var goal = [2]int{}

func init() {
	for i, line := range input {
		for j, cell := range line {
			if cell == empty {
				continue
			}

			pos := [2]int{i, j}
			if cell == start {
				reindeer = pos
				continue
			}

			if cell == end {
				goal = pos
			}

			grid[pos] = cell
		}
	}
}

func run(pos, dir [2]int, point int, path map[[2]int]bool, points map[[2]int]int) {
	if prev, ok := points[pos]; grid[pos] == wall || path[pos] || ok && point >= prev {
		return
	}

	points[pos] = point

	if grid[pos] == end {
		return
	}

	path[pos] = true

	for _, d := range util.Directions {
		nextPoint := point + 1
		if dir != d {
			nextPoint += 1000
			if dir[0] == -d[0] && dir[1] == -d[1] {
				nextPoint += 1000
			}
		}
		run([2]int{pos[0] + d[0], pos[1] + d[1]}, d, nextPoint, path, points)
	}

	path[pos] = false
}

func Part1() int {
	points := map[[2]int]int{}
	run(reindeer, util.Directions[1], 0, map[[2]int]bool{}, points)

	return points[goal]
}

func getPaths(pos [2]int, prevP int, points map[[2]int]int, paths map[[2]int]bool) {
	point, ok := points[pos]
	if grid[pos] == wall || paths[pos] || ok && point >= prevP {
		return
	}

	paths[pos] = true

	if grid[pos] == start {
		return
	}

	for _, dir := range util.Directions {
		getPaths([2]int{pos[0] + dir[0], pos[1] + dir[1]}, point, points, paths)
	}
}

func Part2() int {
	points := map[[2]int]int{}
	run(reindeer, util.Directions[1], 0, map[[2]int]bool{}, points)

	paths := map[[2]int]bool{}
	getPaths(goal, points[goal]+1, points, paths)

	return len(paths)
}
