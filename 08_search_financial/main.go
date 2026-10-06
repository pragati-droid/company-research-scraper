package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fileName := "ghorahi_fy_2081_82.txt"

	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(data)

	keywords := []string{
		"Revenue",
		"Sales",
		"Profit",
		"Net Profit",
		"Total Assets",
		"Total Liabilities",
		"Equity",
		"Earnings Per Share",
		"EPS",
		"Income Statement",
		"Balance Sheet",
	}

	fmt.Println("====================================")
	fmt.Println("FINANCIAL KEYWORD SEARCH")
	fmt.Println("====================================")

	for _, keyword := range keywords {

		fmt.Println()
		fmt.Println("KEYWORD:", keyword)

		// Case-insensitive search
		index := strings.Index(
			strings.ToLower(text),
			strings.ToLower(keyword),
		)

		if index == -1 {
			fmt.Println("Not found")
			continue
		}

		// Show text around the keyword
		start := index - 300
		end := index + 700

		if start < 0 {
			start = 0
		}

		if end > len(text) {
			end = len(text)
		}

		fmt.Println(text[start:end])
		fmt.Println("------------------------------------")
	}
}
