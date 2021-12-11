package grid2d

import "aoc-2021/util/math/vector"

type IntGrid struct {
	Backend GridBackend
}

func (g *IntGrid) Get(position vector.Vec2) int {
	return g.Backend.Get(position).(int)
}

func (g *IntGrid) Set(position vector.Vec2, value int) {
	g.Backend.Set(position, value)
}

func (g *IntGrid) Positions() []vector.Vec2 {
	return g.Backend.Positions()
}

func (g *IntGrid) FilledPositions() []vector.Vec2 {
	return g.Backend.FilledPositions()
}

func (g *IntGrid) Length() int {
	return g.Backend.Length()
}

func (g *IntGrid) Neighbours(position vector.Vec2, includeDiagonals bool) []vector.Vec2 {
	return g.Backend.Neighbours(position, includeDiagonals)
}

func (g *IntGrid) Extents() vector.Vec2 {
	return g.Backend.Extents()
}
