package main

import (
	"fmt"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type Num interface {
	Explode(depth int) bool
	Split() bool
	Parent() *SnailNum
	SetParent(sn *SnailNum)
	Magnitude() int
	Reduce() Num
	Add(rhs Num) *SnailNum
}

type SnailNum struct {
	parent   *SnailNum
	lhs, rhs Num
}

func NewSnailNum(lhs, rhs Num, parent ...*SnailNum) *SnailNum {
	var pp *SnailNum
	if len(parent) >= 1 {
		pp = parent[0]
	}
	sn := SnailNum{
		parent: pp,
	}
	sn.lhs = lhs
	sn.rhs = rhs
	sn.lhs.SetParent(&sn)
	sn.rhs.SetParent(&sn)
	return &sn
}

func (s *SnailNum) String() string {
	return fmt.Sprintf("[%s,%s]", s.lhs, s.rhs)
}

func (s *SnailNum) Magnitude() int {
	return 3*s.lhs.Magnitude() + 2*s.rhs.Magnitude()
}

func (s *SnailNum) findLeftNestedRegNum() *RegNum {
	for current := s; current != nil; current = current.lhs.(*SnailNum) {
		if rn, ok := current.lhs.(*RegNum); ok {
			return rn
		}
	}
	panic("no left nested regnum")
}

func (s *SnailNum) findRightNestedRegNum() *RegNum {
	for current := s; current != nil; current = current.rhs.(*SnailNum) {
		if rn, ok := current.rhs.(*RegNum); ok {
			return rn
		}
	}
	panic("no right nested regnum")
}

func (s *SnailNum) findSiblingRight() (*RegNum, bool) {
	list := []*SnailNum{s}
	prev, current := s, s
	for {
		if current = current.parent; current == nil {
			return nil, false
		}
		list = append(list, current)
		if rn, ok := current.rhs.(*RegNum); ok {
			return rn, true
		}
		if sn, ok := current.rhs.(*SnailNum); ok && sn != prev {
			return sn.findLeftNestedRegNum(), true
		}
		prev = current
	}
}

func (s *SnailNum) findSiblingLeft() (*RegNum, bool) {
	list := []*SnailNum{s}
	prev, current := s, s
	for {
		if current = current.parent; current == nil {
			return nil, false
		}
		list = append(list, current)
		if rn, ok := current.lhs.(*RegNum); ok {
			return rn, true
		}
		if sn, ok := current.lhs.(*SnailNum); ok && sn != prev {
			return sn.findRightNestedRegNum(), true
		}
		prev = current
	}
}

func (s *SnailNum) Explode(depth int) bool {
	if depth == 4 {
		if leftNum, ok := s.findSiblingLeft(); ok {
			leftNum.value += s.lhs.(*RegNum).value
		}
		if rightNum, ok := s.findSiblingRight(); ok {
			rightNum.value += s.rhs.(*RegNum).value
		}

		switch s {
		case s.parent.lhs:
			s.parent.lhs = NewRegNum(0, s.parent)
		case s.parent.rhs:
			s.parent.rhs = NewRegNum(0, s.parent)
		}
		return true
	}

	return s.lhs.Explode(depth+1) || s.rhs.Explode(depth+1) || false
}

func (s *SnailNum) Reduce() Num {
	for {
		if s.Explode(0) {
			continue
		}

		if s.Split() {
			continue
		}

		break
	}
	return s
}

func (s *SnailNum) Add(rhs Num) *SnailNum {
	return NewSnailNum(s, rhs)
}

func (s *SnailNum) Split() bool {
	splitNums := func(value int, parent *SnailNum) *SnailNum {
		v1, v2 := value/2, value/2
		if value%2 != 0 {
			v2++
		}

		return NewSnailNum(NewRegNum(v1), NewRegNum(v2), parent)
	}

	if rn, ok := s.lhs.(*RegNum); ok && rn.value >= 10 {
		s.lhs = splitNums(rn.value, s)
		return true
	} else if !ok && s.lhs.(*SnailNum).Split() {
		return true
	} else if rn, ok := s.rhs.(*RegNum); ok && rn.value >= 10 {
		s.rhs = splitNums(rn.value, s)
		return true
	} else if !ok && s.rhs.(*SnailNum).Split() {
		return true
	}

	return false
}

func (s *SnailNum) Parent() *SnailNum {
	return s.parent
}

func (s *SnailNum) SetParent(p *SnailNum) {
	s.parent = p
}

type RegNum struct {
	parent *SnailNum
	value  int
}

func NewRegNum(value int, parent ...*SnailNum) *RegNum {
	var pp *SnailNum
	if len(parent) >= 1 {
		pp = parent[0]
	}

	return &RegNum{
		value:  value,
		parent: pp,
	}
}

func (r *RegNum) String() string {
	return fmt.Sprintf("%d", r.value)
}

func (r *RegNum) Explode(depth int) bool {
	return false
}

func (r *RegNum) Split() bool {
	return false
}

func (r *RegNum) Parent() *SnailNum {
	return r.parent
}

func (r *RegNum) SetParent(sn *SnailNum) {
	r.parent = sn
}

func (r *RegNum) Magnitude() int {
	return r.value
}

func (r *RegNum) Reduce() Num {
	return r
}
func (r *RegNum) Add(rhs Num) *SnailNum {
	return NewSnailNum(r, rhs)
}

func numDivideIndex(s string) int {
	idx, d := 0, 0
Loop:
	for i, r := range s {
		switch r {
		case '[':
			d++
		case ']':
			d--
		case ',':
			if d == 0 {
				idx = i
				break Loop
			}
		}
	}

	return idx
}

func parse(s string) Num {
	if len(s) == 1 {
		return NewRegNum(util.MustAtoi(s))
	}

	s = s[1 : len(s)-1]
	idx := numDivideIndex(s)
	return NewSnailNum(parse(s[:idx]), parse(s[idx+1:]))
}

func main() {
	input := "2021/Day18/input"
	lines := util.GetFileStrings(input)

	n := parse(lines[0])
	for _, l := range lines[1:] {
		n = n.Add(parse(l)).Reduce()
	}

	max := 0
	for i, l1 := range lines {
		for j, l2 := range lines {
			if i == j {
				continue
			}

			if s := parse(l1).Add(parse(l2)).Reduce().Magnitude(); s > max {
				max = s
			}
		}
	}

	log.Printf("Part 1: Magnitude=%d", n.Magnitude())
	log.Printf("Part 2: Largest sum magnitude=%d", max)
}
