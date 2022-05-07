package api

import (
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
)

// CallHybridAnalysisHash function takes a hash, returns data on it from the hybridanalysis api
func CallHybridAnalysisHash(hash string, response *utils.FrontendResponse2, wg *sync.WaitGroup) {

	// API dokumentasjon https://www.hybrid-analysis.com/docs/api/v2#/Search/post_search_hash

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
		utils.SetGenericError(response)
	}
	defer res.Body.Close()

	fmt.Println("\nStatus paa request", res.Status)
	if res.StatusCode == 200 {

		//fmt.Println("response Status:", res.Status)
		//fmt.Print("Response Headers:", res.Header)
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println("\nBody", string(body))
		//fmt.Println("response Body:", string(body))

		var jsonResponse utils.HybridAnalysishash

		err = json.Unmarshal(body, &jsonResponse)
		if err != nil {
			fmt.Println(string(body))
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
func CallHybridAnalyisUrl(URL string) (VirusTotal utils.FrontendResponse, urlscanio utils.FrontendResponse) {

	fmt.Println("HYBRID URL: ", URL)
	//DENNE FUNKSJONENE KAN SCANNE EN URL MEN DETTE BENYTTER SEG AV VIRUS TOTAL/
	// DETTE ER KANSKJE EN GOD WORK AROUND FOR Å KUNNE BRUKE VT GRATIS SIDEN Hybrid Analysis har lisens.
	// Problem her kan være at dette må inkomporere en "await - 5-15 sekunder
	// om det ikke er noe cachet result på VirusTotal, fordi den maa kjore ny request.".
	// Titter på dette.
	// Vi har CAP på 2000 request i timen hos Hybrid Analyis, dette burde vell holde??? - 200 max i minuttet.
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
	data.Set("no_share_third_party", "true")
	data.Set("allow_community_access", "false")
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
	}

	//var jsonData map[string]interface{}
	var jsonResponse utils.HybridAnalysisURL

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println(err)
	}

	if !jsonResponse.Finished {
		time.Sleep(20 * time.Second) //Får prøve å finne en bedre løsning enn dette men det er det jeg har for now.

		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Ioutil error:", err)
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

func TestHybridAnalyisUrl(URL string, VirusTotal *utils.FrontendResponse2, urlscanio *utils.FrontendResponse2, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("HYBRID URL: ", URL)
	//DENNE FUNKSJONENE KAN SCANNE EN URL MEN DETTE BENYTTER SEG AV VIRUS TOTAL/
	// DETTE ER KANSKJE EN GOD WORK AROUND FOR Å KUNNE BRUKE VT GRATIS SIDEN Hybrid Analysis har lisens.
	// Problem her kan være at dette må inkomporere en "await - 5-15 sekunder
	// om det ikke er noe cachet result på VirusTotal, fordi den maa kjore ny request.".
	// Titter på dette.
	// Vi har CAP på 2000 request i timen hos Hybrid Analyis, dette burde vell holde??? - 200 max i minuttet.
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
	data.Set("no_share_third_party", "true")
	data.Set("allow_community_access", "false")
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

	fmt.Println("response Status:", res.Status)
	if res.StatusCode == http.StatusOK {

		// res.Body.Read("finished") Her skal jeg føre en sjekk som sjekker om "finished = true eller false"

		// Hvis denne er false skal den vente 5 sekunder og kjøre requesten på nytt.
		// Eventuelt om det er en måte å ikke close requesten før den er finished???????

		// Her kan det sjekkes om VirusTotal - Status er Malicious og om Urlscan.io
		// - status er malicious, suspicious, clean etc. også bare returnere denne responsen.

		//fmt.Print("Response Headers:", res.Header)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Ioutil error:", err)
		}

		//var jsonData map[string]interface{}
		var jsonResponse utils.HybridAnalysisURL

		err = json.Unmarshal(body, &jsonResponse)
		if err != nil {
			fmt.Println(err)
		}

		if !jsonResponse.Finished {
			time.Sleep(40 * time.Second) //Får prøve å finne en bedre løsning enn dette men det er det jeg har for now.

			res, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("Ioutil error:", err)
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
		/*
				VirusTotal.Status = jsonResponse.Scanners[0].Status

				// Set the clean value to safe instead for frontend display.
				if VirusTotal.Status == "clean" {
					VirusTotal.Status = "Safe"
				}

				urlscanio.SourceName = jsonResponse.Scanners[1].Name
				urlscanio.Status = jsonResponse.Scanners[1].Status

				fmt.Println("Attempted HybridAnalysisURL output VT:", VirusTotal.SourceName, "   Status:", VirusTotal.Status)
				fmt.Println("\n\nAttempted HybridAnalysisURL output VT:", urlscanio.SourceName, "   Status:", urlscanio.Status)
			} else {
				VirusTotal.SourceName = "VirusTotal"
				VirusTotal.Status = "Error"

				urlscanio.SourceName = "urlscan.io"
				urlscanio.Status = "Error"
			}
		*/
		fmt.Println("WHAT IS THIS \n\n\n", jsonResponse.Finished)
		fmt.Println("URLSCANIO STATUS:", jsonResponse.Scanners[1].Status)

		utils.SetResponeObjectVirusTotal(jsonResponse, VirusTotal)
		utils.SetResponeObjectUrlscanio(jsonResponse, urlscanio)
	} else {
		VirusTotal.SourceName = "VirusTotal"
		VirusTotal.EN.Status = "Error"
		VirusTotal.NO.Status = "Error"

		urlscanio.SourceName = "urlscan.io"
		urlscanio.EN.Status = "Error"
		urlscanio.NO.Status = "Error"

	}
}
