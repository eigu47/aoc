package days

import (
	"fmt"
	"maps"
	"strings"
)

func Day04_1(input []string) string {
	ans := 0
	grid := make(map[[2]int]bool)

	for i, row := range input {
		for j, cell := range row {
			pos := [2]int{i, j}

			if cell == '@' {
				grid[pos] = true
			}
		}
	}

	// for i := range len(test_04) {
	// 	for j := range len(test_04[0]) {
	// 		pos := [2]int{i, j}
	// 		if grid[pos] {
	// 			fmt.Print("@")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	for pos := range grid {
		rolls := 0

		for _, dir := range adjacent {
			curr := [2]int{pos[0] + dir[0], pos[1] + dir[1]}

			if grid[curr] {
				rolls++
				if rolls >= 4 {
					break
				}
			}
		}

		if rolls < 4 {
			ans++
		}
	}

	return fmt.Sprint(ans)
}

func Day04_2(input []string) string {
	ans := 0
	grid := make(map[[2]int]bool)

	for i, row := range input {
		for j, cell := range row {
			pos := [2]int{i, j}

			if cell == '@' {
				grid[pos] = true
			}
		}
	}

	var removeRoll func(grid map[[2]int]bool, removed int) int
	removeRoll = func(grid map[[2]int]bool, removed int) int {
		removedHere := false
		cloneGrid := maps.Clone(grid)

		for pos := range grid {
			rolls := 0

			for _, dir := range adjacent {
				curr := [2]int{pos[0] + dir[0], pos[1] + dir[1]}

				if grid[curr] {
					rolls++
					if rolls >= 4 {
						break
					}
				}
			}

			if rolls < 4 {
				delete(cloneGrid, pos)
				removed++
				removedHere = true
			}
		}

		if removedHere {
			return removeRoll(cloneGrid, removed)
		}

		return removed
	}

	ans = removeRoll(grid, 0)

	// for i := range len(test_04) {
	// 	for j := range len(test_04[0]) {
	// 		pos := [2]int{i, j}
	// 		if removed[pos] {
	// 			fmt.Print("@")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return fmt.Sprint(ans)
}

var test_04 = strings.Split(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`, "\n")

var adjacent = [8][2]int{
	{-1, 0},  // UP
	{-1, 1},  // UP RGT
	{0, 1},   // RGT
	{1, 1},   // DWN RGT
	{1, 0},   // DWN
	{1, -1},  // DWN LFT
	{0, -1},  // LFT
	{-1, -1}, // UP LFT
}
