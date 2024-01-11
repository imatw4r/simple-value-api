package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

/*
Setup graceful exit for application context, incase SIGKILL and SIGINT
*/
func setupTerminationAtSignals(ctx context.Context, cancel context.CancelFunc) {
	shutDownSignals := make(chan os.Signal, 1)
	signal.Notify(shutDownSignals, syscall.SIGINT, syscall.SIGTERM)

	// Listen for SIGINT and SIGKILL and exit gracefully by cancelling parent context
	go func() {
		sig := <-shutDownSignals
		log.Warnf("Received SIGNAL: %s", sig)
		log.Warn("Gracefully terminating application...")
		cancel()
		defer close(shutDownSignals)
	}()
}
