import React from 'react';
import Navbar from './navbar';
import MainInput from './mainInput';
import ttLogo from './img/TT.png'
  
const Upload = () => {
  return (
    <>
	<div className="grid place-items-center">
		
		<Navbar />
		
        <div>
                <img src={ttLogo} className="" alt="NTNU Logo"/>
        </div>
		
		<MainInput />
		
        <div className="container w-full mt-3 mb-3 pl-2 pr-2 sm:pl-36">
            <form className="flex justify-center place-items-center">
                <label className="block m-4">
                    <input type="file" className="block w-full text-3xl text-slate-500
                    file:mr-4 file:py-2 file:px-3
                    file:rounded-full file:border-0
                    file:text-3xl
                    hover:file:bg-blue-500"/>
                </label>
            </form>
        </div>

        
        <div className= "container w-full mt-1.5 mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
            <a href="./investigate">
                <button className="bg-orange-500 p-2 rounded justify-center">Investigate</button>
            </a>
        </div>

    </div>
	</>
  );
};
  
export default Upload;