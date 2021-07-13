/*
Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления, но и информацию об ошибке. В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции. Функция divide(a int, b int) (int, error) получает на вход два числа которые нужно поделить и возвращает результат (int) и ошибку (error). Пакет main уже объявлен, а нужные пакеты уже импортированы!
*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	result, err := divide(a, b)

	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(result)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("sub zero wins")
	}

	return a / b, nil
}
