package main

import (
	"aoc-2021/days/day01"
	"aoc-2021/days/day02"
	"fmt"
	"github.com/alecthomas/kong"
)

var dayEntrypoints = map[string]func(string){
	"1": day01.Day01,
	"2": day02.Day02,
}

var CLI struct {
	Debug bool
	Day   string `arg:""`
	Part  string `arg:""`
}

func main() {
	ctx := kong.Parse(&CLI)
	ctx.Run()
	dayFunc, found := dayEntrypoints[CLI.Day]
	if !found {
		panic(fmt.Sprintf("Unknown day '%s'", CLI.Day))
	}
	dayFunc(CLI.Part)
}
