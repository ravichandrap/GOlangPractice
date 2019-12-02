package main

import "fmt"

func main() {
	day := Monday

	fmt.Println(day.String())
	fmt.Println(day.Weekend())
}


type WeekDays int

const (
	Sunday    WeekDays = 0
	Monday    WeekDays = 1
	Tuesday    WeekDays = 2
	Wednesday WeekDays = 3
	Thursday    WeekDays = 4
	Friday    WeekDays = 5
	Saturday WeekDays = 6
)

func (day WeekDays)Weekend() bool  {
	switch day {
	case Saturday, Sunday:
		return true
	default:
		return false
	}
}

func (day WeekDays)String() string  {
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday"}

	if day < Sunday || day > Saturday {
		return "Unknown"
	}

	return names[day]
}