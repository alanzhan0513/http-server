package main

import "net/http"

type Server interface {
	Route(method string, pattern string, handleFunc func(c *Context))
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler *HttpHandler
}

func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(c *Context)) {
	key := s.handler.key(method, pattern)
	s.handler.handlers[key] = handleFunc
}

func (s *sdkHttpServer) Start(address string) error {
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
		handler: NewHttpHandler(),
	}
}
