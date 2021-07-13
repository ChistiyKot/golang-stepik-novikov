/*
Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X. Если подстроки S нет в строке X - вывести -1
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var xs, s string
	fmt.Scan(&xs)
	fmt.Scan(&s)

	index := strings.Index(xs, s)
	if index >= 0 {
		fmt.Println(index)
	} else {
		fmt.Println(-1)
	}
}
