package main

import (
	"database/sql"
	"fmt"
	"log"

	server "github.com/youlance/user/api/http"
	db "github.com/youlance/user/db/sqlc"
	"github.com/youlance/user/pkg/config"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	DSN := "postgresql://" + config.DBUser + ":" + config.DBPass + "@" + config.DBAddress + "/" + config.DBName + "?sslmode=disable"
	fmt.Print(DSN)

	conn, err := sql.Open(config.DBDriver, DSN)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
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
