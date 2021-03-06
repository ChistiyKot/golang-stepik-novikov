/*
Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается.
Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.
*/
package main

import "fmt"

func main() {
	var x, p, y, year int
	fmt.Scan(&x, &p, &y)

	for year = 0; x <= y; year++ {
		x += x * p / 100
	}

	fmt.Println(year)
}
