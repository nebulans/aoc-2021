package day22

import (
	"aoc-2021/framework"
	"aoc-2021/util/datastructure/stack"
	"aoc-2021/util/math/integer"
	"aoc-2021/util/math/vector"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Cuboid struct {
	Min vector.Vec3
	Max vector.Vec3
}

func (c Cuboid) Size() int {
	return (1 + c.Max.X - c.Min.X) * (1 + c.Max.Y - c.Min.Y) * (1 + c.Max.Z - c.Min.Z)
}

func (c Cuboid) Points() []vector.Vec3 {
	points := make([]vector.Vec3, 0, c.Size())
	for x := c.Min.X; x <= c.Max.X; x++ {
		for y := c.Min.Y; y <= c.Max.Y; y++ {
			for z := c.Min.Z; z <= c.Max.Z; z++ {
				points = append(points, vector.Vec3{x, y, z})
			}
		}
	}
	return points
}

func (c Cuboid) Intersects(other Cuboid) bool {
	return c.Max.X >= other.Min.X && c.Min.X <= other.Max.X &&
		c.Max.Y >= other.Min.Y && c.Min.Y <= other.Max.Y &&
		c.Max.Z >= other.Min.Z && c.Min.Z <= other.Max.Z
}

func (c Cuboid) Contains(other Cuboid) bool {
	return c.Min.X <= other.Min.X && c.Max.X >= other.Max.X &&
		c.Min.Y <= other.Min.Y && c.Max.Y >= other.Max.Y &&
		c.Min.Z <= other.Min.Z && c.Max.Z >= other.Max.Z
}

func (c Cuboid) Intersection(other Cuboid) (bool, Cuboid) {
	if !c.Intersects(other) {
		return false, Cuboid{}
	}
	return true, Cuboid{
		Min: vector.Vec3{
			X: integer.Max(c.Min.X, other.Min.X),
			Y: integer.Max(c.Min.Y, other.Min.Y),
			Z: integer.Max(c.Min.Z, other.Min.Z),
		},
		Max: vector.Vec3{
			X: integer.Min(c.Max.X, other.Max.X),
			Y: integer.Min(c.Max.Y, other.Max.Y),
			Z: integer.Min(c.Max.Z, other.Max.Z),
		},
	}
}

func makeSlices(aMin, aMax, bMin, bMax int) [][2]int {
	if aMin == bMin && aMax == bMax {
		return [][2]int{
			{aMin, aMax},
		}
	} else if aMin == bMin {
		return [][2]int{
			{aMin, integer.Min(aMax, bMax)},
			{integer.Min(aMax, bMax) + 1, integer.Max(aMax, bMax)},
		}
	} else if aMax == bMax {
		return [][2]int{
			{integer.Min(aMin, bMin), integer.Max(aMin, bMin) - 1},
			{integer.Max(aMin, bMin), aMax},
		}
	} else {
		return [][2]int{
			{integer.Min(aMin, bMin), integer.Max(aMin, bMin) - 1},
			{integer.Max(aMin, bMin), integer.Min(aMax, bMax)},
			{integer.Min(aMax, bMax) + 1, integer.Max(aMax, bMax)},
		}
	}
}

func (c Cuboid) Decompose(other Cuboid) []Cuboid {
	out := make([]Cuboid, 0, 27)
	xSlices := makeSlices(c.Min.X, c.Max.X, other.Min.X, other.Max.X)
	ySlices := makeSlices(c.Min.Y, c.Max.Y, other.Min.Y, other.Max.Y)
	zSlices := makeSlices(c.Min.Z, c.Max.Z, other.Min.Z, other.Max.Z)
	for _, x := range xSlices {
		for _, y := range ySlices {
			for _, z := range zSlices {
				d := Cuboid{
					Min: vector.Vec3{x[0], y[0], z[0]},
					Max: vector.Vec3{x[1], y[1], z[1]},
				}
				if c.Contains(d) || other.Contains(d) {
					out = append(out, d)
				}
			}
		}
	}
	return out
}

type CuboidSpace struct {
	regions map[Cuboid]bool
}

func (s *CuboidSpace) Insert(c Cuboid, state bool) {
	insertStack := stack.MakeStack(100)
	insertStack.Push(Instruction{Space: c, State: state})
	for !insertStack.IsEmpty() {
		intersected := false
		i := insertStack.Pop().(Instruction)
		for otherCuboid, otherState := range s.regions {
			if i.Space.Intersects(otherCuboid) {
				intersected = true
				if i.Space == otherCuboid {
					if i.State {
						s.regions[otherCuboid] = i.State
					} else {
						delete(s.regions, otherCuboid)
					}
				} else {
					delete(s.regions, otherCuboid)
					decomposed := i.Space.Decompose(otherCuboid)
					for _, d := range decomposed {
						if i.Space.Contains(d) {
							insertStack.Push(Instruction{Space: d, State: state})
						} else if otherState {
							s.regions[d] = otherState
						}
					}
					break
				}
			}
		}
		if !intersected && i.State {
			s.regions[i.Space] = i.State
		}
	}
}

type Instruction struct {
	Space Cuboid
	State bool
}

type Puzzle struct {
	framework.PuzzleBase
	instructions chan Instruction
}

func (p *Puzzle) Init() {
	p.instructions = make(chan Instruction)
	p.Parts = map[string]func() int{
		"1":  p.initRegionOn,
		"1d": p.initRegionOnDecomposition,
		"2":  p.allOnDecomposition,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	go p.asyncParse(scanner)
}

func (p *Puzzle) asyncParse(scanner *bufio.Scanner) {
	pattern, _ := regexp.Compile("(on|off) x=(-?[0-9]+)..(-?[0-9]+),y=(-?[0-9]+)..(-?[0-9]+),z=(-?[0-9]+)..(-?[0-9]+)")
	parseInt := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}
	for scanner.Scan() {
		line := scanner.Text()
		elements := pattern.FindStringSubmatch(line)
		cuboid := Cuboid{
			Min: vector.Vec3{parseInt(elements[2]), parseInt(elements[4]), parseInt(elements[6])},
			Max: vector.Vec3{parseInt(elements[3]), parseInt(elements[5]), parseInt(elements[7])},
		}
		p.instructions <- Instruction{Space: cuboid, State: elements[1] == "on"}
	}
	close(p.instructions)
}

func (p *Puzzle) initRegionOn() int {
	voxels := map[vector.Vec3]bool{}
	region := Cuboid{
		Min: vector.Vec3{-50, -50, -50},
		Max: vector.Vec3{50, 50, 50},
	}
	for i := range p.instructions {
		intersects, intersection := i.Space.Intersection(region)
		if intersects {
			for _, point := range intersection.Points() {
				voxels[point] = i.State
			}
		}
	}
	onVoxels := 0
	for _, v := range voxels {
		if v {
			onVoxels++
		}
	}
	return onVoxels
}

func (p *Puzzle) initRegionOnDecomposition() int {
	space := &CuboidSpace{regions: map[Cuboid]bool{}}
	region := Cuboid{
		Min: vector.Vec3{-50, -50, -50},
		Max: vector.Vec3{50, 50, 50},
	}
	for instruction := range p.instructions {
		intersects, intersection := instruction.Space.Intersection(region)
		if intersects {
			space.Insert(intersection, instruction.State)
		}
	}
	fmt.Printf("%d regions generated\n", len(space.regions))
	onVoxels := 0
	for r, state := range space.regions {
		if state {
			onVoxels += r.Size()
		}
	}
	return onVoxels
}

func (p *Puzzle) allOnDecomposition() int {
	space := &CuboidSpace{regions: map[Cuboid]bool{}}
	for instruction := range p.instructions {
		space.Insert(instruction.Space, instruction.State)
	}
	fmt.Printf("%d regions generated\n", len(space.regions))
	onVoxels := 0
	for r, state := range space.regions {
		if state {
			onVoxels += r.Size()
		}
	}
	return onVoxels
}
