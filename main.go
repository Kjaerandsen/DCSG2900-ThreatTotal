package main

import (
	"bytes"
	//"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//"net/url"
	// External
	//webrisk "cloud.google.com/go/webrisk/apiv1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"google.golang.org/api/option"
	//webriskpb "google.golang.org/genproto/googleapis/cloud/webrisk/v1"
	//"google.golang.org/api/webrisk/v1"
	//"google.golang.org/api/option"
)

func callAlienVaultAPI(url string) (respone string){
	content, err := ioutil.ReadFile("./APIKey/OTXapikey.txt")
		if err != nil {
			log.Fatal(err)
		}

		// Convert []byte to string and print to screen
		APIKey := string(content)

		getURL := "https://otx.alienvault.com//api/v1/indicators/url/"+url+"/general"

		req, err := http.NewRequest("GET", getURL, nil)
		req.Header.Set("X-OTX-API-KEY", APIKey)

		fmt.Println(req.Header)

		client := &http.Client{}

		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		respone=string(body)

	return
}

func callGoogleApi(url string) (response string){
	
	content, err := ioutil.ReadFile("./APIKey/Apikey.txt")
		if err != nil {
			log.Fatal(err)
		}

		// Convert []byte to string and print to screen
		APIKey := string(content)

		postURL := "https://safebrowsing.googleapis.com/v4/threatMatches:find?key=" + APIKey

		var jsonData = []byte(`
		{
			"client": {
			  "clientId":      "threattotal",
			  "clientVersion": "1.5.2"
			},
			"threatInfo": {
			  "threatTypes":      ["MALWARE", "SOCIAL_ENGINEERING"],
			  "platformTypes":    ["WINDOWS"],
			  "threatEntryTypes": ["URL"],
			  "threatEntries": [
				{"url": "https://` + url + `" },
				{"url": "http://` + url + `"}
			  ]
			}
		}`)

		req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}

		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		//fmt.Println("response Status:", res.Status)
		//fmt.Print("Response Headers:", res.Header)
		body, _ := ioutil.ReadAll(res.Body)
		//fmt.Println("response Body:", string(body))
		response = string(body)

		return 
}

func main() {
	r := gin.Default()

	r.Use(cors.Default())

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
		fmt.Println(c.Query("url"))
		fmt.Println(c.Query("hash"))
		c.JSON(http.StatusOK, `"exampleJsonData":[
			{"name":"John", "occupation":"Chef"},
			{"firstName":"Jane", "occupation":"Doctor"}
		]`)
	})

	// Upload a file TODO
	// figure out routing here, where are we supposted to have/deliver a file?
	// do we make a new route that says "search" instead? discuss this tomorrow
	// https://github.com/gin-gonic/gin#single-file
	r.POST("/uploadFile", func(c *gin.Context) {
		// for a single file
		file, _ := c.FormFile("Filename")
		log.Println(file.Filename)

		// upload file to the specific destination
		c.SaveUploadedFile(file, "/investigate")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	/**
	* Function should gather DATA from public intelligence sources
	* Implementing functionality for OTX, SafeBrowser API.
	*
	 */

	//GOLANG API STUFF:

	r.GET("/public-intelligence", func(c *gin.Context) {
		fmt.Println(c.Query("url"))

		url := c.Query("url")

		safebrowserResponse := callGoogleApi(url)

		otxAlienVaultRespone := callAlienVaultAPI(url)

		fmt.Println("safebrowser response::",safebrowserResponse)

		fmt.Println("ALIENVAULT RESPONSE:::", otxAlienVaultRespone)
		
		

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
