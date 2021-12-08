package day08

import (
	"aoc-2021/util/math/integer"
	"bufio"
	"fmt"
	"math/bits"
	"strings"
)

var segmentMap = map[rune]uint8{
	'a': uint8(0b00000001),
	'b': uint8(0b00000010),
	'c': uint8(0b00000100),
	'd': uint8(0b00001000),
	'e': uint8(0b00010000),
	'f': uint8(0b00100000),
	'g': uint8(0b01000000),
}

type DisplayState struct {
	digits [10]uint8
	output [4]uint8
}

func (d *DisplayState) Decode() int {
	digitMap := map[uint8]int{}
	valueMap := map[int]uint8{}
	// Resolve easy digits
	lengthMap := map[int]int{2: 1, 3: 7, 4: 4, 7: 8}
	for _, digit := range d.digits {
		val, found := lengthMap[bits.OnesCount8(digit)]
		if found {
			digitMap[digit] = val
			valueMap[val] = digit
		}

	}
	// Resolve harder digits
	for _, digit := range d.digits {
		switch bits.OnesCount8(digit) {
		case 5:
			// 2, 3 and 5
			if bits.OnesCount8(digit&valueMap[1]) == 2 {
				digitMap[digit] = 3
				valueMap[3] = digit
			} else if bits.OnesCount8(digit&valueMap[4]) == 3 {
				digitMap[digit] = 5
				valueMap[5] = digit
			} else {
				digitMap[digit] = 2
				valueMap[2] = digit
			}
		case 6:
			if bits.OnesCount8(digit&valueMap[4]) == 4 {
				digitMap[digit] = 9
				valueMap[9] = digit
			} else if bits.OnesCount8(digit&valueMap[1]) == 2 {
				digitMap[digit] = 0
				valueMap[0] = digit
			} else {
				digitMap[digit] = 6
				valueMap[6] = digit
			}
		}
	}
	// Calculate final value
	result := 0
	for i, v := range d.output {
		digit := digitMap[v]
		result += digit * integer.Pow(10, len(d.output)-i-1)
	}
	return result
}

func parseInput(scanner *bufio.Scanner, displays chan<- *DisplayState) {
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, " | ")
		digits := [10]uint8{}
		for i, s := range strings.Split(parts[0], " ") {
			sum := uint8(0)
			for _, r := range s {
				sum += segmentMap[r]
			}
			digits[i] = sum
		}
		output := [4]uint8{}
		for i, s := range strings.Split(parts[1], " ") {
			sum := uint8(0)
			for _, r := range s {
				sum += segmentMap[r]
			}
			output[i] = sum
		}
		state := &DisplayState{
			digits: digits,
			output: output,
		}
		displays <- state
	}
	close(displays)
}

func countUniqueReprInOutput(displays <-chan *DisplayState) int {
	count := 0
	lengths := map[int]int{
		2: 1, // 1
		3: 1, // 7
		4: 1, // 4
		7: 1, // 8
	}
	for display := range displays {
		for _, digit := range display.output {
			count += lengths[bits.OnesCount8(digit)]
		}
	}
	return count
}

func sumOutput(displays <-chan *DisplayState) int {
	s := 0
	for display := range displays {
		s += display.Decode()
	}
	return s
}

var partMap = map[string]func(<-chan *DisplayState) int{
	"1": countUniqueReprInOutput,
	"2": sumOutput,
}

func Day08(part string, input *bufio.Scanner) (string, error) {
	displays := make(chan *DisplayState)
	go parseInput(input, displays)
	result := partMap[part](displays)
	return fmt.Sprintf("%d", result), nil
}
