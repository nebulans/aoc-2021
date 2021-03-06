package day11

import (
	"aoc-2021/framework"
	"aoc-2021/util/input"
	"aoc-2021/util/math/grid2d"
	"aoc-2021/util/math/vector"
	"bufio"
	"fmt"
)

type Puzzle struct {
	framework.PuzzleBase
	grid     *grid2d.IntGrid
	GridImpl string
}

func (p *Puzzle) Init() {
	switch p.GridImpl {
	case "array":
		p.grid = grid2d.MakeIntGrid(grid2d.MakeArrayGrid(vector.Vec2{X: 10, Y: 10}))
	case "map":
		p.grid = grid2d.MakeIntGrid(grid2d.MakeMapGridExtents(0, vector.Vec2{X: 10, Y: 10}))
	}
	p.Parts = map[string]func() int{
		"1": p.countFlashes,
		"2": p.synchronisedFlash,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	values := make([]int, 0, p.grid.Length())
	for scanner.Scan() {
		line := scanner.Text()
		ints := input.SplitInts(line, "")
		values = append(values, ints[:]...)
	}
	for i, pos := range p.grid.Positions() {
		p.grid.Set(pos, values[i])
	}
}

func (p *Puzzle) incrementPoint(pos vector.Vec2) {
	o := p.grid.Get(pos)
	o++
	p.grid.Set(pos, o)
	if o == 10 {
		for _, n := range p.grid.Neighbours(pos, true) {
			p.incrementPoint(n)
		}
	}
}

func flashFormatter(val int) string {
	if val < 10 {
		return fmt.Sprintf("%d", val)
	}
	return "@"
}

func (p *Puzzle) simulateStep() int {
	// Recursively apply increments
	for _, pos := range p.grid.Positions() {
		p.incrementPoint(pos)
	}
	// Display state
	//fmt.Println(p.grid.Format(flashFormatter))
	// Reset all flashing
	flashes := 0
	for _, pos := range p.grid.Positions() {
		o := p.grid.Get(pos)
		if o > 9 {
			p.grid.Set(pos, 0)
			flashes++
		}
	}
	return flashes
}

func (p *Puzzle) countFlashes() int {
	flashes := 0
	//fmt.Println(p.grid.Format(p.grid.DefaultFormatter))
	for i := 0; i < 100; i++ {
		flashes += p.simulateStep()
	}
	return flashes
}

func (p *Puzzle) synchronisedFlash() int {
	flashes := 0
	var i int
	for i = 0; i < 10000; i++ {
		flashes = p.simulateStep()
		if flashes == 100 {
			break
		}
	}
	return i + 1
}
