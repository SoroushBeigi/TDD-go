package websiteracer

import (
	"net/http"
)

func Racer(a, b string) (winner string) {

	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

//struct is used becauses no memory allocation is needed for structcs,
//compared to even a small data type like bool
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
