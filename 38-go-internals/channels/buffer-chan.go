package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2) // Buffered channel with capacity 2

	go func() {
		ch <- 1
		fmt.Println("Sent 1")
		ch <- 2
		fmt.Println("Sent 2")
		ch <- 3 // This will block until another goroutine receives from the channel
		fmt.Println("Sent 3")
	}()

time.Sleep(1 * time.Second)

go func () {
	val := <- ch
	fmt.Println("Value recieved: ", val)
}()

time.Sleep(2 * time.Second);
}