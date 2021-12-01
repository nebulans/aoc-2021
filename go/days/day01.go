package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countIncreases(scanner *bufio.Scanner) int {
	increases := -1
	last := 0
	for scanner.Scan() {
		reading, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Unable to parse int")
		}
		if reading > last {
			increases++
		}
		last = reading
	}
	return increases
}

func countWindowedIncreases(scanner *bufio.Scanner) int {
	increases := 0
	buffer := []int{0, 0, 0}
	position := 0
	for scanner.Scan() {
		reading, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Unable to parse int")
		}
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
	partMap := map[string]func(scanner2 *bufio.Scanner) int{
		"1": countIncreases,
		"2": countWindowedIncreases,
	}
	fmt.Println(partMap[part](scanner))
}
