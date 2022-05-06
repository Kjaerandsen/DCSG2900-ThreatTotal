import React from 'react';
import Navbar from '../components/navbar';
import MainInput from '../components/mainInput';
import ntnuLogo from '../img/ntnuLogoUtenSlagOrd.svg'
import SubNavbar from '../components/subNavbar'
import CookieDisclosure from '../components/cookieDisclosure';

const Upload = () => {
  const fileUpload = e => {
    var formData = new FormData(e.target.form);
    var object={};
    formData.forEach((value, key) => object[key] = value);
    if (object["inputFile"] instanceof File) {
        console.log("loaded file, is in fact a file.")
        console.log(object["inputFile"])

        formData.append('file', object["inputFile"]);
        
        const options = {
          method: 'POST',
          body: formData,
          // If you add this, upload won't work
          // headers: {
          //   'Content-Type': 'multipart/form-data',
          // }
        };
        
        fetch('http://localhost:8081/upload', options);  

        /*fetch('http://localhost:3000/result', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            }
        }).then((response) => response.json()) */
              
    }
    else {
        console.log("this is not a file")
    }
  }
  return (
    <>
	<div className="grid place-items-center">
		
		<Navbar />
        <SubNavbar />
		
        <div className='flex justify-center mt-6 sm:mt-8'>
            <img src={ntnuLogo} className="h-20 sm:h-40 w-auto" alt="NTNU Logo"/>
            <h1 className="text-4xl sm:text-8xl font-bold sm:ml-4 ml-2 pt-2 sm:pt-4 w-auto"> Threat Total </h1>
        </div>
		
		<MainInput />
		
        <div className="container w-full mt-3 mb-3 pl-2 pr-2 sm:pl-36">
            <form action="/result" method="POST" encType="multipart/form-data" className="flex justify-center place-items-center">
                <label className="block m-4">
                    
                    <input type="file" className="block w-full text-3xl text-slate-500
                    file:mr-4 file:py-2 file:px-3
                    file:rounded-full file:border-0
                    file:text-3xl
                    hover:file:bg-blue-500"
                    name = "inputFile"/>

                </label>
                <div className= "container w-full mt-1.5 mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">

                <input type="button" onClick={fileUpload}  value="Investigate" className="bg-orange-500 p-2 rounded justify-center"/>
                </div>
            </form> 
        </div>

            
    <CookieDisclosure />

    </div>
    
	</>
  );
  
};
  
export default Upload;