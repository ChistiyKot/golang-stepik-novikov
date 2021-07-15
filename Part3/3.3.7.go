/*
Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint, которую внутри функции main в дальнейшем можно будет вызвать по имени fn... https://stepik.org/lesson/349644/step/7?unit=333499
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// start
	fn := func(input uint) uint {
		stringResult := ""
		for _, value := range strconv.Itoa(int(input)) {
			currentNumber, _ := strconv.Atoi(string(value))
			if currentNumber%2 == 0 && currentNumber != 0 {
				stringResult += string(value)
			}
		}

		result, _ := strconv.Atoi(stringResult)

		if result == 0 {
			result = 100
		}

		return uint(result)
	}
	// end

	fmt.Println(fn(727178))
}
