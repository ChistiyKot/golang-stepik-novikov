/*
Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника. Нужно проверить, является ли треугольник прямоугольным. Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"
*/
package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)

	if c*c == a*a+b*b {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
