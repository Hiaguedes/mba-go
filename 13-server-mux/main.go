package main

import (
	"net/http"
)

type Blog struct {
	title   string
}

func (b Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to my blog " + b.title + "!"))
}

func main() {
	mux := http.NewServeMux() // NewServeMux is used to create a new HTTP request multiplexer. It returns a new ServeMux instance that can be used to register handlers for different URL paths.

	mux.HandleFunc("/", HelloWorldHandler) // HandleFunc is used to register a handler function for the specified URL path. In this case, it registers the HelloWorldHandler function to handle requests to the root URL ("/").
	mux.Handle("/blog", Blog{ title: "h1.blog" }) // Handle is used to register a handler for the specified URL path. In this case, it registers an instance of the Blog struct to handle requests to the "/blog" URL.
	
	println("Mux1 is running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)

}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

