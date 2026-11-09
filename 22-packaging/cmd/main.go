package main

import (
	"fmt"
	"hiaguedes/packaging/math"

	"github.com/google/uuid" // basta rodar go mod tidy para baixar o pacote e adicionar no go.mod
)

func main() {
	fmt.Println(math.Add(2, 3))

	fmt.Println(uuid.New().String())

}