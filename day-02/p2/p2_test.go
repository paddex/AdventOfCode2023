package p2

import (
	"os"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

func TestP2(t *testing.T) {
	input, err := os.ReadFile("../testinput2")
	if err != nil {
		panic("Can't read file")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := types.App{
		Logger: logger,
		Input:  string(input),
	}

	got := P2(app)
	want := 2286

	if got != want {
		t.Errorf("Ergebnis: %d, Erwartet: %d", got, want)
	}
}
