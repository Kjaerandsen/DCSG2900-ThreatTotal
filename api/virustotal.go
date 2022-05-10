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
		logging.Logerror(unmarshalledBody, "")
	}

	log.Println("here is the test output we maybe want")
	i := 0

	// TODO for later, remove teststruct, as it's only used to put into response, later
	var testStruct = make([]utils.FrontendResponse2, len(vtResponse.Data.Attributes.LastAnalysisResults))

	// iterate through results
	for _, val := range vtResponse.Data.Attributes.LastAnalysisResults {
		//log.Printf("testing, %s, %s", key, val)
		// initialize struct
		// print
		log.Println(val)

		if val.Category == "undetected" ||
			val.Category == "malicious" ||
			val.Category == "malicious" ||
			val.Category == "malicious" {
			// save engine name
			testStruct[i].ID = i + 1
			testStruct[i].SourceName = val.EngineName
			// resolution of AV
			testStruct[i].EN.Status = val.Category

			testStruct[i].EN.Content = vtResponse.Data.Attributes.MeaningfulName
			testStruct[i].EN.Description = vtResponse.Data.Attributes.Magic
			testStruct[i].EN.Tags = vtResponse.Data.Attributes.TypeTag

			//testStruct.EN.Description =

			// can also display the total status (last analysis stats)
			// this is an int ^^ so cant fill it in frontendresponse2
			// question is, do we do it here or later

			i++
		}
	}
	log.Println(testStruct)
	var testStruct2 = make([]utils.FrontendResponse2, i)
	testStruct2 = testStruct[0:(i - 1)]

	totalDanger := vtResponse.Data.Attributes.LastAnalysisStats.Malicious + vtResponse.Data.Attributes.LastAnalysisStats.Suspicious

	response.FrontendResponse = sortDanger(testStruct2, totalDanger, i-totalDanger)
	// IMPORTANT TODO, FIGURE ROUTING

	// Possible to add more cases in the future, for more accurate assessements
	// very realisitc that we need more cases, too narrow for accurate results per now.
	if vtResponse.Data.Attributes.LastAnalysisStats.Malicious == 0 && vtResponse.Data.Attributes.LastAnalysisStats.Suspicious == 0 {
		response.EN.Result = "File is safe."
		response.NO.Result = "Filen er trygg"
		// osv response.EN.Result = fmt.Sprintf("File is considered safe", x av y)
	} else if vtResponse.Data.Attributes.TotalVotes.Malicious > 0 && vtResponse.Data.Attributes.LastAnalysisStats.Suspicious >= 0 {
		response.EN.Result = "File has malicious indicators, consider escalating to the NTNU SOC. "
		response.NO.Result = "Filen har ondsinnede indikatorer, vennligst vurder å eskalere videre til NTNU SOC"
	} else if vtResponse.Data.Attributes.LastAnalysisStats.Harmless > 0 && vtResponse.Data.Attributes.LastAnalysisStats.Malicious == 0 {
		response.EN.Result = "File has been confirmed benign. Further handling of the file is safe"
		response.NO.Result = "Filen er bekreftet godartet, videre håndtering av fil er trygt."
	} else {
		response.EN.Result = "File is suspicious. It is not recommended to further handle this file."
		response.NO.Result = "Filen er mistenkelig. Det anbefales å ikke videre håndtere filen. "
	}

	//var engines int = len(vtResponse.Data.Attributes.LastAnalysisResults)

	utils.SetResultFile(&response, i-1)

	log.Println(response)

	fmt.Println(response)

	return response, nil
}

func sortDanger(values []utils.FrontendResponse2, dangerSize int, safeSize int) []utils.FrontendResponse2 {
	if dangerSize == 0 {
		return values
	}
	var dangerous = make([]utils.FrontendResponse2, dangerSize+1)
	var safe = make([]utils.FrontendResponse2, safeSize+1)
	var i, j = 0, 0

	for l := 0; l < dangerSize+safeSize-1; l++ {
		if values[l].EN.Status == "harmless" || values[l].EN.Status == "undetected" {
			safe[i] = values[l]
			i++
		} else {
			dangerous[j] = values[l]
			j++
		}
	}

	for l := 0; l < dangerSize-1; l++ {
		values[l] = dangerous[l]
		values[l].ID = l
	}

	for l := 0; l < safeSize-1; l++ {
		values[l+dangerSize] = safe[l]
		values[l+dangerSize].ID = l + dangerSize
	}

	return values
}
