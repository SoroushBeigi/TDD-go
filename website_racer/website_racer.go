package websiteracer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, error error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %q and %q", a, b)
	}

}

// struct is used becauses no memory allocation is needed for structcs,
// compared to even a small data type like bool
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
