package service

import (
	"context"
	tsk1 "final/pkg/proto/sync/final-boss/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) DoneTask(ctx context.Context, req *tsk1.DoneTaskRequest) (*tsk1.DoneTaskResponse, error) {
	done, err := s.DB.DoneTask(ctx, req.GetId())
	if err != nil {
		return &tsk1.DoneTaskResponse{Success: false},
			status.Error(codes.Internal, "failed to mark task as done")
	}
	if !done {
		return &tsk1.DoneTaskResponse{Success: false},
			status.Error(codes.NotFound, "failed to find the task")
	}
	return &tsk1.DoneTaskResponse{Success: true}, nil

}
