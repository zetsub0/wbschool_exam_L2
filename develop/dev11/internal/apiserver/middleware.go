package apiserver

import (
	"log"
	"net/http"
)

type Middleware struct {
	next http.Handler
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Printf("- Request from client %s: %s %s", r.RemoteAddr, r.Method, r.URL.Path)

	m.next.ServeHTTP(w, r)
}

func NewMiddleware(next http.Handler) *Middleware {
	return &Middleware{next: next}
}
