package day15

import (
	"fmt"

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
	boxL  = '['
	boxR  = ']'
)

var directions = map[rune][2]int{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
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
	isGrid := true
	for i, line := range input {
		if line == "" {
			isGrid = false
			continue
		}

		for j, cell := range line {
			if isGrid {
				gridSize[0], gridSize[1] = max(gridSize[0], i), max(gridSize[1], j)
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

func moveBoxX(pos map[[2]int]bool, mov [2]int) bool {
	isEmpty := true
	for p := range pos {
		if grid[p] == wall {
			return false
		}

		if _, ok := grid[p]; ok {
			isEmpty = false
		}
	}

	if isEmpty {
		return true
	}

	next := map[[2]int]bool{}
	for p := range pos {
		if grid[p] != boxL && grid[p] != boxR {
			continue
		}

		n := [2]int{p[0] + mov[0], p[1] + mov[1]}
		next[n] = true
		if grid[n] == boxL {
			next[[2]int{n[0], n[1] + 1}] = true
		}

		if grid[n] == boxR {
			next[[2]int{n[0], n[1] - 1}] = true
		}
	}

	if !moveBoxX(next, mov) {
		return false
	}

	for n := range next {
		p := [2]int{n[0] - mov[0], n[1] - mov[1]}
		if pos[p] && (grid[p] == boxL || grid[p] == boxR) {
			grid[n] = grid[p]
			delete(grid, p)
		}
	}

	return true
}

func moveBoxY(pos [2]int, mov [2]int) bool {
	if grid[pos] == wall {
		return false
	}

	if _, ok := grid[pos]; !ok {
		return true
	}

	next := [2]int{pos[0] + mov[0], pos[1] + mov[1]}
	if !moveBoxY(next, mov) {
		return false
	}

	grid[next] = grid[pos]
	delete(grid, pos)

	return true
}

func Part2() int {
	isGrid := true
	for i, line := range input {
		if line == "" {
			isGrid = false
			continue
		}

		for j, cell := range line {
			if isGrid {
				gridSize[0], gridSize[1] = max(gridSize[0], i), max(gridSize[1], j*2)
				if cell == '.' {
					continue
				}

				pos := [2]int{i, j * 2}
				next := [2]int{i, j*2 + 1}
				if cell == empty {
					continue
				}

				if cell == wall {
					grid[pos] = cell
					grid[next] = cell
				} else if cell == box {
					grid[pos] = boxL
					grid[next] = boxR
				} else if cell == robot {
					grid[pos] = cell
					robotPos = pos
				}

				continue
			}

			if mov, ok := directions[cell]; ok {
				movements = append(movements, mov)
			}
		}
	}

	res := 0

	for _, mov := range movements {
		next := [2]int{robotPos[0] + mov[0], robotPos[1] + mov[1]}
		if grid[next] == wall {
			continue
		}

		if grid[next] == boxL || grid[next] == boxR {
			if mov == directions['^'] || mov == directions['v'] {
				pos := map[[2]int]bool{next: true}
				if grid[next] == boxL {
					pos[[2]int{next[0], next[1] + 1}] = true
				}
				if grid[next] == boxR {
					pos[[2]int{next[0], next[1] - 1}] = true
				}

				if !moveBoxX(pos, mov) {
					continue
				}
			}

			if mov == directions['>'] || mov == directions['<'] {
				if !moveBoxY(next, mov) {
					continue
				}
			}
		}

		delete(grid, robotPos)
		robotPos = next
		grid[robotPos] = robot
	}

	for pos, cell := range grid {
		if cell == '[' {
			res += pos[0]*100 + pos[1]
		}
	}

	// fmt.Printf("itineration %+v\n", i)
	for i := range gridSize[0] + 2 {
		for j := range gridSize[1] + 2 {
			pos := [2]int{i, j}
			if cell, ok := grid[pos]; ok {
				fmt.Printf("%+v", string(cell))
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	return res
}
