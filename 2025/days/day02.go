package days

import (
	"fmt"
	"strconv"
	"strings"
)

func Day02_1(input []string) string {
	ans := 0

	for _, idRanges := range strings.Split(input[0], ",") {
		ids := strings.Split(idRanges, "-")
		start, _ := strconv.Atoi(ids[0])
		end, _ := strconv.Atoi(ids[1])

		for id := start; id <= end; id++ {
			idStr := strconv.Itoa(id)
			if len(idStr)%2 != 0 {
				continue
			}

			half := len(idStr) / 2
			if idStr[:half] == idStr[half:] {
				ans += id
			}
		}
	}

	return fmt.Sprint(ans)
}

func Day02_2(input []string) string {
	ans := 0

	for _, idRanges := range strings.Split(input[0], ",") {
		ids := strings.Split(idRanges, "-")
		start, _ := strconv.Atoi(ids[0])
		end, _ := strconv.Atoi(ids[1])

		for id := start; id <= end; id++ {
			idStr := strconv.Itoa(id)
			idLen := len(idStr)
			divisors := getDivisors(idLen)
			invalid := false

			for _, divisor := range divisors {
				pattern := idStr[:divisor]

				for i := divisor; i < idLen; i += divisor {
					if pattern != idStr[i:i+divisor] {
						invalid = false
						break
					}
					invalid = true
				}

				if invalid {
					ans += id
					break
				}
			}
		}
	}

	return fmt.Sprint(ans)
}

func getDivisors(n int) []int {
	n = abs(n)
	divisors := []int{}

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)

			if n/i != i {
				divisors = append(divisors, n/i)
			}
		}
	}

	return divisors
}

func every[T any](s []T, cb func(T) bool) bool {
	for _, v := range s {
		if !cb(v) {
			return false
		}
	}
	return true
}

var test_02 = `
11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
