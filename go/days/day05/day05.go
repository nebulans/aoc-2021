package day05

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Vec2 struct {
	x int
	y int
}

func (vec *Vec2) add(other Vec2) Vec2 {
	return Vec2{vec.x + other.x, vec.y + other.y}
}

type VentLine struct {
	start     Vec2
	end       Vec2
	direction Vec2
}

func step(diff int) int {
	if diff > 0 {
		return 1
	}
	if diff < 0 {
		return -1
	}
	return 0
}

func NewVentLine(start Vec2, end Vec2) *VentLine {
	line := VentLine{start: start, end: end}
	line.direction = Vec2{
		x: step(end.x - start.x),
		y: step(end.y - start.y),
	}
	return &line
}

func (v *VentLine) format() string {
	return fmt.Sprintf("%d,%d -> %d,%d", v.start.x, v.start.y, v.end.x, v.end.y)
}

func (v *VentLine) isAxisAligned() bool {
	return v.direction.x == 0 || v.direction.y == 0
}

func (v *VentLine) points() []Vec2 {
	var points []Vec2
	pos := v.start
	points = append(points, pos)
	for pos != v.end {
		pos = pos.add(v.direction)
		points = append(points, pos)
	}
	return points
}

type Field struct {
	positions map[Vec2]int
	extents   Vec2
	max       int
}

func NewField() *Field {
	return &Field{
		positions: make(map[Vec2]int),
		extents:   Vec2{0, 0},
	}
}

func (f *Field) AddPoint(position Vec2) int {
	f.positions[position]++
	if position.x >= f.extents.x {
		f.extents.x = position.x + 1
	}
	if position.y >= f.extents.y {
		f.extents.y = position.y + 1
	}
	newValue := f.positions[position]
	if f.max < newValue {
		f.max = newValue
	}
	return newValue
}

func (f *Field) Values() []int {
	values := make([]int, f.extents.x*f.extents.y)
	position := 0
	for y := 0; y < f.extents.y; y++ {
		for x := 0; x < f.extents.x; x++ {
			values[position] = f.positions[Vec2{x, y}]
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
		if i%f.extents.x == f.extents.x-1 {
			sep = "\n"
		}
		if v == 0 {
			cells[i] = fmt.Sprintf(".%s", sep)
		} else {
			cells[i] = fmt.Sprintf("%x%s", v, sep)
		}
	}
	return strings.Join(cells, "")
}

func parseInput(scanner *bufio.Scanner, lines chan<- *VentLine) {
	pattern, _ := regexp.Compile("([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)")
	for scanner.Scan() {
		line := scanner.Text()
		elements := pattern.FindStringSubmatch(line)
		startX, _ := strconv.Atoi(elements[1])
		startY, _ := strconv.Atoi(elements[2])
		endX, _ := strconv.Atoi(elements[3])
		endY, _ := strconv.Atoi(elements[4])
		lines <- NewVentLine(Vec2{startX, startY}, Vec2{endX, endY})
	}
	close(lines)
}

func countMultiples(lines <-chan *VentLine, onlyAxisAligned bool) int {
	field := NewField()
	for line := range lines {
		if onlyAxisAligned && !line.isAxisAligned() {
			continue
		}
		for _, point := range line.points() {
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
	fmt.Printf("Field size: %dx%d\n", field.extents.x, field.extents.y)
	return multiples
}

func countAxisAlignedMultiples(lines <-chan *VentLine) int {
	return countMultiples(lines, true)
}

func countAllMultiples(lines <-chan *VentLine) int {
	return countMultiples(lines, false)
}

var partMap = map[string]func(<-chan *VentLine) int{
	"1": countAxisAlignedMultiples,
	"2": countAllMultiples,
}

func Day05(part string, input *bufio.Scanner) (string, error) {
	lines := make(chan *VentLine)
	go parseInput(input, lines)
	result := partMap[part](lines)
	return fmt.Sprintf("%d", result), nil
}
