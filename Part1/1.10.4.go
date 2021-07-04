/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.
*/
package main

import "fmt"

func main() {
	var n, input, sum int

	fmt.Scan(&n)

	for count := 0; count < n; count++ {
		fmt.Scan(&input)

		if input < 100 && input > 9 && input%8 == 0 {
			sum += input
		}
	}

	fmt.Println(sum)
}
