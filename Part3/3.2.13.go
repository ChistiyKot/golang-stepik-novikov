/*
В привычных нам редакторах электронных таблиц присутствует удобное представление числа... https://stepik.org/lesson/348563/step/14?unit=332364
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	numbers := strings.Split(input, ";")
	sum := parse(numbers[0]) / parse(numbers[1])
	result := strconv.FormatFloat(sum, 'f', 4, 64)

	fmt.Println(result)
}

func parse(input string) float64 {
	sanitized := strings.Map(
		func(r rune) rune {
			switch {
			case r == ',':
				return '.'
			case r == ' ' || !unicode.IsNumber(r):
				return -1
			}
			return r
		}, input)

	result, _ := strconv.ParseFloat(sanitized, 0)

	return result
}
