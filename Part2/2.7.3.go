/*
Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var xs string

	fmt.Scan(&xs)

	fmt.Println(strings.Join(strings.Split(xs, ""), "*"))
}
