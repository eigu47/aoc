package main

import (
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
		Run  func([]string) string
	}{
		// {
		// 	Day:  1,
		// 	Part: 2,
		// 	Run:  days.Day01_2,
		// },
		{
			Day:  1,
			Part: 1,
			Run:  days.Day01_1,
		},
	}

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

	idx, _, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	input, err := util.GetInput(2025, days[idx].Day)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	answer := days[idx].Run(input)

	fmt.Printf("Time %v:\n\n%v", time.Since(start), answer)
}
