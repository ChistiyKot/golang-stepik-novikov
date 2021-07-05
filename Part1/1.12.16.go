/*
Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.
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

	count := 0
	for _, element := range data {
		if element > 0 {
			count++
		}
	}

	fmt.Println(count)
}
