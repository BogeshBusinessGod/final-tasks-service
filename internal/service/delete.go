package service

import (
	"context"
	tsk1 "final/pkg/proto/sync/final-boss/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) DeleteTask(ctx context.Context, req *tsk1.DeleteTaskRequest) (*tsk1.DeleteTaskResponse, error) {
	err := s.DB.DeleteTask(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete task")
	}

	return &tsk1.DeleteTaskResponse{}, nil
}
