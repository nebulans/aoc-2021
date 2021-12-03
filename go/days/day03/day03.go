package day03

import (
	"bufio"
	"fmt"
	"math/bits"
)

func parseInput(scanner *bufio.Scanner, readings chan<- uint64) {
	for scanner.Scan() {
		value := uint64(0)
		text := scanner.Text()
		for pos := 0; pos < len(text); pos++ {
			if text[len(text)-pos-1] == '1' {
				value += bits.RotateLeft64(uint64(1), pos)
			}
		}
		readings <- value
	}
	close(readings)
}

func mostFrequentBit(values []uint64) uint64 {
	counts := [64]int{}
	for _, value := range values {
		for pos := 0; pos < 64; pos++ {
			if value&bits.RotateLeft64(uint64(1), 63-pos) > 0 {
				counts[pos]++
			}
		}
	}
	threshold := len(values) / 2
	if len(values)%2 == 0 {
		threshold -= 1
	}
	result := uint64(0)
	for pos := 0; pos < 64; pos++ {
		if counts[pos] > threshold {
			result += bits.RotateLeft64(uint64(1), 63-pos)
		}
	}
	return result
}

func mostFrequentComplementProduct(readings <-chan uint64) int {
	var values []uint64
	leadingZeros := 64
	for reading := range readings {
		values = append(values, reading)
		if bits.LeadingZeros64(reading) < leadingZeros {
			leadingZeros = bits.LeadingZeros64(reading)
		}
	}
	mask := bits.RotateLeft64(uint64(1), 64-leadingZeros) - 1
	ones := mostFrequentBit(values)
	return int(ones * (ones ^ mask))
}

func iterativeFilter(numbers []uint64, filterMask uint64, targetOnes bool) uint64 {
	values := make([]uint64, len(numbers))
	copy(values, numbers)
	oneCounts := mostFrequentBit(values)
	for pos := 0; pos < 64; pos++ {
		mask := bits.RotateLeft64(uint64(1), 63-pos)
		if mask&filterMask == 0 {
			continue
		}
		for i := len(values) - 1; i >= 0; i-- {
			if values[i]&mask != oneCounts&mask == targetOnes {
				values = append(values[:i], values[i+1:]...)
			}
		}
		oneCounts = mostFrequentBit(values)
		if len(values) == 1 {
			return values[0]
		}
	}
	return uint64(0)
}

func iterativeFilterProduct(readings <-chan uint64) int {
	var values []uint64
	leadingZeros := 64
	for reading := range readings {
		values = append(values, reading)
		if bits.LeadingZeros64(reading) < leadingZeros {
			leadingZeros = bits.LeadingZeros64(reading)
		}
	}
	mask := bits.RotateLeft64(uint64(1), 64-leadingZeros) - 1
	ones := iterativeFilter(values, mask, true)
	zeroes := iterativeFilter(values, mask, false)
	return int(ones * zeroes)
}

var partMap = map[string]func(<-chan uint64) int{
	"1": mostFrequentComplementProduct,
	"2": iterativeFilterProduct,
}

func Day03(part string, input *bufio.Scanner) (string, error) {
	readings := make(chan uint64)
	go parseInput(input, readings)
	result := partMap[part](readings)
	return fmt.Sprintf("%d", result), nil
}
