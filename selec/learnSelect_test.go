package selec

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speed of server, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(time.Second * 20)
		fastServer := makeDelayedServer(time.Second * 0)

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL
		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, _ := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(25 * time.Second)
		serverB := makeDelayedServer(21 * time.Second)

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 20*time.Second)

		if err == nil {
			t.Error("expected an error but didn't get")
		}

		defer serverA.Close()
		defer serverB.Close()
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

// func TestRacer(t *testing.T) {
// 	slowUrl := "http://www.facebook.com"
// 	fastUrl := "http://www.quii.dev"
//
// 	want := fastUrl
// 	got := Racer(slowUrl, fastUrl)
//
// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }
