package api

import (
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func UrlIntelligence(c *gin.Context) {
	url := c.Query("url")

	var URLint []byte

	value, err := utils.Conn.Do("GET", url)
	if value == nil {
		if err != nil {
			fmt.Println("Error:" + err.Error())
		}
		fmt.Println("No Cache hit")

		URLint, err = urlSearch(url)
		if err != nil {
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
		}

		// Add the data to the redis backend.
		response, err := utils.Conn.Do("SETEX", url, 60, URLint)
		if err != nil {
			fmt.Println("Error adding data to redis:" + err.Error())
		}

		// Print the response to adding the data (should be "OK")
		fmt.Println(response)

		//fmt.Println("WHERE IS MY CONTENT 2", responseData2)
		// Cache hit
	} else {
		fmt.Println("Cache hit")
		responseBytes, err := json.Marshal(value)
		if err != nil {
			fmt.Println("Error handling redis response:" + err.Error())
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
			// Maybe do another call to delete the key from the database?
		}
		err = json.Unmarshal(responseBytes, &URLint)
		if err != nil {
			fmt.Println("Error handling redis response:" + err.Error())
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
			// Maybe do another call to delete the key from the database?
		}
	}

	c.Data(http.StatusOK, "application/json", URLint)
}

// Makes the api requests used in urlIntelligence
func urlSearch(url string) (data []byte, err error) {
	var wg sync.WaitGroup //Vente gruppe for goroutiner
	var URLint []byte
	var responseData [4]utils.FrontendResponse2

	var p, VirusTotal, urlscanio, alienvault *utils.FrontendResponse2
	p = &responseData[0]
	VirusTotal = &responseData[1]
	urlscanio = &responseData[2]
	alienvault = &responseData[3]

	fmt.Println(url)

	wg.Add(3)
	go TestGoGoogleUrl(url, p, &wg)
	go TestHybridAnalyisUrl(url, VirusTotal, urlscanio, &wg)
	go TestAlienVaultUrl(url, alienvault, &wg)
	wg.Wait()

	//responseData2 := FR122(responseData[:])

	URLint, err = json.Marshal(responseData)
	if err != nil {
		fmt.Println(err)
		return URLint, err
	}

	fmt.Println("WHERE IS MY CONTENT 1", responseData)

	return URLint, nil
}
