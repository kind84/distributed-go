package main

import (
	"fmt"
	"sync"
)

func recv(ch <-chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("Receiving", i)
	}
	wg.Done()
}

func send(ch chan<- int, wg *sync.WaitGroup) {
	fmt.Println("Sending...")
	ch <- 100
	var wg1 sync.WaitGroup
	wg1.Add(2)
	go func(wg1 *sync.WaitGroup) {
		ch <- 200
		wg1.Done()
	}(&wg1)
	go func(wg1 *sync.WaitGroup) {
		ch <- 300
		wg1.Done()
	}(&wg1)
	wg1.Wait()
	close(ch)
	fmt.Println("Sent")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int, 2)
	go recv(ch, &wg)
	go send(ch, &wg)

	wg.Wait()
}
