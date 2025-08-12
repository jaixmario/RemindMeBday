package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func addBirthday() {
	clearScreen()
	fmt.Print("Enter Name: ")
	name := readLine()

	fmt.Print("Enter Day (1-31): ")
	dayStr := readLine()
	day, _ := strconv.Atoi(dayStr)

	fmt.Print("Enter Month (1-12): ")
	monthStr := readLine()
	month, _ := strconv.Atoi(monthStr)

	fmt.Print("Enter Year (optional, press Enter to skip): ")
	yearStr := readLine()
	var year *int
	if yearStr != "" {
		y, _ := strconv.Atoi(yearStr)
		year = &y
	}

	file, _ := os.OpenFile(dataFile, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	var birthdays []Birthday
	json.NewDecoder(file).Decode(&birthdays)

	birthdays = append(birthdays, Birthday{Name: name, Day: day, Month: month, Year: year})

	file.Seek(0, 0)
	file.Truncate(0)
	json.NewEncoder(file).Encode(birthdays)

	fmt.Println("Birthday added successfully!")
}