package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CallAlienVaultUrl function takes a url, returns data on it from the alienvault api
func CallAlienVaultUrl(url string) (response string) {

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

	response = string(body)

	return response
}

// CallAlienVaultHash function takes a hash, returns data on it from the alienvault api
func CallAlienVaultHash(hash string) (response string) {

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
	return response
}
