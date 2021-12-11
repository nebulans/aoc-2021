package grid2d

import (
	"aoc-2021/util/math/vector"
)

type ArrayGrid struct {
	GridGeometry
	values []interface{}
}

func (g *ArrayGrid) Get(position vector.Vec2) interface{} {
	i := position.X + position.Y*g.extents.X
	return g.values[i]
}

func (g *ArrayGrid) Set(position vector.Vec2, value interface{}) {
	i := position.X + position.Y*g.extents.X
	g.values[i] = value
}

func MakeArrayGrid(extents vector.Vec2) *ArrayGrid {
	return &ArrayGrid{
		values:       make([]interface{}, extents.X*extents.Y),
		GridGeometry: GridGeometry{extents},
	}
}
