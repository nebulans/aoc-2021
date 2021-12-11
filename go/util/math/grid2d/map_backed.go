package grid2d

import "aoc-2021/util/math/vector"

type MapGrid struct {
	GridGeometry
	values       map[vector.Vec2]interface{}
	defaultValue interface{}
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

func (g *MapGrid) FilledPositions() []vector.Vec2 {
	filled := make([]vector.Vec2, len(g.values))
	i := 0
	for pos := range g.values {
		filled[i] = pos
		i++
	}
	return filled
}

func MakeMapGrid(defaultValue interface{}) *MapGrid {
	return &MapGrid{
		values:       map[vector.Vec2]interface{}{},
		GridGeometry: GridGeometry{vector.Vec2{X: 0, Y: 0}},
		defaultValue: defaultValue,
	}
}

func MakeMapGridExtents(defaultValue interface{}, extents vector.Vec2) *MapGrid {
	return &MapGrid{
		values:       map[vector.Vec2]interface{}{},
		GridGeometry: GridGeometry{extents},
		defaultValue: defaultValue,
	}
}
