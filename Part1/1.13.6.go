/*
Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.
*/
package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if checkIfTriangle(a, b, c) {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}

func checkIfTriangle(a int, b int, c int) bool {
	return a < b+c &&
		b < a+c &&
		c < a+b
}
