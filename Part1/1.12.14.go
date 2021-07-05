/*
На ввод подаются пять чисел, которые записываются в массив. Однако эта часть программы уже написана.

Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.
*/
package main

import "fmt"

func main() {
	// predefined code
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	// my code
	var max int
	for _, element := range array {
		if element > max {
			max = element
		}
	}

	fmt.Println(max)
}
