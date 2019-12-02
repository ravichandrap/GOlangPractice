package main

import (
	"fmt"
	"time"
)

func main() {
getYearWeekDayfoo()
}

func getYearWeekDayfoo() {


	year, week := time.Now().ISOWeek()
	day := time.Now().Day()
	f := fmt.Sprintf("%d%d%d", year, week, day)
	fmt.Println(f)
}
