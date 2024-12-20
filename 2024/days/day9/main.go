package day9

import (
	"fmt"
	"strconv"

	"github.com/eigu47/aoc2023/util"
)

var intput = util.GetInput(2024, 9)
var disk = []int{}
var last = 0

func init() {
	fileId := 0
	for _, line := range intput {
		for i, rune := range line {
			num, _ := strconv.Atoi(string(rune))
			if i%2 == 0 {
				for range num {
					disk = append(disk, fileId)
				}
				fileId++
				last = len(disk)
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

	for i := 0; i < last; i++ {
		if disk[i] == -1 {
			for j := last - 1; j >= 0 && disk[j] == -1; j-- {
				last = j
			}

			disk[i], disk[last-1] = disk[last-1], disk[i]
			last--
		}
	}

	for i := range last {
		res += i * disk[i]
	}

	for i := range len(disk) {
		if disk[i] != -1 {
			fmt.Printf("%+v", disk[i])
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println()

	return res
}
