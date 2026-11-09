package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha_sigilosa")

	bookHotel(ctx)

}

func bookHotel(ctx context.Context){
	token := ctx.Value("token")

	log.Println(token)
}