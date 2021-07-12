/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".
*/
package main

import "fmt"

func main() {
	var n int
	var result string
	fmt.Scan(&n)

	if n == 1 || (n > 20 && n%10 == 1) {
		result = "korova"
	} else if ((n >= 2 && n <= 4) || n%10 == 2 || n%10 == 3 || n%10 == 4) && (n != 12 && n != 13 && n != 14) {
		result = "korovy"
	} else {
		result = "korov"
	}

	fmt.Printf("%d %s\n", n, result)
}
