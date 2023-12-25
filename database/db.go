package database

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDBConnection() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5432/chatdb?sslmode=disable")
	if err != nil {
		log.Errorf("database connection open error: %v", err)
		return nil, err
	}

	err = migrateDB(db)
	if err != nil {
		log.Errorf("migration error: %v", err)
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

func migrateDB(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Errorf("get driver with instance error: %v", err)
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "postgres", driver)
	if err != nil {
		log.Errorf("new database instance error: %v", err)
		return err
	}
	// if err = m.Down(); err != nil && errors.Is(err, migrate.ErrNoChange){
	// 	log.Errorf("migration down error: %v", err)
	// 	return err
	// }
	err = m.Up()
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Infof("migration up error: %v", err)
		return nil
	}
	return err
}
