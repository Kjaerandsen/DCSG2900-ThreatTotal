# DCSG2900-ThreatTotal:

Threat total is a threat intelligence service which allows you to get a quick overlook over the safety of using a particular website, 
domain or application. 
We retrieve data from the NTNU soc database, as well as external sources such as: .... 

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

Install the dependencies with:

`go mod download`

Run the backend with:

`go run main.go`
