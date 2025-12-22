package days

import (
	"fmt"
	"strings"
)

func Day03_01(input []string) string {
	ans := 0

	for _, banks := range input {
		first := 0

		for i := first; i+1 < len(banks); i++ {
			if int(banks[i]-'0') > int(banks[first]-'0') {
				// fmt.Printf("first %+v, idx %v\n", banks[i]-'0', i)
				first = i
			}
		}

		last := first + 1
		for i := last; i < len(banks); i++ {
			if int(banks[i]-'0') > int(banks[last]-'0') {
				// fmt.Printf("last %+v, idx %v\n", banks[i]-'0', i)
				last = i
			}
		}

		// fmt.Printf("max %v\n", int(banks[first]-'0')*10+int(banks[last]-'0'))
		ans += int(banks[first]-'0')*10 + int(banks[last]-'0')
	}

	return fmt.Sprint(ans)
}

var test_03 = strings.Split(`987654321111111
811111111111119
234234234234278
818181911112111`, "\n")
