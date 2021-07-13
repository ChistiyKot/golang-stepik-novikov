/*
На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Scan(&n)

	fmt.Println(convertToStringOfSquares(n))
}

func convertToStringOfSquares(n int) string {
	var xs string
	var current int

	for {
		if n == 0 {
			break
		}

		current = n % 10
		xs = strconv.Itoa(current*current) + xs
		n = n / 10
	}

	return xs
}
