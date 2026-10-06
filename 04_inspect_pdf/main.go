package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile("ghorahi_fy_2081_82.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	text := string(data)

	fmt.Println("======================================")
	fmt.Println("EXTRACTED PDF TEXT")
	fmt.Println("======================================")

	fmt.Println(text)
}
