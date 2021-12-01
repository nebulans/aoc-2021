package main

import (
	"aoc-2021/days"
	"github.com/alecthomas/kong"
)

var dayEntrypoints = map[string]func(string){
	"1": days.Day01,
}

var CLI struct {
	Debug bool
	Day   string `arg:""`
	Part  string `arg:""`
}

func main() {
	ctx := kong.Parse(&CLI)
	ctx.Run()
	dayEntrypoints[CLI.Day](CLI.Part)
}
