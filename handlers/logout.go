package handlers

import (
	"html/template"
	"net/http"
)

// Logout is a handler that clears a logged in cookie
type Logout struct {
	Name   string
	Path   string
	Domain string
	Next   http.Handler
}

func (l Logout) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     l.Name,
		Value:    "",
		Domain:   l.Domain,
		Path:     l.Path,
		MaxAge:   0,
		HttpOnly: true,
	}
	http.SetCookie(rw, &cookie)
	l.Next.ServeHTTP(rw, r)
}

func LogoutHandler(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("handlers/logout.html"))
	err := tmpl.Execute(rw, nil)
	if err != nil {
		panic(err)
	}
}