import React from 'react'
import Source from './source.js'

// reactjs component that iterates through provided json data and displays it via render
// inspired by https://www.youtube.com/watch?v=9C85o8jIgUU
export default function Sources(props) {

var BG = ""
    return (
        <div>
            <h1 className="text-2xl font-bold">Source data:</h1>
            <div className='bg-yellow-500 bg-red-600 bg-green-600'></div>
            <div className='grid grid-cols-1 p-2 md:grid-cols-2 xl:grid-cols-3'>
            {props.sourceData.map((Data, index ) => {
                if (Data.status === "Safe") {
                    BG = "bg-green-600"
                } else if (Data.status === "Risk") {
                    BG = "bg-red-600"
                } else {
                    BG = "bg-yellow-500"
                }
                return (
                <Source Data = {Data} key = {index} BG = {BG}/>
                )
            })}
            </div>
        </div>
    );
}