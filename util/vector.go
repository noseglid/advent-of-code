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
