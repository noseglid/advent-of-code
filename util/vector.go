package util

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) Set(x, y int) {
	p.X = x
	p.Y = y
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
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
