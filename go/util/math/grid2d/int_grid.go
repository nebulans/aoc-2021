package grid2d

import (
	"aoc-2021/util/math/vector"
	"fmt"
	"strings"
)

type IntGrid struct {
	backend GridBackend
}

func (g *IntGrid) Get(position vector.Vec2) int {
	return g.backend.Get(position).(int)
}

func (g *IntGrid) Set(position vector.Vec2, value int) {
	g.backend.Set(position, value)
}

func (g *IntGrid) Positions() []vector.Vec2 {
	return g.backend.Positions()
}

func (g *IntGrid) FilledPositions() []vector.Vec2 {
	return g.backend.FilledPositions()
}

func (g *IntGrid) Length() int {
	return g.backend.Length()
}

func (g *IntGrid) Neighbours(position vector.Vec2, includeDiagonals bool) []vector.Vec2 {
	return g.backend.Neighbours(position, includeDiagonals)
}

func (g *IntGrid) Extents() vector.Vec2 {
	return g.backend.Extents()
}

func (g *IntGrid) Format(valueFormatter func(int) string) string {
	elems := make([]string, g.Length())
	for i, pos := range g.Positions() {
		val := g.Get(pos)
		formatted := valueFormatter(val)
		if pos.X == g.backend.Extents().X-1 {
			elems[i] = fmt.Sprintf("%s\n", formatted)
		} else {
			elems[i] = formatted
		}
	}
	return strings.Join(elems, "")
}

func (g *IntGrid) DefaultFormatter(val int) string {
	return fmt.Sprintf("%d", val)
}

func (g *IntGrid) HexFormatter(val int) string {
	return fmt.Sprintf("%x", val)
}

func (g *IntGrid) SparseFormatter(val int) string {
	if val == 0 {
		return " "
	}
	return fmt.Sprintf("%d", val)
}

func MakeIntGrid(backend GridBackend) *IntGrid {
	return &IntGrid{backend: backend}
}
