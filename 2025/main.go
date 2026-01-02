package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/eigu47/aoc2025/days"
	"github.com/eigu47/aoc2025/util"
	"github.com/manifoldco/promptui"
)

func main() {
	days := []struct {
		Day  int
		Part int
		Run  func([]string) int
	}{
		{
			Day:  11,
			Part: 2,
			Run:  days.Day11_2,
		},
		{
			Day:  11,
			Part: 1,
			Run:  days.Day11_1,
		},
		// {
		// 	Day:  10,
		// 	Part: 2,
		// 	Run:  days.Day10_2,
		// },
		{
			Day:  10,
			Part: 1,
			Run:  days.Day10_1,
		},
		{
			Day:  9,
			Part: 2,
			Run:  days.Day09_2,
		},
		{
			Day:  9,
			Part: 1,
			Run:  days.Day09_1,
		},
		{
			Day:  8,
			Part: 2,
			Run:  days.Day08_2,
		},
		{
			Day:  8,
			Part: 1,
			Run:  days.Day08_1,
		},
		{
			Day:  7,
			Part: 2,
			Run:  days.Day07_2,
		},
		{
			Day:  7,
			Part: 1,
			Run:  days.Day07_1,
		},
		{
			Day:  6,
			Part: 2,
			Run:  days.Day06_2,
		},
		{
			Day:  6,
			Part: 1,
			Run:  days.Day06_1,
		},
		{
			Day:  5,
			Part: 2,
			Run:  days.Day05_2,
		},
		{
			Day:  5,
			Part: 1,
			Run:  days.Day05_1,
		},
		{
			Day:  4,
			Part: 2,
			Run:  days.Day04_2,
		},
		{
			Day:  4,
			Part: 1,
			Run:  days.Day04_1,
		},
		{
			Day:  3,
			Part: 2,
			Run:  days.Day03_2,
		},
		{
			Day:  3,
			Part: 1,
			Run:  days.Day03_1,
		},
		{
			Day:  2,
			Part: 2,
			Run:  days.Day02_2,
		},
		{
			Day:  2,
			Part: 1,
			Run:  days.Day02_1,
		},
		{
			Day:  1,
			Part: 2,
			Run:  days.Day01_2,
		},
		{
			Day:  1,
			Part: 1,
			Run:  days.Day01_1,
		},
	}

	all := flag.Bool("all", false, "ask all days")
	flag.Parse()
	idx := 0

	if *all {
		items := make([]string, len(days))
		for i, d := range days {
			items[i] = fmt.Sprintf("Day %02d, Part %d", d.Day, d.Part)
		}

		prompt := promptui.Select{
			Label: "Select day",
			Items: items,
			Searcher: func(input string, index int) bool {
				return strings.Contains(
					strings.ToLower(items[index]),
					strings.ToLower(input),
				)
			},
		}

		var err error
		idx, _, err = prompt.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	input, err := util.GetInput(2025, days[idx].Day)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	answer := days[idx].Run(input)

	fmt.Printf("Time %v:\n%v", time.Since(start), answer)
}
