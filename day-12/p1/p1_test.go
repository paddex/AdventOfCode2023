package p1

import (
	"os"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

type appTest struct {
	name  string
	input types.App
	want  int
}

func TestP1(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := types.App{
		Logger: logger,
		Input:  string(input),
	}

	tests := []appTest{
		{"testinput1", app, 21},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := P1(test.input)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

type lineTest struct {
	name  string
	input string
	want  int
}

func TestProcessLine(t *testing.T) {
	tests := []lineTest{
		{"line 1", "???.### 1,1,3", 1},
		{"line 2", ".??..??...?##. 1,1,3", 4},
		{"line 3", "?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"line 4", "????.#...#... 4,1,1", 1},
		{"line 5", "????.######..#####. 1,6,5", 4},
		{"line 6", "?###???????? 3,2,1", 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := processLine(test.input)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
