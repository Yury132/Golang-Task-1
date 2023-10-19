package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Yury132/Golang-Task-1/internal/config"
	transport "github.com/Yury132/Golang-Task-1/internal/transport/http"
	"github.com/Yury132/Golang-Task-1/internal/transport/http/handlers"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		panic(err)
	}

	logger := cfg.Logger()

	//pool, err := cfg.PgPoolConfig()
	//if err != nil {
	//	logger.Fatal().Err(err).Msg("failed to connect to DB")
	//}

	r := handlers.New(logger)
	srv := transport.New("127.0.0.1:8000").WithHandler(r)

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
