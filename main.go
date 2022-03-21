package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	//"net/http"

	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	/*
		TODO SEE
		Perhaps we need a routing to "search" for searching domains, url or file hashes
		then we have another routing for "upload", where we upload files from local machine, and send that

	*/
	r.POST("/", func(c *gin.Context) {
		var outputData []byte
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			http.Error(c.Writer, "Failed to read request", http.StatusInternalServerError)
		}

		var test map[string]interface{}
		err = json.Unmarshal(jsonData, &test)
		if err != nil { // Handled error
			http.Error(c.Writer, "Failed to unmarshal data", http.StatusInternalServerError)
		}
		fmt.Println(test)
		if test["inputText"] == "ntnu.no" {
			outputData, err = json.Marshal("YESYESYESYESYES")
			if err != nil {
				http.Error(c.Writer, "Failed to marshal data", http.StatusInternalServerError)
			}
		} else {
			outputData, err = json.Marshal("NONONONONONO")
			if err != nil {
				http.Error(c.Writer, "Invalid format, please enter a valid domain", http.StatusForbidden)
			}
		}

		c.Data(http.StatusOK, "application/json", outputData)
	})

	r.GET("/result", func(c *gin.Context) {
		fmt.Println(c.Query("inputFile"))
	})

	// Upload a file TODO
	// figure out routing here, where are we supposted to have/deliver a file?
	// do we make a new route that says "search" instead? discuss this tomorrow
	// https://github.com/gin-gonic/gin#single-file
	r.POST("/upload", func(c *gin.Context) {

		// for a single file
		file, _ := c.FormFile("inputFile")
		log.Println(file.Filename)

		// upload file to the specific destination
		c.SaveUploadedFile(file, "/result")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
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
