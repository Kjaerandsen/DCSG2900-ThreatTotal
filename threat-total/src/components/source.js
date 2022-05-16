import React from 'react';
import { useTranslation } from 'react-i18next';
import i18next from '../i18next';

export default function Source(props) {

    //i18next.init({ resources: {} });
    //i18next.addResourceBundle('en', 'sourceEn', translationData.en);
    //i18next.addResourceBundle('no', 'sourceNo', translationData.no);
    
    // Add the translation data from the backend
    if (props.Data !== "") {
    const norsk = props.Data.no
    const english = props.Data.en
    i18next.addResources('en', 'translation', english);
    i18next.addResources('no', 'translation', norsk);
    //i18next.addResource('en', 'translation2', 'status', 'status2', {})
    //console.log ("english: ", english)
    //i18next.addResources({ lng:'no', ns : 'default', any: norsk });
    //console.log("English source x: ", english)
    //console.log( i18next.getDataByLanguage('en'))
    //console.log( i18next.getResource('en', 'translation', 'status'))
    //console.log( i18next.getResource('no', 'translation', 'status'))
    }

    const { t } = useTranslation();
    // If the input is empty return an empty box
    // add a loading animation?
if (props.Data === "") {
    return(
        <div className='bg-white border-2 m-2 border-gray-400 rounded-lg p-1 text-left'>
            <div className='flex'>
                <div className='border-r-2 grid place-items-center pr-1'>
                    <div className={`rounded-full bg-white w-10 h-10`}></div>
                </div>
                <div className="p-1"> 
                    <br></br>
                    <br></br>
                </div>
            </div>
            <div className="">
                <br></br>
                <br></br>
            </div>   
        </div>
    )
} else {
    return (
        <div className='bg-white border-2 m-2 border-gray-400 rounded-lg p-1 text-left'>
            <div className='flex'>
                <div className='border-r-2 grid place-items-center pr-1'>
                    <div className={`rounded-full ${props.BG} w-10 h-10`}></div>
                </div>
                <div className="p-1"> 
                    <h1 className='font-bold'>{t("source")} {props.Data.sourceName}</h1>
                    <p>{t("assessment")} {t("status")}</p>
                </div>
            </div>
            <div className="">
                <p>Tags: {props.Data.en.tags}</p>
                <p>{t("shortForm")} {t("content")}</p>
            </div>   
        </div>
    );
}
    
}