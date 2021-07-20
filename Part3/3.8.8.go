/*
Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
Функция должна называться task().
*/
package main

import "fmt"

func main() {
	ch := make(chan int)

	go task(ch, 1)

	fmt.Println(<-ch)
}

// start
func task(ch chan int, number int) {
	ch <- number + 1
}

// end
