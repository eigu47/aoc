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

	uniqueX := map[int]struct{}{}
	uniqueY := map[int]struct{}{}
	for _, tile := range tiles {
		uniqueX[tile[0]] = struct{}{}
		uniqueY[tile[1]] = struct{}{}
	}

	createCompMap := func(unique map[int]struct{}) ([]int, map[int]int) {
		comp := make([]int, 0, len(unique))
		for i := range unique {
			comp = append(comp, i)
		}
		slices.Sort(comp)

		hMap := make(map[int]int, len(comp))
		for idx, val := range comp {
			hMap[val] = idx
		}

		return comp, hMap
	}

	xComp, xMap := createCompMap(uniqueX)
	yComp, yMap := createCompMap(uniqueY)

	grid := make([][]bool, len(yComp))
	for i := range grid {
		grid[i] = make([]bool, len(xComp))
	}

	for i, tile := range tiles {
		x := xMap[tile[0]]
		y := yMap[tile[1]]
		grid[y][x] = true

		next := tiles[(i+1)%len(tiles)]
		x1, y1 := xMap[tile[0]], yMap[tile[1]]
		x2, y2 := xMap[next[0]], yMap[next[1]]

		if x1 == x2 {
			for y := min(y1, y2); y <= max(y1, y2); y++ {
				grid[y][x1] = true
			}
		} else if y1 == y2 {
			for x := min(x1, x2); x <= max(x1, x2); x++ {
				grid[y1][x] = true
			}
		}
	}

	inside := [2]int{}
found:
	for y, row := range grid {
		for x, cell := range row {
			if cell {
				continue
			}

			cross := 0
			prev := false
			for i := x; i >= 0; i-- {
				if row[i] != prev {
					cross++
				}
				prev = row[i]
			}

			if cross%2 == 1 {

				inside = [2]int{y, x}
				break found
			}
		}
	}

	grid[inside[0]][inside[1]] = true
	queue := [][2]int{inside}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, d := range axes {
			dy, dx := p[0]+d[0], p[1]+d[1]
			if dy < 0 || dy >= len(grid) || dx < 0 || dx >= len(grid[0]) {
				continue
			}
			if !grid[dy][dx] {
				grid[dy][dx] = true
				queue = append(queue, [2]int{dy, dx})
			}
		}
	}

	// for _, row := range grid {
	// 	for _, cell := range row {
	// 		if cell {
	// 			fmt.Printf("#")
	// 		} else {
	// 			fmt.Printf(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Printf("xComp %v, xMap %+v\n", xComp, xMap)
	// fmt.Printf("yComp %v, yMap %+v\n", yComp, yMap)

	slices.SortFunc(pairs, func(a, b *pairArea) int {
		return cmp.Compare(b.area, a.area)
	})

	for _, p := range pairs {
		x1, y1 := xMap[p.pair[0][0]], yMap[p.pair[0][1]]
		x2, y2 := xMap[p.pair[1][0]], yMap[p.pair[1][1]]

		valid := true
		for y := min(y1, y2); y <= max(y1, y2); y++ {
			for x := min(x1, x2); x <= max(x1, x2); x++ {
				if !grid[y][x] {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}

		if valid {
			width := abs(p.pair[0][0]-p.pair[1][0]) + 1
			height := abs(p.pair[0][1]-p.pair[1][1]) + 1
			return width * height
		}
	}

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

var axes = [4][2]int{
	{-1, 0}, // UP
	{0, 1},  // RGT
	{1, 0},  // DWN
	{0, -1}, // LFT
}
