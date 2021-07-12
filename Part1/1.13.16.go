/*
Из натурального числа удалить заданную цифру.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, excluded int

	fmt.Scan(&n, &excluded)

	fmt.Println(filter(n, excluded))
}

func filter(n, excluded int) int {
	var xs string
	var current int

	current = n
	for {
		if n == 0 {
			break
		}

		current = n % 10
		if current != excluded {
			xs = fmt.Sprintf("%d", current) + xs
		}
		n = n / 10
	}

	result, _ := strconv.Atoi(xs)

	return result
}
