/*
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).
*/
package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &data[i])
	}

	for idx, element := range data {
		if idx%2 == 0 {
			fmt.Print(element, " ")
		}
	}
}
