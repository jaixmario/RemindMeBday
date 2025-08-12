package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
	clearScreen()
	showUpcomingBirthdays()

	args := os.Args[1:]
	if len(args) > 0 && args[0] == ".ok" {
		addBirthday()
		return
	}

	for {
		clearScreen()
		showUpcomingBirthdays()

		fmt.Println("\n--- Birthday Reminder ---")
		fmt.Println("1. Add Birthday")
		fmt.Println("2. View All Birthdays")
		fmt.Println("3. View Most Close Birthday")
		fmt.Println("4. Exit")
		fmt.Print("Enter option: ")

		choice := readLine()
		switch choice {
		case "1":
			addBirthday()
		case "2":
			viewBirthdays()
		case "3":
			viewMostCloseBirthday()
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func addBirthday() {
	clearScreen()
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

	data := loadData()
	data = append(data, Birthday{
		Name:  name,
		Day:   day,
		Month: month,
		Year:  year,
	})

	saveData(data)
	fmt.Println("Saved successfully!")
	waitForEnter()
}

func viewBirthdays() {
	clearScreen()
	data := loadData()
	if len(data) == 0 {
		fmt.Println("No birthdays found.")
		waitForEnter()
		return
	}
	fmt.Println("--- Saved Birthdays ---")
	for i, b := range data {
		if b.Year != nil {
			fmt.Printf("%d. %s - %02d/%02d/%d\n", i+1, b.Name, b.Day, b.Month, *b.Year)
		} else {
			fmt.Printf("%d. %s - %02d/%02d\n", i+1, b.Name, b.Day, b.Month)
		}
	}
	waitForEnter()
}

func viewMostCloseBirthday() {
	clearScreen()
	data := loadData()
	if len(data) == 0 {
		fmt.Println("No birthdays found.")
		waitForEnter()
		return
	}

	today := time.Now()
	var closest Birthday
	minDays := 999999

	for _, b := range data {
		thisYear := today.Year()
		birthday := time.Date(thisYear, time.Month(b.Month), b.Day, 0, 0, 0, 0, time.Local)
		if birthday.Before(today) {
			birthday = time.Date(thisYear+1, time.Month(b.Month), b.Day, 0, 0, 0, 0, time.Local)
		}
		days := int(birthday.Sub(today).Hours() / 24)
		if days < minDays {
			minDays = days
			closest = b
		}
	}

	fmt.Println("🎯 Most Close Birthday 🎯")
	if closest.Year != nil {
		fmt.Printf("%s - %02d/%02d/%d (in %d days)\n", closest.Name, closest.Day, closest.Month, *closest.Year, minDays)
	} else {
		fmt.Printf("%s - %02d/%02d (in %d days)\n", closest.Name, closest.Day, closest.Month, minDays)
	}
	waitForEnter()
}

func showUpcomingBirthdays() {
	data := loadData()
	if len(data) == 0 {
		fmt.Println("No upcoming birthdays.")
		return
	}

	today := time.Now()
	upcoming := []Birthday{}

	for _, b := range data {
		thisYear := today.Year()
		birthdayThisYear := time.Date(thisYear, time.Month(b.Month), b.Day, 0, 0, 0, 0, time.Local)

		if birthdayThisYear.Before(today) {
			birthdayThisYear = time.Date(thisYear+1, time.Month(b.Month), b.Day, 0, 0, 0, 0, time.Local)
		}

		diff := birthdayThisYear.Sub(today).Hours() / 24
		if diff >= 0 && diff <= 7 {
			upcoming = append(upcoming, b)
		}
	}

	if len(upcoming) == 0 {
		fmt.Println("No birthdays in the next 7 days.")
		return
	}

	fmt.Println("🎂 Upcoming Birthdays (next 7 days) 🎂")
	for _, b := range upcoming {
		if b.Year != nil {
			fmt.Printf("- %s on %02d/%02d/%d\n", b.Name, b.Day, b.Month, *b.Year)
		} else {
			fmt.Printf("- %s on %02d/%02d\n", b.Name, b.Day, b.Month)
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

func waitForEnter() {
	fmt.Print("\nPress Enter to continue...")
	readLine()
}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}