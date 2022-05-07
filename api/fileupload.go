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

	value, err := utils.Conn.Do("GET", id)
	if value == nil {
		if err != nil {
			fmt.Println("Error:" + err.Error())
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
		response, err := utils.Conn.Do("SETEX", id, 300, fileData)
		if err != nil {
			fmt.Println("Error adding data to redis:" + err.Error())
		}

		fmt.Println(response)

	} else {
		fmt.Println("Cache hit")
		responseBytes, err := json.Marshal(value)
		if err != nil {
			fmt.Println("Error handling redis response:" + err.Error())
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
			// Maybe do another call to delete the key from the database?
		}
		/**
		//var checkData utils.ResultFrontendResponse
		err = json.Unmarshal(responseBytes, &checkdata)
		if err!=nil {
			fmt.Println(string(checkData))
		}
		fmt.Println(string(checkData))
		*/
		err = json.Unmarshal(responseBytes, &fileData)
		if err != nil {
			fmt.Println("Error handling redis response:" + err.Error())
			http.Error(c.Writer, "Failed retrieving api data.", http.StatusInternalServerError)
			return
			// Maybe do another call to delete the key from the database?
		}
	}

	c.Data(http.StatusOK, "application/json", fileData)
}

func uploadFileRetrieveCall(id string) (data []byte, err error) {
	var responseData utils.ResultFrontendResponse
	responseData, err = CallVirusTotal(id)
	if err != nil {
		return nil, err
	}

	fileData, err := json.Marshal(responseData)
	if err != nil {
		fmt.Println(err)
		logging.Logerror(err)
		return nil, err
	}

	return fileData, nil
}

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
