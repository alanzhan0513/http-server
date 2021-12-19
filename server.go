package main

import "net/http"

type Server interface {
	Route(pattern string, handleFunc func(c *Context))
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handleFunc func(c *Context)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		c := NewContext(w, r)
		handleFunc(c)
	})
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}
