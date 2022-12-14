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

// Retrieves data of uploaded file
func UploadFileRetrieve(c *gin.Context) {
	var fileData []byte

	// fetch based on url parameter, file = id
	id := c.Query("file")
	log.Println(id)
	if id == "" {
		log.Println("error, ID is empty")
		logging.Logerrorinfo("Error, ID is empty - Upload")
		http.Error(c.Writer, "No ID provided.", http.StatusInternalServerError)
		return
	}

	value, err := utils.Conn.Do("GET", "file:"+id)
	if value == nil {
		if err != nil {
			fmt.Println("Error:" + err.Error())
			logging.Logerror(err, "ERROR getting file, Fileupload API")
		}
		fmt.Println("No Cache hit")

		// Perform the request
		fileData, err = uploadFileRetrieveCall(id)
		if err != nil {
			log.Println("error, uploadFile id data request.")
			logging.Logerrorinfo("Error, performing api uploadFile id data request.")
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
		}

		// Add the data to the database
		response, err := utils.Conn.Do("SETEX", "file:"+id, utils.CacheDurationFile, fileData)
		if err != nil {
			fmt.Println("Error adding data to redis:" + err.Error())
			logging.Logerror(err, "ERROR adding data to Redis, Fileupload API")
		}

		fmt.Println(response)

	} else {
		fmt.Println("Cache hit")
		responseBytes, err := json.Marshal(value)
		if err != nil {
			fmt.Println("Error handling redis response:" + err.Error())
			logging.Logerror(err, "ERROR in RedisResponse, Fileupload API")
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(responseBytes, &fileData)
		if err != nil {
			fmt.Println("Error handling redis response:" + err.Error())
			logging.Logerror(err, "ERROR handling redis response, fileupload API")
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
		}
	}

	c.Data(http.StatusOK, "application/json", fileData)
}

// Performs request to virustotal, returning a full report of analyzed file
func uploadFileRetrieveCall(id string) (data []byte, err error) {
	var responseData utils.ResultFrontendResponse
	responseData, err = CallVirusTotal(id)
	if err != nil {
		return nil, err
	}

	fileData, err := json.Marshal(responseData)
	if err != nil {
		fmt.Println(err)
		logging.Logerror(err, "")
		return nil, err
	}

	return fileData, nil
}

// Uploads file to VirusTotal and returns an ID reference to the scan results in VirusTotal
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
		logging.Logerror(err, "")
	}
	// copy file locally
	_, err = io.Copy(part, file3)

	if err != nil {
		logging.Logerror(err, "")
	}
	// close writer
	err = writer.Close()

	if err != nil {
		log.Println(err)
		logging.Logerror(err, "")
	}

	// prepare request towards API
	req, err := http.NewRequest("POST", uri, body)

	if err != nil {
		log.Println(err)
		logging.Logerror(err, "")
	}

	// fetch API key
	APIKey := utils.APIKeyVirusTotal

	// add API key to relevant header
	req.Header.Add("X-Apikey", APIKey)

	// dynamically set content type, based on the formdata writer
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// perform the prepared API request
	res, err := http.DefaultClient.Do(req)

	// as long as the request returns 200
	if err != nil {
		log.Println(err)
		logging.Logerror(err, "")
	}

	defer res.Body.Close()

	// read the response
	contents, _ := ioutil.ReadAll(res.Body)

	var jsonResponse utils.VirusTotalUploadID

	// unmarshal contents
	unmarshalledID := json.Unmarshal(contents, &jsonResponse)

	if unmarshalledID != nil {
		log.Println(unmarshalledID)
		logging.Logerror(err, "")
	}

	// fetch ID which is base64 encoded
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
