import React from 'react';
import Navbar from './navbar';
import MainInput from './mainInput';
import ttLogo from './img/TT.png'
  
const Upload = () => {
  return (
    <>
	<div class="grid place-items-center">
		
		<Navbar />
		
        <div class = "bg-white">
                <img src={ttLogo} class="" alt="NTNU Logo"/>
        </div>
		
		<MainInput />
		
        <div class="container w-full mt-3 mb-3 pl-2 pr-2 sm:pl-36">
            <form class="flex justify-center place-items-center">
                <label class="block m-4">
                    <input type="file" class="block w-full text-3xl text-slate-500
                    file:mr-4 file:py-2 file:px-3
                    file:rounded-full file:border-0
                    file:text-3xl
                    hover:file:bg-blue-500"/>
                </label>
            </form>
        </div>

        
        <div class = "container w-full mt-1.5 mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
            <a href="./investigate">
                <button class="bg-orange-500 p-2 rounded justify-center">Investigate</button>
            </a>
        </div>

    </div>
	</>
  );
};
  
export default Upload;