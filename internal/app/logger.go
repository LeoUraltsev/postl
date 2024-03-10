package app

import (
	"log/slog"
	"os"
)

const LocalLevel = "local"
const ProdLevel = "prod"

func SetLogger(level string) *slog.Logger {
	var log *slog.Logger
	switch level {
	case LocalLevel:
		log = slog.New(slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		))
	case ProdLevel:
		log = slog.New(slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		))
	}
	return log
}
