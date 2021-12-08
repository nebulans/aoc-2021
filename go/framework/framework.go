package framework

import (
	"bufio"
	"fmt"
)

type Puzzle interface {
	Init()
	Parse(*bufio.Scanner)
	GetPartFunction(string) func() int
}

type PuzzleBase struct {
	Parts map[string]func() int
}

func (p *PuzzleBase) GetPartFunction(part string) func() int {
	return p.Parts[part]
}

func RunPuzzle(p Puzzle, part string, input *bufio.Scanner) (string, error) {
	p.Init()
	p.Parse(input)
	partFn := p.GetPartFunction(part)
	return fmt.Sprintf("%d", partFn()), nil
}
