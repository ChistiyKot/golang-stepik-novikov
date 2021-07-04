// По данному трехзначному числу определите, все ли его цифры различны.

package main

import "fmt"

func main() {
	var input int

	fmt.Scan(&input)

	var (
		third  int = input % 10
		second int = input % 100 / 10
		first  int = input / 100
	)

	if first != second && second != third && first != third {
		fmt.Printf("YES")
	} else {
		fmt.Printf("NO")
	}
}
