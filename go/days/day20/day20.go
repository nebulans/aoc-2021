package day20

import (
	"aoc-2021/framework"
	"aoc-2021/util/math/grid2d"
	"aoc-2021/util/math/vector"
	"bufio"
	"fmt"
	"math/bits"
)

type InfiniteImage struct {
	grid    *grid2d.IntGrid
	padding int
}

func (i *InfiniteImage) pixelsOn() int {
	s := 0
	for _, p := range i.grid.Positions() {
		s += i.grid.Get(p)
	}
	return s
}

func (i *InfiniteImage) pixelValue(position vector.Vec2) int {
	val := uint16(0)
	for _, o := range []vector.Vec2{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{0, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	} {
		val = bits.RotateLeft16(val, 1)
		v := 0
		p := position.Add(o)
		if !i.grid.ContainsPoint(p) {
			v = i.padding
		} else {
			v = i.grid.Get(p)
		}
		val += uint16(v)
	}
	return int(val)
}

func (i *InfiniteImage) enhance(rule map[int]int) *InfiniteImage {
	newGrid := grid2d.MakeIntGrid(grid2d.MakeArrayGrid(i.grid.Extents()))
	for _, p := range i.grid.Positions() {
		newGrid.Set(p, rule[i.pixelValue(p)])
	}
	newPadding := rule[0]
	if i.padding == 1 {
		newPadding = rule[511]
	}
	return MakeInfiniteImage(newGrid, newPadding)
}

func (i *InfiniteImage) Format() string {
	return i.grid.Format(i.grid.BoolFormatter)
}

func MakeInfiniteImage(grid *grid2d.IntGrid, padding int) *InfiniteImage {
	newGrid := grid2d.MakeIntGrid(grid2d.MakeArrayGrid(grid.Extents().Add(vector.Vec2{2, 2})))
	for _, p := range newGrid.Positions() {
		newGrid.Set(p, padding)
	}
	for _, p := range grid.Positions() {
		newGrid.Set(p.Add(vector.Vec2{1, 1}), grid.Get(p))
	}
	return &InfiniteImage{grid: newGrid, padding: padding}
}

type Puzzle struct {
	framework.PuzzleBase
	image *InfiniteImage
	rule  map[int]int
}

func (p *Puzzle) Init() {
	p.rule = map[int]int{}
	p.Parts = map[string]func() int{
		"1": func() int { return p.onAfter(2) },
		"2": func() int { return p.onAfter(50) },
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	scanner.Scan()
	line := scanner.Text()
	for i, c := range line {
		if c == '#' {
			p.rule[i] = 1
		} else {
			p.rule[i] = 0
		}
	}
	scanner.Scan()
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	grid := grid2d.MakeIntGrid(grid2d.MakeArrayGrid(vector.Vec2{len(lines[0]), len(lines)}))
	for y, l := range lines {
		for x, char := range l {
			if char == '#' {
				grid.Set(vector.Vec2{x, y}, 1)
			}
		}
	}
	p.image = MakeInfiniteImage(grid, 0)
}

func (p *Puzzle) onAfter(iterations int) int {
	fmt.Printf("%s\n\n", p.image.Format())
	for i := 0; i < iterations; i++ {
		p.image = p.image.enhance(p.rule)
	}
	fmt.Printf("%s\n\n", p.image.Format())
	return p.image.pixelsOn()
}
