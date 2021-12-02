package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInput(scanner *bufio.Scanner, readings chan<- int) {
	for scanner.Scan() {
		reading, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Unable to parse int")
		}
		readings <- reading
	}
	close(readings)
}

func countIncreases(readings <-chan int) int {
	increases := -1
	last := 0
	for reading := range readings {
		if reading > last {
			increases++
		}
		last = reading
	}
	return increases
}

func countWindowedIncreases(readings <-chan int) int {
	increases := 0
	buffer := []int{0, 0, 0}
	position := 0
	for reading := range readings {
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

func Day01(part string) {
	scanner := bufio.NewScanner(os.Stdin)
	readings := make(chan int)
	go parseInput(scanner, readings)
	partMap := map[string]func(<-chan int) int{
		"1": countIncreases,
		"2": countWindowedIncreases,
	}
	fmt.Println(partMap[part](readings))
}
