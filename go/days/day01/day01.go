package day01

import (
	"bufio"
	"fmt"
	"strconv"
)

type Puzzle struct {
	readings chan int
	parts    map[string]func() int
}

func (p *Puzzle) Init() {
	p.readings = make(chan int)
	p.parts = map[string]func() int{
		"1": p.countIncreases,
		"2": p.countWindowedIncreases,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	go p.asyncParse(scanner)
}

func (p *Puzzle) asyncParse(scanner *bufio.Scanner) {
	for scanner.Scan() {
		reading, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Unable to parse int")
		}
		p.readings <- reading
	}
	close(p.readings)
}

func (p *Puzzle) countIncreases() int {
	increases := -1
	last := 0
	for reading := range p.readings {
		if reading > last {
			increases++
		}
		last = reading
	}
	return increases
}

func (p *Puzzle) countWindowedIncreases() int {
	increases := 0
	buffer := []int{0, 0, 0}
	position := 0
	for reading := range p.readings {
		if position >= 3 {
			if reading > buffer[position%3] {
				increases++
			}
		}
		buffer[position%3] = reading
		position++
	}
	return increases
}

func (p *Puzzle) Dispatch(part string) (string, error) {
	result := p.parts[part]()
	return fmt.Sprintf("%d", result), nil
}
