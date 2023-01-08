package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sicepat"
	"sicepat/internal/config"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	shutDownTimeout = 10 * time.Second
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	cfg := config.NewConfig()
	server, err := sicepat.New(cfg)
	if err != nil {
		logrus.WithField("error", err.Error()).Error("failed to initialize server")
		return
	}

	go server.Start()

	// wait until killed
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// wait for OS signal
	sig := <-signals

	logrus.WithField("signal", fmt.Sprintf("%#v", sig)).Warn("server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), shutDownTimeout)
	defer cancel()

	server.Shutdown(ctx)
}
