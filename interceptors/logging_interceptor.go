package interceptors

import (
	"context"
	"time"

	log "final/internal/utils/observability"

	"google.golang.org/grpc"
)

func LoggingInterceptor(logger *log.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()

		resp, err := handler(ctx, req)
		duration := time.Since(start)

		if err != nil {
			logger.Error("gRPC request failed",
				"method", info.FullMethod,
				"duration_ms", duration.Milliseconds(),
				"error", err,
			)
		} else {
			logger.Info("gRPC request succeeded",
				"method", info.FullMethod,
				"duration_ms", duration.Milliseconds(),
			)
		}

		return resp, err
	}
}
