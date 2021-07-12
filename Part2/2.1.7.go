/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {
	var current, min int

	for i := 0; i < 4; i++ {
		fmt.Scan(&current)

		if i == 0 || current < min {
			min = current
		}
	}

	return min
}
