package day19

import (
	"aoc-2021/framework"
	"aoc-2021/util/input"
	"aoc-2021/util/math/integer"
	"aoc-2021/util/math/vector"
	"bufio"
	"fmt"
	"strings"
)

type PointCloud struct {
	Points map[vector.Vec3]bool
	Centre vector.Vec3
}

func (c *PointCloud) Add(v vector.Vec3) {
	c.Points[v] = true
}

func (c *PointCloud) Copy() *PointCloud {
	n := MakePointCloud()
	n.Centre = c.Centre
	for k, v := range c.Points {
		n.Points[k] = v
	}
	return n
}

func (c *PointCloud) Rotate(t vector.Transform3) *PointCloud {
	n := MakePointCloud()
	n.Centre = t.Apply(c.Centre)
	for k, v := range c.Points {
		n.Points[t.Apply(k)] = v
	}
	return n
}

func (c *PointCloud) Offset(vec vector.Vec3) *PointCloud {
	n := MakePointCloud()
	n.Centre = c.Centre.Add(vec)
	for k, v := range c.Points {
		n.Points[k.Add(vec)] = v
	}
	return n
}

func (c *PointCloud) CoincidentCount(other *PointCloud) int {
	n := 0
	for k, _ := range c.Points {
		for kk, _ := range other.Points {
			if k == kk {
				n++
			}
		}
	}
	return n
}

func MakePointCloud() *PointCloud {
	return &PointCloud{Points: map[vector.Vec3]bool{}}
}

type Puzzle struct {
	framework.PuzzleBase
	scanners []*PointCloud
}

func (p *Puzzle) Init() {
	p.scanners = make([]*PointCloud, 0)
	p.Parts = map[string]func() int{
		"1": p.distinctPointsAnchor,
		"2": p.maxManhattan,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	scanner.Split(input.BlankLineSplitFunc)
	for scanner.Scan() {
		lines := strings.Split(strings.TrimSpace(scanner.Text()), "\n")
		cloud := MakePointCloud()
		// Discard first line as header
		for _, l := range lines[1:] {
			cloud.Add(input.ParseVec3(l))
		}
		p.scanners = append(p.scanners, cloud)
	}
}

type AlignResult struct {
	Offset      vector.Vec3
	Rotation    vector.Transform3
	Cloud       *PointCloud
	Overlapping int
}

func Align(anchor *PointCloud, other *PointCloud) AlignResult {
	for _, r := range vector.ProperRotations {
		rotated := other.Rotate(r)
		offsetCounts := map[vector.Vec3]int{}
		for i, _ := range anchor.Points {
			for j, _ := range rotated.Points {
				offset := i.Sub(j)
				offsetCounts[offset]++
				if offsetCounts[offset] >= 12 {
					return AlignResult{
						Offset:      offset,
						Rotation:    r,
						Cloud:       rotated.Offset(offset),
						Overlapping: offsetCounts[offset],
					}
				}
			}
		}
	}
	return AlignResult{}
}

func (p *Puzzle) AlignAll() {
	aligned := make([]*PointCloud, 1, len(p.scanners))
	aligned[0] = p.scanners[0].Copy()
	unaligned := make([]*PointCloud, len(p.scanners)-1)
	copy(unaligned, p.scanners[1:])

	for len(unaligned) > 0 {
		fmt.Printf("%d unaligned\n", len(unaligned))
		newUnaligned := make([]*PointCloud, 0, len(unaligned)-1)
		for _, cloud := range unaligned {
			merged := false
			for _, anchor := range aligned {
				ar := Align(anchor, cloud)
				if ar.Overlapping >= 12 {
					aligned = append(aligned, ar.Cloud)
					merged = true
				}
			}
			if !merged {
				newUnaligned = append(newUnaligned, cloud)
			}
		}
		unaligned = newUnaligned
	}
	p.scanners = aligned
}

func (p *Puzzle) distinctPointsAnchor() int {
	p.AlignAll()
	combined := MakePointCloud()
	for _, c := range p.scanners {
		for k, _ := range c.Points {
			combined.Add(k)
		}
	}
	return len(combined.Points)
}

func (p *Puzzle) maxManhattan() int {
	p.AlignAll()
	m := 0
	for i, I := range p.scanners {
		for j, J := range p.scanners {
			if i >= j {
				continue
			}
			offset := I.Centre.Sub(J.Centre)
			manhattan := integer.Abs(offset.X) + integer.Abs(offset.Y) + integer.Abs(offset.Z)
			if manhattan > m {
				m = manhattan
			}
		}
	}

	return m
}
