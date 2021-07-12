/*
Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.
*/
package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 0
	z := 1

	fmt.Println(vote(x, y, z))
}

func vote(x int, y int, z int) int {
	sum := x + y + z

	if sum < 2 {
		return 0
	} else {
		return 1
	}
}
