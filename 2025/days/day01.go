package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var test = strings.Split(`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`, "\n")

func mod(a, n int) int {
	return ((a % n) + n) % n
}

func abs(a int) int {
	return max(a, -a)
}

func Day01_1(input []string) string {
	dial := 50
	ans := 0

	re := regexp.MustCompile(`(\w)(\d+)`)
	for _, val := range input {
		match := re.FindStringSubmatch(val)
		rotation, _ := strconv.Atoi(match[2])

		switch match[1] {
		case "R":
			dial += rotation
		case "L":
			dial -= rotation
		}

		dial = mod(dial, 100)

		// fmt.Printf("dial %+v, ans %+v\n", dial, ans)
		if dial == 0 {
			ans++
		}
	}

	return fmt.Sprintf("%v", ans)
}

func Day01_2(input []string) string {
	dial := 50
	ans := 0

	re := regexp.MustCompile(`(\w)(\d+)`)
	for _, val := range input {
		match := re.FindStringSubmatch(val)
		rotation, _ := strconv.Atoi(match[2])

		// switch match[1] {
		// case "R":
		// 	for range rotation {
		// 		dial = mod(dial+1, 100)
		// 		if dial == 0 {
		// 			ans++
		// 		}
		// 	}
		// case "L":
		// 	for range rotation {
		// 		dial = mod(dial-1, 100)
		// 		if dial == 0 {
		// 			ans++
		// 		}
		// 	}
		// }

		// Time 16.0492ms:
		// 6684
		// fmt.Printf("dial %+v, ans %+v", dial, ans)

		switch match[1] {
		case "R":
			ans += rotation / 100
			mod := rotation % 100
			if dial+mod >= 100 {
				ans++
			}
			dial += rotation
		case "L":
			ans += rotation / 100
			mod := rotation % 100
			if dial != 0 && dial-mod <= 0 {
				ans++
			}
			dial -= rotation
		}
		dial = mod(dial, 100)


		// Time 2.2069ms:
		// fmt.Printf("dial %+v, ans %+v", dial, ans)
	}

	return fmt.Sprintf("%v", ans)
}
