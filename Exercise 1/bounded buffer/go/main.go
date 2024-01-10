package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		// TODO: push real value to buffer
		ch <- i
	}

}

func consumer(ch chan int) {

	time.Sleep(1 * time.Second)
	for {
		i := <-ch //TODO: get real value from buffer
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}

}

func main() {

	// TODO: make a bounded buffer
	ch := make(chan int, 10)

	go consumer(ch)
	go producer(ch)

	select {}
}

/*
The error message "fatal error: all goroutines are asleep - deadlock!" in Go usually occurs when the program encounters
a situation where all goroutines are blocked and unable to proceed.
*/
