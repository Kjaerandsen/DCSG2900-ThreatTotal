package api

import (
	logging "dcsg2900-threattotal/logs"
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// CallAlienVaultUrl function takes a url, returns data on it from the alienvault api
func CallAlienVaultUrl(url string) (response utils.FrontendResponse) {

	//DENNE FUNKSJONEN KAN UTARBEIDES TIL Å BARE RETURNERE MALCICIOUS / SUSPCIOUS OM DET BEFINNER SEG NEVNT I NOEN
	// PULSEES (Problemet her er at ting som er OK kan være i pulse... Må tenke litt her)
	// Convert []byte to string and print to screen
	APIKey := utils.APIKeyOTX

	getURL := "https://otx.alienvault.com//api/v1/indicators/url/" + url + "/general"

	req, err := http.NewRequest("GET", getURL, nil)
	req.Header.Set("X-OTX-API-KEY", APIKey)

	//fmt.Println(req.Header)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR READING JSON DATA", err)
		logging.Logerror(err, "ERROR READING JSON DATA, AlienVault API")
	}

	var jsonResponse utils.AlienVaultURL

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println("UNMARSHAL ERROR:\n\n", err)
		logging.Logerror(err, "Unmarshal error AlienVault API")
	}

	//output:= string(body)
	//fmt.Println(output)
	//fmt.Println("\n\nAMOUNT OF PULSES:::::: ", jsonResponse.PulseInfo.Count)
	if jsonResponse.PulseInfo.Count == 0 {
		response.Status = "Safe"
	} else {
		response.Status = "Risk"
	}

	response.SourceName = "AlienVault"

	//response = string(body)

	return response
}

// CallAlienVaultHash function takes a hash, returns data on it from the alienvault api
func CallAlienVaultHash(hash string, response *utils.FrontendResponse2, wg *sync.WaitGroup) {

	defer wg.Done()
	response.SourceName = "AlienVault"

	APIKey := utils.APIKeyOTX

	getURL := "https://otx.alienvault.com//api/v1/indicators/file/" + hash + "/general"

	req, err := http.NewRequest("GET", getURL, nil)
	req.Header.Set("X-OTX-API-KEY", APIKey)

	//fmt.Println(req.Header)

	client := &http.Client{}

	res, err := client.Do(req)
	//fmt.Println(res.Status)
	//fmt.Print(string(res.Body))
	if err != nil {
		fmt.Println("ERROR IN Request", err)
		logging.Logerror(err, "ERROR IN REQUEST, AlienVault API")
		utils.SetGenericError(response)
	}
	if res.StatusCode == 200 {
		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		if err != nil {
			fmt.Println("ERROR READING JSON DATA", err)
			logging.Logerror(err, "ERROR Reading JSON response, AlienVault API")
			utils.SetGenericError(response)
		}

		var jsonResponse utils.AlienVaultHash

		err = json.Unmarshal(body, &jsonResponse)
		if err != nil {
			utils.SetGenericError(response)
			fmt.Println(err)
		}

		//output:= string(body)
		//fmt.Println(output)
		//fmt.Println("\n\nAMOUNT OF PULSES:::::: ", jsonResponse.PulseInfo.Count)
		utils.SetResponseObjectAlienVaultHash(jsonResponse, response)
	} else {

		response.EN.Content = "We have encountered an error, check if the filehash is a valid filehash."
		response.EN.Status = "ERROR"

		response.NO.Content = "Vi har møtt på en error, sjekk om filhashen er en gyldig filhash."
		response.NO.Status = "ERROR"
	}
	//HER KAN VI SJEKKE OM "VERDICT feltet er" MALICIOUS, SUSPICIOUS ELLER NOE ANNET. OG Bare returnere det.
}

func TestAlienVaultUrl(url string, response *utils.FrontendResponse2, wg *sync.WaitGroup) {

	defer wg.Done()
	//DENNE FUNKSJONEN KAN UTARBEIDES TIL Å BARE RETURNERE MALCICIOUS / SUSPCIOUS OM DET BEFINNER SEG NEVNT I NOEN
	// PULSEES (Problemet her er at ting som er OK kan være i pulse... Må tenke litt her)
	// Convert []byte to string and print to screen
	APIKey := utils.APIKeyOTX

	getURL := "https://otx.alienvault.com//api/v1/indicators/url/" + url + "/general"

	req, err := http.NewRequest("GET", getURL, nil)
	req.Header.Set("X-OTX-API-KEY", APIKey)

	//fmt.Println(req.Header)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR READING JSON DATA", err)
		logging.Logerror(err, "ERROR Reading JSON response, AlienVault API")

	}

	var jsonResponse utils.AlienVaultURL

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println("UNMARSHAL ERROR:\n\n", err)
		logging.Logerror(err, "ERROR unmarshalling, AlienVault URLsearch API")
	}

	/*
		//output:= string(body)
		//fmt.Println(output)
		//fmt.Println("\n\nAMOUNT OF PULSES:::::: ", jsonResponse.PulseInfo.Count)
		if(jsonResponse.PulseInfo.Count == 0){
			response.Status = "Safe"
		}else{
			response.Status = "Risk"
		}

		response.SourceName="AlienVault"
	*/

	//response = string(body)
	utils.SetResponseObjectAlienVault(jsonResponse, response)
}
