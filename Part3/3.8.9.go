/*
Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.
Функция должна называться task2().
*/
package main

import "fmt"

func main() {
	ch := make(chan string)

	go task2(ch, "foo")

	fmt.Println(<-ch)
}

// start
func task2(ch chan string, input string) {
	for i := 0; i < 5; i++ {
		ch <- input + " "
	}
}

// end
