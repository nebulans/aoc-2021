package grid2d

import (
	"aoc-2021/util/math/vector"
)

type ArrayGrid struct {
	values  []interface{}
	Extents vector.Vec2
}

func (g *ArrayGrid) Length() int {
	return g.Extents.X * g.Extents.Y
}

func (g *ArrayGrid) Get(position vector.Vec2) interface{} {
	i := position.X + position.Y*g.Extents.X
	return g.values[i]
}

func (g *ArrayGrid) Set(position vector.Vec2, value interface{}) {
	i := position.X + position.Y*g.Extents.X
	g.values[i] = value
}

func (g *ArrayGrid) ContainsPoint(position vector.Vec2) bool {
	if position.X < 0 || position.Y < 0 || position.X >= g.Extents.X || position.Y >= g.Extents.Y {
		return false
	}
	return true
}

func (g *ArrayGrid) Positions() []vector.Vec2 {
	pos := make([]vector.Vec2, g.Extents.X*g.Extents.Y)
	i := 0
	for y := 0; y < g.Extents.Y; y++ {
		for x := 0; x < g.Extents.X; x++ {
			pos[i] = vector.Vec2{X: x, Y: y}
			i++
		}
	}
	return pos
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
		Extents: extents,
	}
}
