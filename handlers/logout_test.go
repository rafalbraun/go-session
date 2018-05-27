package handlers_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brafales/go-session/handlers"
)

func TestLogout(t *testing.T) {
	expBody := []byte("test!")

	logoutHandler := handlers.Logout{
		Name:   "session",
		Path:   "/",
		Domain: "test.com",
		Next: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Write(expBody)
		}),
	}

	testReq, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Failed to create test request: %v", err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	recorder := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	logoutHandler.ServeHTTP(recorder, testReq)

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

	if cookie.Value != "" {
		t.Errorf("Cookie has the wrong value. Expected empty string, got %v", cookie.Value)
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
}
