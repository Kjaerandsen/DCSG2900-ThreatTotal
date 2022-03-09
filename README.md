# DCSG2900-ThreatTotal:

Threat total is a threat intelligence service which allows you to get a quick overlook over the safety of using a particular website, 
domain or application. 
We retrieve data from the NTNU soc database, as well as external sources such as: .... 

## Team members:

* Johannes Madsen Barstad
* Odin Korsfur Henriksen
* Jonas Kjærandsen
* Peder Andreas Stuen

## About

The application is developed in react.js + tailwindcss for the frontend and golang with the gin web framework for the backend.

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