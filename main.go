package main

import (
	"fmt"
	"os"
)

func main() {
	clearScreen()
	showUpcomingBirthdays()

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "add" {
		addBirthday()
		return
	}

	for {
		clearScreen()
		showUpcomingBirthdays()

		fmt.Println("\n--- Birthday Reminder ---")
		fmt.Println("1. Add Birthday")
		fmt.Println("2. View All Birthdays")
		fmt.Println("3. View Closest Birthday")
		fmt.Println("4. Exit")
		fmt.Print("Enter option: ")

		choice := readLine()
		switch choice {
		case "1":
			addBirthday()
		case "2":
			viewBirthdays()
		case "3":
			showClosestBirthday()
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}