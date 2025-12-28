package days

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day09_1(input []string) int {
	// input = test_09
	ans := 0

	tiles := [][2]int{}
	for _, tile := range input {
		pos := strings.Split(tile, ",")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		tiles = append(tiles, [2]int{x, y})
	}

	type pairArea struct {
		pair [2][2]int
		area int
	}

	pairs := []*pairArea{}
	for i, t1 := range tiles {
		for _, t2 := range tiles[i+1:] {
			pairs = append(pairs, &pairArea{
				pair: [2][2]int{t1, t2},
				area: abs((t1[0] - t2[0] + 1) * (t1[1] - t2[1] + 1)),
			})
		}
	}

	// slicesPrint(pairs)

	maxArea := slices.MaxFunc(pairs, func(a, b *pairArea) int {
		return cmp.Compare(a.area, b.area)
	})

	ans = maxArea.area

	return ans
}

func Day09_2(input []string) int {
	input = test_09
	ans := 0

	return ans
}

var test_09 = strings.Split(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`, "\n")

func slicesPrint[T any](s []T) {
	for i, v := range s {
		fmt.Printf("i: %v, val %+v\n", i, v)
	}
}
