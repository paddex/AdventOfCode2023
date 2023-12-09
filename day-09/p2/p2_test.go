package p2

import (
	"os"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

type valTest struct {
	name  string
	input []int
	want  int
}

type arrTest struct {
	name  string
	input string
	want  []int
}

func TestP2(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := types.App{
		Logger: logger,
		Input:  string(input),
	}

	got := P2(app)
	want := 2

	if got != want {
		t.Errorf("Ergebnis: %d, Erwartet: %d", got, want)
	}
}

func TestGetPreviousValue(t *testing.T) {
	tests := []valTest{
		{"Row 1", []int{0, 3, 6, 9, 12, 15}, -3},
		{"Row 2", []int{1, 3, 6, 10, 15, 21}, 0},
		{"Row 3", []int{10, 13, 16, 21, 30, 45}, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getPrevValue(test.input)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
