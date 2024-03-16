package app

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"postl/internal/config"
	v1 "postl/internal/controller/http/v1"
)

func Run(configPath string) {
	cfg := config.NewConfig(configPath)

	log := SetLogger(cfg.Env)

	log.Info(
		"starting application",
		slog.String("env", cfg.Env),
	)
	h := echo.New()
	v1.New(h)

	srv := &http.Server{
		Addr:         cfg.HttpServer.Address,
		Handler:      h,
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
	}

	log.Info(
		"starting server",
		slog.String(
			"address",
			cfg.HttpServer.Address,
		),
	)

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
}
