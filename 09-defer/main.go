package main

import "fmt"

func main() {
	fmt.Println("Primeira Linha")
	defer fmt.Println("Segunda Linha") // defer é usado para adiar a execução de uma função até que a função que a chamou retorne. Neste caso, a função fmt.Println("Segunda Linha") será executada após a função main retornar, ou seja, após a execução de fmt.Println("Terceira Linha"). O defer é útil para garantir que certas ações sejam realizadas, como liberar recursos ou fechar conexões, mesmo que ocorra um erro ou um retorno antecipado na função.
	fmt.Println("Terceira Linha")

	/* 
		Resposta esperada:

		Primeira Linha
		Terceira Linha
		Segunda Linha
	*/
}