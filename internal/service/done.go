package service

import (
	"context"
	tsk1 "final/pkg/proto/sync/final-boss/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) DoneTask(ctx context.Context, req *tsk1.DoneTaskRequest) (*tsk1.DoneTaskResponse, error) {
	err := s.DB.DoneTask(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to set task status done")
	}

	return &tsk1.DoneTaskResponse{Status: tsk1.Status_STATUS_DONE}, nil
}
