package database

import (
	"database/sql"
	"fmt"
	"hafiztri123/be-sharing-vision/internal/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewDatabase() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		utils.GetMandatoryEnv("DB_USER"),
		utils.GetMandatoryEnv("DB_PWD"),
		utils.GetMandatoryEnv("DB_HOST"),
		utils.GetMandatoryEnv("DB_HOST_PORT"),
		utils.GetMandatoryEnv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	err = runMigrations(db)
	if err != nil {
		panic(err)
	}

	return db

}

func runMigrations(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("could not create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		utils.GetMandatoryEnv("DB_NAME"),
		driver,
	)

	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("could not run migrations: %w", err)
	}

	fmt.Println("Database migrated")

	return nil
}
