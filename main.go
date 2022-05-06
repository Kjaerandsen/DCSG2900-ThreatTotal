package main

import (
	"bytes"
	"context"
	"dcsg2900-threattotal/api"
	"dcsg2900-threattotal/auth"
	"dcsg2900-threattotal/storage"
	"dcsg2900-threattotal/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"sync"
	"dcsg2900-threattotal/logs"

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

// Initialize global variables
func init() {
	var err error

	utils.Ctx = context.Background()

	utils.Config = oauth2.Config{
		ClientID:     os.Getenv("clientId"),
		ClientSecret: os.Getenv("clientSecret"),
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

	oidcConfig := &oidc.Config{
		ClientID: utils.Config.ClientID,
	}

	utils.Verifier = utils.Provider.Verifier(oidcConfig)

	RedisPool := storage.InitPool()
	utils.Conn = RedisPool.Get()
}

func main() {

	r := gin.Default()

	r.Use(cors.Default())

	//_, _ = auth.CodeToToken("")

	// move to init function?

	r.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, hello world, gin.H{
		//	"isSelected": true,
		log.Println("Messsage")
	})

	r.GET("/url-testing", func(c *gin.Context) {

		url := c.Query("url")
		lng := c.Query("lng")

		var wg sync.WaitGroup
		var responseData [4]utils.FrontendResponse2

		if lng != "no" {
			fmt.Println("Language english")
		}

		var p, VirusTotal, urlscanio, alienvault *utils.FrontendResponse2
		p = &responseData[0]
		VirusTotal = &responseData[1]
		urlscanio = &responseData[2]
		alienvault = &responseData[3]

		fmt.Println(url)

		wg.Add(3)
		go api.TestGoGoogleUrl(url, p, &wg)
		go api.TestHybridAnalyisUrl(url, VirusTotal, urlscanio, &wg)
		go api.TestAlienVaultUrl(url, alienvault, &wg)
		wg.Wait()

		//responseData2 := FR122(responseData[:])
		var resultResponse utils.ResultFrontendResponse

		resultResponse.FrontendResponse = responseData[:]

		var setResults *utils.ResultFrontendResponse
		setResults = &resultResponse

		utils.SetResultURL(setResults, len(responseData))

		URLint, err := json.Marshal(resultResponse)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("WHERE IS MY CONTENT 1", responseData)
		//fmt.Println("WHERE IS MY CONTENT 2", responseData2)

		c.Data(http.StatusOK, "application/json", URLint)

	})

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

	r.GET("/login", func(c *gin.Context) {
		code := c.Query("code")
		authenticated, hash := auth.Authenticate(code, "")
		if authenticated {
			fmt.Println("hash is: ", hash)
			c.JSON(http.StatusOK, gin.H{"hash": hash})
		} else {
			http.Error(c.Writer, "Failed authenticating with the code.", http.StatusUnauthorized)
		}
	})

	r.DELETE("/login", func(c *gin.Context) {
		hash := c.Query("userAuth")
		err := auth.Logout(hash)

		if !err {
			http.Error(c.Writer, "Failed logging out the user. Either the user was not logged in, or the login was expired.", http.StatusUnauthorized)
		} else {
			c.JSON(http.StatusOK, gin.H{"Logoutstatus": "Successfull"})
		}
	})

	r.GET("/auth", func(c *gin.Context) {
		auth2 := c.Query("auth")
		authenticated, _ := auth.Authenticate("", auth2)
		if authenticated {
			c.JSON(http.StatusOK, gin.H{"yes": "You are authenticated"})
		} else {
			http.Error(c.Writer, "Authentication is invalid or expired, please try to login again.", http.StatusUnauthorized)
		}
	})

	// TODO: Upload a file
	// figure out routing here, where are we supposted to have/deliver a file?
	// do we make a new route that says "search" instead? discuss this tomorrow
	// inspiration from https://github.com/dutchcoders/go-virustotal/blob/24cc8e6fa329f020c70a3b32330b5743f1ba7971/virustotal.go#L305

	r.POST("/upload", func(c *gin.Context) {

		log.Println("Fileupload worked")
		logging.Loginfo("Fileupload worked")
		

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
			logging.Logerror(err)
		}
		// close writer
		err = writer.Close()

		if err != nil {
			log.Println(err)
			logging.Logerror(err)
		}

		// prepare request towards API
		req, err := http.NewRequest("POST", uri, body)

		if err != nil {
			log.Println(err)
			logging.Logerror(err)
		}

		content, err := ioutil.ReadFile("./APIKey/virusTotal.txt")
		if err != nil {
			//log.Fatal(err)
			fmt.Println(err)
			logging.Logerror(err)
		}

		APIKey := string(content)

		req.Header.Add("X-Apikey", APIKey)
		// error handle here, user should not be able to send requests without api key

		// dynamically set content type, based on the formdata writer
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// perform the prepared API request
		res, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Println(err)
			logging.Logerror(err)
		}

		defer res.Body.Close()

		// så lenge status 200

		// read the response
		contents, _ := ioutil.ReadAll(res.Body)

		var jsonResponse utils.VirusTotalUploadID

		unmarshalledID := json.Unmarshal(contents, &jsonResponse)

		if unmarshalledID != nil {
			log.Println(unmarshalledID)
			logging.Logerror(err)
		}

		encodedID := jsonResponse.Data.ID

		// decode provided values for virustotal report
		decode64, err := base64.RawStdEncoding.DecodeString(encodedID)
		log.Println("decoded here")
		log.Println(string(decode64))

		// extract ID for virustotal report
		trimID := strings.Split((string(decode64)), ":")
		log.Println("TRIMMED")
		log.Println(trimID[0])
		if err != nil {
			log.Println(err)
		}
		// return json object ID
		c.JSON(http.StatusOK, gin.H{"id": trimID[0]})
		// handle error
	})

	r.GET("/upload", func(c *gin.Context) {

		// VT key has been added. REMEMBER TO DEACTIVATE AND CHANGE BEFORE FINAL RELEASE.
		// prepare request towards API

		content, err := ioutil.ReadFile("./APIKey/virusTotal.txt")
		if err != nil {
			//log.Fatal(err)
			fmt.Println(err)
			logging.Logerror(err)
		}
		// Convert []byte to string and print to screen
		APIKey := string(content)
		// remember to change api key, and reference it to a file instead
		// as well as deactivate the key from the account, as it's leaked.

		// fetch based on url parameter, file = id
		id := c.Query("file")
		log.Println(id)
		if id == "" {
			log.Println("error, ID is empty")
			logging.Logerrorinfo("Error, ID is empty - Upload")
		}

		url := fmt.Sprintf("https://www.virustotal.com/api/v3/files/%s", id)
		log.Println(url)
		log.Println(id)

		vtReq, _ := http.NewRequest("GET", url, nil)

		vtReq.Header.Add("Accept", "application/json")

		vtReq.Header.Add("X-Apikey", APIKey)

		vtRes, _ := http.DefaultClient.Do(vtReq)

		//defer vtRes.Body.Close()

		vtBody, _ := ioutil.ReadAll(vtRes.Body)

		log.Println(string(vtBody))

		var vtResponse utils.FileUploadData2

		unmarshalledBody := json.Unmarshal(vtBody, &vtResponse)

		if unmarshalledBody != nil {
			log.Println(unmarshalledBody)
		}

		/*
			test2 := reflect.ValueOf(vtResponse.Data.Attributes.LastAnalysisResults)
			lengthOfTest2 := test2.NumField() */

		var test3 = make([]utils.FrontendResponse4, len(vtResponse.Data.Attributes.LastAnalysisResults))

		//var test [4]utils.FrontendResponse4

		log.Println("here is the test output we maybe want")
		i := 0

		// TODO for later, remove teststruct, as it's only used to put into totalverdict, later
		var testStruct = make([]utils.FrontendResponse2, len(vtResponse.Data.Attributes.LastAnalysisResults))

		// iterate through results
		for _, val := range vtResponse.Data.Attributes.LastAnalysisResults {
			//log.Printf("testing, %s, %s", key, val)
			// initialize struct
			test3[i] = val
			// print
			log.Println(test3[i])

			// save engine name
			testStruct[i].ID = i + 1
			testStruct[i].SourceName = test3[i].EngineName
			// resolution of AV
			testStruct[i].EN.Status = test3[i].Category

			testStruct[i].EN.Content = vtResponse.Data.Attributes.MeaningfulName
			testStruct[i].EN.Description = vtResponse.Data.Attributes.Magic
			testStruct[i].EN.Tags = vtResponse.Data.Attributes.TypeTag

			//testStruct.EN.Description =

			// can also display the total status (last analysis stats)
			// this is an int ^^ so cant fill it in frontendresponse2
			// question is, do we do it here or later

			i++
		}
		log.Println(testStruct)

		var totalVerdict utils.ResultFrontendResponse
		totalVerdict.FrontendResponse = testStruct
		// IMPORTANT TODO, FIGURE ROUTING

		// Possible to add more cases in the future, for more accurate assessements
		// very realisitc that we need more cases, too narrow for accurate results per now.
		if vtResponse.Data.Attributes.LastAnalysisStats.Malicious == 0 && vtResponse.Data.Attributes.LastAnalysisStats.Suspicious == 0 {
			totalVerdict.EN.Result = "File is safe."
			// osv totalVerdict.EN.Result = fmt.Sprintf("File is considered safe", x av y)
		} else if vtResponse.Data.Attributes.TotalVotes.Malicious > 0 && vtResponse.Data.Attributes.LastAnalysisStats.Suspicious >= 0 {
			totalVerdict.EN.Result = "File has malicious indicators, consider escalating. "
		} else if vtResponse.Data.Attributes.LastAnalysisStats.Harmless > 0 && vtResponse.Data.Attributes.LastAnalysisStats.Malicious == 0 {
			totalVerdict.EN.Result = "File has been confirmed benign."
		} else {
			totalVerdict.EN.Result = "File is suspicious."
		}

		var engines int = len(vtResponse.Data.Attributes.LastAnalysisResults)

		utils.SetResultFile(&totalVerdict, engines)

		log.Println("look here")
		log.Println(totalVerdict)

		fmt.Println(totalVerdict)

		// return here

		fileInt, err := json.Marshal(totalVerdict)
		if err != nil {
			fmt.Println(err)
			//c.Data(http.StatusInternalServerError, "application/json", )
		}

		//fmt.Println("WHERE IS MY CONTENT", responseData)

		c.Data(http.StatusOK, "application/json", fileInt)

		// sort list based on what is malicious
		//log.Print(test3)

		// hent resultat via cache

		// total votes feltet virker relevant
		// LAST ANALYSIS STATS - MALICIOUS

	})

	/**
	* Function should gather DATA from public intelligence sources
	* Implementing functionality for OTX, SafeBrowser API.
	*
	 */

	//GOLANG API STUFF:

	/**
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

	})*/

	r.GET("/url-intelligence", func(c *gin.Context) {
		hash := c.Query("userAuth")
		authenticated, _ := auth.Authenticate("", hash)
		if !authenticated {
			c.JSON(http.StatusUnauthorized, gin.H{"authenticated": "You are not authenticated. User login is invalid."})
		} else {
			api.UrlIntelligence(c)
		}
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

		var wg sync.WaitGroup
		hash := c.Query("hash")

		var responseData [2]utils.FrontendResponse2

		var hybridApointer, AlienVaultpointer *utils.FrontendResponse2

		hybridApointer = &responseData[0]
		AlienVaultpointer = &responseData[1]

		wg.Add(2)
		go api.CallHybridAnalysisHash(hash, hybridApointer, &wg)
		go api.CallAlienVaultHash(hash, AlienVaultpointer, &wg)
		wg.Wait()

		var resultResponse utils.ResultFrontendResponse

		resultResponse.FrontendResponse = responseData[:]
		var resultPointer = &resultResponse

		utils.SetResultHash(resultPointer, len(responseData))

		Hashint, err := json.Marshal(resultResponse)
		if err != nil {
			fmt.Println(err)
			logging.Logerror(err)
			//c.Data(http.StatusInternalServerError, "application/json", )
		}

		//fmt.Println("WHERE IS MY CONTENT", responseData)

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
