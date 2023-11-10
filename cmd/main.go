package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"ipfsupload-go/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initial Application initialization
	logger := zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()

	// Try to load env values
	err := godotenv.Load()
	if err != nil {
		logger.Fatal().Msg("Error loading .env file")
	}
	apiInit := api.NewApi("/v1", logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	addr := fmt.Sprintf(":%d", 3000)
	srv := &http.Server{
		Addr:    addr,
		Handler: apiInit.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Err(err).Msg("listen")
		}
	}()
	<-ctx.Done()

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Msg(fmt.Sprintf("Server forced to shutdown: %s", err))
	}
	logger.Info().Msg("exiting server")

	// run the reset of the application and set up the routes
}
