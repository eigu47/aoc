package day14

import (
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
