package p1

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

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

	got := P1(app)
	want := 13

	if got != want {
		t.Errorf("Ergebnis: %d, Erwartet: %d", got, want)
	}
}

func TestProcessLine(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	// logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// app := types.App{
	// 	Logger: logger,
	// 	Input:  string(input),
	// }

	lines := strings.Split(string(input), "\n")

	type test struct {
		name  string
		input string
		want  int
	}

	tests := []test{
		{"Card 1", lines[0], 8},
		{"Card 2", lines[1], 2},
		{"Card 3", lines[2], 2},
		{"Card 4", lines[3], 1},
		{"Card 5", lines[4], 0},
		{"Card 6", lines[5], 0},
		{"Card 7", "Card  7: 88  2 12 | 77  2 12 65", 2},
		{"Card 8", "Card  7: 88 32 12 | 77 32 12 65", 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := processLine(test.input)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
