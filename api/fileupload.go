package api

import (
	"bytes"
	logging "dcsg2900-threattotal/logs"
	"dcsg2900-threattotal/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// to be implemented, functions for displaying fileupload results to frontend

func UploadFile(c *gin.Context) {
	log.Println("Fileupload worked")
	logging.Loginfo("Fileupload worked")

	uri := "https://www.virustotal.com/api/v3/files"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// fetch the file contents
	file2, _ := c.FormFile("file")
	// open the file
	file3, _ := file2.Open()

	// use file contents to fetch file name, associate it with the "file" form header.
	part, err := writer.CreateFormFile("file", file2.Filename)

	if err != nil {
		log.Println(err)
	}
	// copy file locally
	_, err = io.Copy(part, file3)

	if err != nil {
		logging.Logerror(err)
	}
	// close writer
	err = writer.Close()

	if err != nil {
		log.Println(err)
		logging.Logerror(err)
	}

	// prepare request towards API
	req, err := http.NewRequest("POST", uri, body)

	if err != nil {
		log.Println(err)
		logging.Logerror(err)
	}

	content, err := ioutil.ReadFile("./APIKey/virusTotal.txt")
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		logging.Logerror(err)
	}

	APIKey := string(content)

	req.Header.Add("X-Apikey", APIKey)
	// error handle here, user should not be able to send requests without api key

	// dynamically set content type, based on the formdata writer
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// perform the prepared API request
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err)
		logging.Logerror(err)
	}

	defer res.Body.Close()

	// så lenge status 200

	// read the response
	contents, _ := ioutil.ReadAll(res.Body)

	var jsonResponse utils.VirusTotalUploadID

	unmarshalledID := json.Unmarshal(contents, &jsonResponse)

	if unmarshalledID != nil {
		log.Println(unmarshalledID)
		logging.Logerror(err)
	}

	encodedID := jsonResponse.Data.ID

	// decode provided values for virustotal report
	decode64, err := base64.RawStdEncoding.DecodeString(encodedID)
	log.Println("decoded here")
	log.Println(string(decode64))

	// extract ID for virustotal report
	trimID := strings.Split((string(decode64)), ":")
	log.Println("TRIMMED")
	log.Println(trimID[0])
	if err != nil {
		log.Println(err)
	}
	// return json object ID
	c.JSON(http.StatusOK, gin.H{"id": trimID[0]})
	// handle error
}

func UploadFileRetrieve(c *gin.Context) {
	// VT key has been added. REMEMBER TO DEACTIVATE AND CHANGE BEFORE FINAL RELEASE.
	// prepare request towards API

	content, err := ioutil.ReadFile("./APIKey/virusTotal.txt")
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		logging.Logerror(err)
	}
	// Convert []byte to string and print to screen
	APIKey := string(content)
	// remember to change api key, and reference it to a file instead
	// as well as deactivate the key from the account, as it's leaked.

	// fetch based on url parameter, file = id
	id := c.Query("file")
	log.Println(id)
	if id == "" {
		log.Println("error, ID is empty")
		logging.Logerrorinfo("Error, ID is empty - Upload")
	}

	url := fmt.Sprintf("https://www.virustotal.com/api/v3/files/%s", id)
	log.Println(url)
	log.Println(id)

	vtReq, _ := http.NewRequest("GET", url, nil)

	vtReq.Header.Add("Accept", "application/json")

	vtReq.Header.Add("X-Apikey", APIKey)

	vtRes, _ := http.DefaultClient.Do(vtReq)

	//defer vtRes.Body.Close()

	vtBody, _ := ioutil.ReadAll(vtRes.Body)

	log.Println(string(vtBody))

	var vtResponse utils.FileUploadData2

	unmarshalledBody := json.Unmarshal(vtBody, &vtResponse)

	if unmarshalledBody != nil {
		log.Println(unmarshalledBody)
	}

	/*
		test2 := reflect.ValueOf(vtResponse.Data.Attributes.LastAnalysisResults)
		lengthOfTest2 := test2.NumField() */

	var test3 = make([]utils.FrontendResponse4, len(vtResponse.Data.Attributes.LastAnalysisResults))

	//var test [4]utils.FrontendResponse4

	log.Println("here is the test output we maybe want")
	i := 0

	// TODO for later, remove teststruct, as it's only used to put into totalverdict, later
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

	var totalVerdict utils.ResultFrontendResponse
	totalVerdict.FrontendResponse = testStruct
	// IMPORTANT TODO, FIGURE ROUTING

	// Possible to add more cases in the future, for more accurate assessements
	// very realisitc that we need more cases, too narrow for accurate results per now.
	if vtResponse.Data.Attributes.LastAnalysisStats.Malicious == 0 && vtResponse.Data.Attributes.LastAnalysisStats.Suspicious == 0 {
		totalVerdict.EN.Result = "File is safe."
		// osv totalVerdict.EN.Result = fmt.Sprintf("File is considered safe", x av y)
	} else if vtResponse.Data.Attributes.TotalVotes.Malicious > 0 && vtResponse.Data.Attributes.LastAnalysisStats.Suspicious >= 0 {
		totalVerdict.EN.Result = "File has malicious indicators, consider escalating. "
	} else if vtResponse.Data.Attributes.LastAnalysisStats.Harmless > 0 && vtResponse.Data.Attributes.LastAnalysisStats.Malicious == 0 {
		totalVerdict.EN.Result = "File has been confirmed benign."
	} else {
		totalVerdict.EN.Result = "File is suspicious."
	}

	var engines int = len(vtResponse.Data.Attributes.LastAnalysisResults)

	utils.SetResultFile(&totalVerdict, engines)

	log.Println("look here")
	log.Println(totalVerdict)

	fmt.Println(totalVerdict)

	// return here

	fileInt, err := json.Marshal(totalVerdict)
	if err != nil {
		fmt.Println(err)
		//c.Data(http.StatusInternalServerError, "application/json", )
	}

	//fmt.Println("WHERE IS MY CONTENT", responseData)

	c.Data(http.StatusOK, "application/json", fileInt)

	// sort list based on what is malicious
	//log.Print(test3)

	// hent resultat via cache

	// total votes feltet virker relevant
	// LAST ANALYSIS STATS - MALICIOUS
}
