package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {

	file, err := os.Create("ghorahi_reports.csv")

	if err != nil {
		fmt.Println("Error creating CSV:", err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// CSV headings
	writer.Write([]string{
		"Report Name",
		"URL",
		"Type",
	})

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		name := strings.TrimSpace(e.Text)
		url := strings.TrimSpace(e.Attr("href"))

		if name == "" || url == "" {
			return
		}

		// We only want PDF files
		if !strings.Contains(
			strings.ToLower(url),
			".pdf",
		) {
			return
		}

		reportType := "Other"

		lowerName := strings.ToLower(name)

		if strings.Contains(lowerName, "quarter") {
			reportType = "Quarterly Report"
		}

		if strings.Contains(lowerName, "annual") ||
			strings.Contains(lowerName, "fy ") {

			reportType = "Annual Report"
		}

		writer.Write([]string{
			name,
			url,
			reportType,
		})

		fmt.Println("Report:", name)
		fmt.Println("URL:", url)
		fmt.Println("Type:", reportType)
		fmt.Println("--------------------------------")
	})

	err = c.Visit(
		"https://www.ghorahicement.com.np/",
	)

	if err != nil {
		fmt.Println("Error visiting website:", err)
		return
	}

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("SCRAPING COMPLETED")
	fmt.Println("================================")
	fmt.Println("Data saved to: ghorahi_reports.csv")
}
