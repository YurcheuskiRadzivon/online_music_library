package migrator

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type Migrator struct {
	srcDriver source.Driver
}

func NewMigrator(sqlFiles embed.FS, dirName string) (*Migrator, error) {
	driver, err := iofs.New(sqlFiles, dirName)
	if err != nil {
		return nil, fmt.Errorf("error creating source driver: %v", err)
	}
	/*
		// Добавьте логирование для отладки
		files, err := sqlFiles.ReadDir(dirName)
		if err != nil {
			return nil, fmt.Errorf("error reading migration directory: %v", err)
		}
		log.Println("Migration files found:")
		for _, file := range files {
			log.Println(file.Name())
		}
	*/
	return &Migrator{
		srcDriver: driver,
	}, nil
}

func (m *Migrator) ApplyMigrations(db *sql.DB, lgr *logger.Logger) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("unable to create db instance: %v", err)
	}
	lgr.InfoLogger.Println("Db instance has created")

	migrator, err := migrate.NewWithInstance("iofs", m.srcDriver, "postgres", driver)
	if err != nil {
		return fmt.Errorf("unable to create migration instance: %v", err)
	}
	lgr.InfoLogger.Println("Migrator has created")

	defer func() {
		migrator.Close()
	}()

	lgr.DebugLogger.Println("Applying migrations...")

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("unable to apply migration: %v", err)
	}

	lgr.InfoLogger.Println("Migrations applied successfully")
	return nil
}
func (m *Migrator) RollbackMigrations(db *sql.DB, lgr *logger.Logger) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("unable to create db instance: %v", err)
	}
	lgr.InfoLogger.Println("Db instance has created")

	migrator, err := migrate.NewWithInstance("iofs", m.srcDriver, "postgres", driver)
	if err != nil {
		return fmt.Errorf("unable to create migration instance: %v", err)
	}
	lgr.InfoLogger.Println("Migrator has created")
	defer func() {
		migrator.Close()
	}()

	lgr.DebugLogger.Println("Applying migrations...")

	if err := migrator.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("unable to apply migration: %v", err)
	}

	lgr.InfoLogger.Println("Migrations applied successfully")
	return nil
}
