package repository

import (
	"context"
	"final/internal/repository/postgres/sqlc"
)

type DB interface {
	CreateTask(ctx context.Context, arg sqlc.CreateTaskParams) (*sqlc.CreateTaskRow, error)

	ListTasks(ctx context.Context) ([]*sqlc.ListTasksRow, error)
	GetTask(ctx context.Context, id int64) (*sqlc.GetTaskRow, error)

	DoneTask(ctx context.Context, id int64) (bool, error)
	DeleteTask(ctx context.Context, id int64) (bool, error)
}
