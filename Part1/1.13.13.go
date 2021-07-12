/*
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	idx := 0
	for {
		power := int(math.Pow(2, float64(idx)))

		if power > n {
			break
		}

		fmt.Print(power, " ")
		idx++
	}
}
