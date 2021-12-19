package main

import "net/http"

type HttpHandler struct {
	handlers map[string]func(c *Context)
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := h.key(r.Method, r.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(w, r))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	}
}

func (h *HttpHandler) key(method string, pattern string) string {
	return method + "#" + pattern
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{
		handlers: make(map[string]func(c *Context)),
	}
}
