package day10

import (
	"aoc-2021/framework"
	stack2 "aoc-2021/util/datastructure/stack"
	"bufio"
	"fmt"
	"sort"
)

var chunkClosings = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var corruptScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionScores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

type SyntaxLine struct {
	characters  []rune
	corrupt     bool
	corruptChar rune
	completion  []rune
}

func (s *SyntaxLine) Format() string {
	return fmt.Sprintf("%v", s.characters)
}

func (s *SyntaxLine) Process() {
	stack := stack2.MakeStack(len(s.characters))
	for _, c := range s.characters {
		closing, found := chunkClosings[c]
		if found {
			stack.Push(closing)
		} else {
			expected := stack.Pop().(rune)
			if c != expected {
				s.corrupt = true
				s.corruptChar = c
				return
			}
		}
	}
	r := stack.Remaining()
	s.completion = make([]rune, len(r))
	for i, v := range r {
		s.completion[i] = v.(rune)
	}
}

func MakeLine(text string) *SyntaxLine {
	chars := make([]rune, len(text))
	for i, c := range text {
		chars[i] = c
	}
	return &SyntaxLine{
		chars,
		false,
		' ',
		nil,
	}
}

type Puzzle struct {
	framework.PuzzleBase
	lines chan *SyntaxLine
}

func (p *Puzzle) Init() {
	p.lines = make(chan *SyntaxLine)
	p.Parts = map[string]func() int{
		"1": p.corruptScores,
		"2": p.incompleteScores,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	go p.asyncParse(scanner)
}

func (p *Puzzle) asyncParse(scanner *bufio.Scanner) {
	for scanner.Scan() {
		text := scanner.Text()
		p.lines <- MakeLine(text)
	}
	close(p.lines)
}

func (p *Puzzle) corruptScores() int {
	score := 0
	for line := range p.lines {
		line.Process()
		if line.corrupt {
			score += corruptScores[line.corruptChar]
		}
	}
	return score
}

func (p *Puzzle) incompleteScores() int {
	scores := make([]int, 0)
	for line := range p.lines {
		line.Process()
		if !line.corrupt {
			score := 0
			for _, c := range line.completion {
				score *= 5
				score += completionScores[c]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}
