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
	"aoc-2021/framework"
	"bufio"
	"fmt"
	"github.com/alecthomas/kong"
	"os"
	"time"
)

var dayEntrypoints = map[string]func(string, *bufio.Scanner) (string, error){
	"6": day06.Day06,
	"7": day07.Day07,
	"8": day08.Day08,
}

var dayStructs = map[string]framework.Puzzle{
	"1": &day01.Puzzle{},
	"2": &day02.Puzzle{},
	"3": &day03.Puzzle{},
	"4": &day04.Puzzle{},
	"5": &day05.Puzzle{},
}

var CLI struct {
	Debug bool
	Day   string `arg:"" help:"Day to solve"`
	Part  string `arg:"" help:"Part to solve"`
}

func main() {
	kong.Parse(&CLI)
	dayFunc, funcFound := dayEntrypoints[CLI.Day]
	dayStruct, structFound := dayStructs[CLI.Day]
	if !funcFound && !structFound {
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
	} else {
		result, err = dayFunc(CLI.Part, bufio.NewScanner(os.Stdin))
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
