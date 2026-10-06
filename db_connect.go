package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	password := os.Getenv("DB_PASSWORD")

	if password == "" {
		fmt.Println("Please set DB_PASSWORD environment variable.")
		return
	}

	connStr := fmt.Sprintf(
		"host=localhost port=5432 user=postgres password=%s dbname=company_research sslmode=disable",
		password,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}

	fmt.Println("Successfully connected to company_research!")
}