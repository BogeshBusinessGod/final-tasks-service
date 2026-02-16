package service

import (
	"context"
	"final/internal/repository"
	log "final/internal/utils/observability"
	tsk1 "final/pkg/proto/sync/final-boss/v1"
)

type Service interface {
	CreateTask(ctx context.Context, req *tsk1.CreateTaskRequest) (*tsk1.CreateTaskResponse, error)
	ListTasks(ctx context.Context, req *tsk1.ListTasksRequest) (*tsk1.ListTasksResponse, error)
	GetTask(ctx context.Context, req *tsk1.GetTaskRequest) (*tsk1.GetTaskResponse, error)
	DeleteTask(ctx context.Context, req *tsk1.DeleteTaskRequest) (*tsk1.DeleteTaskResponse, error)
	DoneTask(ctx context.Context, req *tsk1.DoneTaskRequest) (*tsk1.DoneTaskResponse, error)
}

type service struct {
	logger *log.Logger
	DB     repository.DB
}

func NewService(DB repository.DB) Service {
	return &service{DB: DB}
}
