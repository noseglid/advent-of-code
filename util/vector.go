package util

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p *Point) Set(x, y int) {
	p.X = x
	p.Y = y
}

func (p Point) Move(dir Dir) Point {
	dx, dy := dir.Deltas()
	return Point{p.X + dx, p.Y + dy}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func Point3DFrom(s string) Point3D {
	var p Point3D
	n, err := fmt.Sscanf(s, "%d,%d,%d", &p.X, &p.Y, &p.Z)
	if err != nil {
		panic(err)
	}
	if n != 3 {
		panic("invalid point 3d: " + s)
	}
	return p
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.X, p.Y, p.Z)
}

func (p Point3D) Manhattan(o Point3D) int {
	return Absolute(p.X-o.X) + Absolute(p.Y-o.Y) + Absolute(p.Z-o.Z)
}

func (p Point3D) Euclidean(o Point3D) float64 {
	return math.Sqrt(float64((p.X-o.X)*(p.X-o.X) + (p.Y-o.Y)*(p.Y-o.Y) + (p.Z-o.Z)*(p.Z-o.Z)))
}
