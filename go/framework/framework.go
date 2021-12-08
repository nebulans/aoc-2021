package framework

import "bufio"

type Puzzle interface {
	Init()
	Parse(*bufio.Scanner)
	Dispatch(string) (string, error)
}

func RunPuzzle(p Puzzle, part string, input *bufio.Scanner) (string, error) {
	p.Init()
	p.Parse(input)
	return p.Dispatch(part)
}
