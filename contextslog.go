package contextslog

import (
	"context"
	"log/slog"
)

var (
	contextType struct{}
	ContextVal  = contextType
)

type Handler struct {
	h slog.Handler

	attrs []slog.Attr
}

func NewContextHandler(h slog.Handler) *Handler {
	return &Handler{
		h: h,

		attrs: make([]slog.Attr, 0),
	}
}

func (h *Handler) Enabled(ctx context.Context, l slog.Level) bool {
	return h.h.Enabled(ctx, l)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandler := h.clone()
	newHandler.addAttrs(attrs)

	return newHandler
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return h.h.WithGroup(name)
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {

	if l := GetFromContext(ctx); l != nil {
		// There was a logger in the context
		if h, ok := l.Handler().(interface{ getAttrs() []slog.Attr }); ok {
			r.AddAttrs(h.getAttrs()...)
		}
	}

	return h.h.Handle(ctx, r)
}

// Some functions we need
func (h *Handler) clone() *Handler {
	return &Handler{
		h:     h.h,
		attrs: h.attrs,
	}
}

func (h *Handler) getAttrs() []slog.Attr {
	return h.attrs
}

func (h *Handler) addAttrs(attrs []slog.Attr) {
	h.attrs = append(h.attrs, attrs...)
}
