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
		{"testinput1", app, 1320},
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

type hashTest struct {
	name  string
	input string
	want  int
}

func TestHashString(t *testing.T) {
	tests := []hashTest{
		{"rn=1", "rn=1", 30},
		{"cm-", "cm-", 253},
		{"qp=3", "qp=3", 97},
		{"cm=2", "cm=2", 47},
		{"rn", "rn", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := hashString(test.input)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
