package main

import (
	"bytes"
	"context"
	"dcsg2900-threattotal/api"
	"dcsg2900-threattotal/storage"
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	// External
	//webrisk "cloud.google.com/go/webrisk/apiv1"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	//"google.golang.org/grpc/status"
	//"google.golang.org/api/option"
	//webriskpb "google.golang.org/genproto/googleapis/cloud/webrisk/v1"
	//"google.golang.org/api/webrisk/v1"
	//"google.golang.org/api/option"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	var err error

	utils.Ctx = context.Background()

	utils.Config = oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://auth.dataporten.no/oauth/authorization",
			TokenURL: "https://auth.dataporten.no/oauth/token",
		},
		RedirectURL: "http://localhost:3000",
		Scopes:      []string{oidc.ScopeOpenID, "email"},
	}

	// Initializing authentication connection
	utils.Provider, err = oidc.NewProvider(utils.Ctx, "https://auth.dataporten.no")
	if err != nil {
		log.Fatal(err)
	}

	// move to init function?
	RedisPool := storage.InitPool()
	utils.Conn = RedisPool.Get()

	r.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, hello world, gin.H{
		//	"isSelected": true,
		log.Println("Messsage")
	})

	r.GET("/url-testing", func(c *gin.Context) {
		api.UrlIntelligence(c)
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

	// TODO: Upload a file
	// figure out routing here, where are we supposted to have/deliver a file?
	// do we make a new route that says "search" instead? discuss this tomorrow
	// inspiration from https://github.com/dutchcoders/go-virustotal/blob/24cc8e6fa329f020c70a3b32330b5743f1ba7971/virustotal.go#L305

	r.POST("/upload", func(c *gin.Context) {

		log.Println("Fileupload worked")

		uri := "https://www.virustotal.com/api/v3/files"
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// fetch the file contents
		file2, _ := c.FormFile("file")
		// open the file
		file3, _ := file2.Open()

		// use file contents to fetch file name, associate it with the "file" form header.
		part, err := writer.CreateFormFile("file", file2.Filename)

		if err != nil {
			log.Println(err)
		}
		// copy file locally
		_, err = io.Copy(part, file3)

		if err != nil {
			log.Println(err)
		}
		// close writer
		err = writer.Close()

		if err != nil {
			log.Println(err)
		}

		// prepare request towards API
		req, err := http.NewRequest("POST", uri, body)

		if err != nil {
			log.Println(err)
		}

		// dynamically set content type, based on the formdata writer
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// remember to change api key, and reference it to a file instead
		// as well as deactivate the key from the account, as it's leaked.
		req.Header.Add("x-apikey", "4062c07a4340e4f8fe5f647412ef936d99d53aa793e1cebfc4b31e43ae801ed0")

		// perform the prepared API request
		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()

		// read the response
		contents, _ := ioutil.ReadAll(res.Body)

		log.Println(res)

		log.Println(string(contents))

	})

	/**
	* Function should gather DATA from public intelligence sources
	* Implementing functionality for OTX, SafeBrowser API.
	*
	 */

	//GOLANG API STUFF:

	r.GET("/public-intelligence", func(c *gin.Context) {
		//fmt.Println(c.Query("url"))

		//url := c.Query("url")

		//Google

		//safebrowserResponse := api.CallGoogleUrl(url)

		//fmt.Println("safebrowser response::", safebrowserResponse.Status)

		//Alienvault
		//var otxAlienVaultRespone [1]utils.FrontendResponse

		//otxAlienVaultRespone[0] = api.CallAlienVaultUrl(url)

		//fmt.Println("safebrowser response::", safebrowserResponse)

		//fmt.Println("ALIENVAULT RESPONSE:::", otxAlienVaultRespone[0].Status)

		filehash := "a7a665a695ec3c0f862a0d762ad55aff6ce6014359647e7c7f7e3c4dc3be81b7"

		filehashAV := api.CallAlienVaultHash(filehash)

		fmt.Println("AlienVAULT FILEHASH LOOKUP::::::::::", filehashAV)

		//Hybrid Analysis:

		//filehashHybrid := "77682670694bb1ab1a48091d83672c9005431b6fc731d6c6deb466a16081f4d1"

		//ResultHybridA := api.CallHybridAnalysisHash(filehashHybrid)

		//fmt.Println("\n\n\n\n\n HYBRID ANALYSIS!!!!::::::::!!!\n\n\n", ResultHybridA)

		/**

		HybridTestURL := "https://testsafebrowsing.appspot.com/s/malware.html"

		ResultURLHybridA := api.CallHybridAnalyisUrl(HybridTestURL)

		fmt.Println("\n\n\n\n\n HYBRID URL:\n\n", ResultURLHybridA)
		*/
	})

	r.GET("/url-intelligence", func(c *gin.Context) {

		url := c.Query("url")
		//lng := c.Query("lng")

		//if lng != "no" {
		//	fmt.Println("Language english")
		//}

		fmt.Println(url)

		var responseData [4]utils.FrontendResponse

		responseData[0] = api.CallGoogleUrl(url)

		responseData[1], responseData[2] = api.CallHybridAnalyisUrl(url)

		responseData[3] = api.CallAlienVaultUrl(url)

		responseData2 := FR122(responseData[:])

		URLint, err := json.Marshal(responseData2)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("WHERE IS MY CONTENT 1", responseData)
		fmt.Println("WHERE IS MY CONTENT 2", responseData2)

		c.Data(http.StatusOK, "application/json", URLint)

		/* Backup of old code

		URLint, err := json.Marshal(responseData)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("WHERE IS MY CONTENT", responseData)

		c.Data(http.StatusOK, "application/json", URLint)
		*/
	})

	r.GET("/url-intelligence2", func(c *gin.Context) {

		responseData := make([]utils.FrontendResponse2, 2)

		responseData[0].SourceName = "Test"
		responseData[1].SourceName = "test2"
		responseData[0].ID = 0
		responseData[1].ID = 1
		responseData[0].EN.Status = "risk1"
		responseData[0].EN.Description = "risk1"
		responseData[0].EN.Content = "risk1"
		responseData[0].NO.Description = "riskno1"
		responseData[0].NO.Status = "riskno1"
		responseData[0].NO.Content = "riskno1"
		responseData[1].EN.Status = "risk2"
		responseData[1].EN.Description = "risk2"
		responseData[1].EN.Content = "risk2"
		responseData[1].NO.Description = "riskno2"
		responseData[1].NO.Status = "riskno2"
		responseData[1].NO.Content = "riskno2"

		URLint, err := json.Marshal(responseData)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("WHERE IS MY CONTENT 2", responseData)

		c.Data(http.StatusOK, "application/json", URLint)

		/* Backup of old code

		URLint, err := json.Marshal(responseData)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("WHERE IS MY CONTENT", responseData)

		c.Data(http.StatusOK, "application/json", URLint)
		*/
	})

	r.GET("/hash-intelligence", func(c *gin.Context) {

		hash := c.Query("hash")

		var Hashint []byte
		var responseData [2]utils.FrontendResponse

		value, err := utils.Conn.Do("GET", hash)
		if value == nil {
			if err != nil {
				fmt.Println("Error:" + err.Error())
			}
			fmt.Println("No Cache hit")
			responseData[0] = api.CallHybridAnalysisHash(hash)

			responseData[1] = api.CallAlienVaultHash(hash)

			Hashint, err = json.Marshal(responseData)
			if err != nil {
				fmt.Println(err)
			}

			response, err := utils.Conn.Do("SETEX", hash, 60, Hashint)
			if err != nil {
				fmt.Println("Error adding data to redis:" + err.Error())
			}
			// Print the response to adding the data (should be "OK"
			fmt.Println(response)

			fmt.Println("WHERE IS MY CONTENT", responseData)
			// Cache hit
		} else {
			fmt.Println("Cache hit")
			responseBytes, err := json.Marshal(value)
			if err != nil {
				fmt.Println("Error handling redis response:" + err.Error())
			}
			err = json.Unmarshal(responseBytes, &Hashint)
			if err != nil {
				fmt.Println("Error handling redis response:" + err.Error())
			}
			fmt.Println("Value is:\n", Hashint)
		}

		c.Data(http.StatusOK, "application/json", Hashint)

	})

	log.Fatal(r.Run(":8081"))
	// These don't do anything, and can't be placed above the line above as they stop the connections prematurely then.
	/*
		conn.Close()      // Close the connection
		redisPool.Close() // Close the pool
	*/
}

// Temporary helper function to create translations from the input

func FR122(input []utils.FrontendResponse) (output []utils.FrontendResponse2) {
	length := len(input)

	fmt.Println("input: ", input)

	output = make([]utils.FrontendResponse2, 4)

	for i := 0; i < length; i++ {
		output[i].ID = input[i].ID
		output[i].SourceName = input[i].SourceName
		output[i].EN.Content = input[i].Content
		output[i].EN.Description = input[i].Description
		output[i].EN.Status = input[i].Status
		//output[i].EN.Tags = input[i].Tags
		output[i].NO.Content = input[i].Content + "norsk"
		output[i].NO.Description = input[i].Description + "norsk"
		output[i].NO.Status = input[i].Status + "norsk"
		//output[i].NO.Tags = input[i].Tags
	}

	fmt.Println()

	return output
}
