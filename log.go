package main

import (
	"fmt"
	"os"

	"gosocket/util"

	"golang.org/x/exp/slog"
)

func SetupLog() {
	handle := new(slog.HandlerOptions)
	handle.AddSource = true
	if util.GlobalConfig.Debug {
		handle.Level = slog.LevelDebug
	}
	fmt.Println(*handle)
	slog.SetDefault(
		slog.New(slog.NewTextHandler(os.Stdout, handle)),
	)
}
