import React, { useEffect } from 'react';
import Navbar from '../components/navbar';
import MainInput from '../components/mainInput';
import ntnuLogo from '../img/ntnuLogoUtenSlagOrd.svg';
import CookieDisclosure from '../components/cookieDisclosure';
import { useTranslation } from 'react-i18next';

function search(e){
    e.preventDefault();
    var formData = new FormData(e.target.form);
    var object={};
    // eslint-disable-next-line
    const regex = /[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()!@:%_\+.~#?&\/\/=]*)/;

    formData.forEach((value, key) => object[key] = value);
    // Also check if empty and return an error, or invalid on both hash and url
    if (object["inputText"].match(regex) ) {
        // if user input matches regex, send to url
        window.location.href= "/result?url="+encodeURIComponent(object["inputText"])
    } else {
        // if user input does not match regex, send to backend and do checks there
        window.location.href= "/result?hash="+encodeURIComponent(object["inputText"])
    }
}

// Possibly cleaner to use an svg image for the headline text
// consider renaming the file? our main function file probably shouldn't contain "test"

function Indextest() {
    const { t } = useTranslation();
    const queryParams = new URLSearchParams(window.location.search);
    const code = queryParams.get('code');

    // Put this on a seperate page with a redirect on completion of the request?
    useEffect(() => {
        if (code != null) {
            fetch('http://localhost:8081/login?code=' + code, {
                method: 'GET',
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json'
                }
            }).then((response) => response.json())
            .then((json) => {
                console.log(json, json.hash)
                localStorage.setItem('userAuth', json.hash)
                window.location.href="/"
            })
            .catch(function(error){
                console.log(error)
                window.location.href="/"
            })
        }
    })

    return (
        <div className="grid place-items-center">
            
            <Navbar />

            <div className='flex justify-center mt-6 sm:mt-8'>
                <img src={ntnuLogo} className="h-20 sm:h-35 md:h-40 w-auto" alt="NTNU Logo"/>
                <h1 className="text-4xl sm:text-6xl md:text-8xl font-bold sm:ml-4 ml-2 pt-2 sm:pt-4 w-auto"> Threat Total </h1>
            </div>
            
            <MainInput IsSearch/>

            <form className='w-full container'>
            <div className= "container w-full mt-3 mb-1.5 pl-2 pr-2 sm:pl-36 sm:pr-36 overflow-hidden">
                <p className="text-left ml-3">
                    {t("inputDescr")}
                </p>
                <div className="p-0 m-2 border-2 border-gray-400 rounded-lg">
                <input className="w-full rounded h-14 p-2 m-0 hover:bg-blue-200" placeholder={t("inputUrl")}
                        type="text" name="inputText">
                </input>
                </div>
            </div>
            <div className= "container w-full mt-1.5 mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
                <a href="/result">
                    <button onClick={search} className="bg-orange-500 p-2 rounded justify-center">{t("investigate")}</button>
                </a>
            </div>
            </form>

            <CookieDisclosure />

        </div>
    )
}

export default Indextest;

