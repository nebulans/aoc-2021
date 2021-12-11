package main

import (
	"aoc-2021/days/day01"
	"aoc-2021/days/day02"
	"aoc-2021/days/day03"
	"aoc-2021/days/day04"
	"aoc-2021/days/day05"
	"aoc-2021/days/day06"
	"aoc-2021/days/day07"
	"aoc-2021/days/day08"
	"aoc-2021/days/day09"
	"aoc-2021/days/day10"
	"aoc-2021/days/day11"
	"aoc-2021/framework"
	"bufio"
	"fmt"
	"github.com/alecthomas/kong"
	"os"
	"time"
)

var dayStructs = map[string]framework.Puzzle{
	"1":  &day01.Puzzle{},
	"2":  &day02.Puzzle{},
	"3":  &day03.Puzzle{},
	"4":  &day04.Puzzle{},
	"5":  &day05.Puzzle{},
	"6":  &day06.Puzzle{},
	"7":  &day07.Puzzle{},
	"8":  &day08.Puzzle{},
	"9":  &day09.Puzzle{},
	"10": &day10.Puzzle{},
	"11": &day11.Puzzle{},
}

var CLI struct {
	Debug bool
	Day   string `arg:"" help:"Day to solve"`
	Part  string `arg:"" help:"Part to solve"`
}

func main() {
	kong.Parse(&CLI)
	dayStruct, structFound := dayStructs[CLI.Day]
	if !structFound {
		_, err := fmt.Fprintf(os.Stderr, "Unknown day '%s'\n", CLI.Day)
		if err != nil {
			panic(err)
		}
		os.Exit(2)
	}
	startTime := time.Now()
	var result string
	var err error
	if structFound {
		result, err = framework.RunPuzzle(dayStruct, CLI.Part, bufio.NewScanner(os.Stdin))
	}
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error reported by day function: %v", err)
		if err != nil {
			panic(err)
		}
		os.Exit(2)
	}
	elapsed := time.Now().Sub(startTime)
	fmt.Printf("Elapsed time: %s\n", elapsed)
	fmt.Printf("\n%s\n", result)
}
