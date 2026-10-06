package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {

	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {

		text := strings.TrimSpace(e.Text)

		fmt.Println("===== GHORAHI CEMENT COMPANY PROFILE =====")
		fmt.Println(text)

	})

	err := c.Visit("https://www.ghorahicement.com.np/about-us-details")

	if err != nil {
		fmt.Println("Error:", err)
	}
}
