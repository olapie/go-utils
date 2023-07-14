package slogctx

import (
	"context"
	"log/slog"
)

type keyType int

const (
	keyLogger keyType = iota
)

func WithLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, keyLogger, l)
}

func GetLogger(ctx context.Context) *slog.Logger {
	l, _ := ctx.Value(keyLogger).(*slog.Logger)
	if l == nil {
		return slog.Default()
	}
	return l
}
