package framework

import (
	"bufio"
	"errors"
	"fmt"
)

type Puzzle interface {
	Init()
	Parse(*bufio.Scanner)
	GetPartFunction(string) (func() int, error)
}

type PuzzleBase struct {
	Parts map[string]func() int
}

func (p *PuzzleBase) GetPartFunction(part string) (func() int, error) {
	partFn, found := p.Parts[part]
	if !found {
		return nil, errors.New(fmt.Sprintf("Part '%s' not found", part))
	}
	return partFn, nil
}

func RunPuzzle(p Puzzle, part string, input *bufio.Scanner) (string, error) {
	p.Init()
	p.Parse(input)
	partFn, err := p.GetPartFunction(part)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", partFn()), nil
}
