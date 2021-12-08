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
	"bufio"
	"fmt"
	"github.com/alecthomas/kong"
	"os"
	"time"
)

var dayEntrypoints = map[string]func(string, *bufio.Scanner) (string, error){
	"1": day01.Day01,
	"2": day02.Day02,
	"3": day03.Day03,
	"4": day04.Day04,
	"5": day05.Day05,
	"6": day06.Day06,
	"7": day07.Day07,
	"8": day08.Day08,
}

var CLI struct {
	Debug bool
	Day   string `arg:"" help:"Day to solve"`
	Part  string `arg:"" help:"Part to solve"`
}

func main() {
	kong.Parse(&CLI)
	dayFunc, found := dayEntrypoints[CLI.Day]
	if !found {
		_, err := fmt.Fprintf(os.Stderr, "Unknown day '%s'\n", CLI.Day)
		if err != nil {
			panic(err)
		}
		os.Exit(2)
	}
	startTime := time.Now()
	result, err := dayFunc(CLI.Part, bufio.NewScanner(os.Stdin))
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error reported by day function\n")
		if err != nil {
			panic(err)
		}
		os.Exit(2)
	}
	elapsed := time.Now().Sub(startTime)
	fmt.Printf("Elapsed time: %s\n", elapsed)
	fmt.Printf("\n%s\n", result)
}
