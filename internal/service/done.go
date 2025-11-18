package service

import (
	"context"
	tsk1 "final/pkg/proto/sync/final/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) DoneTask(ctx context.Context, req *tsk1.DoneTaskRequest) (*tsk1.DoneTaskResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	task, err := s.DB.UpdateTaskStatus(ctx, req.GetId(), "done")
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to mark task as done")
	}

	return &tsk1.DoneTaskResponse{
		Task: &tsk1.Task{
			Id: task.ID,

			Title:   task.Title,
			Content: task.Content.String,
			Status:  task.Status,
			Done:    task.Done,
		},
	}, nil
}
