package websiteracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns faster site", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		if err != nil {
			t.Error("should not get error unless there is a timeout")
		}
	})

	t.Run("returns error if server does not respond within 10 seconds", func(t *testing.T) {
		firstServer := makeDelayedServer(11 * time.Second)
		secondServer := makeDelayedServer(11 * time.Second)

		defer firstServer.Close()
		defer secondServer.Close()

		_, err := Racer(firstServer.URL, secondServer.URL)
		if err == nil {
			t.Error("expected an error on timeout")
		}
	})

}
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
