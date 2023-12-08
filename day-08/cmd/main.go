package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/p1"
	"paddex.net/aoc/p2"
	"paddex.net/aoc/types"
)

func main() {
	var runPart1 bool
	var runPart2 bool
	var debug bool

	flag.BoolVar(&runPart1, "part1", false, "Run Problem 1")
	flag.BoolVar(&runPart2, "part2", false, "Run Problem 2")
	flag.BoolVar(&debug, "debug", false, "Output debugging info")

	flag.Parse()

	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))

	input, err := os.ReadFile("input")
	if err != nil {
		panic("Can't read file")
	}

	app := types.App{
		Logger: logger,
		Input:  string(input),
	}

	if !runPart1 && !runPart2 {
		runPart1 = true
		runPart2 = true
	}

	if runPart1 {
		part1(app)
	}
	if runPart2 {
		part2(app)
	}
}

func part1(app types.App) {
	num := p1.P1(app)
	fmt.Printf("Final result for part 1: %d\n", num)
}

func part2(app types.App) {
	num := p2.P2(app)
	fmt.Printf("Final result for part 2: %d\n", num)
}
