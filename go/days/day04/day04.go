package day04

import (
	"aoc-2021/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type BingoCard struct {
	numbers []int
	seen    []bool
}

func makeCard(numbers []int) BingoCard {
	return BingoCard{
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

func parseInts(strings []string) []int {
	numbers := make([]int, len(strings))
	for i, n := range strings {
		numbers[i], _ = strconv.Atoi(n)
	}
	return numbers
}

func parseInput(scanner *bufio.Scanner) ([]int, []BingoCard) {
	scanner.Split(util.BlankLineSplitFunc)
	// First "field" of input is call order of numbers
	scanner.Scan()
	numbers := parseInts(strings.Split(scanner.Text(), ","))
	// Subsequent fields are bingo cards
	var boards []BingoCard
	for scanner.Scan() {
		line := scanner.Text()
		boardNumbers := parseInts(strings.Fields(line))
		boards = append(boards, makeCard(boardNumbers))
	}
	return numbers, boards
}

func winningScore(numbers []int, cards []BingoCard) int {
	for _, n := range numbers {
		for _, card := range cards {
			card.play(n)
			if card.hasWon() {
				return card.score(n)
			}
		}
	}
	return 0
}

func losingScore(numbers []int, cards []BingoCard) int {
	for _, n := range numbers {
		for i := len(cards) - 1; i >= 0; i-- {
			card := cards[i]
			card.play(n)
			if card.hasWon() {
				if len(cards) == 1 {
					return card.score(n)
				}
				cards = append(cards[:i], cards[i+1:]...)
			}
		}
	}
	return 0
}

var partMap = map[string]func([]int, []BingoCard) int{
	"1": winningScore,
	"2": losingScore,
}

func Day04(part string, input *bufio.Scanner) (string, error) {
	numbers, cards := parseInput(input)
	result := partMap[part](numbers, cards)
	return fmt.Sprintf("%d", result), nil
}
