package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Yury132/Golang-Task-1/internal/client/google"
	"github.com/Yury132/Golang-Task-1/internal/config"
	"github.com/Yury132/Golang-Task-1/internal/service"
	"github.com/Yury132/Golang-Task-1/internal/storage"
	transport "github.com/Yury132/Golang-Task-1/internal/transport/http"
	"github.com/Yury132/Golang-Task-1/internal/transport/http/handlers"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		panic(err)
	}

	logger := cfg.Logger()

	poolCfg, err := cfg.PgPoolConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to DB")
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to db")
	}

	oauthCfg := cfg.SetupConfig()

	googleAPI := google.New(logger)

	strg := storage.New(conn)
	svc := service.New(logger, oauthCfg, googleAPI, strg)
	handler := handlers.New(logger, oauthCfg, svc)
	srv := transport.New(":8080").WithHandler(handler)

	// graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	go func() {
		if err = srv.Run(); err != nil {
			logger.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	<-shutdown
}
