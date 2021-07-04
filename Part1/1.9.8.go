// Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
package main

import "fmt"

func main() {
	var input int

	fmt.Scan(&input)

	var (
		first  = input / 100000 % 10
		second = input / 10000 % 10
		third  = input / 1000 % 10
		fourth = input / 100 % 10
		fifth  = input / 10 % 10
		sixth  = input % 10
	)

	if first+second+third == fourth+fifth+sixth {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
