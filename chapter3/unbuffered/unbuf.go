package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() { ch <- 100 }()

	func() {
		val := <-ch
		fmt.Println(val)
	}()

	func() {
		val := <-ch
		fmt.Println(val)
	}()
}
