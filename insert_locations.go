package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	password := "12345678"

	connStr := fmt.Sprintf(
		"host=localhost port=5432 user=postgres password=%s dbname=company_research sslmode=disable",
		password,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Connected to PostgreSQL!")

	// Remove previously inserted locations
	_, err = db.Exec(`
		DELETE FROM company_locations
		WHERE company_id = 1
	`)
	if err != nil {
		log.Fatal("Error clearing old location data:", err)
	}

	locations := []struct {
		locationType string
		locationName string
		address      string
		source       string
	}{
		{
			"Registered Office",
			"Ghorahi Cement Industry Limited",
			"Thapathali-11, Kathmandu, Nepal",
			"Ghorahi Cement Annual Report FY 2081/82",
		},
		{
			"Plant Site",
			"Ghorahi Cement Industry Limited",
			"Laxminagar VDC, Dang District, Nepal",
			"Ghorahi Cement Annual Report FY 2081/82",
		},
	}

	for _, location := range locations {

		_, err := db.Exec(`
			INSERT INTO company_locations
			(company_id, location_type, location_name, address, source)
			VALUES ($1, $2, $3, $4, $5)
		`,
			1,
			location.locationType,
			location.locationName,
			location.address,
			location.source,
		)

		if err != nil {
			log.Fatal("Error inserting location:", err)
		}

		fmt.Println("Inserted:", location.locationType)
	}

	fmt.Println("Location data inserted successfully!")
}