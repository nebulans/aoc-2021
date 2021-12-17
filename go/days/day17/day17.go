package day17

import (
	"aoc-2021/framework"
	"aoc-2021/util/math/vector"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Area struct {
	Min vector.Vec2
	Max vector.Vec2
}

func (a Area) Contains(point vector.Vec2) bool {
	return a.Min.X <= point.X && a.Max.X >= point.X && a.Min.Y <= point.Y && a.Max.Y >= point.Y
}

type Shot struct {
	position vector.Vec2
	velocity vector.Vec2
}

func (s *Shot) simulateStep() {
	s.position = s.position.Add(s.velocity)
	if s.velocity.X > 0 {
		s.velocity.X -= 1
	}
	s.velocity.Y -= 1
}

func (s *Shot) maxY() int {
	y := s.position.Y
	for y <= s.position.Y {
		y = s.position.Y
		s.simulateStep()
	}
	return y
}

func (s *Shot) hits(target Area) bool {
	for true {
		s.simulateStep()
		if s.position.X > target.Max.X || s.position.Y < target.Min.Y {
			return false
		}
		if target.Contains(s.position) {
			return true
		}
	}
	return false
}

type Puzzle struct {
	framework.PuzzleBase
	target Area
}

func (p *Puzzle) Init() {
	p.Parts = map[string]func() int{
		"1": p.maxHeight,
		"2": p.distinctSolutions,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		pattern, _ := regexp.Compile("x=(-?[0-9]+)..(-?[0-9]+), y=(-?[0-9]+)..(-?[0-9]+)")
		match := pattern.FindStringSubmatch(line)
		minX, _ := strconv.Atoi(match[1])
		maxX, _ := strconv.Atoi(match[2])
		minY, _ := strconv.Atoi(match[3])
		maxY, _ := strconv.Atoi(match[4])
		p.target = Area{
			Min: vector.Vec2{X: minX, Y: minY},
			Max: vector.Vec2{X: maxX, Y: maxY},
		}
	}
}

func (p *Puzzle) maxHeight() int {
	maxYVelocity := -1 * (p.target.Min.Y + 1)
	shot := Shot{
		position: vector.Vec2{X: 0, Y: 0},
		velocity: vector.Vec2{X: 0, Y: maxYVelocity},
	}
	return shot.maxY()
}

func (p *Puzzle) distinctSolutions() int {
	maxYVelocity := -1 * (p.target.Min.Y + 1)
	minYVelocity := p.target.Min.Y
	maxXVelocity := p.target.Max.X
	minXVelocity := 0
	i := 0
	s := 0
	for i = 1; s <= p.target.Min.X; i++ {
		s += i
	}
	minXVelocity = i - 1
	fmt.Printf("X range: %d, %d\n", minXVelocity, maxXVelocity)
	fmt.Printf("Y range: %d, %d\n", minYVelocity, maxYVelocity)
	hits := 0
	misses := 0
	for x := minXVelocity; x <= maxXVelocity; x++ {
		for y := minYVelocity; y <= maxYVelocity; y++ {
			shot := Shot{
				position: vector.Vec2{X: 0, Y: 0},
				velocity: vector.Vec2{X: x, Y: y},
			}
			if shot.hits(p.target) {
				hits++
			} else {
				misses++
			}
		}
	}
	fmt.Printf("%d hits, %d misses\n", hits, misses)
	return hits
}
