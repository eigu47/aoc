package day15

import (
	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 15)

var gridSize = [2]int{0, 0}
var walls = map[[2]int]bool{}
var boxes = map[[2]int]bool{}
var robot = [2]int{}
var movements = [][2]int{}

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
				if cell == '#' {
					walls[pos] = true
				} else if cell == 'O' {
					boxes[pos] = true
				} else if cell == '@' {
					robot = pos
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
	if walls[pos] {
		return pos, false
	}

	if boxes[pos] {
		return nextEmpty([2]int{pos[0] + mov[0], pos[1] + mov[1]}, mov)
	}

	return pos, true
}

func Part1() int {
	res := 0

	for _, mov := range movements {
		next := [2]int{robot[0] + mov[0], robot[1] + mov[1]}

		if walls[next] {
			continue
		}

		if boxes[next] {
			empty, ok := nextEmpty(next, mov)
			if !ok {
				continue
			}

			boxes[empty] = true
			delete(boxes, next)
		}

		robot = next
	}

	for box := range boxes {
		res += box[0]*100 + box[1]
	}

	// for i := range gridSize[0] + 1 {
	// 	for j := range gridSize[1] + 1 {
	// 		pos := [2]int{i, j}
	// 		if walls[pos] {
	// 			fmt.Printf("#")
	// 		} else if boxes[pos] {
	// 			fmt.Printf("O")
	// 		} else if pos == robot {
	// 			fmt.Printf("@")
	// 		} else {
	// 			fmt.Printf(" ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	return res
}
