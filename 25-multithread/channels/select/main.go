package main

import "time"

func main() {
	// select e um switch case mas pra canais

	c1 := make(chan int)
	c2 := make(chan int)

	go func(){
		time.Sleep(time.Second * 1)
		c1 <- 1
	}()

	go func(){
		time.Sleep(time.Second * 3)
		c2 <- 2
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			println("received", msg1)
		case msg2 := <-c2:
			println("received", msg2)
		case <-time.After(time.Second * 2):
			println("timeout")
		// default:
		// 	println("default")
		}
	}



}