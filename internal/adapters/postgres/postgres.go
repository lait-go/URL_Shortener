package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     string `envconfig:"POSTGRES_PORT" required:"true"`
	DB       string `envconfig:"POSTGRES_DB" required:"true"`
}

type Pool struct {
	DB *sqlx.DB
}

func (c *Config) DbKeyInit() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.DB)
}

func New(ctx context.Context, cfg Config) (*Pool, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", cfg.DbKeyInit())
	if err != nil {
		return nil, fmt.Errorf("error connection to bd: %w", err)
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return &Pool{DB: db}, nil
}

func (p *Pool) Close() error {
	return p.DB.Close()
}
