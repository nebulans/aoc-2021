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

type DisplayMapper struct {
	digitMap map[uint8]int
	valueMap map[int]uint8
}

func NewDisplayMapper() *DisplayMapper {
	return &DisplayMapper{
		digitMap: map[uint8]int{},
		valueMap: map[int]uint8{},
	}
}

func (m *DisplayMapper) Map(digit uint8, value int) {
	m.digitMap[digit] = value
	m.valueMap[value] = digit
}

func (m *DisplayMapper) GetDigit(value int) uint8 {
	return m.valueMap[value]
}

func (m *DisplayMapper) GetValue(digit uint8) int {
	return m.digitMap[digit]
}

func (m *DisplayMapper) SharedSegments(digit uint8, value int) int {
	return bits.OnesCount8(digit & m.GetDigit(value))
}

type DisplayState struct {
	digits [10]uint8
	output [4]uint8
}

func (d *DisplayState) Decode() int {
	mapper := NewDisplayMapper()
	// Resolve easy digits
	lengthMap := map[int]int{2: 1, 3: 7, 4: 4, 7: 8}
	for _, digit := range d.digits {
		val, found := lengthMap[bits.OnesCount8(digit)]
		if found {
			mapper.Map(digit, val)
		}

	}
	// Resolve harder digits
	for _, digit := range d.digits {
		switch bits.OnesCount8(digit) {
		case 5:
			// 2, 3 and 5
			if mapper.SharedSegments(digit, 1) == 2 {
				mapper.Map(digit, 3)
			} else if mapper.SharedSegments(digit, 4) == 3 {
				mapper.Map(digit, 5)
			} else {
				mapper.Map(digit, 2)
			}
		case 6:
			// 6, 9 and 0
			if mapper.SharedSegments(digit, 4) == 4 {
				mapper.Map(digit, 9)
			} else if mapper.SharedSegments(digit, 1) == 2 {
				mapper.Map(digit, 0)
			} else {
				mapper.Map(digit, 6)
			}
		}
	}
	// Calculate final value
	result := 0
	for i, v := range d.output {
		digit := mapper.GetValue(v)
		result += digit * integer.Pow(10, len(d.output)-i-1)
	}
	return result
}

func parseDigitSet(line string) []uint8 {
	parts := strings.Split(line, " ")
	digits := make([]uint8, len(parts))
	for i, s := range parts {
		sum := uint8(0)
		for _, r := range s {
			sum += segmentMap[r]
		}
		digits[i] = sum
	}
	return digits
}

func parseInput(scanner *bufio.Scanner, displays chan<- *DisplayState) {
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, " | ")
		digits := [10]uint8{}
		copy(digits[:], parseDigitSet(parts[0]))
		output := [4]uint8{}
		copy(output[:], parseDigitSet(parts[1]))
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
