package main

import (
	"fmt"
	"log"
	"net/http"

	//"net/http"
	"github.com/gin-gonic/gin"
)

/**func main() {

	http.Handle("/", http.FileServer(http.Dir("./Tailwind/html/")))
	log.Fatal(http.ListenAndServe(":80", nil))
}
*/

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./Tailwind/html/*.html")
	// CSS files
	r.Static("/dist", "./Tailwind/dist")
	// Images
	r.Static("/img", "./Tailwind/img")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Generic get request, gets parsed in the RequestHandler function
	r.GET("/:url", func(c *gin.Context) {
		url := c.Param("url")
		RequestHandler(url, c)
	})

	log.Fatal(r.Run(":80"))
}

func RequestHandler(url string, c *gin.Context) {
	fmt.Println("URL IS: " + url + ".")
	if url == "/" {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	} else {
		// TODO: Add a validity test here for the url

		// TODO: Remove trailing slashes and .*

		// TODO: Implement templating? Gin has built in template functionality

		// Display the webpage
		c.HTML(http.StatusOK, url, gin.H{})
	}
}
