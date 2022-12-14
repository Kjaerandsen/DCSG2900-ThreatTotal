package api

import (
	logging "dcsg2900-threattotal/logs"
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)
//URL intelligence function called from main, checks the cache for hits, if miss calls to gather intelligence. 
func UrlIntelligence(c *gin.Context) {
	url := c.Query("url")

	var completeInt bool
	var URLint []byte

	//var URLint utils.APIresponseResult
	value, err := utils.Conn.Do("GET", "url:"+url)
	if value == nil {
		if err != nil {
			fmt.Println("Error:" + err.Error())
			logging.Logerror(err, "Error in cache lookup - Url-intelligence")

		}
		fmt.Println("No Cache hit")

		URLint, err, completeInt = urlSearch(url)
		if err != nil {
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
		}

		// Add the data to the redis backend.
		if completeInt {
			_, err := utils.Conn.Do("SETEX", "url:"+url, utils.CacheDurationUrl, URLint)
			if err != nil {
				fmt.Println("Error adding data to redis:" + err.Error())
				logging.Logerror(err, "Error addding data to redis - Url-intelligence:")
			}
		}

		// Cache hit
	} else {
		fmt.Println("Cache hit")
		responseBytes, err := json.Marshal(value)
		if err != nil {
			fmt.Println("Error handling redis response:" + err.Error())
			logging.Logerror(err, "Error handling redis response - Url-intelligence:")
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return

		}

		err = json.Unmarshal(responseBytes, &URLint)
		if err != nil {
			fmt.Println("Error handling redis response:" + err.Error())
			logging.Logerror(err, "Error handling redis response - Url-intelligence:")
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return

		}
	}

	c.Data(http.StatusOK, "application/json", URLint)
}

// Makes the api requests used in urlIntelligence
func urlSearch(url string) (data []byte, err error, complete bool) {
	var wg sync.WaitGroup //Wait group for go routines
	var URLint []byte
	var responseData [4]utils.FrontendResponse2 //Array of frontend response structs

	var p, VirusTotal, urlscanio, alienvault *utils.FrontendResponse2
	p = &responseData[0]
	VirusTotal = &responseData[1]
	urlscanio = &responseData[2]
	alienvault = &responseData[3]

	wg.Add(3)
	if checkUrlAgainstFilter(url) { //Checks if the URL is in the POC urlfilter.
		go CallGoogleUrl(url, p, &wg) //Calls different functions to contact intelligence sources.
		go CallHybridAnalyisUrl(url, VirusTotal, urlscanio, &wg)
		go CallAlienVaultUrl(url, alienvault, &wg)
	} else { //If URL is in urlfilter, set google to safe as POC (Proof of concept.).
		go giveTrueGoogleUrl(url, p, &wg)
		go CallHybridAnalyisUrl(url, VirusTotal, urlscanio, &wg)
		go CallAlienVaultUrl(url, alienvault, &wg)
	}
	wg.Wait()

	var resultResponse utils.ResultFrontendResponse //Creat new struct that will be sent to frontend.

	resultResponse.FrontendResponse = responseData[:] //Move frontend response structs into resultresponse struct.

	setResults := &resultResponse //Create pointer to resultresponse.

	utils.SetResultURL(setResults, len(responseData)) //Set the result string.

	//FUNCTIONALITY FOR SCREENSHOT OF URLS
	utils.ScreenshotURL(url, setResults) ////

	//fmt.Println(len(resultResponse.Screenshot)) ////Check if screenshot contains anything (Is valid)

	complete = checkIfIntelligenceComplete(resultResponse, len(responseData)) //This runs a check to see if the intelligence is complete
	//If complete is true the intelligence will be cached,
	//If it is not complete the result won't be cached.

	URLint, err = json.Marshal(resultResponse) //Marshal data to be sent to frontend.
	if err != nil {
		fmt.Println(err)
		return URLint, err, complete
	}

	return URLint, nil, complete
}

//Function to check if the intelligence is complete and ready to be cached, returns a complete bool - False = not ready, True = ready.
func checkIfIntelligenceComplete(jsonData utils.ResultFrontendResponse, size int) (complete bool) {
	complete = true

	for i := 0; i <= size-1; i++ {
		if jsonData.FrontendResponse[i].EN.Status == "Awaiting analysis" || jsonData.FrontendResponse[i].EN.Status == "Error" {
			complete = false
		}
	}

	return complete
}

// CheckUrlAgainstFilter function which returns false if an item is whitelisted
func checkUrlAgainstFilter(url string) bool {
	for i := 0; i < len(utils.UrlBlockList); i++ {
		if strings.Contains(url, utils.UrlBlockList[i]) {
			return false
		}
	}
	return true
}

// Function which creates a safe response for the google api, used in combination
// with the filter to demo filter functionality
func giveTrueGoogleUrl(url string, response *utils.FrontendResponse2, wg *sync.WaitGroup) {
	defer wg.Done()

	response.EN.Status = "Safe"
	response.EN.Content = "Google safebrowsing has no data that indicates this is an unsafe URL/Domain"
	response.NO.Status = "Trygg"
	response.NO.Content = "Google Safebrowsing har ingen data som indikerer at dette er en utrygg URL/Domene"
	response.SourceName = "Google SafeBrowsing Api"
}
