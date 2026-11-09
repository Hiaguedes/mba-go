package main

import (
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context();
	log.Println("Request iniciada")

	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		// imprime log no servidor, o stdout
		log.Println("Request processada com sucesso")
		// imorime no browser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done(): // roda quando cancela
		log.Println("Request cancelada pelo cliente")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}