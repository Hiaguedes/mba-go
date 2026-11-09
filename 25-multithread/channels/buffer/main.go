package main

func main() {
	ch := make(chan string, 2) // buffer -> consegue receber mais mensagens de uma vez
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}