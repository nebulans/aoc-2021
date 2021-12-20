package grid2d

import (
	"aoc-2021/util/math/vector"
	"fmt"
	"strings"
)

type IntGrid struct {
	GridBackend
}

func (g *IntGrid) Get(position vector.Vec2) int {
	v := g.GridBackend.Get(position)
	if v == nil {
		return 0
	}
	return v.(int)
}

func (g *IntGrid) Set(position vector.Vec2, value int) {
	g.GridBackend.Set(position, value)
}

func (g *IntGrid) Format(valueFormatter func(int) string) string {
	elems := make([]string, g.Length())
	for i, pos := range g.Positions() {
		val := g.Get(pos)
		formatted := valueFormatter(val)
		if pos.X == g.GridBackend.Extents().X-1 {
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

func (g *IntGrid) BoolFormatter(val int) string {
	if val == 0 {
		return "."
	}
	return "#"
}

func (g *IntGrid) PaddedFormatter(val int) string {
	return fmt.Sprintf("%3d ", val)
}

func MakeIntGrid(backend GridBackend) *IntGrid {
	return &IntGrid{GridBackend: backend}
}
