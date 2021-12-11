package day11

import (
	"aoc-2021/framework"
	"aoc-2021/util/input"
	"aoc-2021/util/math/grid2d"
	"aoc-2021/util/math/vector"
	"bufio"
	"fmt"
	"strings"
)

type Octopus struct {
	energy     int
	flashCount int
}

type Puzzle struct {
	framework.PuzzleBase
	grid *grid2d.ArrayGrid
}

func (p *Puzzle) FormatGrid() string {
	elems := make([]string, p.grid.Length())
	for i, pos := range p.grid.Positions() {
		val := p.grid.Get(pos).(*Octopus).energy
		if pos.X == p.grid.Extents.X-1 {
			elems[i] = fmt.Sprintf("%d\n", val)
		} else {
			elems[i] = fmt.Sprintf("%d", val)
		}
	}
	return strings.Join(elems, "")
}

func (p *Puzzle) Init() {
	p.grid = grid2d.MakeArrayGrid(vector.Vec2{X: 10, Y: 10})
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
		p.grid.Set(pos, &Octopus{energy: values[i], flashCount: 0})
	}
}

func (p *Puzzle) incrementPoint(pos vector.Vec2) {
	o := p.grid.Get(pos).(*Octopus)
	o.energy++
	if o.energy == 10 {
		for _, n := range p.grid.Neighbours(pos, true) {
			p.incrementPoint(n)
		}
	}
}

func (p *Puzzle) simulateStep() int {
	// Recursively apply increments
	for _, pos := range p.grid.Positions() {
		p.incrementPoint(pos)
	}
	// Reset all flashing
	flashes := 0
	for _, pos := range p.grid.Positions() {
		o := p.grid.Get(pos).(*Octopus)
		if o.energy > 9 {
			o.flashCount++
			o.energy = 0
			flashes++
		}
	}
	return flashes
}

func (p *Puzzle) countFlashes() int {
	flashes := 0
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
