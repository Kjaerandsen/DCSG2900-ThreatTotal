package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"dcsg2900-threattotal/utils"
	"encoding/json"
)

/**
* API - Test to check if the URL intelligence endpoint returns HTTP StatusOK when expected
*/

func TestUrlIntelligenceOK(t *testing.T) {

	content, err := ioutil.ReadFile("testauth.txt")
	if err != nil {
		log.Fatal(err)
	}

	auth := string(content)
	url := "http://localhost:8081/url-intelligence?url=testsafebrowsing.appspot.com/s/malware.html&userAuth=" + auth

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error: in request")
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Test failed, error code different from expected 200, received: %d", res.StatusCode)
	}
}

/**
* API test to check whether the url-intelligence endpoint will return 401 Unauthorized when attempting to be accessed without log in.
*/

func TestUrlIntelligenceUnauthorized(t *testing.T) {

	auth := "ThisShouldNotWork"
	url := "http://localhost:8081/url-intelligence?url=testsafebrowsing.appspot.com/s/malware.html&userAuth=" + auth

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error: in request")
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 401 {
		t.Fatalf("Test failed, error code different from expected 401, received: %d", res.StatusCode)
	}
}


/**
* API - Test to check if the hash intelligence endpoint returns HTTP StatusOK when expected
*/
func TestHashIntelligenceOK(t *testing.T){

	content, err := ioutil.ReadFile("testauth.txt")
	if err != nil {
		log.Fatal(err)
	}

	auth := string(content)
	url := "http://localhost:8081/hash-intelligence?hash=a7a665a695ec3c0f862a0d762ad55aff6ce6014359647e7c7f7e3c4dc3be81b7&userAuth=" + auth

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error: in request")
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Test failed, error code different from expected 200, received: %d", res.StatusCode)
	}
}

/**
* API test to check whether the hash-intelligence endpoint will return 401 Unauthorized when attempting to be accessed without log in.
*/
func TestHashIntelligenceUnauthorized(t *testing.T){
	auth := "ThisShouldNotWork"
	url := "http://localhost:8081/hash-intelligence?hash=a7a665a695ec3c0f862a0d762ad55aff6ce6014359647e7c7f7e3c4dc3be81b7&userAuth=" + auth

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error: in request")
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 401 {
		t.Fatalf("Test failed, error code different from expected 401, received: %d", res.StatusCode)
	}
}

/**
* API test to check whether the url-intelligence endpoint returns valid ouput
* This test runs multiple tests, and tests the following:
*
* If status code is 200
* If the data can be unmarshalled to the struct ResultResponse
* If data can be accessed in the struct
* If the first sourceName is "Google Safebrowsing API" as expected
* If there is a screenshot of the requested URL
* If status or content is not set in any of the responses from the intelligence sources
*
*/
func TestUrlIntelligenceValidOutput(t *testing.T) {

	content, err := ioutil.ReadFile("testauth.txt")
	if err != nil {
		log.Fatal(err)
	}

	auth := string(content)
	url := "http://localhost:8081/url-intelligence?url=testsafebrowsing.appspot.com/s/malware.html&userAuth=" + auth

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Error in request")
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Request to URL-Intelligence failed")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Test failed, error code different from expected 200, received: %d", res.StatusCode)
	}


	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error: reading api response")
	}

	var jsonResponse utils.ResultFrontendResponse

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		t.Fatalf("Error unmarshalling, test failed")
	}

	if jsonResponse.FrontendResponse[0].SourceName != "Google SafeBrowsing Api" {
		t.Fatalf("The first sourcename is not Google Safebrowsing API")
	}

	if len(jsonResponse.Screenshot) <= 2 {
		t.Fatalf("Error in screenshot")
	}

	for i := 0; i<=3; i++{ 
		if jsonResponse.FrontendResponse[i].EN.Status == "" {
			t.Fatalf("One status or more statuses are not set in english. Sourcename: %s , content is empty", jsonResponse.FrontendResponse[i].SourceName)
		}
		if jsonResponse.FrontendResponse[i].NO.Status == "" {
			t.Fatalf("One status or more statuses are not set in norwegian, Sourcename: %s is not set.", jsonResponse.FrontendResponse[i].SourceName)
		}
	}

	for i := 0; i<=3; i++{ 
		if jsonResponse.FrontendResponse[i].EN.Content == "" {
			t.Fatalf("One content or more contents are not set in english. Sourcename: %s , content is empty", jsonResponse.FrontendResponse[i].SourceName)
			
		}
		if jsonResponse.FrontendResponse[i].NO.Content == "" {
			t.Fatalf("One status or more contents are not set in norwegian. Sourcename: %s is not set.", jsonResponse.FrontendResponse[i].SourceName)
		}

	}
}

/**
* API test to check whether the url-intelligence endpoint returns valid ouput
* This test runs multiple tests, and tests the following:
*
* If status code is 200
* If the data can be unmarshalled to the struct ResultResponse
* If data can be accessed in the struct
* If the first and second sourceName is "Hybrid Analysis and AlienVault" respectively, as expected
* If status or content is not set in any of the responses from the intelligence sources both in english and norwegian.
* If the status of AlienVault is risk as expected.
*/
func TestHash_IntelligenceValidOutput(t *testing.T) {

	content, err := ioutil.ReadFile("testauth.txt")
	if err != nil {
		log.Fatal(err)
	}

	auth := string(content)
	url := "http://localhost:8081/hash-intelligence?hash=a7a665a695ec3c0f862a0d762ad55aff6ce6014359647e7c7f7e3c4dc3be81b7&userAuth=" + auth

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Error in request")
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Request to Hash-Intelligence failed")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Test failed, error code different from expected 200, received: %d", res.StatusCode)
	}


	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error: reading api response")
	}

	var jsonResponse utils.ResultFrontendResponse

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		t.Fatalf("Error unmarshalling json to struct, test failed")
	}

	if jsonResponse.FrontendResponse[0].SourceName != "Hybrid Analysis" {
		t.Fatalf("The first sourcename is not expected Google Safebrowsing API, output: %s", jsonResponse.FrontendResponse[0].SourceName)
	}

	if jsonResponse.FrontendResponse[1].SourceName != "AlienVault"{
		t.Fatalf("Unexpected sourcename expected AlienVault, reality: %s", jsonResponse.FrontendResponse[1].SourceName)
	}

	for i := 0; i<=1; i++{ 
		if jsonResponse.FrontendResponse[i].EN.Status == "" {
			t.Fatalf("One status or more statuses are not set in english, Sourcename: %s is not set.", jsonResponse.FrontendResponse[i].SourceName)
		}
		if jsonResponse.FrontendResponse[i].NO.Status == "" {
			t.Fatalf("One status or more statuses are not set in norwegian, Sourcename: %s is not set.", jsonResponse.FrontendResponse[i].SourceName)
		}
	}

	for i := 0; i<=1; i++{ 
		if jsonResponse.FrontendResponse[i].EN.Content == "" {
			t.Fatalf("One status or more contents are not set in english, Sourcename: %s is not set.", jsonResponse.FrontendResponse[i].SourceName)
		}
		if jsonResponse.FrontendResponse[i].NO.Content == "" {
			t.Fatalf("One status or more contents are not set in norwegian, Sourcename: %s is not set.", jsonResponse.FrontendResponse[i].SourceName)
		}
	}

	if jsonResponse.FrontendResponse[1].EN.Status != "Risk"{
		t.Fatalf("The status of AlienVault has is not as expected Risk, Status is: %s", jsonResponse.FrontendResponse[1].EN.Status)
	}
}

/**
* This API test checks if an unspecified endpoint in the API returns 404 as expected 
*
*/

func TestNotSpecifiedEndpoint(t *testing.T){

	url := "http://localhost:8081/ThisShouldNotExist"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Error in request")
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Request to Hash-Intelligence failed")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("Staus code did not return 404 as expected, code returned %d", res.StatusCode)
	}
}