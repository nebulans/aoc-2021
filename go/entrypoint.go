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
	"aoc-2021/days/day12"
	"aoc-2021/days/day13"
	"aoc-2021/days/day14"
	"aoc-2021/days/day15"
	"aoc-2021/days/day16"
	"aoc-2021/days/day17"
	"aoc-2021/days/day18"
	"aoc-2021/framework"
	"bufio"
	"fmt"
	"github.com/alecthomas/kong"
	"log"
	"os"
	"time"
)

var dayStructs = map[string]framework.Puzzle{
	"1":   &day01.Puzzle{},
	"2":   &day02.Puzzle{},
	"3":   &day03.Puzzle{},
	"4":   &day04.Puzzle{},
	"5":   &day05.Puzzle{},
	"6":   &day06.Puzzle{},
	"7":   &day07.Puzzle{},
	"8":   &day08.Puzzle{},
	"9":   &day09.Puzzle{},
	"10":  &day10.Puzzle{},
	"11":  &day11.Puzzle{GridImpl: "array"},
	"11m": &day11.Puzzle{GridImpl: "map"},
	"12":  &day12.Puzzle{},
	"13":  &day13.Puzzle{},
	"14":  &day14.Puzzle{},
	"15":  &day15.Puzzle{Impl: "gridCosts"},
	"16":  &day16.Puzzle{},
	"17":  &day17.Puzzle{},
	"18":  &day18.Puzzle{},
}

var CLI struct {
	Debug bool
	Input string
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
	var f *os.File
	if CLI.Input != "" {
		h, err := os.Open(CLI.Input)
		if err != nil {
			log.Fatal(err)
		}
		defer h.Close()
		f = h
	} else {
		f = os.Stdin
	}
	if structFound {
		result, err = framework.RunPuzzle(dayStruct, CLI.Part, bufio.NewScanner(f))
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
