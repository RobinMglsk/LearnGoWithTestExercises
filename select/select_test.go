package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T){
	t.Run("returns the fastest url", func(t *testing.T){
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)
	
		defer slowServer.Close()
		defer fastServer.Close()
	
		slowURL := slowServer.URL
		fastURL := fastServer.URL
	
		want := fastURL
		got, _ := Racer(slowURL, fastURL)
	
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T){
		serverA := makeDelayedServer(11 * time.Microsecond)
		serverB := makeDelayedServer(12 * time.Microsecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 10 * time.Microsecond)

		if err == nil {
			t.Error("expected an error but did't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}