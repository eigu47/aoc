package days

import (
	"slices"
	"strconv"
	"strings"

	"github.com/eigu47/aoc2023/util"
)

func Day5_1() int {
	var input = util.GetInput(2024, 5)

	updates := [][]int{}
	rules := make(map[int][]int)
	res := 0

	isRules := true
	for _, line := range input {
		if line == "" {
			isRules = false
			continue
		}

		if isRules {
			pages := strings.Split(line, "|")
			p1, _ := strconv.Atoi(pages[0])
			p2, _ := strconv.Atoi(pages[1])
			rules[p1] = append(rules[p1], p2)
		} else {
			var update []int
			for _, s := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(s)
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	for _, update := range updates {
		isValid := true
		for i := 0; i+1 < len(update); i++ {
			if !slices.Contains(rules[update[i]], update[i+1]) {
				isValid = false
				break
			}
		}

		if isValid {
			res += update[int(len(update)/2)]
		}
	}

	return res
}
