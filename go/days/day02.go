package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func simpleStep(scanner *bufio.Scanner) int {
	depth := 0
	track := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		displacement, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("Unable to parse int")
		}
		switch parts[0] {
		case "forward":
			track += displacement
		case "down":
			depth += displacement
		case "up":
			depth -= displacement
		default:
			panic("Unrecognised direction")
		}
	}
	return track * depth
}

func aimedStep(scanner *bufio.Scanner) int {
	depth := 0
	track := 0
	aim := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		displacement, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("Unable to parse int")
		}
		switch parts[0] {
		case "forward":
			track += displacement
			depth += displacement * aim
		case "down":
			aim += displacement
		case "up":
			aim -= displacement
		default:
			panic("Unrecognised direction")
		}
	}
	return track * depth
}

func Day02(part string) {
	scanner := bufio.NewScanner(os.Stdin)
	partMap := map[string]func(scanner2 *bufio.Scanner) int{
		"1": simpleStep,
		"2": aimedStep,
	}
	fmt.Println(partMap[part](scanner))
}
