/*
Дано натуральное число N. Выведите его представление в двоичном виде.
*/
package main

import (
	"fmt"
)

func main() {
	var n, r int
	var result string

	fmt.Scan(&n)

	for n > 0 {
		r = n % 2
		n /= 2

		result = fmt.Sprintf("%d", r) + result
	}

	fmt.Println(result)
}
