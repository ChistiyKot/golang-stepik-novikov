/*
Найдите количество минимальных элементов в последовательности.
*/
package main

import "fmt"

func main() {
	var n, current, currentMin, count int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&current)
		if i == 0 {
			currentMin = current
			count++
			continue
		}

		if current < currentMin {
			currentMin = current
			count = 0
		}

		if current == currentMin {
			count++
		}
	}

	fmt.Println(count)
}
