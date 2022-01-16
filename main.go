package main

import (
	"fmt"
)

func main() {
	fmt.Println(isDayInWeek("frf"))
}

// use maps to hold the days of the week
// for range construct
func isDayInWeek(nameOfDay string) bool {
	days := []string{"Sunday", "Monday", "Tuesday", "Wdenesday", "Thursday", "Friday", "Saturday"}
	mapOfDays := make(map[string]int)
	for index, day := range days {
		mapOfDays[day] = index
	}
	fmt.Println(days)
	_, presence := mapOfDays[nameOfDay]
	return presence
}
