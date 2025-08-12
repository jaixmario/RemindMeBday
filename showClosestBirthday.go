package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func showClosestBirthday() {
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
	var closest Birthday
	minDays := 366

	for _, b := range birthdays {
		birthdayDate := time.Date(now.Year(), time.Month(b.Month), b.Day, 0, 0, 0, 0, now.Location())
		if birthdayDate.Before(now) {
			birthdayDate = birthdayDate.AddDate(1, 0, 0)
		}
		daysUntil := int(birthdayDate.Sub(now).Hours() / 24)
		if daysUntil < minDays {
			minDays = daysUntil
			closest = b
		}
	}

	fmt.Printf("Closest Birthday: %s in %d days (%d/%d)\n", closest.Name, minDays, closest.Day, closest.Month)
	fmt.Println("\nPress Enter to return...")
	readLine()
}