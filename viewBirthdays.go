package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func viewBirthdays() {
	clearScreen()
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("No birthdays found.")
		return
	}
	defer file.Close()

	var birthdays []Birthday
	json.NewDecoder(file).Decode(&birthdays)

	fmt.Println("All Birthdays:")
	for _, b := range birthdays {
		if b.Year != nil {
			fmt.Printf("- %s: %02d/%02d/%d\n", b.Name, b.Day, b.Month, *b.Year)
		} else {
			fmt.Printf("- %s: %02d/%02d\n", b.Name, b.Day, b.Month)
		}
	}
	fmt.Println("\nPress Enter to return...")
	readLine()
}