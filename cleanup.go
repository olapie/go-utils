package utils

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func CleanUp(f func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		received := <-c
		slog.Info(received.String())
		f()
	}()
}
