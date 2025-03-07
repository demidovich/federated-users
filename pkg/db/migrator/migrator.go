package migrator

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jmoiron/sqlx"
)

const migrationsDir = "database/migrations"

type migrator struct {
	db        *sqlx.DB
	goMigrate *migrate.Migrate
}

type migrationStatus struct {
	Version uint
	Dirty   bool
	Error   error
}

func NewOrFail(db *sqlx.DB) *migrator {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	migration, err := migrate.NewWithDatabaseInstance("file://"+migrationsDir, "postgres", driver)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return &migrator{
		db:        db,
		goMigrate: migration,
	}
}

func (m *migrator) Up() error {
	tx := m.db.MustBegin()

	err := m.goMigrate.Up()
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return err
}

func (m *migrator) Down() error {
	tx := m.db.MustBegin()

	err := m.goMigrate.Down()
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return err
}
