package main

import (
	"net/http"
)

func main() {
	server := NewHttpServer("hello world")
	server.Route("GET", "/", func(c *Context) {
		c.WriteJson(http.StatusOK, "hello world")
	})
	server.Start(":8080")
}
