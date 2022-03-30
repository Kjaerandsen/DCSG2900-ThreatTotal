package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CallGoogleUrl function takes a url, returns data on it from the google safebrowsing api
func CallGoogleUrl(url string) (response string) {
	//Google API returnerer [] om den ikke kjenner til domenet / URL. Kan bruke dette til å avgjøre om det er malicious eller ikke.

	content, err := ioutil.ReadFile("./APIKey/Apikey.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	APIKey := string(content)

	postURL := "https://safebrowsing.googleapis.com/v4/threatMatches:find?key=" + APIKey

	var jsonData = []byte(`
		{
			"client": {
			  "clientId":      "threattotal",
			  "clientVersion": "1.5.2"
			},
			"threatInfo": {
			  "threatTypes":      ["MALWARE", "SOCIAL_ENGINEERING"],
			  "platformTypes":    ["WINDOWS"],
			  "threatEntryTypes": ["URL"],
			  "threatEntries": [
				{"url": "https://` + url + `" },
				{"url": "http://` + url + `"}
			  ]
			}
		}`)

	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(jsonData))
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
	}
	//fmt.Println("response Body:", string(body))
	response = string(body)

	return response
}
