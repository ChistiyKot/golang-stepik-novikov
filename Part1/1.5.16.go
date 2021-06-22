// Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.

package main

import "fmt"

const (
	DegInHour = 30
	MinInDeg  = 2
)

func main() {
	var input int

	fmt.Scan(&input)

	hours := input / DegInHour
	minutes := (input % DegInHour) * MinInDeg

	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
