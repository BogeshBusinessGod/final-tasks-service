package postgres

import (
	"context"
	"final/internal/repository/postgres/sqlc"
)

func (p *Postgres) CreateTask(ctx context.Context, arg sqlc.CreateTaskParams) (*sqlc.Task, error) {
	p.logger.Debug("CreateTask called", "title", arg.Title)
	return p.queries.CreateTask(ctx, &arg)
}

func (p *Postgres) GetTaskByID(ctx context.Context, id int64) (*sqlc.Task, error) {
	return p.queries.GetTaskByID(ctx, id)
}

func (p *Postgres) ListTasks(ctx context.Context, id int64) ([]*sqlc.Task, error) {
	if id == 0 {
		return p.queries.ListTasks(ctx)
	}
	task, err := p.queries.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return []*sqlc.Task{task}, nil
}

func (p *Postgres) DoneTask(ctx context.Context, id int64) error {
	return p.queries.DoneTask(ctx, id)
}

func (p *Postgres) DeleteTask(ctx context.Context, id int64) error {
	return p.queries.DeleteTask(ctx, id)
}
func (p *Postgres) UpdateTaskStatus(ctx context.Context, id int64, status string) (*sqlc.Task, error) {
	task, err := p.queries.UpdateTaskStatus(ctx, &sqlc.UpdateTaskStatusParams{
		ID:     id,
		Status: status,
	})
	if err != nil {
		return nil, err
	}
	return task, nil
}
