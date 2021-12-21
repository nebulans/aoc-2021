package day21

import (
	"aoc-2021/framework"
	"aoc-2021/util/datastructure/stack"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type DeterministicDice struct {
	faces int
	rolls int
}

func (d *DeterministicDice) Roll() int {
	n := (d.rolls % d.faces) + 1
	d.rolls++
	return n
}

type Player struct {
	position  int
	score     int
	boardSize int
}

func (p *Player) playRound(rolls []int) {
	for _, r := range rolls {
		p.position += r
	}
	p.position = p.position % p.boardSize
	if p.position == 0 {
		p.score += 10 // Position 0 is actually 10 on game board
	} else {
		p.score += p.position
	}
}

func (p *Player) Copy() *Player {
	return &Player{
		position:  p.position,
		score:     p.score,
		boardSize: p.boardSize,
	}
}

type GameResult struct {
	Won           bool
	WinningPlayer int
	Winner        *Player
	Loser         *Player
}

type DiceGame struct {
	players  [2]*Player
	next     int
	rounds   int
	winScore int
}

func (g *DiceGame) PlayRound(rolls []int) GameResult {
	g.rounds++
	g.players[g.next].playRound(rolls)
	result := GameResult{Won: false}
	if g.players[g.next].score >= g.winScore {
		result = GameResult{
			Won:           true,
			WinningPlayer: g.next,
			Winner:        g.players[g.next],
			Loser:         g.players[(g.next+1)%2],
		}
	}
	g.next = (g.next + 1) % 2
	return result
}

func (g *DiceGame) Copy() *DiceGame {
	return &DiceGame{
		players:  [2]*Player{g.players[0].Copy(), g.players[1].Copy()},
		next:     g.next,
		rounds:   g.rounds,
		winScore: g.winScore,
	}
}

func (g *DiceGame) FormatState() string {
	return fmt.Sprintf(
		"Player 1: position %d, score %d; Player 2: position %d, score %d",
		g.players[0].position, g.players[0].score,
		g.players[1].position, g.players[1].score,
	)
}

type Puzzle struct {
	framework.PuzzleBase
	startingGame *DiceGame
}

func (p *Puzzle) Init() {
	p.Parts = map[string]func() int{
		"1": p.deterministicDiceScore,
		"2": p.QuantumDiceWins,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	players := [2]*Player{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		position, _ := strconv.Atoi(parts[1])
		players[i] = &Player{
			score:     0,
			position:  position,
			boardSize: 10,
		}
		i++
	}
	p.startingGame = &DiceGame{players: players}
}

func (p *Puzzle) deterministicDiceScore() int {
	p.startingGame.winScore = 1000
	fmt.Println(p.startingGame.FormatState())
	dice := &DeterministicDice{rolls: 0, faces: 100}
	for true {
		rolls := []int{dice.Roll(), dice.Roll(), dice.Roll()}
		result := p.startingGame.PlayRound(rolls)
		if result.Won {
			fmt.Printf("Losing player scores %d, dice rolled %d times\n", result.Loser.score, dice.rolls)
			return result.Loser.score * dice.rolls
		}
	}
	return 0
}

type GameState struct {
	game  *DiceGame
	count int
}

var quantumRolls = map[int]int{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

func (p *Puzzle) QuantumDiceWins() int {
	p.startingGame.winScore = 21
	game := &GameState{game: p.startingGame.Copy(), count: 1}
	gameStack := stack.MakeStack(200)
	gameStack.Push(game)
	wins, losses := 0, 0
	x := 0
	for !gameStack.IsEmpty() {
		if x%1000000 == 0 {
			fmt.Printf("Iteration %d: stack %d, won %d, lost %d\n", x, gameStack.Len(), wins, losses)
		}
		gs := gameStack.Pop().(*GameState)
		for roll, count := range quantumRolls {
			gg := gs.game.Copy()
			result := gg.PlayRound([]int{roll})
			states := gs.count * count
			if result.Won {
				if result.WinningPlayer == 0 {
					wins += states
				} else {
					losses += states
				}
			} else {
				gameStack.Push(&GameState{game: gg, count: states})
			}
		}
		x++
	}
	return wins
}
