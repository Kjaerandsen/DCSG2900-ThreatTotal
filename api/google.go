package api

import (
	"bytes"
	logging "dcsg2900-threattotal/logs"
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	//"dcsg2900-threattotal/main"
)

//Function to call the Google Safe Browsing API.
//API documentation can be found in: https://developers.google.com/safe-browsing/v4
// Contacted API Endpoint : https://safebrowsing.googleapis.com/v4/threatMatches
func CallGoogleUrl(url string, response *utils.FrontendResponse2, wg *sync.WaitGroup) {
	// Google API returnerer [] om den ikke kjenner til domenet / URL. Kan bruke dette til
	// å avgjøre om det er malicious eller ikke.
	defer wg.Done()

	var httpSearchURL, httpsSearchURL string

	if strings.Contains(url, "https://") {
		httpsSearchURL = url
		container := strings.SplitAfter(url, "https://")

		httpSearchURL = "http://" + container[1]

	} else if strings.Contains(url, "http://") {
		httpSearchURL = url
		container := strings.SplitAfter(url, "http://")
		httpsSearchURL = "https://" + container[1]

		
	} else {
		httpSearchURL = "http://" + url
		httpsSearchURL = "https://" + url
		
	}

	APIKey := utils.APIKeyGoogle

	postURL := "https://safebrowsing.googleapis.com/v4/threatMatches:find?key=" + APIKey

	var jsonData = []byte(`
		{
			"client": {
			  "clientId":      "threattotal",
			  "clientVersion": "1.5.2"
			},
			"threatInfo": {
			  "threatTypes":      ["MALWARE", "SOCIAL_ENGINEERING", "THREAT_TYPE_UNSPECIFIED", "UNWANTED_SOFTWARE","POTENTIALLY_HARMFUL_APPLICATION"],
			  "platformTypes":    ["ANY_PLATFORM"],
			  "threatEntryTypes": ["URL"],
			  "threatEntries": [
				{"url": "`+ httpsSearchURL +`" },
				{"url": "`+ httpSearchURL +`"}
			  ]
			}
		}`)

	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error: reading sending google api request")
		logging.Logerror(err, "ERROR Sending google api request, Google API")

	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: in google api response")
		logging.Logerror(err, "ERROR reading google api response, Google API")
		utils.SetGenericError(response)
		return
	}
	defer res.Body.Close()

	//fmt.Println("response Status:", res.Status)
	//fmt.Print("Response Headers:", res.Header)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: reading google api response")
		logging.Logerror(err, "ERROR reading google api response, Google API")

	}

	var jsonResponse utils.GoogleSafeBrowsing

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println(err)
	}

	utils.SetResponeObjectGoogle(jsonResponse, response)
}
