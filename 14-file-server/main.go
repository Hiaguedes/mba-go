package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // FileServer is used to create a new HTTP handler that serves files from the specified directory. In this case, it serves files from the "./static" directory.
	
	mux := http.NewServeMux()

	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's a blog"))
	})
	
	println("File server is running on http://localhost:8080 running static files")
	log.Fatal(http.ListenAndServe(":8080", mux))
}