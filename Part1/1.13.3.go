/*
Дано трехзначное число. Переверните его, а затем выведите
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Printf("%1d%1d%1d", n%10, n/10%10, n/100%10)
}
