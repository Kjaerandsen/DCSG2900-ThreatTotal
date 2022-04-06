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
	"time"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
)

// CallHybridAnalysisHash function takes a hash, returns data on it from the hybridanalysis api
func CallHybridAnalysisHash(hash string) (response utils.FrontendResponse) {

	// API dokumentasjon https://www.hybrid-analysis.com/docs/api/v2#/Search/post_search_hash

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

	var jsonResponse utils.HybridAnalysishash

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Something cool", jsonResponse.Verdict)

	response.SourceName = "Hybrid Analysis"
	if jsonResponse.Verdict == "malicious" {
		response.Status = "Risk"
		response.Content = "This file is malicious"
		//response.SourceName = jsonResponse.Submissions[0].Filename
	} else if jsonResponse.Verdict == "whitelisted" {
		response.Status = "Safe"
		response.Content = "This file is known to be good"
		//response.SourceName = jsonResponse.Submissions[0].Filename
	} else {
		response.Status = "Safe" //Denne må byttes til at den er ukjent // grå farge elns på frontend.
		response.Content = "This filehash is not known to Hybrid Analysis"
	}

	// Set the filename field if known
	if jsonResponse.Submissions != nil {
		if jsonResponse.Submissions[0].Filename != "" {
			response.Content = response.Content + " " + jsonResponse.Submissions[0].Filename
		}
	}

	return

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

	if jsonResponse.Finished != true {
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

	urlscanio.SourceName = jsonResponse.Scanners[1].Name
	urlscanio.Status = jsonResponse.Scanners[1].Status
	
	fmt.Println("Attempted HybridAnalysisURL output VT:", VirusTotal.SourceName, "   Status:", VirusTotal.Status)
	fmt.Println("\n\nAttempted HybridAnalysisURL output VT:", urlscanio.SourceName, "   Status:", urlscanio.Status)

	return VirusTotal, urlscanio
}
