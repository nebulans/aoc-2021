package framework

import "bufio"

type Puzzle interface {
	Init()
	Parse(*bufio.Scanner)
	Dispatch(string) (string, error)
	Run(string, *bufio.Scanner) (string, error)
}
