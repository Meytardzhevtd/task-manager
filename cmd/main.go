package main

import (
	"database/sql"
	"fmt"
	"log"
	"task-manager/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	c := config.Load()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database!")
}
