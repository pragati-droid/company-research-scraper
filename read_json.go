
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type FinancialYear struct {
    FiscalYear       string  `json:"fiscal_year"`
    Revenue          int64   `json:"revenue_from_operations"`
    GrossProfit      int64   `json:"gross_profit"`
    ProfitOperations int64   `json:"profit_from_operations"`
    NetProfitOrLoss  int64   `json:"net_profit_or_loss"`
    EPS              float64 `json:"earnings_per_share"`
}

type CompanyData struct {
    Company    string          `json:"company"`
    Currency   string          `json:"currency"`
    Financials []FinancialYear `json:"financials"`
}

func main() {
    file, err := os.ReadFile("ghorahi_financials.json")
    if err != nil {
        fmt.Println("Error reading JSON:", err)
        return
    }

    var data CompanyData
    err = json.Unmarshal(file, &data)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    fmt.Println("Company:", data.Company)
    fmt.Println("Currency:", data.Currency)

    for _, year := range data.Financials {
        fmt.Println("------------------------------")
        fmt.Println("Fiscal Year:", year.FiscalYear)
        fmt.Printf("Revenue: NPR %d\n", year.Revenue)
        fmt.Printf("Gross Profit: NPR %d\n", year.GrossProfit)
        fmt.Printf("Operating Profit/Loss: NPR %d\n", year.ProfitOperations)
        fmt.Printf("Net Profit/Loss: NPR %d\n", year.NetProfitOrLoss)
        fmt.Printf("EPS: NPR %.2f\n", year.EPS)
    }
}