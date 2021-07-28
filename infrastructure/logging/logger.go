package logging

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	loggerContextKey = "logger"
)

func New() *logrus.Entry {
	f := new(logrus.TextFormatter)
	f.FullTimestamp = true
	f.TimestampFormat = "02.01.2006 15:04:05"

	logger := logrus.New()
	logger.SetFormatter(f)

	return logrus.NewEntry(logger)
}

// WithLogger wraps the context with the logger.
func WithLogger(ctx context.Context, l *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerContextKey, l)
}

// WithLoggerMiddleware inject the logger into the request context.
func WithLoggerMiddleware(logger *logrus.Entry) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			loggerCtx := WithLogger(r.Context(), logger)
			r = r.WithContext(loggerCtx)

			next.ServeHTTP(w, r)
		})
	}
}

// MustLoggerFromContext gets the logger from context.
func MustLoggerFromContext(ctx context.Context) *logrus.Entry {
	l := ctx.Value(loggerContextKey)
	if l == nil {
		panic("logger is not in context")
	}

	return l.(*logrus.Entry)
}
