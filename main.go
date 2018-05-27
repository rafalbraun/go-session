package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/brafales/go-session/handlers"
)

func main() {
	cookieName := os.Getenv("COOKIE_NAME")
	cookieValue := os.Getenv("COOKIE_VALUE")
	cookieDomain := os.Getenv("COOKIE_DOMAIN")
	cookiePath := os.Getenv("COOKIE_PATH")
	cookieDuration := os.Getenv("COOKIE_DURATION")

	cookieDurationInteger, err := strconv.Atoi(cookieDuration)

	if err != nil {
		panic(err)
	}

	loginHandler := handlers.Login{Name: cookieName, Domain: cookieDomain, Path: cookiePath, Value: cookieValue, Duration: time.Hour * time.Duration(cookieDurationInteger)}
	logoutHandler := handlers.Logout{Name: cookieName, Domain: cookieDomain, Path: cookiePath}

	http.Handle("/login", loginHandler)
	http.Handle("/logout", logoutHandler)

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
	}
}
