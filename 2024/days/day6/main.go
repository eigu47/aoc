package day6

import (
	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 6)
var directions = map[int][2]int{
	0: {-1, 0},
	3: {0, -1},
	1: {0, 1},
	2: {1, 0},
}
var initialPos = [2]int{}
var obstructions = map[[2]int]bool{}
var path = map[[2]int]bool{}

// func move(pos [2]int, dir int) {
// 	if pos[0] < 0 || pos[0] >= len(input) || pos[1] < 0 || pos[1] >= len(input[0]) {
// 		return
// 	}

// 	next := [2]int{pos[0] + directions[dir][0], pos[1] + directions[dir][1]}
// 	if obstructions[next] {
// 		dir = (dir + 1) % 4
// 		next[0], next[1] = pos[0]+directions[dir][0], pos[1]+directions[dir][1]
// 	}

// 	path[pos] = true
// 	move1(next, dir)
// }

func init() {
	for i, line := range input {
		for j, cell := range line {
			if cell == '#' {
				obstructions[[2]int{i, j}] = true
			} else if cell == '^' {
				initialPos[0], initialPos[1] = i, j
				// path[initialPos] = true
			}
		}
	}

	pos := initialPos
	dir := 0
	next := [2]int{}

	for {
		next[0], next[1] = pos[0]+directions[dir][0], pos[1]+directions[dir][1]
		if next[0] < 0 || next[0] >= len(input) || next[1] < 0 || next[1] >= len(input[0]) {
			break
		}

		if obstructions[next] {
			dir = (dir + 1) % 4
			next[0], next[1] = pos[0]+directions[dir][0], pos[1]+directions[dir][1]
		}

		pos = next
		path[pos] = true
	}

	// 	for i, row := range input {
	// 	for j := range row {
	// 		if obs := obstructions[[2]int{i, j}]; obs {
	// 			fmt.Printf("#")
	// 		} else if _, ok := path[[2]int{i, j}]; ok {
	// 			fmt.Printf("X")
	// 		} else {
	// 			fmt.Printf(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}

func Part1() int {
	return len(path)
}

func Part2() int {
	res := 0

	for pathPos := range path {
		if pathPos == initialPos {
			continue
		}

		pos := initialPos
		dir := 0
		newPath := map[[2]int][4]bool{
			// pos: {true, false, false, false},
		}

		obstructions[pathPos] = true
		// newObs := map[[2]int]bool{}
		// for k, v := range obstructions {
		// 	newObs[k] = v
		// }
		// newObs[pathPos] = true

		next := [2]int{}
		for {
			next[0], next[1] = pos[0]+directions[dir][0], pos[1]+directions[dir][1]
			if next[0] < 0 || next[0] >= len(input) || next[1] < 0 || next[1] >= len(input[0]) {
				break
			}

			next[0], next[1] = pos[0]+directions[dir][0], pos[1]+directions[dir][1]
			if obstructions[next] {
				dir = (dir + 1) % 4
				next[0], next[1] = pos[0]+directions[dir][0], pos[1]+directions[dir][1]
			} else {
				pos = next
			}

			if d, ok := newPath[pos]; ok && d[dir] {
				// fmt.Printf("pathPos%+v\n", pathPos)
				// for i, row := range input {
				// 	for j := range row {
				// 		if obstructions[[2]int{i, j}] {
				// 			if pathPos == [2]int{i, j} {
				// 				fmt.Printf("O")
				// 			} else {
				// 				fmt.Printf("#")
				// 			}
				// 		} else if arr, ok := newPath[[2]int{i, j}]; ok {
				// 			if arr[0] {
				// 				fmt.Printf("^")
				// 			} else if arr[1] {
				// 				fmt.Printf(">")
				// 			} else if arr[2] {
				// 				fmt.Printf("V")
				// 			} else if arr[3] {
				// 				fmt.Printf("<")
				// 			}
				// 		} else {
				// 			fmt.Printf(".")
				// 		}
				// 	}
				// 	fmt.Println()
				// }

				res++
				break
			} else if !ok {
				newPath[pos] = [4]bool{}
			}
			pathDir := newPath[pos]
			pathDir[dir] = true
			newPath[pos] = pathDir
		}

		delete(obstructions, pathPos)

	}

	return res
}
