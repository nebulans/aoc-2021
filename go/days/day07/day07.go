package day07

import (
	"aoc-2021/util/input"
	"aoc-2021/util/math/integer"
	"bufio"
	"fmt"
	"math"
)

func constantCost(position int, target int) int {
	return integer.Abs(position - target)
}

func sumBelowCost(position int, target int) int {
	steps := constantCost(position, target)
	return ((2 * steps) * (steps + 1)) / 4
}

func costScan(positions []int, costFunc func(int, int) int) int {
	min := integer.MinSlice(positions)
	max := integer.MaxSlice(positions)
	best := math.MaxInt
	for t := min; t <= max; t++ {
		cost := 0
		for _, v := range positions {
			cost += costFunc(v, t)
		}
		if cost < best {
			best = cost
		}
	}
	return best
}

func parseInput(scanner *bufio.Scanner) []int {
	scanner.Scan()
	return input.SplitInts(scanner.Text(), ",")
}

var partMap = map[string]func(positions []int) int{
	"1": func(p []int) int { return costScan(p, constantCost) },
	"2": func(p []int) int { return costScan(p, sumBelowCost) },
}

func Day07(part string, input *bufio.Scanner) (string, error) {
	positions := parseInput(input)
	result := partMap[part](positions)
	return fmt.Sprintf("%d", result), nil
}
