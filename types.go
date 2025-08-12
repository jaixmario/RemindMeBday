package main

type Birthday struct {
	Name  string `json:"name"`
	Day   int    `json:"day"`
	Month int    `json:"month"`
	Year  *int   `json:"year,omitempty"`
}

const dataFile = "birthdays.json"