package p2

import (
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

func TestP1(t *testing.T) {
	input, err := ioutil.ReadFile("../testinput2")
	if err != nil {
		panic("Can't read file")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := types.App{
		Logger: logger,
		Input:  string(input),
	}

	got := P2(app)
	want := 281

	if got != want {
		t.Errorf("Ergebnis: %d, Erwartet: %d", got, want)
	}
}

func TestMore(t *testing.T) {
	input, err := ioutil.ReadFile("../mytest")
	if err != nil {
		panic("Can't read file")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := types.App{
		Logger: logger,
		Input:  string(input),
	}

	got := P2(app)
	want := 162

	if got != want {
		t.Errorf("Ergebnis: %d, Erwartet: %d", got, want)
	}
}
