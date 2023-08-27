package contextslog

import (
	"context"
	"io"
	"log/slog"
	"os"
	"testing"
)

var happyPath *slog.Logger

func TestMain(m *testing.M) {
	happyPath = slog.New(NewContextHandler(
		slog.NewTextHandler(os.Stdout, nil),
	))

	m.Run()
}

func TestMessage(t *testing.T) {
	var msg []byte = []byte("ExampleMessage")

	l := slog.New(NewContextHandler(
		slog.NewTextHandler(io.Discard, nil),
	))
	l.Info(string(msg))
}

func TestBasicContext(t *testing.T) {
	ctx := AddToContext(context.Background(), happyPath, slog.String("foo", "bar"))

	happyPath.InfoContext(ctx, "Should include values")
}
