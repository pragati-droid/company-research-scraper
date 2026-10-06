package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	url := "https://www.ghorahicement.com.np/assets/financial-report/fy_81_82.pdf"

	fileName := "ghorahi_fy_2081_82.pdf"

	fmt.Println("Downloading report...")
	fmt.Println("URL:", url)

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error downloading:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)

	if response.StatusCode != 200 {
		fmt.Println("Download failed.")
		return
	}

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		fmt.Println("Error saving PDF:", err)
		return
	}

	fmt.Println("================================")
	fmt.Println("DOWNLOAD COMPLETED")
	fmt.Println("================================")
	fmt.Println("Saved as:", fileName)
}
