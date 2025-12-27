package days

import (
	"slices"
	"strconv"
	"strings"
)

func Day05_1(input []string) int {
	// input = test_05
	ans := 0

	mid := slices.Index(input, "")
	freshIds := [][2]int{}

	for _, fresh := range input[:mid] {
		ranges := strings.Split(fresh, "-")
		from, _ := strconv.Atoi(ranges[0])
		to, _ := strconv.Atoi(ranges[1])
		freshIds = append(freshIds, [2]int{from, to})
	}

	for _, val := range input[mid+1:] {
		id, _ := strconv.Atoi(val)

		for _, fresh := range freshIds {
			from := fresh[0]
			to := fresh[1]

			if id >= from && id <= to {
				ans++
				break
			}
		}
	}

	return ans
}

func Day05_2(input []string) int {
	input = test_05
	ans := 0

	mid := slices.Index(input, "")
	freshIds := [][2]int{}

	for _, fresh := range input[:mid] {
		ranges := strings.Split(fresh, "-")
		from, _ := strconv.Atoi(ranges[0])
		to, _ := strconv.Atoi(ranges[1])
		freshIds = append(freshIds, [2]int{from, to})
	}

	slices.SortFunc(freshIds, func(a, b [2]int) int {
		for i := range len(a) {
			if a[i] != b[i] {
				return a[i] - b[i]
			}
		}
		return 0
	})

	curr := -1

	for _, rng := range freshIds {
		from := rng[0]
		to := rng[1]

		if curr >= from {
			from = curr + 1
		}

		if from <= to {
			ans += to - from + 1
		}

		curr = max(curr, to)
	}

	return ans
}

var test_05 = strings.Split(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`, "\n")
