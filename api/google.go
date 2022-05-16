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

// CallGoogleUrl function takes a url, returns data on it from the google safebrowsing api
func CallGoogleUrl(url string) (response utils.FrontendResponse) {
	// Google API returnerer [] om den ikke kjenner til domenet / URL. Kan bruke dette til
	// å avgjøre om det er malicious eller ikke.

	APIKey := utils.APIKeyGoogle

	postURL := "https://safebrowsing.googleapis.com/v4/threatMatches:find?key=" + APIKey

	if strings.Contains(url, "https://") {
		httpsSearchURL := url
		container := strings.SplitAfter(url, "https://")

		httpSearchURL := "http://" + container[1]

		fmt.Println("1 : This is the HTTP URL after splitting and concatinating", httpSearchURL)
		fmt.Println("1 : This is the HTTPs URL after splitting and concatinating", httpsSearchURL)
	} else if strings.Contains(url, "http://") {
		httpSearchURL := url
		container := strings.SplitAfter(url, "http://")
		httpsSearchURL := "http://" + container[1]

		fmt.Println("2 : This is the HTTP URL after splitting and concatinating", httpSearchURL)
		fmt.Println("2 : This is the HTTPs URL after splitting and concatinating", httpsSearchURL)
	} else {
		httpSearchURL := "http://" + url
		httpsSearchURL := "https://" + url
		fmt.Println("3 : This is the HTTP URL after splitting and concatinating", httpSearchURL)
		fmt.Println("4 : This is the HTTPs URL after splitting and concatinating", httpsSearchURL)
	}

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
				{"url": "https://` + url + `" },
				{"url": "http://` + url + `"}
			  ]
			}
		}`)

	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error: reading sending google api request")
		logging.Logerror(err, "ERROR sending REQUEST, Google API")
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
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
		logging.Logerror(err, "ERROR Unmarshalling google api response, Google API")
	}
	output := string(body)
	fmt.Println("BODY::!", output)
	//fmt.Println("ThreatType::::",jsonResponse.Matches[0].ThreatType)
	//fmt.Println("response Body:", string(body))
	if len(jsonResponse.Matches) != 0 {
		response.Content = "This URL has been marked as malicious by Google Safebrowsing, visiting is NOT recommended"
		switch jsonResponse.Matches[0].ThreatType {
		case "MALWARE":
			response.Status = "Risk"

		case "SOCIAL_ENGINEERING":
			response.Status = "Risk"

		case "UNWANTED_SOFTWARE":
			response.Status = "Risk"

		default:
			response.Status = "potentially unsafe"
			response.Content = "This URL has been marked as suspicious, not recommended to visit."
		}
	} else {
		response.Status = "Safe"
		response.Content = "Google safebrowsing has no data that indicates this is an unsafe URL"
	}

	response.SourceName = "Google SafeBrowsing Api"

	return response
}

func TestGoGoogleUrl(url string, response *utils.FrontendResponse2, wg *sync.WaitGroup) {
	// Google API returnerer [] om den ikke kjenner til domenet / URL. Kan bruke dette til
	// å avgjøre om det er malicious eller ikke.
	defer wg.Done()

	var httpSearchURL, httpsSearchURL string

	if strings.Contains(url, "https://") {
		httpsSearchURL = url
		container := strings.SplitAfter(url, "https://")

		httpSearchURL = "http://" + container[1]

		fmt.Println("1 : This is the HTTP URL after splitting and concatinating", httpSearchURL)
		fmt.Println("1 : This is the HTTPs URL after splitting and concatinating", httpsSearchURL)
	} else if strings.Contains(url, "http://") {
		httpSearchURL = url
		container := strings.SplitAfter(url, "http://")
		httpsSearchURL = "https://" + container[1]

		fmt.Println("2 : This is the HTTP URL after splitting and concatinating", httpSearchURL)
		fmt.Println("2 : This is the HTTPs URL after splitting and concatinating", httpsSearchURL)
	} else {
		httpSearchURL = "http://" + url
		httpsSearchURL = "https://" + url
		fmt.Println("3 : This is the HTTP URL after splitting and concatinating", httpSearchURL)
		fmt.Println("4 : This is the HTTPs URL after splitting and concatinating", httpsSearchURL)
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
		panic(err)
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
	output := string(body)
	fmt.Println("BODY::!", output)
	//fmt.Println("ThreatType::::",jsonResponse.Matches[0].ThreatType)
	//fmt.Println("response Body:", string(body))
	/*
		if len(jsonResponse.Matches) != 0 {
			response.Content = "This URL has been marked as malicious by Google Safebrowsing, visiting is NOT recommended"
			switch jsonResponse.Matches[0].ThreatType {
			case "MALWARE":
				response.Status = "Risk"

			case "SOCIAL_ENGINEERING":
				response.Status = "Risk"

			case "UNWANTED_SOFTWARE":
				response.Status = "Risk"

			default:
				response.Status = "potentially unsafe"
				response.Content = "This URL has been marked as suspicious, not recommended to visit."
			}
		} else {
			response.Status = "Safe"
			response.Content = "Google safebrowsing has no data that indicates this is an unsafe URL"
		}

		response.SourceName = "Google SafeBrowsing Api"
	*/
	utils.SetResponeObjectGoogle(jsonResponse, response)
}
