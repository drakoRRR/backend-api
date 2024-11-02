package main

import (
	"backend-api/cmd/api"
	env "backend-api/config"
	"backend-api/db"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 env.Envs.DBUser,
		Passwd:               env.Envs.DBPassword,
		Addr:                 env.Envs.DBAddress,
		DBName:               env.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	server := api.NewApiServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database")
}
