/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
*/
package main

import (
	"fmt"
)

func in(a int, b int) bool {
	if a < 10 {
		return a == b
	}

	if a%10 == b {
		return true
	}

	return in(a/10, b)
}

func main() {
	var first, second int
	fmt.Scan(&first, &second)

	current := first
	var result []int

	for {
		left := current / 10
		right := current % 10

		if in(second, right) {
			result = append(result, right)
		}

		if left == 0 {
			break
		}

		current = left
	}

	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i], " ")
	}
}
