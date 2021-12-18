package day18

import (
	"aoc-2021/framework"
	"bufio"
	"strconv"
)

type SnailfishNode struct {
	Left  *SnailfishNode
	Right *SnailfishNode
	Value int
	Depth int
}

func (n *SnailfishNode) IncrementDepth() {
	n.Depth++
	if !n.IsValue() {
		n.Left.IncrementDepth()
		n.Right.IncrementDepth()
	}
}

func (n *SnailfishNode) Add(other *SnailfishNode) *SnailfishNode {
	result := &SnailfishNode{
		Left:  n,
		Right: other,
		Depth: 0,
	}
	n.IncrementDepth()
	other.IncrementDepth()
	result.ResolveRules()
	return result
}

func (n *SnailfishNode) ResolveRules() {
	changed := true
	for changed {
		changed, _, _ = n.Explode()
		if changed {
			continue
		}
		changed = n.Split()
	}
}

func (n *SnailfishNode) Explode() (bool, int, int) {
	if n.IsValue() {
		return false, 0, 0
	}
	if n.Depth == 4 {
		n.Value = 0
		l, r := n.Left.Value, n.Right.Value
		n.Left, n.Right = nil, nil
		return true, l, r
	} else {
		e, l, r := n.Left.Explode()
		if e {
			n.Right.AddLeft(r)
			return true, l, 0
		}
		e, l, r = n.Right.Explode()
		if e {
			n.Left.AddRight(l)
			return true, 0, r
		}
	}
	return false, 0, 0
}

func (n *SnailfishNode) AddLeft(val int) {
	if n.IsValue() {
		n.Value += val
	} else {
		n.Left.AddLeft(val)
	}
}

func (n *SnailfishNode) AddRight(val int) {
	if n.IsValue() {
		n.Value += val
	} else {
		n.Right.AddRight(val)
	}
}

func (n *SnailfishNode) Split() bool {
	if n.IsValue() {
		if n.Value > 9 {
			m := n.Value % 2
			n.Left = &SnailfishNode{Depth: n.Depth + 1, Value: n.Value / 2}
			n.Right = &SnailfishNode{Depth: n.Depth + 1, Value: n.Value/2 + m}
			return true
		}
	} else {
		return n.Left.Split() || n.Right.Split()
	}
	return false
}

func (n *SnailfishNode) IsValue() bool {
	return n.Left == nil || n.Right == nil
}

func (n *SnailfishNode) Magnitude() int {
	if n.IsValue() {
		return n.Value
	}
	return 3*n.Left.Magnitude() + 2*n.Right.Magnitude()
}

func (n *SnailfishNode) CopyTree() *SnailfishNode {
	if n.IsValue() {
		return &SnailfishNode{Depth: n.Depth, Value: n.Value}
	} else {
		return &SnailfishNode{
			Left:  n.Left.CopyTree(),
			Right: n.Right.CopyTree(),
			Depth: n.Depth,
		}
	}
}

func ParseSnailfishNumber(parent *SnailfishNode, line string, depth int) *SnailfishNode {
	n := &SnailfishNode{Depth: depth}
	splitPoint := 0
	nestingCount := 0
	for i := 1; i < len(line); i++ {
		switch line[i] {
		case '[':
			nestingCount++
		case ']':
			nestingCount--
		case ',':
			if nestingCount == 0 {
				splitPoint = i
				break
			}
		}
	}

	getNode := func(part string) *SnailfishNode {
		if part[0] == '[' {
			return ParseSnailfishNumber(n, part, depth+1)
		} else {
			value, _ := strconv.Atoi(part)
			return &SnailfishNode{
				Value: value,
				Depth: depth + 1,
			}
		}
	}

	n.Left = getNode(line[1:splitPoint])
	n.Right = getNode(line[splitPoint+1 : len(line)-1])
	return n
}

type Puzzle struct {
	framework.PuzzleBase
	numbers chan *SnailfishNode
}

func (p *Puzzle) Init() {
	p.numbers = make(chan *SnailfishNode)
	p.Parts = map[string]func() int{
		"1": p.magnitude,
		"2": p.largestPairMagnitude,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	go p.asyncParse(scanner)
}

func (p *Puzzle) asyncParse(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		p.numbers <- ParseSnailfishNumber(nil, line, 0)
	}
	close(p.numbers)
}

func (p *Puzzle) magnitude() int {
	sum := <-p.numbers
	for n := range p.numbers {
		sum = sum.Add(n)
	}
	return sum.Magnitude()
}

func (p *Puzzle) largestPairMagnitude() int {
	numbers := make([]*SnailfishNode, 0)
	for n := range p.numbers {
		numbers = append(numbers, n)
	}
	best := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			m := numbers[i].CopyTree().Add(numbers[j].CopyTree()).Magnitude()
			if m > best {
				best = m
			}
		}
	}
	return best
}
