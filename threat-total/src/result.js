import React, { useEffect, useState} from "react";
import Navbar from "./navbar";
import Sources from "./sources";

function Result() {
    const queryParams = new URLSearchParams(window.location.search);
    const hash = queryParams.get('hash');
    const url = queryParams.get('url');
    const [JsonData, setJsonData] = useState([""])

    useEffect(() => {
        if (hash != null) {
            console.log({hash})
            // Send an api request to the backend with hash data
            fetch('http://localhost:8081/result?hash=' + hash, {
                method: 'GET',
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json'
                }
            }).then((response) => response.json())
            .then((json) => {
            })
        } else if (url != null){
            // Send an api request to the backend with url data
            fetch('http://localhost:8081/result?url=' + url, {
                method: 'GET',
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json'
                }
            }).then((response) => response.json())
            .then((json) => {
                setJsonData(JSON.parse(json))
            })
        } else {
            // Redirect to error 404 page / 50x for internal issue? or issue diplay?
            setJsonData(JSON.parse("[]"))
            // Show an error message, and show a redirect to search page button
            console.log("Invalid parameter")
        }
        // Need error handling when the backend is unavailable
    }, []);

    console.log(JsonData)

    return (
        <>
        <div className="grid place-items-center">
        
            <Navbar />

        <div className="bg-red-500 container text-center break-words sm:justify-center">
            <h1 className="text-3xl font-bold p-0 mt-8 mb-8 sm:mt-12 sm:mb-12 w-auto">
                Results
            </h1>
            <p className="text-left m-2">
            This page poses a risk or potential risk to visit according to 2/3 of our sources. 
            Please use proper caution and avoid visiting at all if possible.
            <br></br>
            <br></br>
            </p>
            <div className="container">
                
                <Sources sourceData = {JsonData}/>

            </div>
        </div>
            
        <div className= "bg-green-300 container w-full mt-1.5 mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
            <a href="./investigate">
                <button className="bg-orange-500 p-2 rounded justify-center">Submit for Manual Analysis</button>
            </a>
        </div>


        </div> 
        </>
        );
}

export default Result;
