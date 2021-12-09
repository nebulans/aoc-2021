package day06

import (
	"aoc-2021/framework"
	"aoc-2021/util/math/integer"
	"bufio"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type Puzzle struct {
	framework.PuzzleBase
	initialState []int
}

func (p *Puzzle) Init() {
	p.StringParts = map[string]func() string{
		"1": func() string { return p.Simulate(80) },
		"2": func() string { return p.Simulate(256) },
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	scanner.Scan()
	text := scanner.Text()
	parts := strings.Split(text, ",")
	p.initialState = make([]int, len(parts))
	for i, s := range parts {
		v, _ := strconv.Atoi(s)
		p.initialState[i] = v
	}
}

func (p *Puzzle) Simulate(steps int) string {
	ages := [9]*big.Int{}
	for i := 0; i < 9; i++ {
		ages[i] = big.NewInt(0)
	}
	increment := big.NewInt(1)
	for _, s := range p.initialState {
		ages[s].Add(ages[s], increment)
	}
	for i := 0; i < steps; i++ {
		newAges := [9]*big.Int{}
		for a := len(ages) - 1; a >= 0; a-- {
			if a > 0 {
				newAges[a-1] = ages[a]
			} else {
				newAges[8] = ages[0]
				newAges[6].Add(newAges[6], ages[0])
			}
		}
		ages = newAges
	}
	sum := integer.SumBigInt(ages[:])
	digits := fmt.Sprintf("%d", sum)
	return digits
}
