package day15

import (
	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 15)

var gridSize = [2]int{0, 0}
var grid = map[[2]int]rune{}
var robotPos = [2]int{}
var movements = [][2]int{}

const (
	robot = '@'
	wall  = '#'
	box   = 'O'
	empty = '.'
)

var directions = map[rune][2]int{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

func init() {
	isGrid := true
	for i, line := range input {
		if line == "" {
			isGrid = false
			continue
		}

		for j, cell := range line {
			if isGrid {
				gridSize[0], gridSize[1] = max(gridSize[0], i), max(gridSize[1], i)
				if cell == '.' {
					continue
				}

				pos := [2]int{i, j}
				if cell == empty {
					continue
				}

				grid[pos] = cell
				if cell == robot {
					robotPos = pos
				}

				continue
			}

			if mov, ok := directions[cell]; ok {
				movements = append(movements, mov)
			}
		}
	}
}

func nextEmpty(pos [2]int, mov [2]int) ([2]int, bool) {
	if grid[pos] == wall {
		return pos, false
	}

	if grid[pos] == box {
		return nextEmpty([2]int{pos[0] + mov[0], pos[1] + mov[1]}, mov)
	}

	return pos, true
}

func Part1() int {
	res := 0

	for _, mov := range movements {
		next := [2]int{robotPos[0] + mov[0], robotPos[1] + mov[1]}

		if grid[next] == wall {
			continue
		}

		if grid[next] == box {
			empty, ok := nextEmpty(next, mov)
			if !ok {
				continue
			}

			grid[empty] = box
			delete(grid, next)
		}

		delete(grid, robotPos)
		robotPos = next
		grid[robotPos] = robot
	}

	for pos, cell := range grid {
		if cell == 'O' {
			res += pos[0]*100 + pos[1]
		}
	}

	// for i := range gridSize[0] + 1 {
	// 	for j := range gridSize[1] + 1 {
	// 		pos := [2]int{i, j}
	// 		if cell, ok := grid[pos]; ok {
	// 			fmt.Printf("%+v", string(cell))
	// 		} else {
	// 			fmt.Printf(" ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	return res
}
