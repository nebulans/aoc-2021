package day13

import (
	"aoc-2021/framework"
	"aoc-2021/util/input"
	"aoc-2021/util/math/grid2d"
	"aoc-2021/util/math/vector"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Direction int8

const (
	X Direction = iota
	Y
)

var directionMap = map[string]Direction{
	"x": X,
	"y": Y,
}

type FoldInstruction struct {
	Direction Direction
	Value     int
}

type FoldingPaper struct {
	grid2d.GridBackend
}

func (p *FoldingPaper) Get(position vector.Vec2) bool {
	return p.GridBackend.Get(position).(bool)
}

func (p *FoldingPaper) Set(position vector.Vec2, value bool) {
	p.GridBackend.Set(position, value)
}

func (p *FoldingPaper) Fold(fold FoldInstruction) {
	// Transform positions
	for _, position := range p.FilledPositions() {
		if fold.Direction == X {
			if position.X > fold.Value {
				n := vector.Vec2{X: p.Extents().X - position.X - 1, Y: position.Y}
				p.Set(n, true)
			}
		} else {
			if position.Y > fold.Value {
				n := vector.Vec2{X: position.X, Y: p.Extents().Y - position.Y - 1}
				p.Set(n, true)
			}
		}
	}
	// Reduce extents of backing map
	if fold.Direction == X {
		p.SetExtents(vector.Vec2{X: fold.Value, Y: p.Extents().Y})
	} else {
		p.SetExtents(vector.Vec2{X: p.Extents().X, Y: fold.Value})
	}
}

func (p *FoldingPaper) FilledPositions() []vector.Vec2 {
	all := p.GridBackend.FilledPositions()
	out := make([]vector.Vec2, 0, len(all))
	for _, i := range all {
		if p.ContainsPoint(i) {
			out = append(out, i)
		}
	}
	return out
}

func (p *FoldingPaper) Format() string {
	elems := make([]string, p.Length())
	for i, pos := range p.Positions() {
		val := "."
		if p.Get(pos) {
			val = "#"
		}
		if pos.X == p.GridBackend.Extents().X-1 {
			elems[i] = fmt.Sprintf("%s\n", val)
		} else {
			elems[i] = val
		}
	}
	return strings.Join(elems, "")
}

type Puzzle struct {
	framework.PuzzleBase
	points chan vector.Vec2
	folds  chan FoldInstruction
}

func (p *Puzzle) Init() {
	p.points = make(chan vector.Vec2)
	p.folds = make(chan FoldInstruction)
	p.Parts = map[string]func() int{
		"1": func() int { return p.makeFolds(true) },
		"2": func() int { return p.makeFolds(false) },
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	go p.asyncParse(scanner)
}

func (p *Puzzle) asyncParse(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := input.SplitInts(line, ",")
		p.points <- vector.Vec2{X: parts[0], Y: parts[1]}
	}
	close(p.points)
	pattern, _ := regexp.Compile("fold along ([xy])=([0-9]+)")
	for scanner.Scan() {
		line := scanner.Text()
		elements := pattern.FindStringSubmatch(line)
		val, _ := strconv.Atoi(elements[2])
		p.folds <- FoldInstruction{Direction: directionMap[elements[1]], Value: val}
	}
	close(p.folds)
}

func (p *Puzzle) makeFolds(firstOnly bool) int {
	paper := &FoldingPaper{grid2d.MakeMapGrid(false)}
	for point := range p.points {
		paper.Set(point, true)
	}
	for fold := range p.folds {
		paper.Fold(fold)
		if firstOnly {
			break
		}
	}
	fmt.Println(paper.Format())
	return len(paper.FilledPositions())
}
