package day05

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type VentLine struct {
	start [2]int
	end   [2]int
}

func (v *VentLine) format() string {
	return fmt.Sprintf("%d,%d -> %d,%d", v.start[0], v.start[1], v.end[0], v.end[1])
}

func (v *VentLine) isAxisAligned() bool {
	if v.start[0] == v.end[0] {
		return true
	}
	if v.start[1] == v.end[1] {
		return true
	}
	return false
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

func (v *VentLine) points() [][2]int {
	var points [][2]int
	pos := [2]int{v.start[0], v.start[1]}
	points = append(points, pos)
	for pos != v.end {
		pos[0] += step(v.end[0] - pos[0])
		pos[1] += step(v.end[1] - pos[1])
		points = append(points, pos)
	}
	return points
}

type Field struct {
	positions map[[2]int]int
	extents   [2]int
	max       int
}

func NewField() *Field {
	return &Field{
		positions: make(map[[2]int]int),
		extents:   [2]int{0, 0},
	}
}

func (f *Field) AddPoint(position [2]int) int {
	f.positions[position]++
	if position[0] >= f.extents[0] {
		f.extents[0] = position[0] + 1
	}
	if position[1] >= f.extents[1] {
		f.extents[1] = position[1] + 1
	}
	newValue := f.positions[position]
	if f.max < newValue {
		f.max = newValue
	}
	return newValue
}

func (f *Field) Values() []int {
	values := make([]int, f.extents[0]*f.extents[1])
	position := 0
	for y := 0; y < f.extents[1]; y++ {
		for x := 0; x < f.extents[0]; x++ {
			values[position] = f.positions[[2]int{x, y}]
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
		if i%f.extents[0] == f.extents[0]-1 {
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
		lines <- &VentLine{
			start: [2]int{startX, startY},
			end:   [2]int{endX, endY},
		}
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
	fmt.Printf("Higest count: %d\n", field.max)
	fmt.Printf("Field size: %dx%d\n", field.extents[0], field.extents[1])
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
