package api

import (
	logging "dcsg2900-threattotal/logs"
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func HashIntelligence(c *gin.Context) {

	var wg sync.WaitGroup
	hash := c.Query("hash")

	var responseData [2]utils.FrontendResponse2

	var hybridApointer, AlienVaultpointer *utils.FrontendResponse2

	hybridApointer = &responseData[0]
	AlienVaultpointer = &responseData[1]

	wg.Add(2)
	go CallHybridAnalysisHash(hash, hybridApointer, &wg)
	go CallAlienVaultHash(hash, AlienVaultpointer, &wg)
	wg.Wait()

	var resultResponse utils.ResultFrontendResponse

	resultResponse.FrontendResponse = responseData[:]
	var resultPointer = &resultResponse

	utils.SetResultHash(resultPointer, len(responseData))

	Hashint, err := json.Marshal(resultResponse)
	if err != nil {
		fmt.Println(err)
		logging.Logerror(err)
		//c.Data(http.StatusInternalServerError, "application/json", )
	}

	//fmt.Println("WHERE IS MY CONTENT", responseData)

	c.Data(http.StatusOK, "application/json", Hashint)

}
