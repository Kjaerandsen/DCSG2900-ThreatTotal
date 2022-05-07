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
