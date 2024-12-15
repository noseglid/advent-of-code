package util

type Grid [][]rune

type Dir int

const (
	N = iota
	NE
	E
	SE
	S
	SW
	W
	NW
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

func (g *Grid) Switch(ax, ay, bx, by int) {
	gg := [][]rune(*g)
	gg[by][bx], gg[ay][ax] = gg[ay][ax], gg[by][bx]
}
