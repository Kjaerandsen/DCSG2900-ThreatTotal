import React from 'react';
import Navbar from '../components/navbar';
import CookieDisclosure from '../components/cookieDisclosure';
  
const Investigate = () => {
  return (
    <>
	<div className="grid place-items-center">
		
		<Navbar />
        
        <div className= "container w-full h-auto">
        <h1 className="text-center text-3xl font-bold pt-12">Investigation:</h1>
            <div className="bg-gray-200 grid grid-cols-1 sm:grid-cols-1 sm:mr-36 sm:ml-36 mt-12 mb-6">
            <div className="flex justify-center place-items-center w-full h-full">
                <div className="rounded border-2 border-gray-900 w-full bg-white text-left"> 
                    <p className="p-1">https://www.ntnu.edu</p>
                </div>
            </div>
            </div>
        </div>

        <div className= "container w-full mt-3 mb-1.5 pl-2 pr-2 sm:pl-36 sm:pr-36 overflow-hidden">
            
            <p className="text-lime-500">
                <br/>No results found for: https://www.ntnu.edu<br/><br/>
                Reputation score: 0<br/><br/>
            </p>
            <p>
                If you find this URL suspicious and we do not have any reputation about this domain you can escalate 
                this URL for manual analysis by clicking the button below.<br/><br/>
            </p>
        </div>
    
        <div className= "bg-white container w-full mt-1.5 mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
            <a href="/investigate" className="text-white font-semibold">
                <button className="bg-red-600 hover:bg-red-400 p-2 rounded border-black border-2 justify-center">
                    Send to manual analysis
                </button>
            </a>
        </div>

        <CookieDisclosure />

    </div>
	</>
  );
};
  
export default Investigate;