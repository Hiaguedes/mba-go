package main

import "fmt"

func main() {
	evento := []string{"ev1", "ev2", "ev3", "ev4", "ev5"}

	fmt.Println("evento[:1]", evento[:1])
	evento = append(evento[:1], evento[2:]...) // tiro o segundo elemento (ev2)

	fmt.Println(evento)
}