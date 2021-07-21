/*
Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).
https://stepik.org/lesson/345547/step/15?unit=329291
*/
package main

import "sync"

func main() {
}

type Worker struct {
	sync.Mutex
	sync.WaitGroup
}

func (w *Worker) Invoke(fn func(int) int, input int, result *int) {
	w.Add(1)
	go func() {
		y := fn(input)
		w.Lock()
		*result += y
		w.Unlock()
		w.Done()
	}()
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		results := make([]int, n)
		var worker Worker

		for i := range results {
			worker.Invoke(fn, <-in1, &results[i])
			worker.Invoke(fn, <-in2, &results[i])
		}

		worker.Wait()

		for i := range results {
			out <- results[i]
		}
	}()
}
