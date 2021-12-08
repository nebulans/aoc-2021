package day03

import (
	"bufio"
	"fmt"
	"math/bits"
)

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

type Puzzle struct {
	readings chan uint64
	parts    map[string]func() int
}

func (p *Puzzle) Init() {
	p.readings = make(chan uint64)
	p.parts = map[string]func() int{
		"1": p.mostFrequentComplementProduct,
		"2": p.iterativeFilterProduct,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	go p.asyncParse(scanner)
}

func (p *Puzzle) asyncParse(scanner *bufio.Scanner) {
	for scanner.Scan() {
		value := uint64(0)
		text := scanner.Text()
		for pos := 0; pos < len(text); pos++ {
			if text[len(text)-pos-1] == '1' {
				value += bits.RotateLeft64(uint64(1), pos)
			}
		}
		p.readings <- value
	}
	close(p.readings)
}

func (p *Puzzle) mostFrequentComplementProduct() int {
	var values []uint64
	leadingZeros := 64
	for reading := range p.readings {
		values = append(values, reading)
		if bits.LeadingZeros64(reading) < leadingZeros {
			leadingZeros = bits.LeadingZeros64(reading)
		}
	}
	mask := bits.RotateLeft64(uint64(1), 64-leadingZeros) - 1
	ones := mostFrequentBit(values)
	return int(ones * (ones ^ mask))
}

func (p *Puzzle) iterativeFilterProduct() int {
	var values []uint64
	leadingZeros := 64
	for reading := range p.readings {
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

func (p *Puzzle) Dispatch(part string) (string, error) {
	result := p.parts[part]()
	return fmt.Sprintf("%d", result), nil
}
