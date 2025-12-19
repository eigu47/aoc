package main

import (
	"fmt"
	"log"

	"github.com/eigu47/aoc2025/days"
	"github.com/manifoldco/promptui"
)

func main() {
	_days := []func(){
		days.Day01_1,
	}

	items := make([]string, len(_days))
	for i := range _days {
		items[i] = fmt.Sprintf("Day%02d", i+1)
	}

	prompt := promptui.Select{
		Label: "Select day",
		Items: items,
	}

	index, _, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	_days[index]()
}
