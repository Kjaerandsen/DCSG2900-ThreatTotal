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
			response, err := utils.Conn.Do("SETEX", "url:"+url, utils.CacheDurationUrl, URLint)
			if err != nil {
				fmt.Println("Error adding data to redis:" + err.Error())
				logging.Logerror(err, "Error addding data to redis - Url-intelligence:")
			}

			// Print the response to adding the data (should be "OK")
			fmt.Println("Bool is true")
			fmt.Println(response)
		}
		//fmt.Println("WHERE IS MY CONTENT 2", responseData2)
		// Cache hit
	} else {
		fmt.Println("Cache hit")
		responseBytes, err := json.Marshal(value)
		if err != nil {
			fmt.Println("Error handling redis response:" + err.Error())
			logging.Logerror(err, "Error handling redis response - Url-intelligence:")
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
			// Maybe do another call to delete the key from the database?
		}
		/**
		//var checkData utils.ResultFrontendResponse
		err = json.Unmarshal(responseBytes, &checkdata)
		if err!=nil {
			fmt.Println(string(checkData))
		}
		fmt.Println(string(checkData))
		*/
		err = json.Unmarshal(responseBytes, &URLint)
		if err != nil {
			fmt.Println("Error handling redis response:" + err.Error())
			logging.Logerror(err, "Error handling redis response - Url-intelligence:")
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
			// Maybe do another call to delete the key from the database?
		}
	}

	c.Data(http.StatusOK, "application/json", URLint)
}

// Makes the api requests used in urlIntelligence
func urlSearch(url string) (data []byte, err error, complete bool) {
	var wg sync.WaitGroup //Vente gruppe for goroutiner
	var URLint []byte
	var responseData [4]utils.FrontendResponse2

	var p, VirusTotal, urlscanio, alienvault *utils.FrontendResponse2
	p = &responseData[0]
	VirusTotal = &responseData[1]
	urlscanio = &responseData[2]
	alienvault = &responseData[3]

	wg.Add(2)
	fmt.Println(url)
	if checkUrlAgainstFilter(url) {
		wg.Add(1)
		go TestGoGoogleUrl(url, p, &wg)
		go TestHybridAnalyisUrl(url, VirusTotal, urlscanio, &wg)
		go TestAlienVaultUrl(url, alienvault, &wg)
	} else {
		go giveTrueGoogleUrl(url, p)
		go TestHybridAnalyisUrl(url, VirusTotal, urlscanio, &wg)
		go TestAlienVaultUrl(url, alienvault, &wg)
	}
	wg.Wait()

	var resultResponse utils.ResultFrontendResponse

	resultResponse.FrontendResponse = responseData[:]

	setResults := &resultResponse

	utils.SetResultURL(setResults, len(responseData))

	
	//TESTING FUNCTIONALITY FOR SCREENSHOT OF URLS
	utils.ScreenshotURL(url, setResults)			////
													////
	fmt.Println(len(resultResponse.Screenshot))		////
	
	complete = checkIfIntelligenceComplete(resultResponse, len(responseData)) //This runs a check to see if the intelligence is complete
	//If complete is true the intelligence will be cached,
	//If it is not complete the result won't be cached.

	URLint, err = json.Marshal(resultResponse)
	if err != nil {
		fmt.Println(err)
		return URLint, err, complete
	}

	//fmt.Println("WHERE IS MY CONTENT 1", responseData)

	return URLint, nil, complete
}

func checkIfIntelligenceComplete(jsonData utils.ResultFrontendResponse, size int) (complete bool) {
	complete = true

	for i := 0; i <= size-1; i++ {
		if jsonData.FrontendResponse[i].EN.Status == "Awaiting analysis" || jsonData.FrontendResponse[i].EN.Status == "Error" {
			complete = false
		}
	}

	return complete
}

func checkUrlAgainstFilter(url string) bool {
	for i := 0; i < len(utils.UrlBlockList); i++ {
		if strings.Contains(url, utils.UrlBlockList[i]) {
			return false
		}
	}
	return true
}

func giveTrueGoogleUrl(url string, response *utils.FrontendResponse2) {
	response.EN.Status = "Safe"
	response.EN.Content = "Google safebrowsing has no data that indicates this is an unsafe URL/Domain"
	response.NO.Status = "Trygg"
	response.NO.Content = "Google Safebrowsing har ingen data som indikerer at dette er en utrygg URL/Domene"
	response.SourceName = "Google SafeBrowsing Api"
}
