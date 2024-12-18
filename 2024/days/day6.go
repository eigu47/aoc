package days

import (
	"github.com/eigu47/aoc2023/util"
)

func Day6_1() int {
	input := util.GetInput(2024, 6)
	result := map[[2]int]bool{}

	DIRS := map[int][2]int{
		0: {-1, 0},
		3: {0, -1},
		1: {0, 1},
		2: {1, 0},
	}

	dir := 0
	pos := [2]int{}
	obs := map[[2]int]bool{}

	for i, line := range input {
		for j, cell := range line {
			if cell == '#' {
				obs[[2]int{i, j}] = true
			} else if cell == '^' {
				// pos = [2]int{i, j}
				pos[0], pos[1] = i, j
				result[pos] = true
			}
		}
	}

	next := [2]int{}
	for {
		// next := [2]int{pos[0] + DIRS[dir][0], pos[1] + DIRS[dir][1]}
		next[0], next[1] = pos[0] + DIRS[dir][0], pos[1] + DIRS[dir][1]
		if next[0] < 0 || next[0] >= len(input) || next[1] < 0 || next[1] >= len(input[0]) {
			break
		}

		if _, ok := obs[next]; ok {
			dir = (dir + 1) % 4
			next[0], next[1] = pos[0] + DIRS[dir][0], pos[1] + DIRS[dir][1]
		}

		pos = next
		result[pos] = true
	}

	// for i, row := range input {
	// 	for j := range row {
	// 		if obs := obs[[2]int{i, j}]; obs {
	// 			fmt.Printf("#")
	// 		} else if _, ok := result[[2]int{i, j}]; ok {
	// 			fmt.Printf("X")
	// 			// fmt.Printf(".")
	// 		} else {
	// 			fmt.Printf(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return len(result)
}
