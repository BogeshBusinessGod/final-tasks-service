package interceptors

import (
	"context"

	log "final/internal/utils/observability"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func WithValidation(logger *log.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		protoMsg, ok := req.(proto.Message)
		if !ok {
			logger.Warn("not a proto message", "method", info.FullMethod)
			return nil, status.Errorf(codes.InvalidArgument, "not valid proto message")
		}

		// Если тип реализует Validate(), вызываем его
		if v, ok := protoMsg.(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				logger.Warn("validation failed",
					"method", info.FullMethod,
					"err", err.Error(),
				)
				return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
			}
		}

		// Всё ок — продолжаем
		return handler(ctx, req)
	}
}
