package main

import "fmt"

type msg struct {
	ID    int
	value string
}

func handleIntChan(intChan <-chan int, done chan<- int) {
	for i := 0; i < 6; i++ {
		fmt.Println(<-intChan)
	}
	done <- 0
}

func handleMsgChan(msgChan <-chan msg, done chan<- int) {
	for i := 0; i < 6; i++ {
		fmt.Println(<-msgChan)
	}
	done <- 0
}

func main() {
	intChan := make(chan int)
	done := make(chan int)

	go func() {
		intChan <- 9
		intChan <- 2
		intChan <- 3
		intChan <- 7
		close(intChan)
	}()
	go handleIntChan(intChan, done)

	msgChan := make(chan msg)
	go func() {
		for i := 1; i < 5; i++ {
			msgChan <- msg{
				ID:    i,
				value: fmt.Sprintf("VALUE-%v", i),
			}
		}
		close(msgChan)
	}()
	go handleMsgChan(msgChan, done)

	<-done
	<-done

	// intChan <- 100
}
