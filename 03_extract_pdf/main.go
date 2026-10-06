package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ledongthuc/pdf"
)

func main() {

	fileName := "ghorahi_fy_2081_82.pdf"

	// Open PDF
	file, reader, err := pdf.Open(fileName)

	if err != nil {
		fmt.Println("Error opening PDF:", err)
		return
	}

	defer file.Close()

	fmt.Println("PDF opened successfully!")
	fmt.Println("Number of pages:", reader.NumPage())

	var allText strings.Builder

	// Read every page
	for pageNumber := 1; pageNumber <= reader.NumPage(); pageNumber++ {

		page := reader.Page(pageNumber)

		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)

		if err != nil {
			fmt.Println(
				"Error reading page",
				pageNumber,
				":",
				err,
			)
			continue
		}

		textLength := len(strings.TrimSpace(text))

fmt.Printf("Page %d: extracted %d characters\n", pageNumber, textLength)

allText.WriteString(text)
allText.WriteString("\n")
	}

	// Save extracted text
	outputFile := "ghorahi_fy_2081_82.txt"

	err = os.WriteFile(
		outputFile,
		[]byte(allText.String()),
		0644,
	)

	if err != nil {
		fmt.Println("Error saving text:", err)
		return
	}

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("PDF EXTRACTION COMPLETED")
	fmt.Println("================================")
	fmt.Println("Text saved to:", outputFile)
}
