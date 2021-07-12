/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumInt(1, 2, 3))
}

func sumInt(xs ...int) (int, int) {
	sum := 0
	for _, x := range xs {
		sum += x
	}

	return len(xs), sum
}
