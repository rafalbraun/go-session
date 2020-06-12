package main

import (
	"go-session/handlers"
	"net/http"
	"time"
)

func main() {
	cookieName := "session"
	cookieValue := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	cookieDomain := ""
	cookiePath := "/"
	cookieDurationInteger := 520
	cookieExpires := time.Now().Add(120 * time.Second)

	loginHandler := handlers.Login{Name: cookieName,
		Domain: cookieDomain,
		Path:   cookiePath,
		Value:  cookieValue,
		MaxAge: cookieDurationInteger,
		Expires: cookieExpires,
		Next:   http.HandlerFunc(handlers.LoginHandler),
	}

	logoutHandler := handlers.Logout{
		Name:   cookieName,
		Domain: cookieDomain,
		Path:   cookiePath,
		Next:   http.HandlerFunc(handlers.LogoutHandler),
	}

	indexHandler := handlers.Index{
		IsLoggedIn: false,
	}

	http.Handle("/index", indexHandler)
	http.Handle("/login", loginHandler)
	http.Handle("/logout", logoutHandler)

	address := "localhost:8080"

	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
	}
}
