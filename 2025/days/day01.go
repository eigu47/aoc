package days

import (
	"fmt"
	"regexp"
	"strconv"
)

func mod(a, n int) int {
	r := a % n
	if r < 0 {
		r += n
	}
	return r
}

func Day01_1(input []string) string {
	dial := 50
	ans := 0

	// 	test := strings.Split(`L68
	// L30
	// R48
	// L5
	// R60
	// L55
	// L1
	// L99
	// R14
	// L82`, "\n")

	re := regexp.MustCompile(`(\w)(\d+)`)
	for _, val := range input {
		match := re.FindStringSubmatch(val)
		rotation, _ := strconv.Atoi(match[2])

		switch match[1] {
		case "R":
			dial = mod((dial + rotation), 100)
		case "L":
			dial = mod((dial - rotation), 100)
		}

		// fmt.Printf("pos %+v\n", dial)
		if dial == 0 {
			ans++
		}
	}

	return fmt.Sprintf("%v", ans)
}

func Day01_2(input []string) string {
	return "day1_2"
}
