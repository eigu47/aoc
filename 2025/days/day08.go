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

	type box struct {
		pos    [3]int
		parent *box
		size   int
	}

	boxes := []*box{}
	for _, position := range input {
		pos := strings.Split(position, ",")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		z, _ := strconv.Atoi(pos[2])

		b := &box{
			pos:  [3]int{x, y, z},
			size: 1,
		}
		b.parent = b
		boxes = append(boxes, b)
	}

	type pairDist struct {
		pair [2]*box
		dist float64
	}

	pairs := []pairDist{}
	for i, b1 := range boxes {
		for _, b2 := range boxes[i+1:] {
			pairs = append(pairs, pairDist{
				pair: [2]*box{b1, b2},
				dist: getDistanceVector(b1.pos, b2.pos),
			})
		}
	}

	slices.SortFunc(pairs, func(a, b pairDist) int {
		return cmp.Compare(a.dist, b.dist)
	})

	var find func(*box) *box
	find = func(b *box) *box {
		if b.parent != b {
			b.parent = find(b.parent)
		}
		return b.parent
	}

	union := func(a, b *box) {
		rootA := find(a)
		rootB := find(b)
		if rootA == rootB {
			return
		}

		rootA.parent = rootB
		rootB.size += rootA.size
		rootA.size = 0
	}

	for i := 0; i < 1000; i++ {
		pair := pairs[i].pair
		union(pair[0], pair[1])
	}

	slices.SortFunc(boxes, func(a, b *box) int {
		return cmp.Compare(b.size, a.size)
	})

	for i := 0; i < 3; i++ {
		ans *= boxes[i].size
	}

	return ans
}

func Day08_2(input []string) int {
	input = test_08
	ans := 0

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

type unionFind struct {
	parent []int
}

func newUnionFind(len int) *unionFind {
	parent := make([]int, len)
	for i := range len {
		parent[i] = i
	}
	return &unionFind{parent: parent}
}

func (u *unionFind) find(i int) int {
	if u.parent[i] != i {
		u.parent[i] = u.find(u.parent[i])
	}

	return u.parent[i]
}

func (u *unionFind) union(a, b int) {
	rootA := u.find(a)
	rootB := u.find(b)
	u.parent[rootA] = rootB
}
