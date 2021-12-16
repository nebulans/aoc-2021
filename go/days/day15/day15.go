package day15

import (
	"aoc-2021/framework"
	"aoc-2021/util/input"
	"aoc-2021/util/math/grid2d"
	"aoc-2021/util/math/integer"
	"aoc-2021/util/math/vector"
	"bufio"
	"fmt"
)

type BestCostGrid struct {
	cellCosts        *grid2d.IntGrid
	cumulativeCosts  *grid2d.IntGrid
	totalComparisons int
}

func (g *BestCostGrid) calculatePosition(point vector.Vec2) bool {
	neighbours := g.cellCosts.Neighbours(point, false)
	costs := make([]int, len(neighbours))
	for i, n := range neighbours {
		g.totalComparisons++
		costs[i] = g.cumulativeCosts.Get(n) + g.cellCosts.Get(point)
	}
	best := integer.MinSlice(costs)
	if best < g.cumulativeCosts.Get(point) && best < g.cornerValue() {
		g.cumulativeCosts.Set(point, best)
		return true
	}
	return false
}

func (g *BestCostGrid) recalculate() int {
	changes := 0
	for _, point := range g.cellCosts.Positions() {
		changed := g.calculatePosition(point)
		if changed {
			changes++
		}
	}
	return changes
}

func (g *BestCostGrid) cornerValue() int {
	corner := vector.Vec2{X: g.cumulativeCosts.Extents().X - 1, Y: g.cumulativeCosts.Extents().Y - 1}
	return g.cumulativeCosts.Get(corner)
}

func (g *BestCostGrid) solve() int {
	iterations := 0
	changes := 1
	for changes > 0 {
		iterations++
		changes = g.recalculate()
	}
	fmt.Printf("%d recalculation iterations\n", iterations)
	fmt.Printf("%d comparisons\n", g.totalComparisons)
	return g.cornerValue()
}

func MakeBestCostGrid(costs *grid2d.IntGrid) *BestCostGrid {
	cumulative := grid2d.MakeIntGrid(grid2d.MakeArrayGrid(costs.Extents()))
	for _, pos := range cumulative.Positions() {
		if pos == (vector.Vec2{X: 0, Y: 0}) {
			cumulative.Set(pos, 0)
		} else {
			cumulative.Set(pos, 10000)
		}
	}
	return &BestCostGrid{cellCosts: costs, cumulativeCosts: cumulative}
}

type Puzzle struct {
	framework.PuzzleBase
	grid *grid2d.IntGrid
}

func (p *Puzzle) Init() {
	p.Parts = map[string]func() int{
		"1": p.lowestScore,
		"2": p.lowestScoreTiled,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	x := 0
	y := 0
	allVals := make([]int, 0)
	for scanner.Scan() {
		y++
		line := scanner.Text()
		parts := input.SplitInts(line, "")
		x = len(parts)
		allVals = append(allVals, parts[:]...)
	}
	p.grid = grid2d.MakeIntGrid(grid2d.MakeArrayGrid(vector.Vec2{X: x, Y: y}))
	for i, pos := range p.grid.Positions() {
		p.grid.Set(pos, allVals[i])
	}
}

func (p *Puzzle) lowestScore() int {
	costs := MakeBestCostGrid(p.grid)
	return costs.solve()
}

func (p *Puzzle) lowestScoreTiled() int {
	tiledCosts := grid2d.MakeIntGrid(grid2d.MakeArrayGrid(p.grid.Extents().Mul(5)))
	for _, pos := range tiledCosts.Positions() {
		source := vector.Vec2{
			X: pos.X % p.grid.Extents().X,
			Y: pos.Y % p.grid.Extents().X,
		}
		manhattan := (pos.X / p.grid.Extents().X) + (pos.Y / p.grid.Extents().Y)
		value := p.grid.Get(source) + manhattan
		if value > 9 {
			value -= 9
		}
		tiledCosts.Set(pos, value)
	}
	costs := MakeBestCostGrid(tiledCosts)
	return costs.solve()
}
