package service

import (
	"context"
	tsk1 "final/pkg/proto/sync/final-boss/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) DeleteTask(ctx context.Context, req *tsk1.DeleteTaskRequest) (*tsk1.DeleteTaskResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err := s.DB.UpdateTaskStatus(ctx, req.GetId(), "deleted")
	if err != nil {
		return &tsk1.DeleteTaskResponse{
			Success: false,
		}, status.Error(codes.Internal, "failed to delete task")
	}

	return &tsk1.DeleteTaskResponse{
		Success: true,
	}, nil
}
