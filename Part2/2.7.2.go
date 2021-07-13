/*
На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(math.Sqrt(float64(a*a) + float64(b*b)))
}
