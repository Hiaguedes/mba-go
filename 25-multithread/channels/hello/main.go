package main

import "fmt"

// thread principal
func main() {
	// channel
	/*
		Fazer comunicacao entre threads
		Seguranca para um thread saber o momento em que ela pode trabalhar com um determinado dado
		
	*/
	
	channel := make(chan string) // canal vazio

	// thread 2
	go func() {
		channel <- "Ola Mundo!" // canal cheio

	}()

	// thread 1

	msg := <-channel // canal esvazia
	fmt.Println(msg) // printo mensagem so apos canal esvaziar

}