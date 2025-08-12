package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
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

	if len(birthdays) == 0 {
		fmt.Println("No birthdays found.")
		return
	}

	now := time.Now()
	fmt.Println("All Birthdays:")

	for _, b := range birthdays {
		// Calculate next birthday
		birthdayThisYear := time.Date(now.Year(), time.Month(b.Month), b.Day, 0, 0, 0, 0, now.Location())
		if birthdayThisYear.Before(now) {
			birthdayThisYear = birthdayThisYear.AddDate(1, 0, 0)
		}
		daysLeft := int(birthdayThisYear.Sub(now).Hours() / 24)

		// Age calculation
		ageText := ""
		if b.Year != nil {
			age := now.Year() - *b.Year
			if birthdayThisYear.Year() > now.Year() {
				age = age 
			}
			ageText = fmt.Sprintf(", Age: %d", age)
		}

		// Display
		if b.Year != nil {
			fmt.Printf("- %s: %02d/%02d/%d | %d days left%s\n", b.Name, b.Day, b.Month, *b.Year, daysLeft, ageText)
		} else {
			fmt.Printf("- %s: %02d/%02d | %d days left\n", b.Name, b.Day, b.Month, daysLeft)
		}
	}

	fmt.Println("\nPress Enter to return...")
	readLine()
}