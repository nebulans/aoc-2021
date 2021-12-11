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
}
