package day8

import (
	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 8)
var antinodeGrid = map[[2]int]bool{}
var antennaMap = map[rune][][2]int{}

func init() {
	for i, line := range input {
		for j, rune := range line {
			if rune != '.' {
				pos := [2]int{i, j}
				antennaMap[rune] = append(antennaMap[rune], pos)
			}
		}
	}
}

func Part1() int {
	for _, antennas := range antennaMap {
		for i := 0; i+1 < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				dx := antennas[i][0] - antennas[j][0]
				dy := antennas[i][1] - antennas[j][1]

				antinodePos := [2][2]int{
					{antennas[i][0] + dx, antennas[i][1] + dy},
					{antennas[j][0] - dx, antennas[j][1] - dy},
				}

				for _, pos := range antinodePos {
					if pos[0] >= 0 && pos[0] < len(input) && pos[1] >= 0 && pos[1] < len(input[0]) {
						antinodeGrid[pos] = true
					}
				}
			}
		}
	}

	// fmt.Printf("antinodeGrid %+v\n", antinodeGrid)
	// for i, line := range input {
	// 	for j, rune := range line {
	// 		if antinodeGrid[[2]int{i, j}] {
	// 			fmt.Printf("#")
	// 		} else {
	// 			fmt.Printf("%s", string(rune))
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return len(antinodeGrid)
}
