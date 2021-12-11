package grid2d

import "aoc-2021/util/math/vector"

type MapGrid struct {
	values       map[vector.Vec2]interface{}
	extents      vector.Vec2
	defaultValue interface{}
}

func (g *MapGrid) Length() int {
	return g.extents.X * g.extents.Y
}

func (g *MapGrid) Extents() vector.Vec2 {
	return g.extents
}

func (g *MapGrid) Get(position vector.Vec2) interface{} {
	val, found := g.values[position]
	if found {
		return val
	}
	return g.defaultValue
}

func (g *MapGrid) Set(position vector.Vec2, value interface{}) {
	g.values[position] = value
	if position.X > g.extents.X {
		g.extents.X = position.X
	}
	if position.Y > g.extents.Y {
		g.extents.Y = position.Y
	}
}

func (g *MapGrid) ContainsPoint(position vector.Vec2) bool {
	if position.X < 0 || position.Y < 0 || position.X >= g.extents.X || position.Y >= g.extents.Y {
		return false
	}
	return true
}

func (g *MapGrid) Positions() []vector.Vec2 {
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

func (g *MapGrid) FilledPositions() []vector.Vec2 {
	filled := make([]vector.Vec2, len(g.values))
	i := 0
	for pos := range g.values {
		filled[i] = pos
		i++
	}
	return filled
}

func (g *MapGrid) Neighbours(point vector.Vec2, includeDiagonal bool) []vector.Vec2 {
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

func MakeMapGrid(defaultValue interface{}) *MapGrid {
	return &MapGrid{
		values:       map[vector.Vec2]interface{}{},
		extents:      vector.Vec2{X: 0, Y: 0},
		defaultValue: defaultValue,
	}
}
