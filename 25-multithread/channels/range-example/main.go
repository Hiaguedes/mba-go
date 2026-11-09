package main

import "fmt"

func publish(ch chan int) { // enche o canal
	for i := range 10 {
		ch <- i
	}

	// passar isso pra nao ter deadlock, o reader ficara esperando um valor que nunca vira
	close(ch)
}

func reader(ch chan int) { // esvazia o canl
	for i := range ch {
		fmt.Printf("Received %d\n", i)
	}
}

func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch) // ele nao roda em background, o publish sim.  o reader fica aguardando o canal ficar cheio pra rodar
}