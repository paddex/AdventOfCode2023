package p2

import (
	"os"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

type p2test struct {
	name  string
	input types.App
	want  int
}

type symbolTest struct {
	name  string
	input []pipe
	want  string
}

func TestP2(t *testing.T) {
	input, err := os.ReadFile("../testinput2")
	if err != nil {
		panic("Can't read file")
	}
	input2, err := os.ReadFile("../testinput2-2")
	if err != nil {
		panic("Can't read file")
	}
	input3, err := os.ReadFile("../testinput2-3")
	if err != nil {
		panic("Can't read file")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := types.App{
		Logger: logger,
		Input:  string(input),
	}
	app2 := types.App{
		Logger: logger,
		Input:  string(input2),
	}
	app3 := types.App{
		Logger: logger,
		Input:  string(input3),
	}

	tests := []p2test{
		{"Test 1", app, 4},
		{"Test 2", app2, 8},
		{"Test 3", app3, 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := P2(test.input)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestFindStartingSymbol(t *testing.T) {
	input, err := os.ReadFile("../testinput2")
	if err != nil {
		panic("Can't read file")
	}
	input2, err := os.ReadFile("../testinput2-2")
	if err != nil {
		panic("Can't read file")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := types.App{
		Logger: logger,
		Input:  string(input),
	}
	app2 := types.App{
		Logger: logger,
		Input:  string(input2),
	}

	b1 := getBoard(app.Input)
	s1 := findStart(b1)
	ps1 := findStartPipes(b1, s1)
	b2 := getBoard(app2.Input)
	s2 := findStart(b2)
	ps2 := findStartPipes(b2, s2)

	tests := []symbolTest{
		{"Test 1", ps1, "F"},
		{"Test 2", ps2, "F"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := findStartingSymbol(test.input)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
