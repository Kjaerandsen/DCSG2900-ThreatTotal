package api

import (
	logging "dcsg2900-threattotal/logs"
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CallVirusTotal returns data on a file id from virustotal
func CallVirusTotal(id string) (response utils.ResultFrontendResponse, err error) {
	// VT key has been added. REMEMBER TO DEACTIVATE AND CHANGE BEFORE FINAL RELEASE.
	// prepare request towards API
	// Convert []byte to string and print to screen
	APIKey := utils.APIKeyVirusTotal
	// remember to change api key, and reference it to a file instead
	// as well as deactivate the key from the account, as it's leaked.
	if id == "" {
		log.Println("error, ID is empty")
		logging.Logerrorinfo("Error, ID is empty - Upload")
	}

	url := fmt.Sprintf("https://www.virustotal.com/api/v3/files/%s", id)
	log.Println(url)
	log.Println(id)

	vtReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	vtReq.Header.Add("Accept", "application/json")

	vtReq.Header.Add("X-Apikey", APIKey)

	vtRes, err := http.DefaultClient.Do(vtReq)
	if err != nil {
		return
	}

	//defer vtRes.Body.Close()

	vtBody, err := ioutil.ReadAll(vtRes.Body)
	if err != nil {
		return
	}

	log.Println(string(vtBody))

	var vtResponse utils.FileUploadData2
	unmarshalledBody := json.Unmarshal(vtBody, &vtResponse)

	if unmarshalledBody != nil {
		log.Println(unmarshalledBody)
	}

	var test3 = make([]utils.FrontendResponse4, len(vtResponse.Data.Attributes.LastAnalysisResults))

	log.Println("here is the test output we maybe want")
	i := 0

	// TODO for later, remove teststruct, as it's only used to put into response, later
	var testStruct = make([]utils.FrontendResponse2, len(vtResponse.Data.Attributes.LastAnalysisResults))

	// iterate through results
	for _, val := range vtResponse.Data.Attributes.LastAnalysisResults {
		//log.Printf("testing, %s, %s", key, val)
		// initialize struct
		test3[i] = val
		// print
		log.Println(test3[i])

		// save engine name
		testStruct[i].ID = i + 1
		testStruct[i].SourceName = test3[i].EngineName
		// resolution of AV
		testStruct[i].EN.Status = test3[i].Category

		testStruct[i].EN.Content = vtResponse.Data.Attributes.MeaningfulName
		testStruct[i].EN.Description = vtResponse.Data.Attributes.Magic
		testStruct[i].EN.Tags = vtResponse.Data.Attributes.TypeTag

		//testStruct.EN.Description =

		// can also display the total status (last analysis stats)
		// this is an int ^^ so cant fill it in frontendresponse2
		// question is, do we do it here or later

		i++
	}
	log.Println(testStruct)

	response.FrontendResponse = testStruct
	// IMPORTANT TODO, FIGURE ROUTING

	// Possible to add more cases in the future, for more accurate assessements
	// very realisitc that we need more cases, too narrow for accurate results per now.
	if vtResponse.Data.Attributes.LastAnalysisStats.Malicious == 0 && vtResponse.Data.Attributes.LastAnalysisStats.Suspicious == 0 {
		response.EN.Result = "File is safe."
		// osv response.EN.Result = fmt.Sprintf("File is considered safe", x av y)
	} else if vtResponse.Data.Attributes.TotalVotes.Malicious > 0 && vtResponse.Data.Attributes.LastAnalysisStats.Suspicious >= 0 {
		response.EN.Result = "File has malicious indicators, consider escalating. "
	} else if vtResponse.Data.Attributes.LastAnalysisStats.Harmless > 0 && vtResponse.Data.Attributes.LastAnalysisStats.Malicious == 0 {
		response.EN.Result = "File has been confirmed benign."
	} else {
		response.EN.Result = "File is suspicious."
	}

	var engines int = len(vtResponse.Data.Attributes.LastAnalysisResults)

	utils.SetResultFile(&response, engines)

	log.Println("look here")
	log.Println(response)

	fmt.Println(response)

	return response, nil
}
