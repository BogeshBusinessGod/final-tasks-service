package repository

import (
	"context"
	"final/internal/repository/postgres/sqlc"
)

type DB interface {
	CreateTask(ctx context.Context, arg sqlc.CreateTaskParams) (*sqlc.Task, error)
	GetTaskByID(ctx context.Context, id int64) (*sqlc.Task, error)
	ListTasks(ctx context.Context, id int64) ([]*sqlc.Task, error)
	DoneTask(ctx context.Context, id int64) error
	DeleteTask(ctx context.Context, id int64) error
	UpdateTaskStatus(ctx context.Context, id int64, status string) (*sqlc.Task, error)
}
