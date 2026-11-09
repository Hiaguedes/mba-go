package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack() // Print the stack trace for debugging purposes
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Recovered from panic: " + err.(string)))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("something went wrong")
	})

	mux.HandleFunc("/recover", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Recovered from panic: " + err.(string)))
			}
		}()
		panic("something went wrong")
	})

	log.Println("Server is running on http://localhost:8080")
	if err :=http.ListenAndServe(":8080", recoverMiddleware(mux)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}