/*
Найдите самое большее число на отрезке от a до b, кратное 7 .
*/
package main

import "fmt"

func main() {
	var a, b, max int
	found := false
	fmt.Scan(&a, &b)

	for current := b; current >= a; current-- {
		if current%7 == 0 {
			max = current
			found = true
			break
		}
	}

	if found {
		fmt.Println(max)
	} else {
		fmt.Println("NO")
	}
}
