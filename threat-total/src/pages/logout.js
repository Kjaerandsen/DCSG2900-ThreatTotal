import React from 'react';
import Navbar from '../components/navbar';
import CookieDisclosure from '../components/cookieDisclosure';
import { useTranslation } from 'react-i18next';
import { t } from 'i18next';

const Logout = () => {

  return (
    <>
	<div className="grid place-items-center">
		
		<Navbar />
		
        <div className='flex justify-center mt-6 sm:mt-8'>
            <h1 className="text-4xl sm:text-8xl font-bold sm:ml-4 ml-2 pt-2 sm:pt-4 w-auto"> Threat Total </h1>
        </div>
		
        <div className="container w-full mt-3 mb-3 pl-2 pr-2 sm:pl-36">
            <h1>{t("logoutMessage")}</h1>
            <br></br>
        </div>

        <div>
            <button className='bg-gray-400 border-2 rounded-lg p-2'>{t("loginButton")}</button>
        </div>
    <CookieDisclosure />

    </div>
    
	</>
  );
  
};
  
export default Logout;