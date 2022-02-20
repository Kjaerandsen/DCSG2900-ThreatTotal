import React from 'react';
import Navbar from './navbar';
import MainInput from './mainInput';
import ttLogo from './img/TT.png'
  
const Indextest = () => {
  return (
    <>
	<div class="grid place-items-center">
		
		<Navbar />
		
        <div>
            <a href="/" class="w-full">
                <img src={ttLogo} class="h-auto" alt="Threat Total Logo"/>
            </a>
        </div>
		
		<MainInput />
		
        <div class = "container w-full mt-3 mb-1.5 pl-2 pr-2 sm:pl-36 sm:pr-36 overflow-hidden">
            <form class="p-0 m-2 border-2 border-gray-400 rounded-lg">
            <input class="w-full rounded h-14 p-2 m-0 hover:bg-blue-200" placeholder="https://www.ntnu.no"></input>
            </form>
        </div>

        <div class = "container w-full mt-1.5 mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
            <a href="/investigate">
                <button class="bg-orange-500 p-2 rounded justify-center">Investigate</button>
            </a>
        </div>
    </div>
	</>
  );
};
  
export default Indextest;