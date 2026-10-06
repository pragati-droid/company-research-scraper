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

	// Remove previously inserted management data
	_, err = db.Exec(`
		DELETE FROM company_management
		WHERE company_id = 1
	`)
	if err != nil {
		log.Fatal("Error clearing old management data:", err)
	}

	management := []struct {
		name     string
		position string
		source   string
	}{
		{
			"Aditya Sanghai",
			"Managing Director",
			"Ghorahi Cement Annual Report FY 2081/82",
		},
		{
			"Devendra Sinha",
			"Executive Director",
			"Ghorahi Cement Annual Report FY 2081/82",
		},
		{
			"Shankar Lal Agiwal",
			"General Manager Finance",
			"Ghorahi Cement Annual Report FY 2081/82",
		},
	}

	for _, person := range management {

		_, err := db.Exec(`
			INSERT INTO company_management
			(company_id, person_name, position, source)
			VALUES ($1, $2, $3, $4)
		`,
			1,
			person.name,
			person.position,
			person.source,
		)

		if err != nil {
			log.Fatal("Error inserting management member:", err)
		}

		fmt.Println("Inserted:", person.name)
	}

	fmt.Println("Management data inserted successfully!")
}