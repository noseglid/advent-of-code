package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

type P struct {
	X, Y, Z int
}

func (p P) Len() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
}

type Vec struct {
	I       int
	P, V, A P
}

func checkCollissions(vecs []Vec) []int {
	collides := []int{}
	occupied := map[P][]int{}

	for _, v := range vecs {
		occupied[v.P] = append(occupied[v.P], v.I)
	}

	for _, o := range occupied {
		if len(o) > 1 {
			collides = append(collides, o...)
		}
	}

	return collides
}

func simulate(vecs []Vec) {
	for i := range vecs {
		vecs[i].V.X += vecs[i].A.X
		vecs[i].V.Y += vecs[i].A.Y
		vecs[i].V.Z += vecs[i].A.Z
		vecs[i].P.X += vecs[i].V.X
		vecs[i].P.Y += vecs[i].V.Y
		vecs[i].P.Z += vecs[i].V.Z
	}
}

func main() {
	lines := util.GetFileStrings("2017/Day20/input")
	vecs := []Vec{}

	for i, l := range lines {
		var vec Vec
		fmt.Sscanf(l, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>", &vec.P.X, &vec.P.Y, &vec.P.Z, &vec.V.X, &vec.V.Y, &vec.V.Z, &vec.A.X, &vec.A.Y, &vec.A.Z)
		vec.I = i
		vecs = append(vecs, vec)
	}

	slices.SortFunc(vecs, func(lhs, rhs Vec) int {
		return int(lhs.A.Len() - rhs.A.Len())
	})

	fmt.Printf("Closest to origin (part1): %d\n", vecs[0].I)

	for i := 0; i < 20000; i++ {
		simulate(vecs)
		collisions := checkCollissions(vecs)
		for _, I := range collisions {
			for i, v := range vecs {
				if v.I == I {
					vecs = append(vecs[:i], vecs[i+1:]...)
					break
				}
			}
		}
	}

	fmt.Printf("Number left after collisions (part2): %d\n", len(vecs))

}
