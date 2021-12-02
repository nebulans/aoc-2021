package day02

import (
	"bufio"
	"fmt"
	"os"
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

func parseInput(scanner *bufio.Scanner, instructions chan<- Instruction) {
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
		instructions <- Instruction{direction, displacement}
	}
	close(instructions)
}

func simpleStep(instructions <-chan Instruction) int {
	depth := 0
	track := 0
	for instruction := range instructions {
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

func aimedStep(instructions <-chan Instruction) int {
	depth := 0
	track := 0
	aim := 0
	for instruction := range instructions {
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

func Day02(part string) {
	scanner := bufio.NewScanner(os.Stdin)
	instructions := make(chan Instruction)
	go parseInput(scanner, instructions)
	partMap := map[string]func(<-chan Instruction) int{
		"1": simpleStep,
		"2": aimedStep,
	}
	fmt.Println(partMap[part](instructions))
}
