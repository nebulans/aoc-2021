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
	Run(string) (string, error)
}

type PuzzleBase struct {
	Parts       map[string]func() int
	StringParts map[string]func() string
}

func (p *PuzzleBase) Run(part string) (string, error) {
	intFn, found := p.Parts[part]
	if found {
		return fmt.Sprintf("%d", intFn()), nil
	}
	stringFn, found := p.StringParts[part]
	if found {
		return stringFn(), nil
	}
	return "", errors.New(fmt.Sprintf("Part '%s' not found", part))
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
	return p.Run(part)
}
