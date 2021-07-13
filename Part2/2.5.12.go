/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
*/
package main

import (
	"fmt"
	"unicode"
)

func main() {
	var xs string

	fmt.Scan(&xs)

	if isValid(xs) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func isValid(xs string) bool {
	if len(xs) < 5 {
		return false
	}

	for _, r := range xs {
		if unicode.Is(unicode.Latin, r) || unicode.IsDigit(r) {
			continue
		} else {
			return false
		}
	}
	return true
}
