/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var xs, result string

	fmt.Scan(&xs)

	for _, value := range xs {
		if strings.Count(xs, string(value)) == 1 {
			result += string(value)
		}
	}

	fmt.Println(result)
}
