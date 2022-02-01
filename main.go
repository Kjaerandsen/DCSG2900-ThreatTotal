package main

import (
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
    r.LoadHTMLFiles("./Tailwind/html/index.html")
	

	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", gin.H{

		})
	})
	log.Fatal(r.Run())
	}