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

// CallAlienVaultHash function takes a hash, returns data on it from the alienvault api
func CallAlienVaultHash(hash string, response *utils.FrontendResponse2, wg *sync.WaitGroup) {

	defer wg.Done()
	response.SourceName = "AlienVault"	//Adds sourcename

	APIKey := utils.APIKeyOTX		//Gets API key

	getURL := "https://otx.alienvault.com//api/v1/indicators/file/" + hash + "/general"		//Sets the endpoint URL

	req, err := http.NewRequest("GET", getURL, nil)
	req.Header.Set("X-OTX-API-KEY", APIKey)

	client := &http.Client{}

	res, err := client.Do(req)
	
	if err != nil {
		fmt.Println("ERROR IN Request", err)
		logging.Logerror(err, "ERROR IN REQUEST, AlienVault API")
		utils.SetGenericError(response)
	}
	if res.StatusCode == 200 {		//Checks Statuscode IF ok, continue. 
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

		utils.SetResponseObjectAlienVaultHash(jsonResponse, response)
	} else {

		response.EN.Content = "We have encountered an error, check if the filehash is a valid filehash."
		response.EN.Status = "ERROR"

		response.NO.Content = "Vi har møtt på en error, sjekk om filhashen er en gyldig filhash."
		response.NO.Status = "ERROR"
	}
}

func CallAlienVaultUrl(url string, response *utils.FrontendResponse2, wg *sync.WaitGroup) {
	defer wg.Done()
	
	APIKey := utils.APIKeyOTX

	getURL := "https://otx.alienvault.com//api/v1/indicators/url/" + url + "/general"

	req, err := http.NewRequest("GET", getURL, nil)
	req.Header.Set("X-OTX-API-KEY", APIKey)

	

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

	utils.SetResponseObjectAlienVault(jsonResponse, response)
}
