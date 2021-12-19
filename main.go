package main

import (
	"io"
	"net/http"
)

func main() {
	server := NewHttpServer("hello world")
	server.Route("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello World")
	})
	server.Start(":8080")
}
