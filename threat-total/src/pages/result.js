import React, { useEffect, useState} from "react";
import Navbar from "../components/navbar";
import Sources from "../components/sources";
import SubNavbar from '../components/subNavbar'
import { useTranslation } from 'react-i18next';
import CookieDisclosure from "../components/cookieDisclosure";
import { Oval } from 'react-loader-spinner';
import i18next from '../i18next';
import Screenshot from "../components/displayScreenshot";

function Result() {
    const queryParams = new URLSearchParams(window.location.search);
    const hash = queryParams.get('hash');
    const url = queryParams.get('url');
    const file = queryParams.get('file');
    const [JsonData, setJsonData] = useState([""])
    const [Err, setErr] = useState(false)
    const { t } = useTranslation();
    const [isLoading, setIsLoading] = useState(false);
    const userAuth = localStorage.getItem('userAuth')
    const [ScreenshotProvided, setScreenshotProvided] = useState(false)
    
    // Perform api request to the backend
    useEffect(() => {
        if (userAuth != null) {
            if (hash != null) {
                console.log({hash})
                setIsLoading(true);
                // Send an api request to the backend with hash data
                fetch(process.env.REACT_APP_BACKEND_URL+'/hash-intelligence?hash=' + hash + "&userAuth=" + userAuth, {
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
            } else if (url != null ){
                setIsLoading(true);
                // Send an api request to the backend with url data
                fetch(process.env.REACT_APP_BACKEND_URL+'/url-intelligence?url=' + url + "&userAuth=" + userAuth, {
                    method: 'GET',
                    headers: {
                        Accept: 'application/json',
                        'Content-Type': 'application/json'
                    }
                }).then((response) => response.json())
                .then((json) => {
                    // If the authentication response is invalid remove the authentication
                    // locally and redirect to the login page.
                    if (json.authenticated !== undefined){
                        localStorage.removeItem('userAuth')
                        window.location.href="/login"
                    } else {
                        setJsonData(json);
                        setIsLoading(false);
                        if (json.screenshot !== null) {
                            setScreenshotProvided(true)
                        }
                    }
                })
                .catch(function(error){
                    console.log(error)
                    setErr(true)
                    setIsLoading(false)
                })
            } else if (file != null){
                setIsLoading(true);
                // Send an api request to the backend with upload file id
                fetch(process.env.REACT_APP_BACKEND_URL+'/upload?file=' + file + "&userAuth=" + userAuth, {
                    method: 'GET',
                    headers: {
                        Accept: 'application/json',
                        'Content-Type': 'application/json'
                    }
                }).then((response) => response.json())
                .then((json) => {
                    // If the authentication response is invalid remove the authentication
                    // locally and redirect to the login page.
                    if (json.authenticated !== undefined){
                        localStorage.removeItem('userAuth')
                        window.location.href="/login"
                    } else {
                        setJsonData(json);
                        setIsLoading(false)
                    }
                })
                .catch(function(error){
                    console.log(error)
                    setErr(true)
                    setIsLoading(false)
                })
            } else {
                // If no valid search is provided redirect to the home page
                window.location.href= "/"
            }
        } else {
            // If a valid parameter is sent, but there is no login redirect to login page
            window.location.href="/login"
        }
    }, [file, hash, url, userAuth]);

    // Add the translation data if the backend request gave data
    if (JsonData !== undefined){
        i18next.addResources('en', 'translation', JsonData.EN);
        i18next.addResources('no', 'translation', JsonData.NO);
    }

    const renderResult = (
        <div className="container text-center break-words sm:justify-center">
        <div className={ScreenshotProvided ? "grid p-2 md:grid-cols-2 xl:grid-cols-3" : "grid p-2"}>
            <div className="xl:col-span-2">
            <h1 className="text-4xl font-bold p-0 mb-8 sm:mt-12 sm:mb-12 w-auto">
                {t("resultTitle")}
            </h1>
            <p className="text-middle text-2xl m-2 pl-2 pr-2 sm:pl-16 sm:pr-16 xl:pl-36 xl:pr-36">
                {t("Result")}
                <br></br><br></br>
                <br></br>
            </p>
            </div>
            <div>
            <Screenshot screenshot = {JsonData.Screenshot} />
            </div>
        </div>
            <div className="container">
                <Sources sourceData = {JsonData} err = {Err}/>
            </div>
        </div>
    );
    
    return (
        <div className="grid place-items-center">
        
            <Navbar />
            <SubNavbar page="resultPage"/>

            <div className="pt-10 pb-10">
            {isLoading ? 
                <>
                <div className="flex justify-center place-items-center">
                <Oval height="100" width="100" color="grey" className="m-auto"/>
                </div>
                <div><p>Loading could take up to a minute.</p></div>
                </> : renderResult}
            </div>
        <div className= "container w-full mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
                <button onClick={() => EscalateAnalysis(url, userAuth, hash)} className="bg-orange-500 p-2 rounded justify-center">{t("manualAnalysisBtn")}</button>
        </div>

        <CookieDisclosure />

        </div>
        );
}

// Function which escalates a request to manual analysis by sending an api request to the backend.
function EscalateAnalysis(url, userAuth, filehash){

    fetch(process.env.REACT_APP_BACKEND_URL+'/escalate?url=' + url +"&result=" + window.location.href + "&userAuth=" + userAuth + "&hash=" + filehash, {
        method: 'GET',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json'
        }
    }).then((response) => response.json())
    .then((json) => {
        // If the authentication response is invalid remove the authentication
        // locally and redirect to the login page.
        if (json.authenticated !== undefined){
            localStorage.removeItem('userAuth')
            window.location.href="/login"
        }
    })
    window.alert("An email has been sent informing you about the escalation to manual analysis, futher contact will be made by email")
}

export default Result;
