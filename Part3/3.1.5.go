/*
Внутри функции main (объявлять функцию не нужно) необходимо написать программу: https://stepik.org/lesson/345543/step/5?unit=329288
*/
package main

import (
	"fmt"
)

func main() {
	// start
	var n int
	var cache = map[int]int{}

	for i := 0; i < 10; i++ {
		fmt.Scan(&n)

		if _, ok := cache[n]; !ok {
			cache[n] = work(n)
		}

		fmt.Print(cache[n], " ")
	}
	// end
}

func work(n int) int {
	return n
}
