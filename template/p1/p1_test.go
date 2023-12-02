package p1

import (
	"os"
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
	want := 142

	if got != want {
		t.Errorf("Ergebnis: %d, Erwartet: %d", got, want)
	}
}
