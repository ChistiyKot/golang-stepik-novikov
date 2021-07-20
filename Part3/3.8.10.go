/*
Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий этап конвейера только если оно отличается от того, что пришло ранее... https://stepik.org/lesson/360357/step/10?unit=344766
*/
package main

import "fmt"

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)

	go func() {
		defer close(inputStream)

		for _, r := range "112334456" {
			inputStream <- string(r)
		}
	}()

	for x := range outputStream {
		fmt.Print(x)
	}
}

// start
func removeDuplicates(inputStream chan string, outputStream chan string) {
	var prev string
	for value := range inputStream {
		if value == prev {
			continue
		}

		outputStream <- value
		prev = value
	}

	close(outputStream)
}

// end
