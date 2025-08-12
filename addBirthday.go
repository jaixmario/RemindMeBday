package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addBirthday() {
	clearScreen()

	// Name validation (cannot be empty)
	var name string
	for {
		fmt.Print("Enter Name: ")
		name = strings.TrimSpace(readLine())
		if name != "" {
			break
		}
		fmt.Println("Name cannot be empty. Please try again.")
	}

	// Day validation
	var day int
	for {
		fmt.Print("Enter Day (1-31): ")
		dayStr := strings.TrimSpace(readLine())
		d, err := strconv.Atoi(dayStr)
		if err == nil && d >= 1 && d <= 31 {
			day = d
			break
		}
		fmt.Println("Invalid day. Please enter a number between 1 and 31.")
	}

	// Month validation
	var month int
	for {
		fmt.Print("Enter Month (1-12): ")
		monthStr := strings.TrimSpace(readLine())
		m, err := strconv.Atoi(monthStr)
		if err == nil && m >= 1 && m <= 12 {
			month = m
			break
		}
		fmt.Println("Invalid month. Please enter a number between 1 and 12.")
	}

	// Year validation (optional)
	var year *int
	for {
		fmt.Print("Enter Year (optional, press Enter to skip): ")
		yearStr := strings.TrimSpace(readLine())
		if yearStr == "" {
			break
		}
		y, err := strconv.Atoi(yearStr)
		if err == nil && y > 0 {
			year = &y
			break
		}
		fmt.Println("Invalid year. Please enter a valid number or leave empty.")
	}

	// Open or create file
	file, err := os.OpenFile(dataFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read existing birthdays
	var birthdays []Birthday
	json.NewDecoder(file).Decode(&birthdays)

	// Append new birthday
	birthdays = append(birthdays, Birthday{Name: name, Day: day, Month: month, Year: year})

	// Save to file
	file.Seek(0, 0)
	file.Truncate(0)
	json.NewEncoder(file).Encode(birthdays)

	fmt.Println("✅ Birthday added successfully!")
}