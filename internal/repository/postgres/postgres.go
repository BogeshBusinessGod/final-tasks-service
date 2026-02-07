package postgres

import (
	"context"
	"embed"
	"final/internal/repository/postgres/sqlc"
	log "final/internal/utils/observability"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // 👈 обязательный импорт драйвера
	_ "github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"

	"final/internal/config"

	"final/migrations/migrator"
	pgmigrations "final/migrations/postgres"
)

type Postgres struct {
	connPool *pgxpool.Pool
	queries  *sqlc.Queries
	logger   *log.Logger // 🟢 теперь твой логгер
	cfg      *config.Postgres
}

func NewPostgres(ctx context.Context, logger *log.Logger, cfg *config.Postgres) (*Postgres, error) {
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.SSLMode,
	)

	pool, err := pgxpool.New(ctx, conn)
	if err != nil {
		logger.Error("connection to postgres failed", err)
		return nil, err
	}
	logger.Info("PostgreSQL connection pool initialized", "component", "postgres")

	migrats := []embed.FS{pgmigrations.FS}
	for _, fs := range migrats {
		if err := migrator.DoMigrate(fs, conn); err != nil {
			logger.Error("migration failed", err)
			return nil, err
		}
	}
	logger.Info("migrations applied successfully", "component", "postgres")

	return &Postgres{
		connPool: pool,
		queries:  sqlc.New(pool),
		logger:   logger,
		cfg:      cfg,
	}, nil
}

func (p *Postgres) Close() { p.connPool.Close() }

func (p *Postgres) Q() *sqlc.Queries { return p.queries }
