package day05

import (
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
		extents:   vector.Vec2{0, 0},
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
			values[position] = f.positions[vector.Vec2{x, y}]
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

func parseInput(scanner *bufio.Scanner, lines chan<- *vector.EndpointLine) {
	pattern, _ := regexp.Compile("([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)")
	for scanner.Scan() {
		line := scanner.Text()
		elements := pattern.FindStringSubmatch(line)
		startX, _ := strconv.Atoi(elements[1])
		startY, _ := strconv.Atoi(elements[2])
		endX, _ := strconv.Atoi(elements[3])
		endY, _ := strconv.Atoi(elements[4])
		lines <- vector.NewEndpointLine(vector.Vec2{startX, startY}, vector.Vec2{endX, endY})
	}
	close(lines)
}

func countMultiples(lines <-chan *vector.EndpointLine, onlyAxisAligned bool) int {
	field := NewField()
	for line := range lines {
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

func countAxisAlignedMultiples(lines <-chan *vector.EndpointLine) int {
	return countMultiples(lines, true)
}

func countAllMultiples(lines <-chan *vector.EndpointLine) int {
	return countMultiples(lines, false)
}

var partMap = map[string]func(<-chan *vector.EndpointLine) int{
	"1": countAxisAlignedMultiples,
	"2": countAllMultiples,
}

func Day05(part string, input *bufio.Scanner) (string, error) {
	lines := make(chan *vector.EndpointLine)
	go parseInput(input, lines)
	result := partMap[part](lines)
	return fmt.Sprintf("%d", result), nil
}
