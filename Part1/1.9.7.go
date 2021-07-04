// Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

package main

import "fmt"

func main() {
	var input int

	fmt.Scan(&input)

	if input < 10 {
		fmt.Println(input)
	} else if input < 99 {
		fmt.Println(input / 10)
	} else if input < 999 {
		fmt.Println(input / 100)
	} else if input < 9999 {
		fmt.Println(input / 1000)
	} else if input == 10000 {
		fmt.Println(1)
	}
}
