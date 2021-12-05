package vector

import (
	"aoc-2021/util/math/integer"
	"fmt"
)

type EndpointLine struct {
	Start     Vec2
	End       Vec2
	Direction Vec2
}

func NewEndpointLine(start Vec2, end Vec2) *EndpointLine {
	direction := Vec2{
		integer.UnitStep(end.X - start.X),
		integer.UnitStep(end.Y - start.Y),
	}
	return &EndpointLine{
		Start:     start,
		End:       end,
		Direction: direction,
	}
}

func (line *EndpointLine) Format() string {
	return fmt.Sprintf("(%d,%d) -> (%d,%d)", line.Start.X, line.Start.Y, line.End.X, line.End.Y)
}

func (line *EndpointLine) Points() []Vec2 {
	length := integer.Max(integer.Abs(line.End.X-line.Start.X), integer.Abs(line.End.Y-line.Start.Y))
	points := make([]Vec2, 0, length)
	pos := line.Start
	points = append(points, pos)
	for pos != line.End {
		pos = pos.Add(line.Direction)
		points = append(points, pos)
	}
	return points
}
