package postgres

import (
	"context"
	"errors"
	"final/internal/models"
	"final/internal/repository/postgres/sqlc"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) CreateTask(ctx context.Context, arg sqlc.CreateTaskParams) (*sqlc.CreateTaskRow, error) {
	p.logger.Debug("repo CreateTask", "title", arg.Title, "status", arg.Status)
	res, err := p.queries.CreateTask(ctx, &arg)
	if err != nil {
		p.logger.Error("repo CreateTask failed", err)
		return nil, err
	}
	return res, nil
}

func (p *Postgres) GetTask(ctx context.Context, id int64) (*sqlc.GetTaskRow, error) {
	p.logger.Debug("repo GetTask", "id", id)

	res, err := p.queries.GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			p.logger.Debug("repo GetTask not found", "id", id)
			return nil, err
		}

		p.logger.Error("repo GetTask failed", err, "id", id)
		return nil, err
	}

	return res, nil
}

func (p *Postgres) ListTasks(ctx context.Context) ([]*sqlc.ListTasksRow, error) {
	p.logger.Debug("repo ListTasks")

	res, err := p.queries.ListTasks(ctx)
	if err != nil {
		p.logger.Error("repo ListTasks failed", err)
		return nil, err
	}

	return res, nil
}

func (p *Postgres) DoneTask(ctx context.Context, id int64) error {
	p.logger.Debug("repo DoneTask", "id", id, "status", models.StatusDone)

	rows, err := p.queries.DoneTask(ctx, &sqlc.DoneTaskParams{
		ID:     id,
		Status: models.StatusDone,
	})
	if err != nil {
		p.logger.Error("repo DoneTask failed", err, "id", id)
		return err
	}

	if rows == 0 {
		return ErrTaskNotFound
	}

	return nil
}

var ErrTaskNotFound = errors.New("failed to find the task")

func (p *Postgres) DeleteTask(ctx context.Context, id int64) error {
	rows, err := p.queries.DeleteTask(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrTaskNotFound
	}
	return nil
}
