package handlers

import (
	"net/http"
	"time"
)

// Logout is a handler that clears a logged in cookie
type Logout struct {
	Name   string
	Path   string
	Domain string
	Next   http.Handler
}

func (l Logout) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	cookie := http.Cookie{Name: l.Name, Value: "", Domain: l.Domain, Path: l.Path, Expires: expiration, HttpOnly: true}
	http.SetCookie(rw, &cookie)
	l.Next.ServeHTTP(rw, r)
}
