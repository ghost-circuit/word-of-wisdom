package interceptor

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Logging is a gRPC interceptor that logs the details of each RPC call.
func Logging(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	start := time.Now()

	defer func() {
		duration := time.Since(start)

		var reqSize, respSize int
		if protoReq, ok := req.(proto.Message); ok {
			reqSize = proto.Size(protoReq)
		}

		if protoResp, ok := resp.(proto.Message); ok {
			respSize = proto.Size(protoResp)
		}

		code := errorCode(err).String()

		slog.Info("rpc call",
			slog.String("method", info.FullMethod),
			slog.Duration("duration", duration),
			slog.Int("request_size", reqSize),
			slog.Int("response_size", respSize),
			slog.String("code", code),
		)
	}()

	resp, err = handler(ctx, req)

	return
}

func errorCode(err error) codes.Code {
	if err == nil {
		return codes.OK
	}

	errStatus, ok := status.FromError(err)
	if !ok {
		return codes.Unknown
	}

	return errStatus.Code()
}
