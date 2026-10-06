package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	password := os.Getenv("PGPASSWORD")

	connStr := fmt.Sprintf(
		"host=localhost port=5432 user=postgres password=%s dbname=company_research sslmode=disable",
		password,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}

	fmt.Println("Connected to PostgreSQL!")

	// Insert Ghorahi Cement profile
	_, err = db.Exec(`
		INSERT INTO company_profiles (
			company_id,
			company_name,
			industry,
			country,
			location,
			established_year,
			products_services,
			website
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`,
		1,
		"Ghorahi Cement Industry Limited",
		"Cement Manufacturing",
		"Nepal",
		"Ghorahi, Dang, Nepal",
		"2009",
		"Cement and clinker manufacturing",
		"",
	)

	if err != nil {
		fmt.Println("Error inserting company profile:", err)
		return
	}

	fmt.Println("Ghorahi Cement profile inserted successfully!")
}