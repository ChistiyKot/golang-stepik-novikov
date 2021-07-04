/*
Последовательность состоит из натуральных чисел и завершается числом 0.
Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
*/
package main

import "fmt"

func main() {
	var n, count, max int

	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n > max {
			max = n
			count = 0
		}

		if n == max {
			count++
		}
	}

	fmt.Println(count)
}
