package main

import (
	"database/sql"
	"log"

	"errors"

	server "github.com/youlance/user/api/http"
	db "github.com/youlance/user/db/sqlc"
	"github.com/youlance/user/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	DSN := "postgresql://" + config.DBUser + ":" + config.DBPass + "@" + config.DBAddress + "/" + config.DBName + "?sslmode=disable"

	conn, err := sql.Open(config.DBDriver, DSN)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	if err := migrateUp(conn, "file://migration"); err != nil {
		panic(err)
	}

	store := db.New(conn)
	server, err := server.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func migrateUp(db *sql.DB, migrationsPath string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	}

	return nil
}
