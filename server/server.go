package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Arimeka/image-transformer"
	"github.com/Arimeka/image-transformer/configuration"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func main() {
	var (
		logger *zap.Logger
		err    error
	)

	if configuration.Environment() == "development" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatal("unable initialize logger", err)
	}

	r := transformer.NewRoutes()

	server := fasthttp.Server{
		Handler:              r.Handler,
		ReadTimeout:          1 * time.Second,
		WriteTimeout:         5 * time.Second,
		MaxKeepaliveDuration: 1 * time.Minute,
		MaxRequestBodySize:   20 * 1024 * 1024,
	}

	go func() {
		err = server.ListenAndServe(configuration.Bind())
		logger.Fatal("Server crash", zap.Error(err))
	}()

	logger.Info("Listening", zap.String("bind", configuration.Bind()))

	sigc := make(chan os.Signal)
	signal.Notify(sigc, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGSTOP)
	<-sigc
	err = server.Shutdown()
	logger.Info("Shutdown", zap.Error(err))
}
