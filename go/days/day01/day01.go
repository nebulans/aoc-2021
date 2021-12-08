package day01

import (
	"bufio"
	"fmt"
	"strconv"
)

type Puzzle struct {
	readings chan int
}

func (p *Puzzle) Init() {
	p.readings = make(chan int)
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
	result := 0
	switch part {
	case "1":
		result = p.countIncreases()
	case "2":
		result = p.countWindowedIncreases()
	}
	return fmt.Sprintf("%d", result), nil
}

func (p *Puzzle) Run(part string, scanner *bufio.Scanner) (string, error) {
	p.Init()
	p.Parse(scanner)
	return p.Dispatch(part)
}
