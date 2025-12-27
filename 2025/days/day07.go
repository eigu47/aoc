package days

import (
	"strings"
)

func Day07_1(input []string) int {
	// input = test_07
	ans := 0

	tachyon := make(map[[2]int]bool)

	for y, row := range input {
		for x, cell := range row {
			pos := [2]int{y, x}
			if cell == 'S' {
				tachyon[addPos(pos, dirs["dwn"])] = true
			}

			if tachyon[pos] && y+1 < len(input) {
				if input[y+1][x] == '^' {
					tachyon[addPos(pos, dirs["dwnLft"])] = true
					tachyon[addPos(pos, dirs["dwnRgt"])] = true
					ans++
				} else {
					tachyon[addPos(pos, dirs["dwn"])] = true
				}
			}
		}
	}

	// for y, row := range input {
	// 	for x, cell := range row {
	// 		if tachyon[[2]int{y, x}] {
	// 			fmt.Print("|")
	// 		} else {
	// 			fmt.Print(string(cell))
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return ans
}

func Day07_2(input []string) int {
	// input = test_07
	ans := 0

	memo := make(map[[2]int]int)

	var traverse func(pos [2]int) int
	traverse = func(pos [2]int) int {
		if pos[0] >= len(input) {
			return 1
		}

		if v, ok := memo[pos]; ok {
			return v
		}

		res := 0
		if input[pos[0]][pos[1]] == '^' {
			res = traverse(addPos(pos, dirs["dwnLft"])) + traverse(addPos(pos, dirs["dwnRgt"]))
		} else {
			res = traverse(addPos(pos, dirs["dwn"]))
		}

		memo[pos] = res
		return res
	}

	ans = traverse([2]int{0, strings.Index(input[0], "S")})

	return ans
}

var test_07 = strings.Split(`.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`, "\n")

var dirs = map[string][2]int{
	"up":     {-1, 0},
	"upRgt":  {-1, 1},
	"rgt":    {0, 1},
	"dwnRgt": {1, 1},
	"dwn":    {1, 0},
	"dwnLft": {1, -1},
	"lft":    {0, -1},
	"upLft":  {-1, -1},
}

func addPos(a, b [2]int) [2]int {
	return [2]int{a[0] + b[0], a[1] + b[1]}
}
