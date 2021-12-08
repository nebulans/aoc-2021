package day05

import (
	"aoc-2021/framework"
	"aoc-2021/util/math/vector"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Field struct {
	positions map[vector.Vec2]int
	extents   vector.Vec2
	max       int
}

func NewField() *Field {
	return &Field{
		positions: make(map[vector.Vec2]int),
		extents:   vector.Vec2{X: 0, Y: 0},
	}
}

func (f *Field) AddPoint(position vector.Vec2) int {
	f.positions[position]++
	if position.X >= f.extents.X {
		f.extents.X = position.X + 1
	}
	if position.Y >= f.extents.Y {
		f.extents.Y = position.Y + 1
	}
	newValue := f.positions[position]
	if f.max < newValue {
		f.max = newValue
	}
	return newValue
}

func (f *Field) Values() []int {
	values := make([]int, f.extents.X*f.extents.Y)
	position := 0
	for y := 0; y < f.extents.Y; y++ {
		for x := 0; x < f.extents.X; x++ {
			values[position] = f.positions[vector.Vec2{X: x, Y: y}]
			position++
		}
	}
	return values
}

func (f *Field) FilledValues() []int {
	values := make([]int, len(f.positions))
	i := 0
	for _, v := range f.positions {
		values[i] = v
		i++
	}
	return values
}

func (f *Field) Format() string {
	values := f.Values()
	cells := make([]string, len(values))
	for i, v := range values {
		sep := ""
		if i%f.extents.X == f.extents.X-1 {
			sep = "\n"
		}
		if v == 0 {
			cells[i] = fmt.Sprintf(".%s", sep)
		} else {
			cells[i] = fmt.Sprintf("%X%s", v, sep)
		}
	}
	return strings.Join(cells, "")
}

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
	field := NewField()
	for line := range p.lines {
		if onlyAxisAligned {
			if line.Direction.X != 0 && line.Direction.Y != 0 {
				continue
			}
		}
		for _, point := range line.Points() {
			field.AddPoint(point)
		}
	}
	multiples := 0
	for _, v := range field.FilledValues() {
		if v > 1 {
			multiples++
		}
	}
	//fmt.Println(field.Format())
	fmt.Printf("Higest count: %d\n", field.max)
	fmt.Printf("Field size: %dx%d\n", field.extents.X, field.extents.Y)
	return multiples
}
