package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	distance  int
}

func parseInput(scanner *bufio.Scanner, instructions chan<- instruction) {
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		displacement, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("Unable to parse int")
		}
		instructions <- instruction{parts[0], displacement}
	}
	close(instructions)
}

func simpleStep(instructions <-chan instruction) int {
	depth := 0
	track := 0
	for i := range instructions {
		switch i.direction {
		case "forward":
			track += i.distance
		case "down":
			depth += i.distance
		case "up":
			depth -= i.distance
		default:
			panic("Unrecognised direction")
		}
	}
	return track * depth
}

func aimedStep(instructions <-chan instruction) int {
	depth := 0
	track := 0
	aim := 0
	for i := range instructions {
		switch i.direction {
		case "forward":
			track += i.distance
			depth += i.distance * aim
		case "down":
			aim += i.distance
		case "up":
			aim -= i.distance
		default:
			panic("Unrecognised direction")
		}
	}
	return track * depth
}

func Day02(part string) {
	scanner := bufio.NewScanner(os.Stdin)
	instructions := make(chan instruction)
	go parseInput(scanner, instructions)
	partMap := map[string]func(<-chan instruction) int{
		"1": simpleStep,
		"2": aimedStep,
	}
	fmt.Println(partMap[part](instructions))
}
