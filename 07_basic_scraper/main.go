package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {

	// Create CSV file
	file, err := os.Create("ghorahi_links.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	// Create CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write column headings
	writer.Write([]string{
		"Title",
		"URL",
		"Type",
	})

	// Create Colly scraper
	c := colly.NewCollector()

	// Find all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		title := strings.TrimSpace(e.Text)
		url := strings.TrimSpace(e.Attr("href"))

		// Ignore useless links
		if url == "" ||
			url == "#" ||
			strings.HasPrefix(url, "javascript:") {
			return
		}

		// Determine type
		linkType := "Webpage"

		if strings.Contains(strings.ToLower(url), ".pdf") {
			linkType = "PDF"
		}

		// Save to CSV
		writer.Write([]string{
			title,
			url,
			linkType,
		})

	})

	// Visit website
	err = c.Visit("https://www.ghorahicement.com.np/")

	if err != nil {
		fmt.Println("Error visiting website:", err)
		return
	}

	fmt.Println("Scraping completed!")
	fmt.Println("Data saved to ghorahi_links.csv")
}
package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {

	c := colly.NewCollector()

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println("H1:", e.Text)
	})

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		fmt.Println("H2:", e.Text)
	})

	c.OnHTML("h3", func(e *colly.HTMLElement) {
		fmt.Println("H3:", e.Text)
	})

	err := c.Visit("https://www.ghorahicement.com.np/about-us-details")

	if err != nil {
		fmt.Println("Error:", err)
	}
}