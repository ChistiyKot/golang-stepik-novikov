/*
Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток
*/
package main

import (
	"fmt"
)

func main() {
	var k, h, m int
	fmt.Scanf("%d", &k)

	h = k / 3600
	k %= 3600
	m = k / 60

	fmt.Printf("It is %d hours %d minutes.\n", h, m)
}
