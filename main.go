package main

import (
	"log"
	//"net/http"

	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"fmt"
	"encoding/json"
	"io/ioutil"
	//"log"
)

/**func main() {

	http.Handle("/", http.FileServer(http.Dir("./Tailwind/html/")))
	log.Fatal(http.ListenAndServe(":80", nil))
}
*/

func main() {
	r := gin.Default()
	
	r.Use(cors.Default())

	
	
	//r.LoadHTMLGlob("./Tailwind/html/**/*.html")
	// CSS files
	//r.Static("/dist", "./Tailwind/dist")
	// Images
	//r.Static("/img", "./Tailwind/img")

	r.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, hello world, gin.H{
		//	"isSelected": true,
		log.Println("Messsage")
	})

	/**
	r.POST("/searchreputation", func(c *gin.Context){
		//data := c.PostForm("submitted")
		reqData, err := ioutil.ReadAll(c.Request.Body)
		var data interface{}

		err = json.Unmarshal(reqData, &data)
	
		if err!=nil{

		}
		else{

		c.JSON(http.StatusOK, data)
		}
		
	})
	*/

	r.POST("/upload", func(c *gin.Context) {
        var outputData []byte
        jsonData, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            // Handle error
        }
        var test map[string]interface{}
        err = json.Unmarshal(jsonData, &test)
        if err != nil {
            // Handle error
        }
        fmt.Println(test)
        if test["inputText"] == "ntnu.no" {
            outputData, err = json.Marshal("YESYESYESYESYES")
            if err != nil {
                // Handle error
            }
        } else {
            outputData, err = json.Marshal("NONONONONONO")
            if err != nil {
                // Handle error
            }
        }

        c.Data(http.StatusOK, "application/json", outputData)
    })
	/*
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"isSelected": false,
		})
	})

	r.GET("/investigate", func(c *gin.Context) {
		c.HTML(http.StatusOK, "investigate.html", gin.H{})
	})
	

	/*
		// Generic get request, gets parsed in the RequestHandler function
		r.GET("/:url", func(c *gin.Context) {
			url := c.Param("url")
			RequestHandler(url, c)
		})
	*/

	log.Fatal(r.Run(":8081"))
}

/*
func RequestHandler(url string, c *gin.Context) {
	fmt.Println("URL IS: " + url + ".")
	if url == "favicon.ico" {
		return
	}
	// TODO: Add a validity test here for the url
	if url == "upload.html" {
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"isSelected": false,
		})
		return
	}

	// TODO: Remove trailing slashes and .*

	// TODO: Implement templating? Gin has built in template functionality

	// Display the webpage
	c.HTML(http.StatusOK, url, gin.H{})
}
*/
