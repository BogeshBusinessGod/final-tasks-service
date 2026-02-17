package service

import (
	"context"
	"final/internal/models"
	"final/internal/repository/postgres/sqlc"
	tsk1 "final/pkg/proto/sync/final-boss/v1"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"
)

func (s *service) CreateTask(ctx context.Context, req *tsk1.CreateTaskRequest) (*tsk1.CreateTaskResponse, error) {

	task, err := s.DB.CreateTask(ctx, sqlc.CreateTaskParams{
		Title:   req.GetTitle(),
		Content: req.GetContent(),
		Status:  models.StatusNew,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create task")
	}

	return &tsk1.CreateTaskResponse{
		Task: &tsk1.Task{
			Id:      task.ID,
			Title:   task.Title,
			Content: task.Content,
			Status:  tsk1.Status_STATUS_NEW,
		},
	}, nil
}
