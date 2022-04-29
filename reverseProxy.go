package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

//// Inspiration taken from: https://blog.joshsoftware.com/2021/05/25/simple-and-powerful-reverseproxy-in-go/

// NewProxy takes target host and creates a reverse proxy
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	URL, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(URL), nil
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	// All traffic to "http://localhost:8081" goes through this proxy-server at "http://localhost:8080" first
	proxy, err := NewProxy("http://localhost:8081")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", ProxyRequestHandler(proxy))
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
