/*
Дано трехзначное число. Найдите сумму его цифр.
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Print(n%10 + n/10%10 + n/100%10)
}
