package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) { // uma tarefa que leva 1 segundo pra acabar, quando um acaba outro comeca de maneira sequencial
	for i := range 10 {
		fmt.Printf("%d - Task %s is running \n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	// Thread 2
	go task("A", &waitGroup)
	// Thread 3 - somente por colocar go na frente
	go task("B", &waitGroup)

	go func(){
		for i:= range(5){
			fmt.Printf("%d - Task %s is running \n", i, "anonymous calling")
			time.Sleep(1 * time.Second)
			waitGroup.Done()

		}
	}()

	// sair - encerra as threads

	// time.Sleep(15 * time.Second) // eu espero o tempo que for pra realizar as operacoes das threads - nao e muito bom

	waitGroup.Wait() // pra isso se passa as wait groups
}