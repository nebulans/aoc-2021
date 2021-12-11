package day05

import (
	"aoc-2021/framework"
	"aoc-2021/util/math/grid2d"
	"aoc-2021/util/math/vector"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Puzzle struct {
	framework.PuzzleBase
	lines chan *vector.EndpointLine
}

func (p *Puzzle) Init() {
	p.lines = make(chan *vector.EndpointLine)
	p.Parts = map[string]func() int{
		"1": func() int { return p.countMultiples(true) },
		"2": func() int { return p.countMultiples(false) },
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	go p.asyncParse(scanner)
}

func (p *Puzzle) asyncParse(scanner *bufio.Scanner) {
	pattern, _ := regexp.Compile("([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)")
	for scanner.Scan() {
		line := scanner.Text()
		elements := pattern.FindStringSubmatch(line)
		startX, _ := strconv.Atoi(elements[1])
		startY, _ := strconv.Atoi(elements[2])
		endX, _ := strconv.Atoi(elements[3])
		endY, _ := strconv.Atoi(elements[4])
		p.lines <- vector.NewEndpointLine(vector.Vec2{X: startX, Y: startY}, vector.Vec2{X: endX, Y: endY})
	}
	close(p.lines)
}

func (p *Puzzle) countMultiples(onlyAxisAligned bool) int {
	field := grid2d.MakeIntGrid(grid2d.MakeMapGrid(0))
	for line := range p.lines {
		if onlyAxisAligned {
			if line.Direction.X != 0 && line.Direction.Y != 0 {
				continue
			}
		}
		for _, point := range line.Points() {
			field.Set(point, field.Get(point)+1)
		}
	}
	multiples := 0
	for _, p := range field.FilledPositions() {
		if field.Get(p) > 1 {
			multiples++
		}
	}
	fmt.Printf("Field size: %dx%d\n", field.Extents().X, field.Extents().Y)
	fmt.Println(field.Format(field.SparseFormatter))
	return multiples
}
