package service

import (
	"context"
	"final/internal/repository/postgres/sqlc"
	tsk1 "final/pkg/proto/sync/final-boss/v1"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) CreateTask(ctx context.Context, req *tsk1.CreateTaskRequest) (*tsk1.CreateTaskResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	task, err := s.DB.CreateTask(ctx, sqlc.CreateTaskParams{
		Title:   req.GetTitle(),
		Content: pgtype.Text{String: req.GetContent(), Valid: true},
		Status:  "new",
		Done:    false,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create task")
	}

	return &tsk1.CreateTaskResponse{
		Task: &tsk1.Task{
			Id:      task.ID,
			Title:   task.Title,
			Content: task.Content.String,
			Status:  task.Status,
			Done:    task.Done,
		},
	}, nil
}
