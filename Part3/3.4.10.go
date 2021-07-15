/*
Пришло время для задач, где вы сможете применить полученные знания на практике... https://stepik.org/lesson/350788/step/10?unit=334666
*/
package main

import (
	"fmt"
)

func main() {
	value1, value2, operation := readTask()
	val1, ok := value1.(float64)

	if !ok {
		printError(value1)
		return
	}

	val2, ok := value2.(float64)

	if !ok {
		printError(value2)
		return
	}

	op, ok := operation.(string)
	if !ok || (op != "/" && op != "+" && op != "-" && op != "*") {
		fmt.Println("недопустимая операция")
		return
	}

	fmt.Printf("%.4f", run(op, val1, val2))
}

func run(operation string, value1, value2 float64) float64 {
	switch operation {
	case "+":
		return value1 + value2
	case "-":
		return value1 - value2
	case "*":
		return value1 * value2
	case "/":
		return value1 / value2
	}

	panic("Undefined operation.")
}

func printError(value interface{}) {
	fmt.Printf("value=%v: %T", value, value)
}

func readTask() (value1, value2, operation interface{}) {
	return 2.0, 2.2, "+"
}
