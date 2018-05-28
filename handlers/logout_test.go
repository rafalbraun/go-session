package handlers_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brafales/go-session/handlers"
)

func TestLogout(t *testing.T) {
	expBody := []byte("test!")
	now := time.Now().UTC()
	expire := func() time.Time {
		return now
	}

	logoutHandler := handlers.Logout{
		Name:    "session",
		Path:    "/",
		Domain:  "test.com",
		Expires: expire,
		Next: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Write(expBody)
		}),
	}

	testReq, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Failed to create test request: %v", err)
	}

	recorder := httptest.NewRecorder()
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

	if !compareCookieTimes(cookie.Expires.UTC(), now.UTC()) {
		t.Errorf("Cookie has the wrong expiry time. Expected %v, got %v", now.UTC(), cookie.Expires.UTC())
	}
}
