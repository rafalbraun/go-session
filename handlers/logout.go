package handlers

import (
	"net/http"
	"time"
)

// Logout is a handler that clears a logged in cookie
type Logout struct {
	Name    string
	Path    string
	Domain  string
	Expires func() time.Time
	Next    http.Handler
}

func (l Logout) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: l.Name, Value: "", Domain: l.Domain, Path: l.Path, Expires: l.Expires(), HttpOnly: true}
	http.SetCookie(rw, &cookie)
	l.Next.ServeHTTP(rw, r)
}
