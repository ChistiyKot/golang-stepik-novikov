/*
По данным числам, определите количество чисел, которые равны нулю.
*/
package main

import "fmt"

func main() {
	var n, input, count int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&input)

		if input == 0 {
			count++
		}
	}

	fmt.Println(count)
}
