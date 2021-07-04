/*
Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
*/
package main

import "fmt"

func main() {
	var n, c, d, result int

	fmt.Scan(&n, &c, &d)

	for current := 1; current <= n; current++ {
		if current%c == 0 && current%d != 0 {
			result = current
			break
		}
	}

	if result != 0 {
		fmt.Println(result)
	}
}
