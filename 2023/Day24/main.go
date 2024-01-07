package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type Vec3 struct {
	X, Y, Z float64
}
type hail struct {
	p Vec3
	v Vec3
}

// 20, 19, 15 @  1, -5, -3
func parseHail(s string) hail {
	var h hail
	_, err := fmt.Sscanf(s, "%f, %f, %f @ %f, %f, %f", &h.p.X, &h.p.Y, &h.p.Z, &h.v.X, &h.v.Y, &h.v.Z)
	if err != nil {
		panic(err)
	}
	return h
}

func equation(h hail) (float64, float64) {
	k := h.v.Y / h.v.X
	return k, h.p.Y - k*h.p.X
}

func intersects(h1, h2 hail, min, max float64) (bool, float64, float64) {
	k1, m1 := equation(h1)
	k2, m2 := equation(h2)
	if k1 == k2 {
		// parallell
		return false, 0, 0
	}

	xi := (m2 - m1) / (k1 - k2)
	yi := k1*xi + m1
	t1 := (xi - h1.p.X) / h1.v.X
	t2 := (xi - h2.p.X) / h2.v.X
	if t1 < 0 || t2 < 0 {
		return false, 0, 0
	}

	return xi > min && yi > min && xi < max && yi < max, h1.p.Z + h1.v.Z*t1, h2.p.Z + h2.v.Z*t2
}

func p2(hails []hail) {
	V := 500.0

	h1, h2 := hails[0], hails[1]

	for vx := -V; vx <= V; vx++ {
		for vy := -V; vy <= V; vy++ {
			for vz := -V; vz <= V; vz++ {
				A, a, B, b, C, c, D, d := h1.p.X, h1.v.X-vx, h1.p.Y, h1.v.Y-vy, h2.p.X, h2.v.X-vx, h2.p.Y, h2.v.Y-vy
				t := (d*(C-A) - c*(D-B)) / ((a * d) - (b * c))
				x := h1.p.X + h1.v.X*t - vx*t
				y := h1.p.Y + h1.v.Y*t - vy*t
				z := h1.p.Z + h1.v.Z*t - vz*t

				hit := true
				for i := 0; i < len(hails); i++ {
					h := hails[i]
					var k float64
					if h.v.X != vx {
						k = (x - h.p.X) / (h.v.X - vx)
					} else if h.v.Y != vy {
						k = (y - h.p.Y) / (h.v.Y - vy)
					} else if h.v.Z != vz {
						k = (z - h.p.Z) / (h.v.Z - vz)
					} else {
						continue
					}
					if (x+k*vx != h.p.X+k*h.v.X) || (y+k*vy != h.p.Y+k*h.v.Y) || (z+k*vz != h.p.Z+k*h.v.Z) {
						hit = false
						break
					}
				}

				if hit {
					fmt.Printf("Sum of position (part2): %d\n", int(x+y+z))
					return
				}

			}
		}
	}
}

func main() {
	lines := util.GetFileStrings("2023/Day24/input")
	// min, max := 7.0, 27.0
	min, max := 200000000000000.0, 400000000000000.0

	var hails []hail
	for _, s := range lines {
		h := parseHail(s)
		hails = append(hails, h)
	}

	n := 0
	for i := 0; i < len(hails); i++ {
		for j := i + 1; j < len(hails); j++ {
			if ok, _, _ := intersects(hails[i], hails[j], min, max); ok {
				n++
			}
		}
	}
	fmt.Printf("num intersections (part1): %d\n", n)

	p2(hails)
}
