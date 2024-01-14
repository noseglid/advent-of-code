package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type moon struct {
	x, y, z    int
	vx, vy, vz int
}

func (m moon) Potential() int {
	return (util.Absolute(m.x) + util.Absolute(m.y) + util.Absolute(m.z)) *
		(util.Absolute(m.vx) + util.Absolute(m.vy) + util.Absolute(m.vz))
}

func applyGravity(moons []moon) {
	for i := range moons {
		for j := range moons {
			if moons[i].x < moons[j].x {
				moons[i].vx++
			}
			if moons[i].x > moons[j].x {
				moons[i].vx--
			}
			if moons[i].y < moons[j].y {
				moons[i].vy++
			}
			if moons[i].y > moons[j].y {
				moons[i].vy--
			}
			if moons[i].z < moons[j].z {
				moons[i].vz++
			}
			if moons[i].z > moons[j].z {
				moons[i].vz--
			}
		}
	}
}

func tick(moons []moon) {
	for i := range moons {
		moons[i].x += moons[i].vx
		moons[i].y += moons[i].vy
		moons[i].z += moons[i].vz
	}
}

func print(moons []moon) {
	for _, m := range moons {
		fmt.Println(m)
	}
}

func hash(moons []moon) (string, string, string) {
	return fmt.Sprintf("%d,%d,%d,%d-%d,%d,%d,%d", moons[0].x, moons[1].x, moons[2].x, moons[3].x, moons[0].vx, moons[1].vx, moons[2].vx, moons[3].vx),
		fmt.Sprintf("%d,%d,%d,%d-%d,%d,%d,%d", moons[0].y, moons[1].y, moons[2].y, moons[3].y, moons[0].vy, moons[1].vy, moons[2].vy, moons[3].vy),
		fmt.Sprintf("%d,%d,%d,%d-%d,%d,%d,%d", moons[0].z, moons[1].z, moons[2].z, moons[3].z, moons[0].vz, moons[1].vz, moons[2].vz, moons[3].vz)
}

func main() {

	lines := util.GetFileStrings("2019/Day12/input")

	moons := []moon{}
	for _, l := range lines {
		var x, y, z int
		if _, err := fmt.Sscanf(l, "<x=%d, y=%d, z=%d>", &x, &y, &z); err != nil {
			panic(err)
		}
		moons = append(moons, moon{
			x: x, y: y, z: z,
		})
	}
	moons2 := append([]moon{}, moons...)

	for i := 0; i < 1000; i++ {
		applyGravity(moons)
		tick(moons)
	}

	potential := 0
	for _, m := range moons {
		potential += m.Potential()
	}

	fmt.Printf("Potential (part1): %d\n", potential)

	seenx, seeny, seenz := map[string]bool{}, map[string]bool{}, map[string]bool{}

	n := 0
	xrep, yrep, zrep := 0, 0, 0
	for {
		hx, hy, hz := hash(moons2)
		if xrep == 0 && seenx[hx] {
			xrep = n
		}
		seenx[hx] = true

		if yrep == 0 && seeny[hy] {
			yrep = n
		}
		seeny[hy] = true

		if zrep == 0 && seenz[hz] {
			zrep = n
		}
		seenz[hz] = true

		if xrep != 0 && yrep != 0 && zrep != 0 {
			break
		}

		applyGravity(moons2)
		tick(moons2)
		n++
	}

	fmt.Printf("Time until universe repeats (part2): %d\n", util.LCM(xrep, yrep, zrep))

}
