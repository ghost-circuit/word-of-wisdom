package interceptor

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Recover is a gRPC interceptor that recovers from panics and logs the error details.
func Recover(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (result any, err error) {
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("%v", r)

			slog.Error("recovered from panic",
				slog.String("error", msg),
				slog.String("method", info.FullMethod),
			)

			result = nil
			err = status.Error(codes.Internal, msg)
		}
	}()

	return handler(ctx, req)
}
