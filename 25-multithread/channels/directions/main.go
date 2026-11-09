package main

import "fmt"

// chan<- vai receber o resultado
// pode entender como entrar no channel
// receive only
func receive(name string, ch chan<- string) {
	ch <- name
}

// <-chan significa que vai entregar o resultado
// pode entender como sair do channel
// send only
func read(data <-chan string) {
	fmt.Println(<-data)
}

// thread 1
func main() {
	ch := make(chan string)
	go receive("Hello", ch)
	
	read(ch)
}