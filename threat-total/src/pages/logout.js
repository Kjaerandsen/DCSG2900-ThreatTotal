import React from 'react';
import Navbar from '../components/navbar';
import CookieDisclosure from '../components/cookieDisclosure';
import SubNavbar from '../components/subNavbar'
import { useTranslation } from 'react-i18next';
import ntnuLogo from '../img/ntnuLogoUtenSlagOrd.svg';
import { t } from 'i18next';

const Logout = () => {

  return (
    <>
	<div className="grid place-items-center">
		
		<Navbar />
		<SubNavbar page="logoutPage"/>
        
        <div className='flex justify-center mt-6 sm:mt-8'>
            <img src={ntnuLogo} className="h-20 sm:h-35 md:h-40 w-auto" alt="NTNU Logo"/>
            <h1 className="text-4xl sm:text-6xl md:text-8xl font-bold sm:ml-4 ml-2 pt-2 sm:pt-4 w-auto"> Threat Total </h1>
        </div>

        <div className='container pt-6 pb-6 sm:pt-12 sm:pb-8 pl-2 pr-2 sm:pl-16 sm:pr-16 xl:pl-36 xl:pr-36'>
            <h1 className='text-center text-2xl'>{t("logoutMessage")}</h1>
            <br></br>
        </div>

        <div>
            <a href="https://auth.dataporten.no/oauth/authorization?
                    client_id=b798a173-514b-4f52-ad14-2838dacde3f1&
                    response_type=code&
                    redirect_uri=http://localhost:3000&
                    scope=openid&
                    state=whatever&
                    secret=198cb677-d113-446c-807e-8d9ad81ab8e0">
            <button className='bg-blue-400 border-2 rounded-lg p-2'>{t("loginButton")}</button>
            </a>
        </div>

    <CookieDisclosure />

    </div>
    
	</>
  );
  
};
  
export default Logout;