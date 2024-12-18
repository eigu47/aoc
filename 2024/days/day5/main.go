package day5

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 5)

var rules = make(map[int][]int)
var updates = [][]int{}

func init() {
	isRules := true
	for _, line := range input {
		if line == "" {
			isRules = false
			continue
		}

		if isRules {
			pages := strings.Split(line, "|")
			rgt, _ := strconv.Atoi(pages[0])
			lft, _ := strconv.Atoi(pages[1])
			rules[rgt] = append(rules[rgt], lft)
		} else {
			var update []int
			for _, u := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(u)
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}
}

func Part1() int {
	res := 0

	for _, update := range updates {
		isValid := true
		for i := 0; i+1 < len(update); i++ {
			if !slices.Contains(rules[update[i]], update[i+1]) {
				isValid = false
				break
			}
		}

		if isValid {
			res += update[len(update)/2]
		}
	}

	return res
}

func Part2() int {
	res := 0

	for _, update := range updates {
		isValid := true
		for i := 0; i+1 < len(update); i++ {
			if !slices.Contains(rules[update[i]], update[i+1]) {
				sort.Slice(update, func(i, j int) bool {
					return slices.Contains(rules[update[i]], update[j])
				})
				isValid = false

				break
			}
		}

		if !isValid {
			res += update[len(update)/2]
		}
	}

	return res
}
