package api

import (
	logging "dcsg2900-threattotal/logs"
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// CallHybridAnalysisHash function takes a hash, returns data on it from the hybridanalysis api
// API endpoint documentation https://www.hybrid-analysis.com/docs/api/v2#/Search/post_search_hash
func CallHybridAnalysisHash(hash string, response *utils.FrontendResponse2, wg *sync.WaitGroup) {

	defer wg.Done()

	response.SourceName = "Hybrid Analysis"

	APIKey := utils.APIKeyHybridAnalysis

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
		fmt.Println("Request error HybridA", err)
		logging.Logerror(err, "Error in request to hybridAnalysis")

		utils.SetGenericError(response)
	}
	defer res.Body.Close()

	fmt.Println("\nStatus paa request", res.Status)
	if res.StatusCode == 200 {

		body, _ := ioutil.ReadAll(res.Body)

		var jsonResponse utils.HybridAnalysishash

		err = json.Unmarshal(body, &jsonResponse)
		if err != nil {
			fmt.Println(err)
			if len(string(body)) == 2 { //If this statement is true it means that the request
				//is sucessful but it cant be unmarshalled because it returns empty
				//It returns empty because HybridAnalysis does not have any information on the hash.
				utils.SetResponseObjectHybridAnalysisHash(jsonResponse, response) //This function will then parse this as unknown.
			} else {
				utils.SetGenericError(response) //If it did not return empty but still failed it means something else went wrong,											//Returning an error
			}
			return //Returning
		}

		utils.SetResponseObjectHybridAnalysisHash(jsonResponse, response)
	} else {
		utils.SetGenericError(response)
	}
}

// CallHybridAnalyisUrl function takes a url, returns data on it from the hybridanalysis api
// https://www.hybrid-analysis.com/docs/api/v2#/Quick%20Scan/post_quick_scan_url Documentation for contacted endpoint
func CallHybridAnalyisUrl(URL string) (VirusTotal utils.FrontendResponse, urlscanio utils.FrontendResponse) {

	APIKey := utils.APIKeyHybridAnalysis

	postURL := "https://www.hybrid-analysis.com/api/v2/quick-scan/url"

	data := url.Values{}
	data.Set("scan_type", "all")                //What type of scan to perform
	data.Set("url", URL)                        //Sets URL to search
	data.Set("no_share_third_party", "true")    //Makes the search not accessible to 3-rd party others
	data.Set("allow_community_access", "false") //Does not share search with community

	req, err := http.NewRequest("POST", postURL, strings.NewReader(data.Encode())) //Sets the new request.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("api-key", APIKey)              //Set API key
	req.Header.Set("User-Agent", "Falcon Sandbox") //Set USER-AGENT, just to bypass user-agent check - See documentation on API

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error in request")
		logging.Logerror(err, "Error in request Hybrid Analysis")
	}
	defer res.Body.Close()

	// res.Body.Read("finished") Her skal jeg føre en sjekk som sjekker om "finished = true eller false"

	// Hvis denne er false skal den vente 5 sekunder og kjøre requesten på nytt.
	// Eventuelt om det er en måte å ikke close requesten før den er finished???????

	// Her kan det sjekkes om VirusTotal - Status er Malicious og om Urlscan.io
	// - status er malicious, suspicious, clean etc. også bare returnere denne responsen.

	//fmt.Println("response Status:", res.Status)
	//fmt.Print("Response Headers:", res.Header)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Ioutil error:", err)
		logging.Logerror(err, "Ioutil error HybridAnalysis: ")

	}

	//var jsonData map[string]interface{}
	var jsonResponse utils.HybridAnalysisURL

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println(err)
	}

	if !jsonResponse.Finished {
		time.Sleep(20 * time.Second)

		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Ioutil error:", err)
			logging.Logerror(err, "Ioutil error HybridAnalysis: ")

		}

		var jsonResponse utils.HybridAnalysisURL

		err = json.Unmarshal(body, &jsonResponse)
		if err != nil {
			fmt.Println(err)
		}
	}

	VirusTotal.SourceName = jsonResponse.Scanners[0].Name
	VirusTotal.Status = jsonResponse.Scanners[0].Status

	// Set the clean value to safe instead for frontend display.
	if VirusTotal.Status == "clean" {
		VirusTotal.Status = "Safe"
	}

	urlscanio.SourceName = jsonResponse.Scanners[1].Name
	urlscanio.Status = jsonResponse.Scanners[1].Status

	fmt.Println("Attempted HybridAnalysisURL output VT:", VirusTotal.SourceName, "   Status:", VirusTotal.Status)
	fmt.Println("\n\nAttempted HybridAnalysisURL output VT:", urlscanio.SourceName, "   Status:", urlscanio.Status)

	return VirusTotal, urlscanio
}

//Function to perform request to the Hybrid Analysis API for URL and domain intelligence.
// https://www.hybrid-analysis.com/docs/api/v2#/Quick%20Scan/post_quick_scan_url Documentation on used API endpoint.

func TestHybridAnalyisUrl(URL string, VirusTotal *utils.FrontendResponse2, urlscanio *utils.FrontendResponse2, wg *sync.WaitGroup) {

	defer wg.Done()

	APIKey := utils.APIKeyHybridAnalysis

	postURL := "https://www.hybrid-analysis.com/api/v2/quick-scan/url"

	data := url.Values{}
	data.Set("scan_type", "all")                //Sets the scan type.
	data.Set("url", URL)                        //Sets the URL to be searched
	data.Set("no_share_third_party", "true")    //Does not share search with 3rd party
	data.Set("allow_community_access", "false") //Sets it so that search is not shared with community.

	req, err := http.NewRequest("POST", postURL, strings.NewReader(data.Encode())) //Creates new post request
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")            //Sets required content type
	req.Header.Set("api-key", APIKey)                                              //Adds the API key
	req.Header.Set("User-Agent", "Falcon Sandbox")                                 //Sets user agent to falcon sandbox, to bypass user agent check.

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err, "Error in request to Hybrid Analysis - URL endpoint. ")
		logging.Logerror(err, "Error in request to Hybrid Analysis - URL")
		utils.SetGenericError(VirusTotal)
		utils.SetGenericError(urlscanio)
		return
	}
	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	if res.StatusCode == http.StatusOK {

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Ioutil error:", err)
			logging.Logerror(err, "Ioutil error HybridAnalysis: ")

		}

		var jsonResponse utils.HybridAnalysisURL

		err = json.Unmarshal(body, &jsonResponse)
		if err != nil {
			fmt.Println(err)
		}

		if !jsonResponse.Finished {
			time.Sleep(40 * time.Second) //In case the analysis is not finished, we wait 40 seconds to perform a new request.

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err, "Error in request to Hybrid Analysis - URL endpoint. ")
				logging.Logerror(err, "Error in request to Hybrid Analysis - URL")
				utils.SetGenericError(VirusTotal)
				utils.SetGenericError(urlscanio)
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("Ioutil error:", err)
				logging.Logerror(err, "Ioutil error HybridAnalysis: ")

			}

			var jsonResponse utils.HybridAnalysisURL

			err = json.Unmarshal(body, &jsonResponse)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println(jsonResponse)
		VirusTotal.SourceName = jsonResponse.Scanners[0].Name
		urlscanio.SourceName = jsonResponse.Scanners[1].Name

		utils.SetResponseObjectVirusTotal(jsonResponse, VirusTotal)
		utils.SetResponseObjectUrlscanio(jsonResponse, urlscanio)
	} else if res.StatusCode == http.StatusBadRequest {	//Added a special check here to see if the domain does not ecist

		body, err := ioutil.ReadAll(res.Body)		//If body can not be read, default to generic error
		if err != nil {
			fmt.Println("Ioutil error:", err)
			logging.Logerror(err, "Ioutil error HybridAnalysis: ")
			VirusTotal.SourceName = "VirusTotal"
			urlscanio.SourceName = "urlscan.io"

			utils.SetGenericError(VirusTotal)
			utils.SetGenericError(urlscanio)
		}

		var jsonResponse utils.HybridAnalysisBadRequest

		err = json.Unmarshal(body, &jsonResponse)		//If json data can not be unmarshaled default to generic error struct
		if err != nil {
			fmt.Println(err)
			logging.Logerror(err, "Ioutil error HybridAnalysis: ")
			VirusTotal.SourceName = "VirusTotal"
			urlscanio.SourceName = "urlscan.io"

			utils.SetGenericError(VirusTotal)
			utils.SetGenericError(urlscanio)
		}
		if jsonResponse.Message == "Failed to download file: domain does not exist" {	//If message contains this, it means domain does not exist
			VirusTotal.SourceName = "VirusTotal"
			VirusTotal.EN.Status = "Safe"
			VirusTotal.EN.Content = "Domain does not exist"

			VirusTotal.NO.Status = "Trygg"
			VirusTotal.NO.Content = "Domenet eksisterer ikke"

			urlscanio.SourceName = "urlscan.io"
			urlscanio.EN.Status = "Safe"
			urlscanio.EN.Content = "Domain does not exist"

			urlscanio.NO.Status = "Trygg"
			urlscanio.NO.Content = "Domenet eksisterer ikke"
		}

	} else {
		VirusTotal.SourceName = "VirusTotal"
		VirusTotal.EN.Status = "Error"
		VirusTotal.NO.Status = "Error"

		urlscanio.SourceName = "urlscan.io"
		urlscanio.EN.Status = "Error"
		urlscanio.NO.Status = "Error"
	}
}
