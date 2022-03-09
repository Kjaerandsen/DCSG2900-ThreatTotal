import React, { useEffect} from "react";
import Navbar from "./navbar";
import { useParams } from "react-router-dom";

// look in the url, url decode and write to client

function Result() {
    
    const queryParams = new URLSearchParams(window.location.search);
    const hash = queryParams.get('hash');
    const url = queryParams.get('url');

    useEffect(() => {
        console.log({hash})
        console.log({url})
    });

    return (
        <>
        <div className="grid place-items-center">
        
            <Navbar />


        <div className="bg-red-500 container text-center break-words sm:justify-center">
            <h1 className="text-3xl font-bold p-0 mt-8 mb-8 sm:mt-12 sm:mb-12 w-auto">
                Results
            </h1>
            <p className="w-auto green-400">
                
                display results here
            </p>
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
