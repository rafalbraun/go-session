package handlers

import (
	"html/template"
	"net/http"
)

type Index struct {
	IsLoggedIn bool
}

func (l Index) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	msg, err := r.Cookie("session")
	data := Index{IsLoggedIn: false}
	if msg != nil && msg.Value != "" {
		data.IsLoggedIn = true
	}
	tmpl := template.Must(template.ParseFiles("handlers/index.html"))
	err = tmpl.Execute(rw, data)
	if err != nil {
		panic(err)
	}
}
