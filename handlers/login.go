package handlers

import (
	"net/http"
	"time"
)

// Login is a handler that sets a logged in cookie
type Login struct {
	Name     string
	Value    string
	Path     string
	Domain   string
	Duration time.Duration
	Next     http.Handler
}

func (l Login) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(l.Duration)
	cookie := http.Cookie{Name: l.Name, Value: l.Value, Domain: l.Domain, Path: l.Path, Expires: expiration, HttpOnly: true}
	http.SetCookie(rw, &cookie)
	l.Next.ServeHTTP(rw, r)
}
