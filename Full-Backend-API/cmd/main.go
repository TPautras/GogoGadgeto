package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"gitub.com/TPautras/ecom/cmd/api"
	"gitub.com/TPautras/ecom/config"
	"gitub.com/TPautras/ecom/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAdress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("Connected to database successfully")
}
