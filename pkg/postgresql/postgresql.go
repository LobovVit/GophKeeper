// Package postgresql - included functions for init SQL connections
package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewConn - function for opening new SQL connection
func NewConn(ctx context.Context, dsn, mgrDir string) (*sql.DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("DB open: %w", err)
	}
	err = conn.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}
	err = Migrate(mgrDir, conn)
	if err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}
	return conn, nil
}

// Migrate - function for apply migration scripts
func Migrate(migrationDir string, db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("migrate driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:%s", migrationDir),
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("migrate new: %w", err)
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrate up: %w", err)
	}
	return nil
}
