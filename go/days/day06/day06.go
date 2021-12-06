package day06

import (
	"aoc-2021/util/math/integer"
	"bufio"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func parseInput(scanner *bufio.Scanner) []int {
	scanner.Scan()
	text := scanner.Text()
	parts := strings.Split(text, ",")
	out := make([]int, len(parts))
	for i, s := range parts {
		v, _ := strconv.Atoi(s)
		out[i] = v
	}
	return out
}

func simulate(initialState []int, steps int) string {
	ages := [9]*big.Int{}
	for i := 0; i < 9; i++ {
		ages[i] = big.NewInt(0)
	}
	increment := big.NewInt(1)
	for _, s := range initialState {
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

var partMap = map[string]func([]int) string{
	"1": func(ages []int) string { return simulate(ages, 80) },
	"2": func(ages []int) string { return simulate(ages, 256) },
}

func Day06(part string, input *bufio.Scanner) (string, error) {
	timers := parseInput(input)
	result := partMap[part](timers)
	return result, nil
}
