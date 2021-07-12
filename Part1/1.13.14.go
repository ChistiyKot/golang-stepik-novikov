/*
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.
*/
package main

import (
	"fmt"
)

func main() {
	var a, n, fn, fn1, fn2, result int

	fmt.Scan(&a)

	n = 2
	fn = 1
	fn1 = 1
	fn2 = 0

	for {
		if a == fn {
			result = n
			break
		}

		if fn > a {
			result = -1
			break
		}

		fn2 = fn1
		fn1 = fn
		fn = fn2 + fn1
		n++
	}

	fmt.Println(result)
}
