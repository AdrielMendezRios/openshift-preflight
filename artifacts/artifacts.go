// Package artifacts provides functionality for writing artifact files in configured
// artifacts directory. This package operators with a singleton directory variable that can be
// changed and reset. It provides simple functionality that can be accessible from
// any calling library.
package artifacts

import (
	"context"
	"io"

	"github.com/sirupsen/logrus"
)

const DefaultArtifactsDir = "artifacts"

// ContextWithWriter adds ArtifactWriter w to the context ctx.
func ContextWithWriter(ctx context.Context, w ArtifactWriter) context.Context {
	return context.WithValue(ctx, artifactWriterContextKey, w)
}

// WriterFromContext returns the writer from the context, or nil.
func WriterFromContext(ctx context.Context) ArtifactWriter {
	w := ctx.Value(artifactWriterContextKey)
	if writer, ok := w.(ArtifactWriter); ok {
		return writer
	}

	return nil
}

// contextKey is a key used to store/retrieve ArtifactsWriter in/from context.Context.
type contextKey string

const artifactWriterContextKey contextKey = "ArtifactWriter"

// ArtifactWriter is the functionality required by all implementations.
type ArtifactWriter interface {
	WriteFile(filename string, contents io.Reader) (fullpathToFile string, err error)
}

// logrusContextKey is a key used to store/retrieve logrus.Logger in/from context.Context.
type logrusContextKey string

const logrusLoggerContextKey logrusContextKey = "LogrusLogger"

// ContextWithLogrusLogger adds logrus.Logger l to the context ctx.
func ContextWithLogrusLogger(ctx context.Context, l *logrus.Logger) context.Context {
	return context.WithValue(ctx, logrusLoggerContextKey, l)
}

// LogrusLoggerFromContext returns the logrus.Logger from the context, or nil.
func LogrusLoggerFromContext(ctx context.Context) *logrus.Logger {
	l := ctx.Value(logrusLoggerContextKey)
	if logger, ok := l.(*logrus.Logger); ok {
		return logger
	}
	return nil
}
