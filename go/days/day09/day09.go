package day09

import (
	"aoc-2021/framework"
	"aoc-2021/util/input"
	"aoc-2021/util/math/integer"
	"aoc-2021/util/math/vector"
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type HeightMap struct {
	heights []int
	extents vector.Vec2
}

func (m *HeightMap) Get(pos vector.Vec2) int {
	i := pos.X + pos.Y*m.extents.X
	return m.heights[i]
}

func (m *HeightMap) Contains(pos vector.Vec2) bool {
	if pos.X < 0 || pos.Y < 0 || pos.X >= m.extents.X || pos.Y >= m.extents.Y {
		return false
	}
	return true
}

func (m *HeightMap) Positions() []vector.Vec2 {
	pos := make([]vector.Vec2, m.extents.X*m.extents.Y)
	i := 0
	for y := 0; y < m.extents.Y; y++ {
		for x := 0; x < m.extents.X; x++ {
			pos[i] = vector.Vec2{x, y}
			i++
		}
	}
	return pos
}

func (m *HeightMap) Neighbours(pos vector.Vec2) []vector.Vec2 {
	n := make([]vector.Vec2, 0, 4)
	vecs := [4]vector.Vec2{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	for _, v := range vecs {
		p := pos.Add(v)
		if m.Contains(p) {
			n = append(n, p)
		}
	}
	return n
}

func (m *HeightMap) Minima() []vector.Vec2 {
	minima := make([]vector.Vec2, 0)
	for _, position := range m.Positions() {
		height := m.Get(position)
		neighbourHeights := make([]int, 0, 4)
		for _, neighbourPosition := range m.Neighbours(position) {
			neighbourHeights = append(neighbourHeights, m.Get(neighbourPosition))
		}
		if height < integer.MinSlice(neighbourHeights) {
			minima = append(minima, position)
		}
	}
	return minima
}

func (m *HeightMap) Format() string {
	elems := make([]string, m.extents.X*m.extents.Y)
	for i, pos := range m.Positions() {
		val := m.Get(pos)
		if pos.X == m.extents.X-1 {
			elems[i] = fmt.Sprintf("%d\n", val)
		} else {
			elems[i] = fmt.Sprintf("%d", val)
		}
	}
	return strings.Join(elems, "")
}

func (m *HeightMap) basinPoints(start vector.Vec2) []vector.Vec2 {
	points := make([]vector.Vec2, 0, len(m.heights))
	points = append(points, start)
	for i := 0; i < len(points); i++ {
		pos := points[i]
		for _, n := range m.Neighbours(pos) {
			if m.Get(n) == 9 {
				continue
			}
			found := false
			for _, p := range points {
				if p == n {
					found = true
				}
			}
			if !found {
				points = append(points, n)
			}
		}
	}
	return points
}

type Puzzle struct {
	framework.PuzzleBase
	heights *HeightMap
}

func (p *Puzzle) Init() {
	p.heights = &HeightMap{heights: make([]int, 0), extents: vector.Vec2{}}
	p.Parts = map[string]func() int{
		"1": p.riskLevelSum,
		"2": p.fromMinimaLargestBasins,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	h := 0
	for scanner.Scan() {
		line := scanner.Text()
		vals := input.SplitInts(line, "")
		for _, v := range vals {
			p.heights.heights = append(p.heights.heights, v)
		}
		p.heights.extents.X = len(vals)
		h++
	}
	p.heights.extents.Y = h
}

func (p *Puzzle) riskLevelSum() int {
	score := 0
	for _, pos := range p.heights.Minima() {
		score += 1 + p.heights.Get(pos)
	}
	return score
}

func (p *Puzzle) fromMinimaLargestBasins() int {
	minima := p.heights.Minima()
	scores := make([]int, len(minima))
	for i, pos := range minima {
		scores[i] = len(p.heights.basinPoints(pos))
	}
	sort.Ints(scores)
	return scores[len(scores)-1] * scores[len(scores)-2] * scores[len(scores)-3]
}
