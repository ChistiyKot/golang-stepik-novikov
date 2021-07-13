/*
На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является паллиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	xs := []rune(strings.TrimSuffix(text, "\n"))

	if isCaseSensitivePalindrome(xs) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

func isCaseSensitivePalindrome(xs []rune) bool {
	length := len(xs)
	right := length - 1
	for left := 0; left < length/2; left++ {
		if xs[left] != xs[right] {
			return false
		}

		right--
	}

	return true
}
