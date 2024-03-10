package app

import (
	"log/slog"
	"postl/internal/config"
)

func Run(configPath string) {
	cfg := config.NewConfig(configPath)

	log := SetLogger(cfg.Env)

	log.Info(
		"starting application",
		slog.String("env", cfg.Env),
	)
}
