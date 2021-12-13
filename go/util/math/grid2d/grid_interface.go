package grid2d

import "aoc-2021/util/math/vector"

type GridBackend interface {
	Get(position vector.Vec2) interface{}
	Set(position vector.Vec2, value interface{})
	Positions() []vector.Vec2
	FilledPositions() []vector.Vec2
	Length() int
	Neighbours(position vector.Vec2, includeDiagonals bool) []vector.Vec2
	Extents() vector.Vec2
	SetExtents(vector.Vec2)
	ContainsPoint(vector.Vec2) bool
}

type GridGeometry struct {
	extents vector.Vec2
}

func (g *GridGeometry) Length() int {
	return g.extents.X * g.extents.Y
}

func (g *GridGeometry) Extents() vector.Vec2 {
	return g.extents
}

func (g *GridGeometry) SetExtents(e vector.Vec2) {
	g.extents = e
}

func (g *GridGeometry) ContainsPoint(position vector.Vec2) bool {
	if position.X < 0 || position.Y < 0 || position.X >= g.extents.X || position.Y >= g.extents.Y {
		return false
	}
	return true
}

func (g *GridGeometry) Positions() []vector.Vec2 {
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

func (g *GridGeometry) FilledPositions() []vector.Vec2 {
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

func (g *GridGeometry) Neighbours(point vector.Vec2, includeDiagonal bool) []vector.Vec2 {
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
