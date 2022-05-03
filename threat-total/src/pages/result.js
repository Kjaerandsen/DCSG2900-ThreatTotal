import React, { useEffect, useState} from "react";
import Navbar from "../components/navbar";
import Sources from "../components/sources";
import { useTranslation } from 'react-i18next';
import CookieDisclosure from "../components/cookieDisclosure";
import { Oval } from 'react-loader-spinner';
import i18next from '../i18next';

function Result() {
    const queryParams = new URLSearchParams(window.location.search);
    const hash = queryParams.get('hash');
    const url = queryParams.get('url');
    const [JsonData, setJsonData] = useState([""])
    const [Err, setErr] = useState(false)
    const { t } = useTranslation();
    const [isLoading, setIsLoading] = useState(false);
    
    useEffect(() => {
        if (hash != null) {
            console.log({hash})
            setIsLoading(true);
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
                setIsLoading(false)
            })
            .catch(function(error){
                console.log(error)
                setErr(true)
                setIsLoading(false)
            })
        } else if (url != null){
            setIsLoading(true);
            // Send an api request to the backend with url data
            fetch('http://localhost:8081/url-intelligence?url=' + url, {
                method: 'GET',
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json'
                }
            }).then((response) => response.json())
            .then((json) => {
                setJsonData(json);
                setIsLoading(false)
            })
            .catch(function(error){
                console.log(error)
                setErr(true)
                setIsLoading(false)
            })
        } else {
            // Redirect to index if no parameters are provided
            //window.location.href= "/"
            // Send an api request to the backend with url data
            setIsLoading(true);
            fetch('http://localhost:8081/result', {
                method: 'GET',
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json'
                }
            }).then((response) => response.json())
            .then((json) => {
                setJsonData(json)
                setIsLoading(false)
            })
            .catch(function(error){
                console.log(error)
                setErr(true)
                setIsLoading(false)
            })
        }
        // Need error handling when the backend is unavailable
    }, []);

    console.log(JsonData)
    if (JsonData != undefined){
        i18next.addResources('en', 'translation', JsonData.EN);
        i18next.addResources('no', 'translation', JsonData.NO);
    }

    const renderResult = (
        <div className="container text-center break-words sm:justify-center">
        
            <h1 className="text-4xl font-bold p-0 mb-8 sm:mt-12 sm:mb-12 w-auto">
                {t("resultTitle")}
            </h1>
            <p className="text-middle text-2xl m-2 pl-2 pr-2 sm:pl-16 sm:pr-16 xl:pl-36 xl:pr-36">
                {t("Result")}
                <br></br><br></br>
                <br></br>
            </p>
            <div className="container">
                
                <Sources sourceData = {JsonData} err = {Err}/>

            </div>
        </div>
    );
    
    return (
        <div className="grid place-items-center">
        
            <Navbar />

            <div className="pt-10 pb-10">
            {isLoading ? <Oval height="100" width="100" color="grey"/> : renderResult}
            </div>
            
        <div className= "container w-full mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
            <a href="./investigate">
                <button className="bg-orange-500 p-2 rounded justify-center">{t("manualAnalysisBtn")}</button>
            </a>
        </div>

        <CookieDisclosure />

        </div>
        );
}

export default Result;
