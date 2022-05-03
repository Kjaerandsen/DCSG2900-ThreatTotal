import React from 'react'
import Source from './source.js'
import { useTranslation } from 'react-i18next';

// reactjs component that iterates through provided json data and displays it via render
// inspired by https://www.youtube.com/watch?v=9C85o8jIgUU
export default function Sources(props) {

    const { t } = useTranslation();
    var BG = ""
    // Checks if an error has occured while connecting to the backend, if true return an error message
    if (props.err || props.sourceData == undefined || props.sourceData == "") {
        return (
            <div className='bg-white border-2 border-gray-400 rounded-lg p-2 m-4'>
                <h1>{t("backendError")}</h1>
            </div>
        );
    // Else return the source data
    } else {
        console.log ("sourceData: " + props.sourceData)
        return (
            <div className="bg-gray-200">
                <h1 className="text-2xl font-bold">{t("sourceTitle")}</h1>
                <div className='bg-yellow-500 bg-red-600 bg-green-600'></div>
                <div className='grid grid-cols-1 p-2 md:grid-cols-2 xl:grid-cols-3'>
                {props.sourceData.ResponseData.map((Data, index ) => {
                    if (Data.en.status === "Safe") {
                        BG = "bg-green-600"
                    } else if (Data.en.status === "Risk") {
                        BG = "bg-red-600"
                    } else {
                        BG = "bg-yellow-500"
                    }
                    // Ha en "Assessment guide" som forklarer fargene på siden / på "about" siden?
                    return (
                    <Source Data = {Data} key = {index} BG = {BG}/>
                    )
                })}
                </div>
            </div>
        );
    }
}