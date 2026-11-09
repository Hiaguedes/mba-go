package main

import (
	"io"
	"net/http"
)

func main() {
	req,err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	println("Status code:", req.StatusCode) // Print the status code of the response.
	println("Status:", req.Status) // Print the status of the response.
	println("Content length:", req.ContentLength) // Print the content length of the response.

	res, err := io.ReadAll(req.Body) // ReadAll is used to read the entire response body. It takes the response body as an argument and returns the content as a byte slice and an error if any.

	if err != nil {
		panic(err)
	}

	println("Response body:", string(res)) // Print the response body. The byte slice is converted to a string for printing.

	defer req.Body.Close()
}