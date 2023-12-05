package types

import "golang.org/x/exp/slog"

type App struct {
	Logger *slog.Logger
	Input  string
}
