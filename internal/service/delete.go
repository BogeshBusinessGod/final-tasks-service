package service

import (
	"context"
	tsk1 "final/pkg/proto/sync/final-boss/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) DeleteTask(ctx context.Context, req *tsk1.DeleteTaskRequest) (*tsk1.DeleteTaskResponse, error) {
	deleted, err := s.DB.DeleteTask(ctx, req.GetId())
	if err != nil {
		return &tsk1.DeleteTaskResponse{Success: false},
			status.Error(codes.Internal, "failed to delete task")
	}
	if !deleted {
		return &tsk1.DeleteTaskResponse{Success: false},
			status.Error(codes.NotFound, "task not found")
	}
	return &tsk1.DeleteTaskResponse{Success: true}, nil
}
