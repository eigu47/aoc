package days

import (
	"regexp"
	"strconv"
	"strings"
)

func Day10_1(input []string) int {
	// input = test_10
	ans := 0

	type machine struct {
		lights  int
		buttons []int
		joltage []int
	}

	// bitmask all the values
	re := regexp.MustCompile(`\[(.+)\] ((?:\(.+\) )+)\{(.+)\}`)
	btnRe := regexp.MustCompile(`\(([\d+,]+)\)`)
	machines := []*machine{}
	for _, m := range input {
		match := re.FindStringSubmatch(m)

		lights := 0
		for i, l := range match[1] {
			if l == '#' {
				lights |= 1 << i
			}
		}

		buttons := []int{}
		for _, btnMatch := range btnRe.FindAllStringSubmatch(match[2], -1) {
			btnPress := 0
			for _, btn := range strings.Split(btnMatch[1], ",") {
				i, _ := strconv.Atoi(btn)
				btnPress |= 1 << i
			}

			buttons = append(buttons, btnPress)
		}

		joltage := []int{}
		for _, jStr := range strings.Split(match[3], ",") {
			j, _ := strconv.Atoi(jStr)
			joltage = append(joltage, j)
		}

		machines = append(machines, &machine{
			lights:  lights,
			buttons: buttons,
			joltage: joltage,
		})
	}

	// for _, m := range machines {
	// 	fmt.Printf("m %+v\n", m)
	// }

	type queue struct {
		lights int
		depth  int
	}

	// BFS for each machine from 0 to lights
minPress:
	for _, machine := range machines {
		visited := map[int]bool{0: true}
		q := []queue{{0, 0}}
		for len(q) > 0 {
			curr := q[0]
			q = q[1:]

			if curr.lights == machine.lights {
				ans += curr.depth
				continue minPress
			}

			for _, n := range machine.buttons {
				next := curr.lights ^ n
				if !visited[next] {
					visited[next] = true
					q = append(q, queue{next, curr.depth + 1})
				}
			}
		}
	}

	return ans
}

func Day10_2(input []string) int {
	input = test_10
	ans := 0

	return ans
}

var test_10 = strings.Split(`[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`, "\n")
