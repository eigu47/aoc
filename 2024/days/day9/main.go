package day9

import (
	"strconv"

	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 9)

func Part1() int {
	res := 0
	var disk = []int{}

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

func Part2() int {
	res := 0
	files := map[int][]int{}
	freeSpace := map[int][]int{}

	curr := 0
	for _, line := range input {
		for i, rune := range line {
			num, _ := strconv.Atoi(string(rune))
			if i%2 == 0 {
				files[i/2] = []int{curr, curr + num}
			} else {
				freeSpace[i/2] = []int{curr, curr + num}
			}
			curr += num
		}
	}

	for rgt := len(input[0]) / 2; rgt >= 0; rgt-- {
		file := files[rgt]

		for lft := 0; lft < len(input[0])/2; lft++ {
			free, ok := freeSpace[lft]
			if !ok {
				continue
			}

			if free[0] > file[0] {
				break
			}

			freeSize := free[1] - free[0]
			fileSize := file[1] - file[0]
			if freeSize < fileSize {
				continue
			}

			file[0], file[1] = free[0], free[0]+fileSize
			free[0] += fileSize

			if free[1]-free[0] <= 0 {
				delete(freeSpace, lft)
			}
		}
	}

	// disk := map[int]int{}
	// for idx, file := range files {
	// 	for i := file[0]; i < file[1]; i++ {
	// 		disk[i] = idx
	// 	}
	// }
	// fmt.Printf("disk %+v\n", disk)
	// for i := 0; i < len(input[0]); i++ {
	// 	if val, ok := disk[i]; ok {
	// 		fmt.Printf("%+v", val)
	// 	} else {
	// 		fmt.Printf(".")
	// 	}
	// }
	// fmt.Println()

	for idx, file := range files {
		for i := file[0]; i < file[1]; i++ {
			res += idx * i
		}
	}

	return res
}
