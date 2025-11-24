package migrations

import (
	"crudl/internal/adapters/postgres"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)


func RunMigrate(path string, cfg postgres.Config) error {
	key := cfg.Source
	if key == "" {
		return fmt.Errorf("migration database connection string is empty")
	}
	
	mig, err := migrate.New(fmt.Sprintf("file://%s", path), key)
	if err != nil {
		return fmt.Errorf("error to run migration: %w", err)
	}

	if err = mig.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error to up migration: %w", err)
	}

	return nil
}
