package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}

	jsonVar := bytes.NewBuffer([]byte(`{name: "Hiago"}`)) // o terceiro parametro e um reader ou buffer pro post conseguir aplicar o post
	
	resp, err := c.Post("http://google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}