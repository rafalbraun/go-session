package handlers_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-session/handlers"
)

func TestLogin(t *testing.T) {
	expBody := []byte("test!")

	loginHandler := handlers.Login{
		Name:   "session",
		Value:  "logged in",
		Path:   "/",
		Domain: "test.com",
		MaxAge: 60,
		Next: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Write(expBody)
		}),
	}

	testReq, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Failed to create test request: %v", err)
	}

	recorder := httptest.NewRecorder()
	loginHandler.ServeHTTP(recorder, testReq)
	response := recorder.Result()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Failed to read the response: %v", err)
	}

	if !bytes.Equal(bodyBytes, expBody) {
		t.Errorf("Unexpected response: %v", err)
	}

	cookies := response.Cookies()

	if len(cookies) != 1 {
		t.Error("Response returned more than one cookie")
	}

	cookie := cookies[0]

	if cookie.Value != "logged in" {
		t.Errorf("Cookie has the wrong value. Expected %v, got %v", "logged in", cookie.Value)
	}

	if cookie.Domain != "test.com" {
		t.Errorf("Cookie has the wrong domain. Expected %v, got %v", "test.com", cookie.Domain)
	}

	if cookie.Name != "session" {
		t.Errorf("Cookie has the wrong name. Expected %v, got %v", "session", cookie.Name)
	}

	if cookie.Path != "/" {
		t.Errorf("Cookie has the wrong domain. Expected %v, got %v", "/", cookie.Path)
	}

	if cookie.MaxAge != 60 {
		t.Errorf("Cookie has the wrong max age. Expected %v, got %v", 60, cookie.MaxAge)
	}
}
