import React from "react";
import Data from "../src/testData/data.json"

// reactjs component that iterates through json file and displays it via render
// inspired from https://www.youtube.com/watch?v=9C85o8jIgUU
const Display = () => {
    
    return (
        <div>
        <h1>hello world</h1>
        {Data.map((Data, index ) => {
            return <div>
                <h1>{Data.title}</h1>
                <p>{Data.content}</p>
            </div>
        })}
        </div>
    );

   
}

export default Display;
