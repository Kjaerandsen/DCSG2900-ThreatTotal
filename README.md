# DCSG2900-ThreatTotal:

Threat total is a threat intelligence service which allows you to get a quick overlook over the safety of using a particular website, 
domain or application. 

## Team members:

* Johannes Madsen Barstad
* Odin Korsfur Henriksen
* Jonas Kjærandsen
* Peder Andreas Stuen

## About:

The application is developed in react.js + tailwindcss for the frontend and golang with the gin web framework for the backend.
The backend is in the root directory of the project.
While the frontend is located in the "threat-total" folder.

## Data sources:
- Google Safebrowsing Api
- HybridAnalysis Falcon Sandbox Public Api
- Alienvault Open Threat Exchange Api
- VirustTotal Api


### Frontend structure:
Relevant files to look at lie in the src directory which is split into:
- The base directory containing the app files.
- The pages folder containing the different pages.
- The img folder containing images
- The components folder containing components used on the different pages.

### Backend structure:
- The main file is called main.go.
- Api functions are located in the "api" folder under the "api" package
- Utilitites such as structs and constants are located in the "utils" folder under the "utils" package


# Development:

## Frontend:

Go to the threat-total folder

To install the dependencies run:

`npm i`

To start the development server run:

`npm start`

To generate css while working run:

`npx tailwindcss -i ./src/input.css -o ./src/output.css --watch`

## Backend:

Open the main folder.

Install the dependencies and run the backend with:

`go run main.go`

For the backend to run you need to set the following environemnt variables:

"clientId": Set to the clientId of your feide application.

"clientSecret": Set to the clientSecret of your feide application.

"feideRedirectUrl": Set to your feide redirect url.

"APIKeyOTX": Set to your OTX api key.

"APIKeyHybridAnalysis": Set to your Hybrid Analysis api key.

"APIKeyVirusTotal": Set to your VirusTotal api key.

"APIKeyGoogle": Set to your Google api key.

"redisPassword": Set to your redis instance password

"redisUrl": Set to your redis instance url in the format "ip:port"

For the frontend you need to create a `.env` file under the `threat-total` folder. It should contain the following on seperate lines:

"REACT_APP_FEIDELOGINURL=": Set to the login url to feide, including a client_id, redirect_uri, response_type and a state. For example:
For example: https://auth.dataporten.no/oauth/authorization?client_id=YourIDHere&response_type=code&redirect_uri=YourRedirectHere&scope=openid&state=whatever

"REACT_APP_BACKEND_URL=": Set to the port and ip of your backend. For example for localhost: https://127.0.0.1:8081

And have a redis instance up and running with the following config:

Password set to the password defined in your environment variable.
`CONFIG SET requirepass "your password here"`
If your redis instance is running on a different ip address you also need to set
protected mode to no.
Which can be done through the redis-cli with `CONFIG SET protected-mode no`

***TESTING***
This project has implemented API and Unit testing through main_test.go

The implemented tests are maninly for the url - and hash - intelligence endpoints and test various functionality contained in the endpoint as well as expected return values.

There is also a test implemented to test wether an unspecified endpoint returns code 404 - Not found. 

***TEST OVERVIEW***

Test function: TestUrlIntelligenceOK
What it does: This function tests if the url-intelligence endpoint returns the expected status code 200 when called as logged in user.

Test function: TestUrlIntelligenceUnauthorized
What it does: This function tests if the url intelligence endpoints returns the expected status code 401 when called without the user log in.

Test function: TestHashIntelligenceOK
What it does: This function tests if the hash-intelligence endpoint returns the expected status code 200 when called as logged in user.

Test function: TestHashIntelligenceUnauthorized
What it does: This function tests if the hash intelligence endpoints returns the expected status code 401 when called without the user log in.

Test function: TestUrlIntelligenceValidOutput
What it does: 
* API test to check whether the url-intelligence endpoint returns valid ouput
* This test runs multiple tests, and tests the following:

* If status code is 200
* If the data can be unmarshalled to the struct ResultResponse
* If data can be accessed in the struct
* If the first sourceName is "Google Safebrowsing API" as expected
* If there is a screenshot of the requested URL
* If status or content is not set in any of the responses from the intelligence sources both in english and norwegian


Test function: TestHash_IntelligenceValidOutput
What it does:
* API test to check whether the hash-intelligence endpoint returns valid ouput
* This test runs multiple tests, and tests the following:
*
* If status code is 200
* If the data can be unmarshalled to the struct ResultResponse
* If data can be accessed in the struct
* If the first and second sourceName is "Hybrid Analysis and AlienVault" respectively, as expected
* If status or content is not set in any of the responses from the intelligence sources both in english and norwegian.
* If the status of AlienVault is risk as expected.


Test function: TestNotSpecifiedEndpoint
What it does:  This API test checks if an unspecified endpoint in the API returns 404 as expected 

**HOW TO USE TESTING ***

To use testing:

1. Create a file named testauth.txt in directory /dcsg2900-threattotal
2. Start the backend main.go file using go run main.go
3. Start the frontend webserver using npm start in directory /dcsg2900-threattotal/threat-total
4. Log to the application at localhost:3000
5. Maneuver to the local storage tab of your webbrowserver in google chrome it is located under application -> local storage.
6. Copy the value of the userAuth token stored in local storage.
7. Paste the value of the userAuth token into the .txt file created named testauth.txt.
8. Run the tests using command go test -v from the /dcsg2900-threattotal directory. 