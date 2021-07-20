/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельных горутинах вызвать функцию work() 10 раз и дождаться результатов выполнения вызванных функций.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// start
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			work()
			defer wg.Done()
		}()
	}

	wg.Wait()
	// end
}

func work() {
	id := rand.Intn(100)
	fmt.Printf("Горутина #%d начала выполнение \n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Горутина #%d завершила выполнение \n", id)
}
