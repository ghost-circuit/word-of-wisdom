package interceptor

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

var (
	// loggerOpts are the options for the loggerFunc interceptor.
	loggerOpts = []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}

	// loggerFunc is the loggerFunc for the loggerFunc interceptor.
	loggerFunc = logging.LoggerFunc(
		func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
			slog.Log(ctx, slog.Level(lvl), msg, fields...)
		},
	)

	// Logger is the logger interceptor.
	Logger = logging.UnaryServerInterceptor(loggerFunc, loggerOpts...)
)
