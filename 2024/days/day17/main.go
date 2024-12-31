package day17

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/eigu47/aoc2023/util"
)

var input = util.GetInput(2024, 17)
var a = 0
var b = 0
var c = 0
var program = []int{}

func init() {
	rgx := regexp.MustCompile(`\d+`)
	aNum, _ := strconv.Atoi(rgx.FindString(input[0]))
	a = aNum
	bNum, _ := strconv.Atoi(rgx.FindString(input[1]))
	b = bNum
	cNum, _ := strconv.Atoi(rgx.FindString(input[2]))
	c = cNum

	for _, prog := range strings.Split(regexp.MustCompile(`[\d,]+`).FindString(input[4]), ",") {
		progNum, _ := strconv.Atoi(prog)
		program = append(program, progNum)
	}
}

func Part1() string {
	res := []string{}

	var operands = func(op int) int {
		switch op {
		case 0:
			return 0
		case 1:
			return 1
		case 2:
			return 2
		case 3:
			return 3
		case 4:
			return a
		case 5:
			return b
		case 6:
			return c
		default:
			return 0
		}
	}

	opcodes := map[int]func(int, int) (idx int){
		0: func(combo int, idx int) int {
			a = a / int(math.Pow(float64(2), float64(operands(combo))))
			idx += 2
			return idx
		},
		1: func(literal int, idx int) int {
			b = b ^ literal
			idx += 2
			return idx
		},
		2: func(combo int, idx int) int {
			b = operands(combo) % 8
			idx += 2
			return idx
		},
		3: func(literal int, idx int) int {
			if a == 0 {
				idx += 2
				return idx
			}
			return literal
		},
		4: func(_ int, idx int) int {
			b = b ^ c
			idx += 2
			return idx
		},
		5: func(combo int, idx int) int {
			idx += 2
			res = append(res, fmt.Sprint(operands(combo)%8))
			return idx
		},
		6: func(combo int, idx int) int {
			b = a / int(math.Pow(float64(2), float64(operands(combo))))
			idx += 2
			return idx
		},
		7: func(combo int, idx int) int {
			c = a / int(math.Pow(float64(2), float64(operands(combo))))
			idx += 2
			return idx
		},
	}

	for i := 0; i < len(program); {
		i = opcodes[program[i]](program[i+1], i)
	}

	return strings.Join(res, ",")
}
