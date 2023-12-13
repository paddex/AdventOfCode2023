package p2

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
		{"testinput1", app, 400},
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

type processTest struct {
	name  string
	input string
	want  int
}

func TestProcess(t *testing.T) {
	inp1 := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`

	inp2 := `#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

	tests := []processTest{
		{"pattern 1", inp1, 300},
		{"pattern 2", inp2, 100},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := processPattern(test.input)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
