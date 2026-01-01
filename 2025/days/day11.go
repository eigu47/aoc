package days

import (
	"strings"
)

func Day11_1(input []string) int {
	// input = test_11
	ans := 0

	type device struct {
		label  string
		output []*device
	}

	devices := make(map[string]*device, len(input))
	for _, row := range input {
		fields := strings.FieldsFunc(row, func(r rune) bool {
			return r == ' ' || r == ':'
		})

		output := make([]*device, len(fields[1:]))
		for i, label := range fields {
			if _, ok := devices[label]; !ok {
				devices[label] = &device{
					label: label,
				}
			}

			if i != 0 {
				output[i-1] = devices[label]
			}
		}

		devices[fields[0]].output = output
	}

	// for _, d := range devices {
	// 	fmt.Printf("d %+v\n", d)
	// }

	var dfs func(d *device, visited map[*device]bool, result int) int
	dfs = func(d *device, visited map[*device]bool, result int) int {
		if d == devices["out"] {
			return result + 1
		}

		visited[d] = true
		for _, out := range d.output {
			result = dfs(out, visited, result)
		}

		return result
	}

	ans = dfs(devices["you"], map[*device]bool{}, 0)

	return ans
}

func Day11_2(input []string) int {
	input = test_11
	ans := 0

	return ans
}

var test_11 = strings.Split(`aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`, "\n")
