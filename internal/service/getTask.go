package service

import (
	"context"
	"errors"
	"final/internal/conv"
	"final/internal/models"
	tsk1 "final/pkg/proto/sync/final-boss/v1"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetTask(ctx context.Context, req *tsk1.GetTaskRequest) (*tsk1.GetTaskResponse, error) {
	task, err := s.DB.GetTask(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "task not found")
		}
		return nil, status.Error(codes.Internal, "failed to get task")
	}

	return &tsk1.GetTaskResponse{
		Task: &tsk1.Task{
			Id:      task.ID,
			Title:   task.Title,
			Content: task.Content,
			Status:  conv.TaskStatusToProto(task.Status),
			Done:    task.Status == models.StatusDone,
		},
	}, nil
}
