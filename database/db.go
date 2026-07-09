package database

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect(host, port, user, password, databaseName string) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		databaseName,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return err
	}

	DB = db
	return migrateUp(DB)
}

func ShutdownDatabase() error {
	return DB.Close()
}

func migrateUp(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {

		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "postgres", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
