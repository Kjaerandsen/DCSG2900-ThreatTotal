package main

import (
	//"context"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"net/url"
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


func CheckURLHybridAnalyis(URL string) (response string){


/**
	type scanners struct{
		name string
		status string
		error_message string
		progress int
		total int
		positives int
		percent int
		anti_virus_results string
	}
*/

	type expectedOutput struct {
		All struct {
		submission_type string `json:"submission_type"`
		id string `json:"id"`
		sha256 string `json:"sha256"`
		scanners map[int]interface{} `json:"scanners"`
		whitelist map[string]interface{} `json:"whitelist"`
		reports map[string]interface{} `json:"reports"`
		finished bool `json:"finished"`
		}
	}

	fmt.Println("HYBRID URL: ", URL)
	//DENNE FUNKSJONENE KAN SCANNE EN URL MEN DETTE BENYTTER SEG AV VIRUS TOTAL/ DETTE ER KANSKJE EN GOD WORK AROUND FOR Å KUNNE BRUKE VT GRATIS SIDEN Hybrid Analysis har lisens.
	//Problem her kan være at dette må inkomporere en "await - 5-15 sekunder om det ikke er noe cachet result på VirusTotal, fordi den maa kjore ny request.".
	//Titter på dette. 
	//Vi har CAP på 2000 request i timen hos Hybrid Analyis, dette burde vell holde??? - 200 max i minuttet. 
	// https://www.hybrid-analysis.com/docs/api/v2#/Quick%20Scan/post_quick_scan_url Dokumentasjon for dette API endpointet.

	content, err := ioutil.ReadFile("./APIKey/HybridAnalysisAPI.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Convert []byte to string and print to screen
    APIKey := string(content)

    postURL := "https://www.hybrid-analysis.com/api/v2/quick-scan/url"

    data := url.Values{}
    data.Set("scan_type", "all")
    data.Set("url", URL)
    data.Set("no_share_third_party","true")
    data.Set("allow_community_access","false")
    //data.Set("submit_name","")

    req, err := http.NewRequest("POST", postURL, strings.NewReader(data.Encode()))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("api-key", APIKey)
    req.Header.Set("User-Agent", "Falcon Sandbox")

    client := &http.Client{}

    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

	//res.Body.Read("finished") Her skal jeg føre en sjekk som sjekker om "finished = true eller false"
	
	//Hvis denne er false skal den vente 5 sekunder og kjøre requesten på nytt.
	//Eventuelt om det er en måte å ikke close requesten før den er finished???????
	

	//Her kan det sjekkes om VirusTotal - Status er Malicious og om Urlscan.io - status er malicious, suspicious, clean etc. også bare returnere denne responsen. 


    //fmt.Println("response Status:", res.Status)
    //fmt.Print("Response Headers:", res.Header)
    body, err := ioutil.ReadAll(res.Body)
	if err!= nil{
		fmt.Println("Ioutil error:", err)
	}

	//var jsonData map[string]interface{}
	var jsonData expectedOutput

	err = json.Unmarshal(body, &jsonData.All)
	if err!=nil{
		panic(err)
	}

	fmt.Println("\n\nScanners", jsonData.All.id)

	fmt.Println("\n\n\n\n\nEverything", jsonData.All)

	//ttstr := jsonData["scanners"].(map[string]interface{})["VirusTotal"].(string)

	//fmt.Println("\n\n\nTTSTR ER:", ttstr)

	//json.Unmarshal([]byte(body), &jsonData)

	//fmt.Println("Status is:", jsonData)

    //fmt.Println("response Body:", string(body))
    response = string(body)

    return
	
}



func CheckFileHashHybridAnalysis(hash string) (response string){

	//API dokumentasjon https://www.hybrid-analysis.com/docs/api/v2#/Search/post_search_hash
	
	content, err := ioutil.ReadFile("./APIKey/HybridAnalysisAPI.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	APIKey := string(content)

	postURL := "https://www.hybrid-analysis.com/api/v2/search/hash"

	data := url.Values{}
    data.Set("hash", hash)
    

	req, err := http.NewRequest("POST", postURL, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("api-key", APIKey)
	req.Header.Set("User-Agent", "Falcon Sandbox")

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

func LookUpFileHashAlienVault(hash string) (response string){

	content, err := ioutil.ReadFile("./APIKey/OTXapikey.txt")
	if err != nil {
		log.Fatal(err)
	}
	APIKey := string(content)

	getURL := "https://otx.alienvault.com//api/v1/indicators/file/" + hash + "/analysis"

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

	response = string(body)
	
	//HER KAN VI SJEKKE OM "VERDICT feltet er" MALICIOUS, SUSPICIOUS ELLER NOE ANNET. OG Bare returnere det. 
	return
}


func callAlienVaultAPI(url string) (respone string) {

	//DENNE FUNKSJONEN KAN UTARBEIDES TIL Å BARE RETURNERE MALCICIOUS / SUSPCIOUS OM DET BEFINNER SEG NEVNT I NOEN PULSEES (Problemet her er at ting som er OK kan være i pulse... Må tenke litt her)
	content, err := ioutil.ReadFile("./APIKey/OTXapikey.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	APIKey := string(content)

	getURL := "https://otx.alienvault.com//api/v1/indicators/url/" + url + "/general"

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

	respone = string(body)

	return
}

func callGoogleApi(url string) (response string) {

	//Google API returnerer [] om den ikke kjenner til domenet / URL. Kan bruke dette til å avgjøre om det er malicious eller ikke. 

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

		safebrowserResponse := callGoogleApi(url)

		//Alienvault

		otxAlienVaultRespone := callAlienVaultAPI(url)

		fmt.Println("safebrowser response::", safebrowserResponse)

		fmt.Println("ALIENVAULT RESPONSE:::", otxAlienVaultRespone)

		filehash:= "a7a665a695ec3c0f862a0d762ad55aff6ce6014359647e7c7f7e3c4dc3be81b7"

		filehashAV := LookUpFileHashAlienVault(filehash)

		fmt.Println("AlienVAULT FILEHASH LOOKUP::::::::::",filehashAV)

		//Hybrid Analysis: 

		filehashHybrid:="77682670694bb1ab1a48091d83672c9005431b6fc731d6c6deb466a16081f4d1"

		ResultHybridA:=CheckFileHashHybridAnalysis(filehashHybrid);

		fmt.Println("\n\n\n\n\n HYBRID ANALYSIS!!!!::::::::!!!\n\n\n",ResultHybridA);

		HybridTestURL := "https://testsafebrowsing.appspot.com/s/malware.html"

		ResultURLHybridA:=CheckURLHybridAnalyis(HybridTestURL)

		fmt.Println("\n\n\n\n\n HYBRID URL:\n\n", ResultURLHybridA)

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
