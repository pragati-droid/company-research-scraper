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

	// Remove previously inserted board data for Ghorahi Cement
	_, err = db.Exec(`
		DELETE FROM company_board
		WHERE company_id = 1
	`)
	if err != nil {
		log.Fatal("Error clearing old board data:", err)
	}

	// Insert board members
	boardMembers := []struct {
		name     string
		position string
		source   string
	}{
		{
			"Purushottam Lal Sanghai",
			"Chairman & Director",
			"Ghorahi Cement Annual Report FY 2081/82",
		},
		{
			"Trilokchand Agrawal",
			"Director",
			"Ghorahi Cement Annual Report FY 2081/82",
		},
		{
			"Aditya Sanghai",
			"Managing Director",
			"Ghorahi Cement Annual Report FY 2081/82",
		},
		{
			"Anju Sanghai",
			"Director",
			"Ghorahi Cement Annual Report FY 2081/82",
		},
		{
			"Vivek Agrawal",
			"Director",
			"Ghorahi Cement Annual Report FY 2081/82",
		},
		{
			"Bimal Prasad Dhakal",
			"Independent Director",
			"Ghorahi Cement Annual Report FY 2081/82",
		},
		{
			"Milan Sigdel",
			"Public Director",
			"Ghorahi Cement Annual Report FY 2081/82",
		},
	}

	for _, member := range boardMembers {

		_, err := db.Exec(`
			INSERT INTO company_board
			(company_id, person_name, position, source)
			VALUES ($1, $2, $3, $4)
		`,
			1,
			member.name,
			member.position,
			member.source,
		)

		if err != nil {
			log.Fatal("Error inserting board member:", err)
		}

		fmt.Println("Inserted:", member.name)
	}

	fmt.Println("Board data inserted successfully!")
}