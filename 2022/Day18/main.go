package main

import (
	"fmt"
	"math"

	"github.com/noseglid/advent-of-code/util"
)

type blob struct {
	droplets []droplet
}

func (b blob) Contains(d droplet) bool {
	for _, dd := range b.droplets {
		if dd == d {
			return true
		}
	}
	return false
}

func (b blob) DropletBelongs(d droplet) bool {
	for _, dd := range b.droplets {
		if dd == d {
			return true
		}
		if dd.Adjacent(d) {
			return true
		}
	}
	return false
}

func (b blob) Adjacent(b2 blob) bool {
	for _, d1 := range b.droplets {
		for _, d2 := range b2.droplets {
			if d1.Adjacent(d2) {
				return true
			}
		}
	}
	return false
}

type droplet struct {
	x, y, z int
}

func (d1 droplet) Adjacent(d2 droplet) bool {
	if d1 == d2 {
		return false
	}
	dx, dy, dz := util.Absolute(d1.x-d2.x), util.Absolute(d1.y-d2.y), util.Absolute(d1.z-d2.z)
	return dx == 0 && dy == 0 && dz == 1 || dx == 0 && dy == 1 && dz == 0 || dx == 1 && dy == 0 && dz == 0
}

func surface(droplets []droplet) int {
	n := 0
	for _, d1 := range droplets {
		freeSides := 6
		for _, d2 := range droplets {
			if d1.Adjacent(d2) {
				freeSides--
			}
		}
		n += freeSides
	}

	return n
}

type cuboid struct {
	x, y, z int
	w, h, d int
}

func bounds(droplets []droplet) cuboid {
	minx, maxx, miny, maxy, minz, maxz := math.MaxInt, 0, math.MaxInt, 0, math.MaxInt, 0
	for _, d := range droplets {
		minx = util.Min(minx, d.x)
		maxx = util.Max(maxx, d.x)
		miny = util.Min(miny, d.y)
		maxy = util.Max(maxy, d.y)
		minz = util.Min(minz, d.z)
		maxz = util.Max(maxz, d.z)
	}
	return cuboid{minx, miny, minz, maxx - minx, maxy - miny, maxz - minz}
}

func main() {
	var droplets []droplet

	// input := []string{
	// 	"1,1,1", "2,1,1", "3,1,1",
	// 	"1,2,1", "2,2,1", "3,2,1",
	// 	"1,3,1", "2,3,1", "3,3,1",
	// 	"1,1,2", "2,1,2", "3,1,2",
	// 	"1,2,2", "3,2,2",
	// 	"1,3,2", "2,3,2", "3,3,2",
	// 	"1,1,3", "2,1,3", "3,1,3",
	// 	"1,2,3", "2,2,3", "3,2,3",
	// 	"1,3,3", "2,3,3", "3,3,3",
	// }
	// input := util.GetFileStrings("2022/Day18/sample")
	input := util.GetFileStrings("2022/Day18/input")

	for _, l := range input {
		var d droplet
		if _, err := fmt.Sscanf(l, "%d,%d,%d", &d.x, &d.y, &d.z); err != nil {
			panic(err)
		}

		droplets = append(droplets, d)
	}
	fmt.Printf("Surface area (part1): %d\n", surface(droplets))

	bounds := bounds(droplets)
	lavaBlob := blob{droplets: droplets}
	waterBlobs := []blob{}
	fmt.Printf("bounds=%v\n", bounds)

	fmt.Printf("building blobs...\n")
	for x := bounds.x - 1; x <= bounds.w+2; x++ {
		for y := bounds.y - 1; y <= bounds.h+2; y++ {
		C:
			for z := bounds.z - 1; z <= bounds.d+2; z++ {
				dd := droplet{x, y, z}
				if lavaBlob.Contains(dd) {
					continue C
				}

				for _, blob := range waterBlobs {
					if blob.Contains(dd) {
						blob.droplets = append(blob.droplets, dd)
						continue C
					}
				}
				// no existing blob, create new blob
				blob := blob{}
				blob.droplets = append(blob.droplets, dd)
				waterBlobs = append(waterBlobs, blob)
			}
		}
	}

	fmt.Printf("merging blobs...\n")
Outer:
	for {
		fmt.Printf("%d blobs left\n", len(waterBlobs))
		for k, w1 := range waterBlobs {
			for i, w2 := range waterBlobs {
				if k == i {
					continue
				}
				if w1.Adjacent(w2) {
					waterBlobs[k].droplets = append(w1.droplets, w2.droplets...)
					waterBlobs = append(waterBlobs[:i], waterBlobs[i+1:]...)
					continue Outer
				}
			}
		}
		break
	}

	largestBlobIndex, ll := 0, len(waterBlobs[0].droplets)
	for i := 1; i < len(waterBlobs); i++ {
		if len(waterBlobs[i].droplets) > ll {
			ll = len(waterBlobs[i].droplets)
			largestBlobIndex = i
		}
	}
	waterBlobs = append(waterBlobs[:largestBlobIndex], waterBlobs[largestBlobIndex+1:]...)

	interiorSize := 0
	for _, w := range waterBlobs {
		interiorSize += surface(w.droplets)
	}

	fmt.Printf("Exterior surface area (part2): %d\n", surface(lavaBlob.droplets)-interiorSize)
}
