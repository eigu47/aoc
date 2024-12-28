package day14

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 14)

const (
	// sizeX = 11
	// sizeY = 7
	sizeX = 101
	sizeY = 103
)

func Part1() int {
	upLft := 0
	upRgt := 0
	dwnLft := 0
	dwnRgt := 0

	for _, line := range input {
		nums := regexp.MustCompile(`-?\d+`).FindAllString(line, 4)
		posX, _ := strconv.Atoi(nums[0])
		posY, _ := strconv.Atoi(nums[1])
		velX, _ := strconv.Atoi(nums[2])
		velY, _ := strconv.Atoi(nums[3])

		newX := ((posX+velX*100)%sizeX + sizeX) % sizeX
		newY := ((posY+velY*100)%sizeY + sizeY) % sizeY

		if newX < sizeX/2 && newY < sizeY/2 {
			upLft++
		} else if newX < sizeX/2 && newY > sizeY/2 {
			upRgt++
		} else if newX > sizeX/2 && newY < sizeY/2 {
			dwnLft++
		} else if newX > sizeX/2 && newY > sizeY/2 {
			dwnRgt++
		}
	}

	return upLft * upRgt * dwnLft * dwnRgt
}

func Part2() int {
	grid := make([][]bool, sizeY)
	for i := range grid {
		grid[i] = make([]bool, sizeX)
	}

	type robot struct {
		pos [2]int
		vel [2]int
	}

	robots := []*robot{}

	for _, line := range input {
		nums := regexp.MustCompile(`-?\d+`).FindAllString(line, 4)
		posX, _ := strconv.Atoi(nums[0])
		posY, _ := strconv.Atoi(nums[1])
		velX, _ := strconv.Atoi(nums[2])
		velY, _ := strconv.Atoi(nums[3])

		robots = append(robots, &robot{
			[2]int{posX, posY},
			[2]int{velX, velY},
		})

		grid[posY][posX] = true
	}

	file, err := os.Create("./days/day14/output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file.Close()
	}()

	for i := range 10000 {
		for _, r := range robots {
			grid[r.pos[1]][r.pos[0]] = false

			r.pos[0] = ((r.pos[0]+r.vel[0])%sizeX + sizeX) % sizeX
			r.pos[1] = ((r.pos[1]+r.vel[1])%sizeY + sizeY) % sizeY

			grid[r.pos[1]][r.pos[0]] = true
		}

		buffer := bytes.Buffer{}
		buffer.WriteString(fmt.Sprintf("second %d\n", i+1))
		for _, row := range grid {
			for _, cell := range row {
				if cell {
					buffer.WriteString("●")
				} else {
					buffer.WriteString(" ")
				}
			}
			buffer.WriteString("\n")
		}
		buffer.WriteString("\n")

		_, err := file.Write(buffer.Bytes())
		if err != nil {
			log.Fatal(err)
		}
	}

	// search for ●●●●●●●●●
	// res is 8050
	return 0
}
