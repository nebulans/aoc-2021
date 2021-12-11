package grid2d

import (
	"aoc-2021/util/math/vector"
	"fmt"
	"testing"
)

func TestArrayGrid_AlignedNeighbours(t *testing.T) {
	var tests = []struct {
		pos        vector.Vec2
		neighbours []vector.Vec2
	}{
		{vector.Vec2{1, 1}, []vector.Vec2{
			{0, 1},
			{1, 0},
			{1, 2},
			{2, 1},
		}},
		{vector.Vec2{2, 0}, []vector.Vec2{
			{1, 0},
			{2, 1},
			{3, 0},
		}},
		{vector.Vec2{4, 4}, []vector.Vec2{
			{3, 4},
			{4, 3},
		}},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Neighbours of %v", tt.pos)
		t.Run(testName, func(t *testing.T) {
			g := MakeArrayGrid(vector.Vec2{5, 5})
			ns := g.Neighbours(tt.pos, false)
			if len(ns) != len(tt.neighbours) {
				t.Errorf("expected %d neighbours, got %d", len(tt.neighbours), len(ns))
			}
			for _, e := range tt.neighbours {
				found := false
				for _, n := range ns {
					if n == e {
						found = true
					}
				}
				if !found {
					t.Errorf("expected neighbour %v not found in %v", e, ns)
				}
			}
		})
	}
}

func TestArrayGrid_DiagonalNeighbours(t *testing.T) {
	var tests = []struct {
		pos        vector.Vec2
		neighbours []vector.Vec2
	}{
		{vector.Vec2{1, 1}, []vector.Vec2{
			{0, 0},
			{0, 1},
			{0, 2},
			{1, 0},
			{1, 2},
			{2, 0},
			{2, 1},
			{2, 2},
		}},
		{vector.Vec2{2, 0}, []vector.Vec2{
			{1, 0},
			{1, 1},
			{2, 1},
			{3, 1},
			{3, 0},
		}},
		{vector.Vec2{4, 4}, []vector.Vec2{
			{3, 3},
			{3, 4},
			{4, 3},
		}},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Neighbours of %v", tt.pos)
		t.Run(testName, func(t *testing.T) {
			g := MakeArrayGrid(vector.Vec2{5, 5})
			ns := g.Neighbours(tt.pos, true)
			if len(ns) != len(tt.neighbours) {
				t.Errorf("expected %d neighbours, got %d", len(tt.neighbours), len(ns))
			}
			for _, e := range tt.neighbours {
				found := false
				for _, n := range ns {
					if n == e {
						found = true
					}
				}
				if !found {
					t.Errorf("expected neighbour %v not found in %v", e, ns)
				}
			}
		})
	}
}
