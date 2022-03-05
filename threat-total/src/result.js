import React from "react";
import Navbar from "./navbar";

// look in the url, url decode and write to client

const Result = () => {
	return (
	<>
    <div class="bg-gray-200 grid place-items-center">
    
        <Navbar />


    <div class = "bg-red-300 container text-center break-words sm:justify-center">
        <h1 class="text-3xl font-bold p-0 mt-8 mb-8 sm:mt-12 sm:mb-12 w-auto">
            Results
        </h1>
        <p class ="w-auto green-400">
            
            display results here
        </p>
    </div>
        
    <div class = "bg-green-300 container w-full mt-1.5 mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
        <a href="./investigate">
            <button class="bg-orange-500 p-2 rounded justify-center">Submit for Manual Analysis</button>
        </a>
    </div>

  </div> 
  </>
  );
};

export default Result;
