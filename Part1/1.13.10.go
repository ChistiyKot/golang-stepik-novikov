/*
Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр, на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации. Этот процесс повторяется до тех пор, пока не будет получена одна цифра.
По данному числу определите его цифровой корень.
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(calcDigitalRoot(n))
}

func calcDigitalRoot(n int) int {
	root := sumOfDigits(n)

	if root > 9 {
		root = calcDigitalRoot(root)
	}

	return root
}

func sumOfDigits(n int) int {
	sum := 0
	for {
		if n == 0 {
			break
		}

		sum += n % 10
		n = n / 10
	}

	return sum
}
