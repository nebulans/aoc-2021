package main

import (
	"aoc-2021/days/day01"
	"aoc-2021/days/day02"
	"bufio"
	"fmt"
	"github.com/alecthomas/kong"
	"os"
)

var dayEntrypoints = map[string]func(string, *bufio.Scanner) (string, error){
	"1": day01.Day01,
	"2": day02.Day02,
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
	result, err := dayFunc(CLI.Part, bufio.NewScanner(os.Stdin))
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error reported by day function\n")
		if err != nil {
			panic(err)
		}
		os.Exit(2)
	}
	fmt.Println(result)
}
