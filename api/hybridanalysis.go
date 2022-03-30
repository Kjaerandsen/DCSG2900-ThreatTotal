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
)

// CallHybridAnalysisHash function takes a hash, returns data on it from the hybridanalysis api
func CallHybridAnalysisHash(hash string) (response string) {

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
	response = string(body)

	return

}

// CallHybridAnalyisUrl function takes a url, returns data on it from the hybridanalysis api
func CallHybridAnalyisUrl(URL string) (response string) {

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
	var jsonData utils.HybridAnalysisUrl

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n\nScanners", jsonData.Id)

	fmt.Println("\n\n\n\n\nEverything", jsonData.Scanners[0])

	//ttstr := jsonData["scanners"].(map[string]interface{})["VirusTotal"].(string)

	//fmt.Println("\n\n\nTTSTR ER:", ttstr)

	//json.Unmarshal([]byte(body), &jsonData)

	//fmt.Println("Status is:", jsonData)

	//fmt.Println("response Body:", string(body))
	response = string(body)

	return response
}
