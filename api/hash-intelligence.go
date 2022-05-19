package api

import (
	logging "dcsg2900-threattotal/logs"
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

func HashIntelligence(c *gin.Context) {

	var hashInt []byte
	var err error
	var completeInt bool

	hash := strings.TrimSpace(c.Query("hash"))

	value, err := utils.Conn.Do("GET", "hash:"+hash)
	if value == nil {
		if err != nil {
			fmt.Println("Error:" + err.Error())
			logging.Logerror(err, "Error in retrieving cache - hash-intelligence")
		}
		fmt.Println("No Cache hit")

		// Perform the request
		hashInt, err, completeInt = hashSearch(hash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error ": "Invalid response from third party API's."})
			return
		}

		if completeInt {
			// Add the data to the database
			response, err := utils.Conn.Do("SETEX", "hash:"+hash, utils.CacheDurationHash, hashInt)
			if err != nil {
				fmt.Println("Error adding data to redis:" + err.Error())
				logging.Logerror(err, "Error adding data to redis, hash-intelligence")

			}

			fmt.Println(response)
		}

	} else {

		fmt.Println("Cache hit")
		responseBytes, err := json.Marshal(value)
		if err != nil {
			fmt.Println("Error handling redis response:" + err.Error())
			logging.Logerror(err, "ERROR handling redis response, hash-intelligence")
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(responseBytes, &hashInt)
		if err != nil {
			fmt.Println("Error handling redis response:" + err.Error())
			logging.Logerror(err, "Error unmarshalling response, hash-intelligence")
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
		}
	}

	c.Data(http.StatusOK, "application/json", hashInt)
}

func hashSearch(hash string) (data []byte, err error, complete bool) {

	var wg sync.WaitGroup
	var responseData [2]utils.FrontendResponse2

	var hybridApointer, AlienVaultpointer *utils.FrontendResponse2

	hybridApointer = &responseData[0]
	AlienVaultpointer = &responseData[1]

	wg.Add(2)
	go CallHybridAnalysisHash(hash, hybridApointer, &wg)	//Threaded calls to hybrid analysis and Alien vault intelligence gathering. 
	go CallAlienVaultHash(hash, AlienVaultpointer, &wg)
	wg.Wait()			//Wait till both functions are complete. 

	var resultResponse utils.ResultFrontendResponse		//Creates a new datastructure to add data to. 

	resultResponse.FrontendResponse = responseData[:]
	var resultPointer = &resultResponse

	utils.SetResultHash(resultPointer, len(responseData))		//Calls function to set the result string based on data. 

	complete = checkIfIntelligenceCompleteHash(resultResponse, len(responseData))		//Checks if data is complete & ready to be cached. 
	hashInt, err := json.Marshal(resultResponse)
	if err != nil {
		fmt.Println(err)
		logging.Logerror(err, "")
		return nil, err, complete
	}

	return hashInt, nil, complete		//Returns the intelligence, a nil error, and a bool if the intelligence is complete. 
}


//This function checks if the intelligence is complete. 
func checkIfIntelligenceCompleteHash(jsonData utils.ResultFrontendResponse, size int) (complete bool) {
	complete = true

	for i := 0; i <= size-1; i++ {
		if jsonData.FrontendResponse[i].EN.Status == "Awaiting analysis" || jsonData.FrontendResponse[i].EN.Status == "Error" {
			complete = false
		}
	}

	return complete
}
