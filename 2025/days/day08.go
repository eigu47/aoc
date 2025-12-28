package days

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day08_1(input []string) int {
	// input = test_08
	ans := 1

	boxes := []*d8_box{}
	for _, position := range input {
		pos := strings.Split(position, ",")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		z, _ := strconv.Atoi(pos[2])

		b := &d8_box{
			pos:  [3]int{x, y, z},
			size: 1,
		}
		b.parent = b
		boxes = append(boxes, b)
	}

	pairs := []d8_pairDist{}
	for i, b1 := range boxes {
		for _, b2 := range boxes[i+1:] {
			pairs = append(pairs, d8_pairDist{
				pair: [2]*d8_box{b1, b2},
				dist: getDistanceVector(b1.pos, b2.pos),
			})
		}
	}

	slices.SortFunc(pairs, func(a, b d8_pairDist) int {
		return cmp.Compare(a.dist, b.dist)
	})

	for i := 0; i < 1000; i++ {
		pair := pairs[i].pair
		d8_union(pair[0], pair[1])
	}

	slices.SortFunc(boxes, func(a, b *d8_box) int {
		return cmp.Compare(b.size, a.size)
	})

	for i := 0; i < 3; i++ {
		ans *= boxes[i].size
	}

	return ans
}

func Day08_2(input []string) int {
	// input = test_08
	ans := 0

	boxes := []*d8_box{}
	for _, position := range input {
		pos := strings.Split(position, ",")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		z, _ := strconv.Atoi(pos[2])

		b := &d8_box{
			pos:  [3]int{x, y, z},
			size: 1,
		}
		b.parent = b
		boxes = append(boxes, b)
	}

	pairs := []d8_pairDist{}
	for i, b1 := range boxes {
		for _, b2 := range boxes[i+1:] {
			pairs = append(pairs, d8_pairDist{
				pair: [2]*d8_box{b1, b2},
				dist: getDistanceVector(b1.pos, b2.pos),
			})
		}
	}

	slices.SortFunc(pairs, func(a, b d8_pairDist) int {
		return cmp.Compare(a.dist, b.dist)
	})

	for _, p := range pairs {
		d8_union(p.pair[0], p.pair[1])

		lastCon := slices.ContainsFunc(boxes, func(b *d8_box) bool {
			return b.size == len(boxes)
		})

		if lastCon {
			ans = p.pair[0].pos[0] * p.pair[1].pos[0]
			break
		}
	}

	// for _, b := range boxes {
	// 	fmt.Printf("pos %+v, size %v\n", b.pos, b.size)
	// }

	return ans
}

var test_08 = strings.Split(`162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`, "\n")

func getDistanceVector(a, b [3]int) float64 {
	dx := float64(a[0] - b[0])
	dy := float64(a[1] - b[1])
	dz := float64(a[2] - b[2])
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type d8_box struct {
	pos    [3]int
	parent *d8_box
	size   int
}

type d8_pairDist struct {
	pair [2]*d8_box
	dist float64
}

func d8_find(b *d8_box) *d8_box {
	if b.parent != b {
		b.parent = d8_find(b.parent)
	}
	return b.parent
}

func d8_union(a, b *d8_box) {
	rootA := d8_find(a)
	rootB := d8_find(b)
	if rootA == rootB {
		return
	}

	rootA.parent = rootB
	rootB.size += rootA.size
	rootA.size = 0
}
