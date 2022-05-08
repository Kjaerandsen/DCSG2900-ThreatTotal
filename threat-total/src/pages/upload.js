import React from 'react';
import Navbar from '../components/navbar';
import MainInput from '../components/mainInput';
import ntnuLogo from '../img/ntnuLogoUtenSlagOrd.svg'
import SubNavbar from '../components/subNavbar'
import CookieDisclosure from '../components/cookieDisclosure';
import { useTranslation } from 'react-i18next';

const Upload = () => {
  const userAuth = localStorage.getItem('userAuth')
  const { t } = useTranslation();

  const fileUpload = e => {
    var formData = new FormData(e.target.form);
    var object={};
    formData.forEach((value, key) => object[key] = value);
    if (object["inputFile"] instanceof File) {
        if (userAuth !== null) {
        console.log("loaded file, is in fact a file.")
        console.log(object["inputFile"])

        formData.append('file', object["inputFile"]);
        
        const options = {
          credentials: 'same-origin',
          method: 'POST',
          body: formData,
          // If you add this, upload won't work
          // headers: {
          //   'Content-Type': 'multipart/form-data',
          // }
        };
        
        // forward ID only not object
        fetch(process.env.REACT_APP_BACKEND_URL+'/upload?userAuth=' + userAuth, options)
        .then(response => response.json())
        .then((json) => {
            if (json.authenticated !== undefined){
                localStorage.removeItem('userAuth')
                window.location.href="/login"
            } else {
                window.location.href= "/result?file="+encodeURIComponent(json.id)
            }
        })
        // error handle for non-successful requests
        .catch(function(error){
            console.log(error)        
        })
        } else {
            // If the user is not logged in redirect to the login screen
            window.location.href="/login"
        }    
    }
    else {
        console.log("this is not a file")
        window.alert("Invalid file")
    }
  }


  return (
    <>
	<div className="grid place-items-center">
		
		<Navbar />
        <SubNavbar page="uploadPage"/>
		
        <div className='flex justify-center mt-6 sm:mt-8'>
            <img src={ntnuLogo} className="h-20 sm:h-40 w-auto" alt="NTNU Logo"/>
            <h1 className="text-4xl sm:text-8xl font-bold sm:ml-4 ml-2 pt-2 sm:pt-4 w-auto"> Threat Total </h1>
        </div>
		
		<MainInput />
		
        <div className="container w-full mt-3 mb-3 pl-2 pr-2 sm:pl-36 sm:pr-36">
            <p className="text-left ml-3">
                {t("uploadDescr")}
            </p>
            <form action="/result" method="POST" encType="multipart/form-data">
                <div className= "flex w-full mt-1.5 mb-3 justify-center place-items-center cursor-pointer">
                    <input type="file" className="text-3xl text-slate-500
                    file:mr-4 file:py-2 file:px-3
                    file:rounded-full file:border-0
                    file:text-3xl
                    hover:file:bg-blue-500
                    cursor-pointer"
                    name="inputFile"/>
                </div>

                <div className="mt-1.5 mb-3 flex w-full sm:pl-36 sm:pr-36 justify-center place-items-center overflow-hidden">
                <br></br>
                <br></br>
                <br></br>
                <input type="button" onClick={fileUpload}  value={t("investigate")} className="bg-orange-500 p-2 rounded justify-center cursor-pointer"/>
                </div>
            </form> 
        </div>

            
    <CookieDisclosure />

    </div>
    
	</>
  );
  
};
  
export default Upload;