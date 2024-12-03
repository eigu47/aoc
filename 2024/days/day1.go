package days

import (
	"fmt"
	"log"

	"github.com/eigu47/aoc2023/util"
)

func Part1() int {
	input, err := util.NewInput()
	if err != nil {
		log.Fatalln(err)
	}

	data, err := input.Strings(2024, 1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%v\n", data)

	return 0
}
