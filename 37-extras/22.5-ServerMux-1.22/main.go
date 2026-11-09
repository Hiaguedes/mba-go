package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /books/{id}", GetBookHandler)
	mux.HandleFunc("POST /books/{id}", GetBookHandler)
	mux.HandleFunc("GET /books/dir/{d...}", BooksPathHandler)
	mux.HandleFunc("GET /books/{$}", BooksHandler) // exato - ou seja so /books
	mux.HandleFunc("GET /books/precedence/latest", BooksPrecedenceHandler)
	mux.HandleFunc("GET /books/precedence/{x}", BooksPrecedence2Handler)
	// mux.HandleFunc("GET /books/{s}", BooksPrecedenceHandler)
	// mux.HandleFunc("GET /{s}/latest", BooksPrecedence2Handler)
	log.Println("Server is running on http://localhost:9000")
	http.ListenAndServe(":9000", mux)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Book " + id))
}

func BooksPathHandler(w http.ResponseWriter, r *http.Request) {
	dirpath := r.PathValue("d") // Access captured directory path segments as slice
	fmt.Fprintf(w, "Accessing directory path: %s\n", dirpath)
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books"))
}

func BooksPrecedenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence"))
}

func BooksPrecedence2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence 2"))
}
