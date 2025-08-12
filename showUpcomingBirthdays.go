package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func showUpcomingBirthdays() {
	file, err := os.Open(dataFile)
	if err != nil {
		return
	}
	defer file.Close()

	var birthdays []Birthday
	if err := json.NewDecoder(file).Decode(&birthdays); err != nil {
		return
	}

	now := time.Now()
	fmt.Println("Upcoming Birthdays:")
	for _, b := range birthdays {
		birthdayDate := time.Date(now.Year(), time.Month(b.Month), b.Day, 0, 0, 0, 0, now.Location())
		if birthdayDate.Before(now) {
			birthdayDate = birthdayDate.AddDate(1, 0, 0)
		}
		daysUntil := int(birthdayDate.Sub(now).Hours() / 24)
		if daysUntil <= 30 {
			fmt.Printf("- %s in %d days (%d/%d)\n", b.Name, daysUntil, b.Day, b.Month)
		}
	}
}