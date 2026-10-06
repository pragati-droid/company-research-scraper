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

	c.OnHTML("h4", func(e *colly.HTMLElement) {
		fmt.Println("H4:", e.Text)
	})

	err := c.Visit(
		"https://www.ghorahicement.com.np/about-us-details",
	)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
