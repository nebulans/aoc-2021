package day04

import (
	"aoc-2021/framework"
	"aoc-2021/util/input"
	"bufio"
	"fmt"
	"strings"
)

type BingoCard struct {
	numbers []int
	seen    []bool
}

func makeCard(numbers []int) *BingoCard {
	return &BingoCard{
		numbers: numbers,
		seen:    make([]bool, len(numbers)),
	}
}

func (card *BingoCard) play(number int) bool {
	for i, v := range card.numbers {
		if v == number {
			card.seen[i] = true
			return true
		}
	}
	return false
}

func (card *BingoCard) winningMasks() [][]int {
	masks := make([][]int, 10)
	for i := 0; i < 5; i++ {
		row := make([]int, 5)
		col := make([]int, 5)
		for j := 0; j < 5; j++ {
			row[j] = i*5 + j
			col[j] = j*5 + i
		}
		masks[i*2] = row
		masks[i*2+1] = col
	}
	return masks
}

func (card *BingoCard) wins(mask []int) bool {
	for _, m := range mask {
		if card.seen[m] != true {
			return false
		}
	}
	return true
}

func (card *BingoCard) hasWon() bool {
	for _, mask := range card.winningMasks() {
		if card.wins(mask) {
			return true
		}
	}
	return false
}

func (card *BingoCard) score(lastCall int) int {
	sum := 0
	for i, n := range card.numbers {
		if card.seen[i] == false {
			sum += n
		}
	}
	return sum * lastCall
}

func (card *BingoCard) format() string {
	parts := make([]string, 25)
	for i, v := range card.numbers {
		sep := " "
		if i%5 == 4 {
			sep = "\n"
		}
		if card.seen[i] {
			parts[i] = fmt.Sprintf("(%2d)%s", v, sep)
		} else {
			parts[i] = fmt.Sprintf(" %2d %s", v, sep)
		}
	}
	return strings.Join(parts, "")
}

type Puzzle struct {
	framework.PuzzleBase
	calls []int
	cards []*BingoCard
}

func (p *Puzzle) Init() {
	p.Parts = map[string]func() int{
		"1":   p.winningScore,
		"2":   p.losingScore,
		"2am": p.losingArrayMask,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	scanner.Split(input.BlankLineSplitFunc)
	// First "field" of input is call order of numbers
	scanner.Scan()
	p.calls = input.SplitInts(scanner.Text(), ",")
	// Subsequent fields are bingo cards
	for scanner.Scan() {
		line := scanner.Text()
		boardNumbers := input.ParseInts(strings.Fields(line))
		p.cards = append(p.cards, makeCard(boardNumbers))
	}
}

func (p *Puzzle) winningScore() int {
	for _, n := range p.calls {
		for _, card := range p.cards {
			card.play(n)
			if card.hasWon() {
				return card.score(n)
			}
		}
	}
	return 0
}

func (p *Puzzle) losingScore() int {
	for _, n := range p.calls {
		for i := len(p.cards) - 1; i >= 0; i-- {
			card := p.cards[i]
			card.play(n)
			if card.hasWon() {
				if len(p.cards) == 1 {
					return card.score(n)
				}
				p.cards = append(p.cards[:i], p.cards[i+1:]...)
			}
		}
	}
	return 0
}

func (p *Puzzle) losingArrayMask() int {
	mask := make([]bool, len(p.cards))
	wins := 0
	for _, n := range p.calls {
		for i := 0; i < len(p.cards); i++ {
			if mask[i] {
				continue
			}
			card := p.cards[i]
			card.play(n)
			if card.hasWon() {
				wins++
				if wins == len(p.cards) {
					return card.score(n)
				}
				mask[i] = true
			}
		}
	}
	return 0
}
