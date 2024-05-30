package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/shubhamku044/ecomm/cmd/api"
	"github.com/shubhamku044/ecomm/config"
	"github.com/shubhamku044/ecomm/db"
)

func main() {
	db, err := db.NewMYSQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPass,
		Addr:                 config.Envs.DBAddr,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatalf("error while connecting to db: %v", err)
	}

	initStorage(db)

	server := api.NewAPIServer("localhost:8080", db)

	if err := server.Run(); err != nil {
		log.Fatalf("error while running server: %v", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("error while pinging db: %v", err)
	}
	log.Println("Connected to db")
}
