/*
Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.
*/
package main

import (
	"fmt"
)

func main() {
	var xs string
	var max int32

	fmt.Scan(&xs)

	for _, value := range xs {
		if value > max {
			max = value
		}
	}

	fmt.Println(string(max))
}
