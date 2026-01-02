package days

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Day12_1(input []string) int {
	// input = test_12
	ans := 0

	type present struct {
		shape [3][3]bool
		area  int
	}

	type region struct {
		size     [2]int
		presents map[*present]int
	}

	presents := []*present{}
	regions := []*region{}
	re := regexp.MustCompile(`^\d+:`)
	for i := 0; i < len(input); i++ {
		row := input[i]
		if row == "" {
			continue
		}

		if re.MatchString(row) {
			shape := [3][3]bool{}
			area := 0
			for y := 0; y < 3; y++ {
				for x, cell := range input[i+y+1] {
					if cell == '#' {
						shape[y][x] = true
						area++
					}
				}
			}
			presents = append(presents, &present{shape, area})
			i += 3
			continue
		}

		splits := strings.FieldsFunc(row, func(r rune) bool {
			return unicode.IsSpace(r) || r == 'x' || r == ':'
		})

		size := [2]int{}
		regionPresents := map[*present]int{}
		for i, numS := range splits {
			num, _ := strconv.Atoi(numS)
			if i < 2 {
				size[i] = num
			} else {
				regionPresents[presents[i-2]] = num
			}
		}
		regions = append(regions, &region{size, regionPresents})
	}

	for _, region := range regions {
		total := 0
		for present, count := range region.presents {
			total += (present.area * count)
		}

		if total < region.size[0]*region.size[1] {
			ans++
		}

		fmt.Printf(
			"size %+v, maxArea %+v, presents %+v, total %+v\n",
			region.size,
			region.size[0]*region.size[1],
			slices.Collect(maps.Values(region.presents)),
			total,
		)
	}

	return ans
}

func Day12_2(input []string) int {
	input = test_12
	ans := 0

	return ans
}

var test_12 = strings.Split(`0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`, "\n")
