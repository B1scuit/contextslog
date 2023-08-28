package contextslog

import (
	"context"
	"log/slog"
	"sync"
)

type contextType struct{ s string }

var (
	contextVal = contextType{"context"}
	mutexVal   = contextType{"mutex"}
)

func AddToContext(ctx context.Context, logger *slog.Logger, attrs ...any) context.Context {
	m, ok := ctx.Value(mutexVal).(*sync.RWMutex)
	if !ok {
		m = &sync.RWMutex{}
		ctx = context.WithValue(ctx, mutexVal, m)
	}

	m.Lock()
	defer m.Unlock()

	return context.WithValue(ctx, contextVal, logger.With(attrs...))
}

func GetFromContext(ctx context.Context) *slog.Logger {
	m, mok := ctx.Value(mutexVal).(*sync.RWMutex)
	h, hok := ctx.Value(contextVal).(*slog.Logger)

	if !hok || !mok {
		ctx = AddToContext(ctx, slog.Default())
		return GetFromContext(ctx)
	}

	m.RLock()
	defer m.RUnlock()

	return h
}
