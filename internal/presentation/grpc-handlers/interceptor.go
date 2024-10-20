package grpc_handlers

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
)

// ConvertErrorInterceptor converts errors to gRPC status errors.
func ConvertErrorInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	resp, err = handler(ctx, req)
	if err == nil {
		return
	}

	slog.Warn("convert error interceptor",
		slog.String("method", info.FullMethod),
		slog.Any("error", err),
	)

	err = convertError(err)
	return
}
