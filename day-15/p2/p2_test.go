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

	tests := []appTest{
		{"testinput1", app, 145},
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

type labelTest struct {
	name  string
	input string
	label string
	op    string
	flen  int
}

func TestGetLabel(t *testing.T) {
	tests := []labelTest{
		{"rn=1", "rn=1", "rn", "=", 1},
		{"cm-", "cm-", "cm", "-", 0},
		{"qp=3", "qp=3", "qp", "=", 3},
		{"cm=2", "cm=2", "cm", "=", 2},
		{"rn", "rn", "rn", "", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			name, op, flen := getLabelName(test.input)

			if name != test.label || op != test.op || flen != test.flen {
				t.Errorf("TEST: %s: got %v, want %v, got %v, want %v, got %v, want %v", test.name, name, test.label, op, test.op, flen, test.flen)
			}
		})
	}
}
