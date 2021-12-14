package day14

import (
	"aoc-2021/framework"
	"bufio"
	"regexp"
	"sort"
)

type PolymerChain struct {
	pairs map[[2]rune]int
	head  rune
	tail  rune
}

func (c *PolymerChain) SimulateStep(rules map[[2]rune]rune) {
	newPairs := map[[2]rune]int{}
	for pair, count := range c.pairs {
		i := rules[pair]
		newPairs[[2]rune{pair[0], i}] += count
		newPairs[[2]rune{i, pair[1]}] += count
	}
	c.pairs = newPairs
}

func (c *PolymerChain) RuneCounts() map[rune]int {
	counts := map[rune]int{}
	for pair, count := range c.pairs {
		counts[pair[0]] += count
		counts[pair[1]] += count
	}
	counts[c.head]++
	counts[c.tail]++
	for r, c := range counts {
		counts[r] = c / 2
	}
	return counts
}

func MakePolymerChain(elements []rune) *PolymerChain {
	pairs := map[[2]rune]int{}
	for i := 0; i < len(elements)-1; i++ {
		pairs[[2]rune{elements[i], elements[i+1]}]++
	}
	return &PolymerChain{
		pairs: pairs,
		head:  elements[0],
		tail:  elements[len(elements)-1],
	}
}

type Puzzle struct {
	framework.PuzzleBase
	polymerTemplate []rune
	rules           map[[2]rune]rune
}

func (p *Puzzle) Init() {
	p.polymerTemplate = make([]rune, 0)
	p.rules = map[[2]rune]rune{}
	p.Parts = map[string]func() int{
		"1": func() int { return p.elementCountDifference(10) },
		"2": func() int { return p.elementCountDifference(40) },
	}
}

func getFirstRune(s string) rune {
	return []rune(s)[0]
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	scanner.Scan()
	template := scanner.Text()
	p.polymerTemplate = []rune(template)
	scanner.Scan()

	pattern, _ := regexp.Compile("([A-Z])([A-Z]) -> ([A-Z])")
	for scanner.Scan() {
		line := scanner.Text()
		match := pattern.FindStringSubmatch(line)
		p.rules[[2]rune{getFirstRune(match[1]), getFirstRune(match[2])}] = getFirstRune(match[3])
	}
}

func (p *Puzzle) elementCountDifference(iterations int) int {
	chain := MakePolymerChain(p.polymerTemplate)
	for i := 0; i < iterations; i++ {
		chain.SimulateStep(p.rules)
	}
	runeCounts := chain.RuneCounts()
	counts := make([]int, 0, len(runeCounts))
	for _, v := range runeCounts {
		counts = append(counts, v)
	}
	sort.Ints(counts)
	return counts[len(counts)-1] - counts[0]
}
