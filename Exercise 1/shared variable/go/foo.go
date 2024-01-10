// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func numberServer(incrementCh chan int, decrementCh chan int) {
	for {
		select {
		case <-incrementCh:
			i++
		case <-decrementCh:
			i--
		}
	}
}

func incrementing(ch chan int, done chan int) {
	//TODO: increment i 1000000 times
	for j := 0; j < 1000001; j++ {
		ch <- 1
	}
	done <- 1

}

func decrementing(ch chan int, done chan int) {
	//TODO: decrement i 1000000 times
	for j := 0; j < 1000000; j++ {
		ch <- 1
	}
	done <- 1

}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	incrementCh := make(chan int, 1)
	decrementCh := make(chan int, 1)
	incrementDone := make(chan int, 1)
	decrementDone := make(chan int, 1)

	go incrementing(incrementCh, incrementDone)
	go decrementing(decrementCh, decrementDone)
	go numberServer(incrementCh, decrementCh)

	<-incrementDone
	<-decrementDone // Once both signals have been received, the main goroutine continues to the next lines of code

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)

	Println("The magic number is:", i)
}
