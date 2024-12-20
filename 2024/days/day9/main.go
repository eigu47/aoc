package day9

import (
	"strconv"

	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 9)
var disk = []int{}

func init() {
	for _, line := range input {
		for i, rune := range line {
			num, _ := strconv.Atoi(string(rune))
			if i%2 == 0 {
				for range num {
					disk = append(disk, i/2)
				}
			} else {
				for range num {
					disk = append(disk, -1)
				}
			}
		}
	}
}

func Part1() int {
	res := 0

	for lft, rgt := 0, len(disk)-1; lft < rgt; lft++ {
		if disk[lft] == -1 {
			for disk[rgt] == -1 {
				rgt--
			}

			disk[lft], disk[rgt] = disk[rgt], disk[lft]
		}
	}

	for i := 0; i < len(disk) && disk[i] != -1; i++ {
		res += i * disk[i]
	}

	// for i := range len(disk) {
	// 	if disk[i] != -1 {
	// 		fmt.Printf("%+v", disk[i])
	// 	} else {
	// 		fmt.Printf(".")
	// 	}
	// }
	// fmt.Println()

	return res
}
