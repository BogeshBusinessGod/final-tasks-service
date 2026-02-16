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

func (p *Postgres) DoneTask(ctx context.Context, id int64) (bool, error) {
	p.logger.Debug("repo DoneTask", "id", id, "status", models.StatusDone)

	rows, err := p.queries.DoneTask(ctx, &sqlc.DoneTaskParams{
		ID:     id,
		Status: models.StatusDone,
	})
	if err != nil {
		p.logger.Error("repo DoneTask failed", err, "id", id)
		return false, err
	}

	return rows > 0, nil
}

func (p *Postgres) DeleteTask(ctx context.Context, id int64) (bool, error) {
	rows, err := p.queries.DeleteTask(ctx, id) // rows int64
	if err != nil {
		return false, err
	}
	return rows > 0, nil
}
