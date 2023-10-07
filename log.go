package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func SetupLog() {
	slog.SetDefault(
		slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})),
	)
}
