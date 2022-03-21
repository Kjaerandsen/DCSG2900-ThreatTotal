import React from 'react';
import Navbar from './Navbar';
import MainInput from './mainInput';
import ttLogo from './img/TT.png'

  

const Upload = () => {

  const fileUpload = e => {
    var formData = new FormData(e.target.form);
    var object={};
    formData.forEach((value, key) => object[key] = value);
    if (object["inputFile"] instanceof File) {
        console.log("loaded file, is in fact a file.")
        console.log(object["inputFile"])
        // TODO, can find file, though dont know where to forward it
    }
    else {
        console.log("this is not a file")
    }
  }
  return (
    <>
	<div className="grid place-items-center">
		
		<Navbar />
		
        <div>
                <img src={ttLogo} className="" alt="NTNU Logo"/>
        </div>
		
		<MainInput />
		
        <div className="container w-full mt-3 mb-3 pl-2 pr-2 sm:pl-36">
            <form action="/result" method="POST" encType="multipart/form-data" className="flex justify-center place-items-center">
                <label className="block m-4">
                    
                    <input type="file" onChange={fileUpload} className="block w-full text-3xl text-slate-500
                    file:mr-4 file:py-2 file:px-3
                    file:rounded-full file:border-0
                    file:text-3xl
                    hover:file:bg-blue-500"
                    name = "inputFile"/>
                    <input type ="submit" value="Submit"/>
                </label>
                
            </form>
        </div>

        
        <div className= "container w-full mt-1.5 mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
            <a href="./result">
                <button onClick={fileUpload} className="bg-orange-500 p-2 rounded justify-center">Investigate</button>
                
            </a>
        </div>

    </div>
    
	</>
  );
  
};
  
export default Upload;