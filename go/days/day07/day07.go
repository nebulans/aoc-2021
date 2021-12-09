package day07

import (
	"aoc-2021/framework"
	"aoc-2021/util/input"
	"aoc-2021/util/math/integer"
	"bufio"
	"math"
)

func constantCost(position int, target int) int {
	return integer.Abs(position - target)
}

func sumBelowCost(position int, target int) int {
	steps := constantCost(position, target)
	return ((2 * steps) * (steps + 1)) / 4
}

type Puzzle struct {
	framework.PuzzleBase
	positions []int
}

func (p *Puzzle) Init() {
	p.Parts = map[string]func() int{
		"1": func() int { return p.CostScan(constantCost) },
		"2": func() int { return p.CostScan(sumBelowCost) },
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	scanner.Scan()
	p.positions = input.SplitInts(scanner.Text(), ",")
}

func (p *Puzzle) CostScan(costFunc func(int, int) int) int {
	min := integer.MinSlice(p.positions)
	max := integer.MaxSlice(p.positions)
	best := math.MaxInt
	for t := min; t <= max; t++ {
		cost := 0
		for _, v := range p.positions {
			cost += costFunc(v, t)
		}
		if cost < best {
			best = cost
		}
	}
	return best
}
