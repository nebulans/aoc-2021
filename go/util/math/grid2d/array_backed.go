package grid2d

import (
	"aoc-2021/util/math/vector"
)

type ArrayGrid struct {
	values  []interface{}
	extents vector.Vec2
}

func (g *ArrayGrid) Length() int {
	return g.extents.X * g.extents.Y
}

func (g *ArrayGrid) Extents() vector.Vec2 {
	return g.extents
}

func (g *ArrayGrid) Get(position vector.Vec2) interface{} {
	i := position.X + position.Y*g.extents.X
	return g.values[i]
}

func (g *ArrayGrid) Set(position vector.Vec2, value interface{}) {
	i := position.X + position.Y*g.extents.X
	g.values[i] = value
}

func (g *ArrayGrid) ContainsPoint(position vector.Vec2) bool {
	if position.X < 0 || position.Y < 0 || position.X >= g.extents.X || position.Y >= g.extents.Y {
		return false
	}
	return true
}

func (g *ArrayGrid) Positions() []vector.Vec2 {
	pos := make([]vector.Vec2, g.extents.X*g.extents.Y)
	i := 0
	for y := 0; y < g.extents.Y; y++ {
		for x := 0; x < g.extents.X; x++ {
			pos[i] = vector.Vec2{X: x, Y: y}
			i++
		}
	}
	return pos
}

func (g *ArrayGrid) FilledPositions() []vector.Vec2 {
	return g.Positions()
}

var axisNeighbours = [4]vector.Vec2{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

var diagonalNeighbours = [4]vector.Vec2{
	{1, 1},
	{1, -1},
	{-1, -1},
	{-1, 1},
}

func (g *ArrayGrid) Neighbours(point vector.Vec2, includeDiagonal bool) []vector.Vec2 {
	vecs := axisNeighbours[:]
	if includeDiagonal {
		vecs = append(vecs, diagonalNeighbours[:]...)
	}
	n := make([]vector.Vec2, 0, len(vecs))
	for _, v := range vecs {
		p := point.Add(v)
		if g.ContainsPoint(p) {
			n = append(n, p)
		}
	}
	return n
}

func MakeArrayGrid(extents vector.Vec2) *ArrayGrid {
	return &ArrayGrid{
		values:  make([]interface{}, extents.X*extents.Y),
		extents: extents,
	}
}
