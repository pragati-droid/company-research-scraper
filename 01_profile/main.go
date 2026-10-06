package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type CompanyProfile struct {
	CompanyName string   `json:"company_name"`
	Group       string   `json:"group"`
	Established string   `json:"established"`
	Location    string   `json:"location"`
	Description string   `json:"description"`
	Journey     []string `json:"journey"`
}

func main() {

	profile := CompanyProfile{
		CompanyName: "Ghorahi Cement Industry",
		Group:       "Triveni Group",
		Established: "2009",
		Location:    "Ghorahi, Dang, Nepal",
	}

	c := colly.NewCollector()

	// Find the About Us section
	c.OnHTML("body", func(e *colly.HTMLElement) {

		bodyText := e.Text

		// Extract company description
		start := strings.Index(
			bodyText,
			"Ghorahi Cement Industry is a subsidiary",
		)

		end := strings.Index(
			bodyText,
			"Our Journey",
		)

		if start != -1 && end != -1 && end > start {

			description := bodyText[start:end]

			profile.Description = strings.TrimSpace(description)
		}
	})

	// Extract the journey years
	c.OnHTML("h4", func(e *colly.HTMLElement) {

		year := strings.TrimSpace(e.Text)

		if year == "2009" ||
			year == "2013" ||
			year == "2015" ||
			year == "2020" ||
			year == "2021" {

			profile.Journey = append(
				profile.Journey,
				year,
			)
		}
	})

	// Visit website
	err := c.Visit(
		"https://www.ghorahicement.com.np/about-us-details",
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Display results
	fmt.Println("======================================")
	fmt.Println("GHORAHI CEMENT COMPANY PROFILE")
	fmt.Println("======================================")

	fmt.Println("Company:", profile.CompanyName)
	fmt.Println("Group:", profile.Group)
	fmt.Println("Established:", profile.Established)
	fmt.Println("Location:", profile.Location)

	fmt.Println("\nDescription:")
	fmt.Println(profile.Description)

	fmt.Println("\nJourney Years:")
	for _, year := range profile.Journey {
		fmt.Println("-", year)
	}

	// Save as JSON
	file, err := os.Create("ghorahi_profile.json")

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(profile)

	if err != nil {
		fmt.Println("Error saving JSON:", err)
		return
	}

	fmt.Println("\n--------------------------------------")
	fmt.Println("Data saved to ghorahi_profile.json")
	fmt.Println("--------------------------------------")
}
