/*
Вам необходимо написать функцию calculator следующего вида:
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
https://stepik.org/lesson/345547/step/13?unit=329291
*/
package main

import "fmt"

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})

	r := calculator(ch1, ch2, stop)
	//ch1 <- 2
	//ch2 <- 2
	//close(stop)
	fmt.Println(<-r)
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	// start
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)
		select {
		case <-stopChan:
			return
		case val := <-firstChan:
			resultChan <- val * val
		case val := <-secondChan:
			resultChan <- val * 3
		}
	}()

	return resultChan
	// end
}
