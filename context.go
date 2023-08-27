package contextslog

import (
	"context"
	"log/slog"
)

var (
	contextType struct{}
	contextVal  = contextType
)

func AddToContext(ctx context.Context, logger *slog.Logger, attrs ...any) context.Context {
	return context.WithValue(ctx, contextVal, logger.With(attrs...))
}

func GetFromContext(ctx context.Context) *slog.Logger {
	h := ctx.Value(contextVal)

	if h == nil {
		return slog.Default()
	}

	if l, ok := h.(*slog.Logger); ok {
		return l
	}

	slog.InfoContext(ctx, "type in context is %T not *slog.Logger", h)
	return nil
}
