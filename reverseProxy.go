package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

//// Inspiration taken from: https://le-gall.bzh/post/go/a-reverse-proxy-in-go-using-gin/

// proxy is a reverse proxy which routes traffic to "remote" through localhost:8080
func proxy(c *gin.Context) {
	remote, err := url.Parse("http://localhost:3000")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	// Sets request parameters
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}

	// Start proxy
	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	r := gin.Default()

	// Define catch all path
	r.Any("/*proxyPath", proxy)

	// Start server on port 8080
	log.Fatal(r.Run(":8080"))
}
