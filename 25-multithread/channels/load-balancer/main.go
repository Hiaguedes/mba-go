package main

import (
	"fmt"
	"time"
)

func worker(id int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", id, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	qtyWorkers := 10000
	
	
	for i := range(qtyWorkers){
		go worker(i , data)
	}

	for i := range(100000){
		data <- i
	}
}