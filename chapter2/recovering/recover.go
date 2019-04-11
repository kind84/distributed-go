package main

import (
	"fmt"
	"sync"
)

func simpleFunc(index int, wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()
	defer wg.Done()
	fmt.Println("Attempting x/(x-10) where x = ", index, " answer is : ", index/(index-10))
}

func main() {
	var wg sync.WaitGroup
	wg.Add(40)

	for i := 0; i < 40; i++ {
		go func(j int) {
			simpleFunc(j, &wg)
		}(i)
	}

	wg.Wait()
}
