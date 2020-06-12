package handlers

import (
	"html/template"
	"net/http"
	"time"
)

// Login is a handler that sets a logged in cookie
type Login struct {
	Name   string
	Value  string
	Path   string
	Domain string
	MaxAge int
	Expires time.Time
	Next   http.Handler
}

func (l Login) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     l.Name,
		Value:    l.Value,
		Domain:   l.Domain,
		Path:     l.Path,
		Expires:  l.Expires,
		MaxAge:   l.MaxAge,
		HttpOnly: true,
	}
	http.SetCookie(rw, &cookie)
	l.Next.ServeHTTP(rw, r)
}

func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("handlers/login.html"))
	err := tmpl.Execute(rw, nil)
	if err != nil {
		panic(err)
	}
}