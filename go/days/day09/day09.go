package day09

import (
	"aoc-2021/framework"
	"aoc-2021/util/input"
	"aoc-2021/util/math/grid2d"
	"aoc-2021/util/math/integer"
	"aoc-2021/util/math/vector"
	"bufio"
	"sort"
)

type Puzzle struct {
	framework.PuzzleBase
	heights *grid2d.IntGrid
}

func (p *Puzzle) Init() {
	p.Parts = map[string]func() int{
		"1": p.riskLevelSum,
		"2": p.fromMinimaLargestBasins,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	x := 0
	y := 0
	all := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		vals := input.SplitInts(line, "")
		all = append(all, vals[:]...)
		x = len(vals)
		y++
	}
	p.heights = &grid2d.IntGrid{Backend: grid2d.MakeArrayGrid(vector.Vec2{X: x, Y: y})}
	for i, pos := range p.heights.Positions() {
		p.heights.Set(pos, all[i])
	}
}

func (p *Puzzle) Minima() []vector.Vec2 {
	minima := make([]vector.Vec2, 0)
	for _, position := range p.heights.Positions() {
		height := p.heights.Get(position)
		neighbourHeights := make([]int, 0, 4)
		for _, neighbourPosition := range p.heights.Neighbours(position, false) {
			neighbourHeights = append(neighbourHeights, p.heights.Get(neighbourPosition))
		}
		if height < integer.MinSlice(neighbourHeights) {
			minima = append(minima, position)
		}
	}
	return minima
}

func (p *Puzzle) basinPoints(start vector.Vec2) []vector.Vec2 {
	points := make([]vector.Vec2, 0, p.heights.Length())
	points = append(points, start)
	for i := 0; i < len(points); i++ {
		pos := points[i]
		for _, n := range p.heights.Neighbours(pos, false) {
			if p.heights.Get(n) == 9 {
				continue
			}
			found := false
			for _, p := range points {
				if p == n {
					found = true
				}
			}
			if !found {
				points = append(points, n)
			}
		}
	}
	return points
}

func (p *Puzzle) riskLevelSum() int {
	score := 0
	for _, pos := range p.Minima() {
		score += 1 + p.heights.Get(pos)
	}
	return score
}

func (p *Puzzle) fromMinimaLargestBasins() int {
	minima := p.Minima()
	scores := make([]int, len(minima))
	for i, pos := range minima {
		scores[i] = len(p.basinPoints(pos))
	}
	sort.Ints(scores)
	return scores[len(scores)-1] * scores[len(scores)-2] * scores[len(scores)-3]
}
