package p1

import (
	"os"
	"strings"
	"testing"

	"golang.org/x/exp/slog"
	"paddex.net/aoc/types"
)

type test struct {
	name  string
	input int
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

	got := P1(app)
	want := 35

	if got != want {
		t.Errorf("Ergebnis: %d, Erwartet: %d", got, want)
	}
}

func TestGetSoil(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	parts := strings.Split(string(input), "\n\n")
	seedToSoil := getMap(parts[1])

	tests := []test{
		{"Seed 79", 79, 81},
		{"Seed 14", 14, 14},
		{"Seed 55", 55, 57},
		{"Seed 13", 13, 13},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getSoil(test.input, seedToSoil)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestGetFertilizer(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	parts := strings.Split(string(input), "\n\n")
	soilToFertilizer := getMap(parts[2])

	tests := []test{
		{"Seed 79", 81, 81},
		{"Seed 14", 14, 53},
		{"Seed 55", 57, 57},
		{"Seed 13", 13, 52},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getFertilizer(test.input, soilToFertilizer)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestGetWater(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	parts := strings.Split(string(input), "\n\n")
	fertilizerToWater := getMap(parts[3])

	tests := []test{
		{"Seed 79", 81, 81},
		{"Seed 14", 53, 49},
		{"Seed 55", 57, 53},
		{"Seed 13", 52, 41},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getWater(test.input, fertilizerToWater)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestGetLight(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	parts := strings.Split(string(input), "\n\n")
	waterToLight := getMap(parts[4])

	tests := []test{
		{"Seed 79", 81, 74},
		{"Seed 14", 49, 42},
		{"Seed 55", 53, 46},
		{"Seed 13", 41, 34},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getLight(test.input, waterToLight)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestGetTemperature(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	parts := strings.Split(string(input), "\n\n")
	lightToTemperature := getMap(parts[5])

	tests := []test{
		{"Seed 79", 74, 78},
		{"Seed 14", 42, 42},
		{"Seed 55", 46, 82},
		{"Seed 13", 34, 34},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getTemperature(test.input, lightToTemperature)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestGetHumidity(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	parts := strings.Split(string(input), "\n\n")
	temperatureToHumidity := getMap(parts[6])

	tests := []test{
		{"Seed 79", 78, 78},
		{"Seed 14", 42, 43},
		{"Seed 55", 82, 82},
		{"Seed 13", 34, 35},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getHumidity(test.input, temperatureToHumidity)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestGetLocation(t *testing.T) {
	input, err := os.ReadFile("../testinput1")
	if err != nil {
		panic("Can't read file")
	}

	parts := strings.Split(string(input), "\n\n")
	humidityToLocation := getMap(parts[7])

	tests := []test{
		{"Seed 79", 78, 82},
		{"Seed 14", 43, 43},
		{"Seed 55", 82, 86},
		{"Seed 13", 35, 35},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getLocation(test.input, humidityToLocation)

			if got != test.want {
				t.Errorf("TEST: %s: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
