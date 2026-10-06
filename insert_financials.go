package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Financial struct {
	FiscalYear       string
	Revenue          int64
	GrossProfit      int64
	ProfitOperations int64
	NetProfitOrLoss  int64
	EPS              float64
}

func main() {

	password := "12345678"

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

	data := []Financial{
		{
			FiscalYear:       "2081/82",
			Revenue:          5052530375,
			GrossProfit:      1017267378,
			ProfitOperations: 133106904,
			NetProfitOrLoss:  -610378179,
			EPS:              -13.36,
		},
		{
			FiscalYear:       "2080/81",
			Revenue:          4477970279,
			GrossProfit:      308529938,
			ProfitOperations: -474378742,
			NetProfitOrLoss:  -1447080360,
			EPS:              -31.68,
		},
	}

	for _, year := range data {

		_, err := db.Exec(`
			INSERT INTO financials (
				company_id,
				fiscal_year,
				revenue_from_operations,
				gross_profit,
				profit_from_operations,
				net_profit_or_loss,
				earnings_per_share
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			ON CONFLICT (company_id, fiscal_year) DO NOTHING
		`,
			1,
			year.FiscalYear,
			year.Revenue,
			year.GrossProfit,
			year.ProfitOperations,
			year.NetProfitOrLoss,
			year.EPS,
		)

		if err != nil {
			fmt.Println("Error inserting", year.FiscalYear, ":", err)
			return
		}

		fmt.Println("Inserted:", year.FiscalYear)
	}

	fmt.Println("Financial data inserted successfully!")
}