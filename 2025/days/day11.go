package days

import (
	"strings"
)

func Day11_1(input []string) int {
	// input = test_11
	ans := 0

	devices := make(map[string][]string, len(input))
	for _, row := range input {
		fields := strings.FieldsFunc(row, func(r rune) bool {
			return r == ' ' || r == ':'
		})

		devices[fields[0]] = fields[1:]
	}

	// for _, d := range devices {
	// 	fmt.Printf("d %+v\n", d)
	// }

	var dfs func(device string, visited map[string]bool, result int) int
	dfs = func(device string, visited map[string]bool, result int) int {
		if device == "out" {
			return result + 1
		}

		visited[device] = true
		for _, out := range devices[device] {
			result = dfs(out, visited, result)
		}

		return result
	}

	ans = dfs("you", map[string]bool{}, 0)
	return ans
}

func Day11_2(input []string) int {
	// input = test_11_2
	ans := 0

	devices := make(map[string][]string, len(input))
	for _, row := range input {
		fields := strings.FieldsFunc(row, func(r rune) bool {
			return r == ' ' || r == ':'
		})

		devices[fields[0]] = fields[1:]
	}

	getCountPath := func(device string) int {
		visited := map[string]bool{}
		type state struct {
			device string
			dac    bool
			fft    bool
		}
		memo := map[state]int{}

		var dfs func(string) int
		dfs = func(device string) int {
			s := state{
				device,
				visited["dac"],
				visited["fft"],
			}

			if val, ok := memo[s]; ok {
				return val
			}

			if device == "out" {
				if s.dac && s.fft {
					return 1
				}
				visited[device] = false
				return 0
			}

			visited[device] = true
			total := 0
			for _, output := range devices[device] {
				if !visited[output] {
					total += dfs(output)
				}
			}

			visited[device] = false
			memo[s] = total
			return total
		}

		return dfs(device)
	}

	ans = getCountPath("svr")
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

var test_11_2 = strings.Split(`svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`, "\n")
