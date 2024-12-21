package util

type Grid [][]rune

type Dir int
type RelDir int

const (
	N Dir = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func (d Dir) String() string {
	switch d {
	case N:
		return "N"
	case NE:
		return "NE"
	case E:
		return "E"
	case SE:
		return "SE"
	case S:
		return "S"
	case SW:
		return "SW"
	case W:
		return "W"
	case NW:
		return "NW"
	}
	panic("bad dir")
}

const (
	Right RelDir = iota
	Left
)

func (d Dir) Deltas() (int, int) {
	switch d {
	case N:
		return 0, -1
	case NE:
		return 1, -1
	case E:
		return 1, 0
	case SE:
		return 1, 1
	case S:
		return 0, 1
	case SW:
		return -1, 1
	case W:
		return -1, 0
	case NW:
		return -1, 1
	}
	panic("bad dir")
}

func (d Dir) Turn(rd RelDir) Dir {
	switch d {
	case N:
		if rd == Left {
			return W
		} else {
			return E
		}
	case E:
		if rd == Left {
			return N
		} else {
			return S
		}
	case S:
		if rd == Left {
			return E
		} else {
			return W
		}
	case W:
		if rd == Left {
			return S
		} else {
			return N
		}
	}
	panic("Bad dir")
}

func VDir(r rune) Dir {
	switch r {
	case '^':
		return N
	case '>':
		return E
	case 'v':
		return S
	case '<':
		return W
	}

	panic("bad r")
}

func (g Grid) Print() {
	PrintRuneGrid(g)
}

func (g Grid) Find(r rune) (int, int) {
	for y, row := range g {
		for x, cell := range row {
			if r == cell {
				return x, y
			}
		}
	}
	panic("no find")
}

func (g Grid) Get(x, y int) rune {
	return g[y][x]
}
func (g Grid) Set(x, y int, r rune) {
	g[y][x] = r
}

func (g Grid) ShortestPath(start, end Point, isMovable func(x, y int, r rune) bool) (int, []Point, bool) {
	nodes := g.ShortestPathNodes(start, end, isMovable)

	var path []Point
	current := start
Outer:
	for current != end {

		path = append(path, current)

		for _, m := range []Dir{N, W, E, S} {
			rx, ry := g.GetMove(current.X, current.Y, m)
			p := Point{rx, ry}
			if v, ok := nodes[p]; ok && v == nodes[current]-1 {
				current = p
				continue Outer
			}
		}

		// No path
		return 0, nil, false
	}
	path = append(path, end)

	return nodes[start], path, true
}

func (g Grid) ShortestPathNodes(start, end Point, isMovable func(x, y int, r rune) bool) map[Point]int {
	unvisited := []Point{end}
	nodes := map[Point]int{
		end: 0,
	}

	for len(unvisited) > 0 {
		n := unvisited[0]
		for _, m := range []Dir{N, E, S, W} {
			rx, ry := g.GetMove(n.X, n.Y, m)
			pp := Point{rx, ry}
			if g.InBounds(rx, ry) && (pp == start || isMovable(rx, ry, g.Get(rx, ry))) {
				in := Point{rx, ry}
				if v, ok := nodes[in]; !ok {
					nodes[in] = nodes[n] + 1
					unvisited = append(unvisited, in)
				} else if nodes[n]+1 < v {
					nodes[in] = v
				}
			}
		}
		unvisited = unvisited[1:]
	}

	return nodes
}

func (g Grid) Each(f func(x, y int, r rune)) {
	for y, row := range g {
		for x, cell := range row {
			f(x, y, cell)
		}
	}
}

func (g Grid) GetMove(x, y int, d Dir) (int, int) {
	switch d {
	case N:
		return x, y - 1
	case NE:
		return x + 1, y - 1
	case E:
		return x + 1, y
	case SE:
		return x + 1, y + 1
	case S:
		return x, y + 1
	case SW:
		return x - 1, y + 1
	case W:
		return x - 1, y
	case NW:
		return x - 1, y - 1
	}

	panic("bad dir")
}

func (g Grid) InBounds(x, y int) bool {
	return y >= 0 && y < len(g) && x >= 0 && x < len(g[y])
}

func (g *Grid) Switch(ax, ay, bx, by int) {
	gg := [][]rune(*g)
	gg[by][bx], gg[ay][ax] = gg[ay][ax], gg[by][bx]
}
