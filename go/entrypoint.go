package main

import (
	"aoc-2021/days"
	"fmt"
	"github.com/alecthomas/kong"
)

var dayEntrypoints = map[string]func(string){
	"1": days.Day01,
	"2": days.Day02,
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
