import React, { useEffect, useState} from "react";
import Navbar from "../components/navbar";
import Sources from "../components/sources";
import { useTranslation } from 'react-i18next';
import i18n from "../i18next";

function Result() {
    const queryParams = new URLSearchParams(window.location.search);
    const hash = queryParams.get('hash');
    const url = queryParams.get('url');
    const [JsonData, setJsonData] = useState([""])
    const [Err, setErr] = useState(false)
    const { t } = useTranslation();
    
    useEffect(() => {
        if (hash != null) {
            console.log({hash})
            // Send an api request to the backend with hash data
            fetch('http://localhost:8081/hash-intelligence?hash=' + hash, {
                method: 'GET',
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json'
                }
            }).then((response) => response.json())
            .then((json) => {
                setJsonData(json)
                //setJsonData(JSON.parse(json))
            })
            .catch(function(error){
                console.log(error)
                setErr(true)
            })
        } else if (url != null){
            // Send an api request to the backend with url data
            fetch('http://localhost:8081/url-intelligence?url=' + url, {
                method: 'GET',
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json'
                }
            }).then((response) => response.json())
            .then((json) => {
                setJsonData(json)
            })
            .catch(function(error){
                console.log(error)
                setErr(true)
            })
        } else {
            // Redirect to index if no parameters are provided
            //window.location.href= "/"
            // Send an api request to the backend with url data
            fetch('http://localhost:8081/result', {
                method: 'GET',
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json'
                }
            }).then((response) => response.json())
            .then((json) => {
                setJsonData(json)
            })
            .catch(function(error){
                console.log(error)
                setErr(true)
            })
        }
        // Need error handling when the backend is unavailable
    }, []);

    console.log(JsonData)

    return (
        <div className="grid place-items-center">
        
            <Navbar />

        <div className="container text-center break-words sm:justify-center">
            <h1 className="text-3xl font-bold p-0 mt-8 mb-8 sm:mt-12 sm:mb-12 w-auto">
                {t("resultTitle")}
            </h1>
            <p className="text-left m-2 pl-2 pr-2 sm:pl-16 sm:pr-16 xl:pl-36 xl:pr-36">
            This page poses a risk or potential risk to visit according to 2/3 of our sources. 
            Please use proper caution and avoid visiting at all if possible. 
            <br></br><br></br>Forandre denne til en api respons?
            <br></br>
            <br></br>
            </p>
            <div className="container">
                
                <Sources sourceData = {JsonData} err = {Err}/>

            </div>
        </div>
            
        <div className= "container w-full mt-1.5 mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
            <a href="./investigate">
                <button className="bg-orange-500 p-2 rounded justify-center">{t("manualAnalysisBtn")}</button>
            </a>
        </div>


        </div>
        );
}

export default Result;
