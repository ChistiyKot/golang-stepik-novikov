/*
На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)
*/
package main

import (
	"fmt"
)

func main() {
	var xs, result string
	fmt.Scan(&xs)

	for idx, value := range xs {
		if idx%2 != 0 {
			result += string(value)
		}
	}

	fmt.Println(result)
}
