package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Birthday struct {
	Name  string `json:"name"`
	Day   int    `json:"day"`
	Month int    `json:"month"`
	Year  *int   `json:"year,omitempty"`
}

const dataFile = "birthdays.json"

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == ".ok" {
		addBirthday()
		return
	}

	for {
		fmt.Println("\n--- Birthday Reminder ---")
		fmt.Println("1. Add Birthday")
		fmt.Println("2. View All Birthdays")
		fmt.Println("3. Exit")
		fmt.Print("Enter option: ")

		choice := readLine()
		switch choice {
		case "1":
			addBirthday()
		case "2":
			viewBirthdays()
		case "3":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func addBirthday() {
	name := ""
	for {
		fmt.Print("Enter name: ")
		name = strings.TrimSpace(readLine())
		if name != "" {
			break
		}
		fmt.Println("Name cannot be empty.")
	}

	day := readInt("Enter day (1-31): ", 1, 31)
	month := readInt("Enter month (1-12): ", 1, 12)

	fmt.Print("Enter year (optional, press enter to skip): ")
	yearInput := strings.TrimSpace(readLine())
	var year *int
	if yearInput != "" {
		yearVal, err := strconv.Atoi(yearInput)
		currentYear := time.Now().Year()
		if err != nil || yearVal < 1900 || yearVal > currentYear {
			fmt.Println("Invalid year. Skipping.")
		} else {
			year = &yearVal
		}
	}

	// Load existing data
	data := loadData()

	// Add new birthday
	data = append(data, Birthday{
		Name:  name,
		Day:   day,
		Month: month,
		Year:  year,
	})

	// Save
	saveData(data)
	fmt.Println("Saved successfully!")
}

func viewBirthdays() {
	data := loadData()
	if len(data) == 0 {
		fmt.Println("No birthdays found.")
		return
	}
	fmt.Println("\n--- Saved Birthdays ---")
	for i, b := range data {
		if b.Year != nil {
			fmt.Printf("%d. %s - %d/%d/%d\n", i+1, b.Name, b.Day, b.Month, *b.Year)
		} else {
			fmt.Printf("%d. %s - %d/%d\n", i+1, b.Name, b.Day, b.Month)
		}
	}
}

func loadData() []Birthday {
	file, err := os.Open(dataFile)
	if err != nil {
		return []Birthday{}
	}
	defer file.Close()

	var data []Birthday
	json.NewDecoder(file).Decode(&data)
	return data
}

func saveData(data []Birthday) {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(data)
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readInt(prompt string, min, max int) int {
	for {
		fmt.Print(prompt)
		input := readLine()
		val, err := strconv.Atoi(input)
		if err == nil && val >= min && val <= max {
			return val
		}
		fmt.Printf("Please enter a valid number between %d and %d.\n", min, max)
	}
}