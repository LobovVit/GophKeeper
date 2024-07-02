// Package postgresql - included functions for init SQL connections
package postgresql

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewConn - function for opening new SQL connection
func NewConn(ctx context.Context, dsn string) (*sql.DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("DB open: %w", err)
	}
	err = conn.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}
	err = Migrate(conn)
	if err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}
	return conn, nil
}

//go:embed migrations/*.sql
var fs embed.FS

// Migrate - function for apply migration scripts
func Migrate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("migrate driver: %w", err)
	}
	d, err := iofs.New(fs, "migrations") // Get migrations from folder
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		return fmt.Errorf("migrate new: %w", err)
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrate up: %w", err)
	}
	return nil
}
