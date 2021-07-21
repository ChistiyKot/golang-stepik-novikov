/*
Вам необходимо написать функцию calculator следующего вида:
func calculator(arguments <-chan int, done <-chan struct{}) <-chan int
https://stepik.org/lesson/345547/step/14?unit=329291
*/
package main

import "fmt"

func main() {
	args := make(chan int)
	done := make(chan struct{})

	r := calculator(args, done)
	args <- 2
	args <- 4
	//done <- struct{}{}

	fmt.Println(<-r)
}

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	// start
	resultChan := make(chan int)

	go func() {
		result := 0
		for {
			select {
			case <-done:
				resultChan <- result
				close(resultChan)
			case arg := <-arguments:
				result += arg
			}
		}
	}()

	return resultChan
	// end
}
