package main

import (
	"dcsg2900-threattotal/api"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// External
	//webrisk "cloud.google.com/go/webrisk/apiv1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"google.golang.org/grpc/status"
	//"google.golang.org/api/option"
	//webriskpb "google.golang.org/genproto/googleapis/cloud/webrisk/v1"
	//"google.golang.org/api/webrisk/v1"
	//"google.golang.org/api/option"
)

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
		c.JSON(http.StatusOK, `[{"id":1,"sourceName":"Threat Total","content":"Unsafe: potentially unwanted software.","tags":["PUA","Windows","Social Engineering","URL"],"description":"Potentially unwanted software, might be used for lorem ipsum dolor sin amet.","status":"Potentially unsafe"},{"id":2,"sourceName":"Google safebrowsing","content":"Unsafe: Malware.","tags":["Malware","Windows","URL"],"description":"Malware found at he location, might be used for lorem ipsum dolor sin amet.","status":"Risk"},{"id":3,"sourceName":"Source 3","content":"Safe: No known risks at this location.","tags":["URL","Safe"],"description":"No known risks at this location. The data source has no information on this url.","status":"Safe"}]`)
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

	/**
	* Function should gather DATA from public intelligence sources
	* Implementing functionality for OTX, SafeBrowser API.
	*
	 */

	//GOLANG API STUFF:

	r.GET("/public-intelligence", func(c *gin.Context) {
		fmt.Println(c.Query("url"))

		url := c.Query("url")

		//Google

		safebrowserResponse := api.CallGoogleApi(url)

		//Alienvault

		otxAlienVaultRespone := api.CallAlienVaultAPI(url)

		fmt.Println("safebrowser response::", safebrowserResponse)

		fmt.Println("ALIENVAULT RESPONSE:::", otxAlienVaultRespone)

		filehash := "a7a665a695ec3c0f862a0d762ad55aff6ce6014359647e7c7f7e3c4dc3be81b7"

		filehashAV := api.LookUpFileHashAlienVault(filehash)

		fmt.Println("AlienVAULT FILEHASH LOOKUP::::::::::", filehashAV)

		//Hybrid Analysis:

		filehashHybrid := "77682670694bb1ab1a48091d83672c9005431b6fc731d6c6deb466a16081f4d1"

		ResultHybridA := api.CheckFileHashHybridAnalysis(filehashHybrid)

		fmt.Println("\n\n\n\n\n HYBRID ANALYSIS!!!!::::::::!!!\n\n\n", ResultHybridA)

		HybridTestURL := "https://testsafebrowsing.appspot.com/s/malware.html"

		ResultURLHybridA := api.CheckURLHybridAnalyis(HybridTestURL)

		fmt.Println("\n\n\n\n\n HYBRID URL:\n\n", ResultURLHybridA)

	})

	log.Fatal(r.Run(":8081"))
}
