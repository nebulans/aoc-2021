package day02

import (
	"aoc-2021/framework"
	"bufio"
	"strconv"
	"strings"
)

type Direction int8

const (
	directionForward Direction = iota
	directionDown
	directionUp
)

type Instruction struct {
	direction Direction
	distance  int
}

type Puzzle struct {
	framework.PuzzleBase
	instructions chan Instruction
}

func (p *Puzzle) Init() {
	p.instructions = make(chan Instruction)
	p.Parts = map[string]func() int{
		"1": p.simpleStep,
		"2": p.aimedStep,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	go p.asyncParse(scanner)
}

func (p *Puzzle) asyncParse(scanner *bufio.Scanner) {
	directionMap := map[string]Direction{
		"forward": directionForward,
		"down":    directionDown,
		"up":      directionUp,
	}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		displacement, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("Unable to parse int")
		}
		direction, found := directionMap[parts[0]]
		if !found {
			panic("Unable to parse direction")
		}
		p.instructions <- Instruction{direction, displacement}
	}
	close(p.instructions)
}

func (p *Puzzle) simpleStep() int {
	depth := 0
	track := 0
	for instruction := range p.instructions {
		switch instruction.direction {
		case directionForward:
			track += instruction.distance
		case directionDown:
			depth += instruction.distance
		case directionUp:
			depth -= instruction.distance
		}
	}
	return track * depth
}

func (p *Puzzle) aimedStep() int {
	depth := 0
	track := 0
	aim := 0
	for instruction := range p.instructions {
		switch instruction.direction {
		case directionForward:
			track += instruction.distance
			depth += instruction.distance * aim
		case directionDown:
			aim += instruction.distance
		case directionUp:
			aim -= instruction.distance
		}
	}
	return track * depth
}
