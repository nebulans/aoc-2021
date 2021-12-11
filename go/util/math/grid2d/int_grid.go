package grid2d

import "aoc-2021/util/math/vector"

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

func MakeIntGrid(backend GridBackend) *IntGrid {
	return &IntGrid{backend: backend}
}
