import React from 'react'

// reactjs component that iterates through provided json data and displays it via render
// inspired by https://www.youtube.com/watch?v=9C85o8jIgUU
export default function Sources(props) {

    // Missing changing colour according to assessment, and tags as sub components?
return (
    <div>
        <h1 className="text-2xl font-bold">Source data:</h1>
        <div className='grid grid-cols-1 p-2 md:grid-cols-2 xl:grid-cols-3'>
        {props.sourceData.map((Data, index ) => {
        return <div key={index} className='bg-white border-2 m-2 border-gray-400 rounded-lg p-1 text-left'>
            <div className='flex'> 
                <div className='border-r-2 grid place-items-center pr-1'>
                    <div className={`rounded-full bg-red-600 w-10 h-10`}></div>
                </div>
                <div className="p-1"> 
                    <h1 className='font-bold'>Source: {Data.sourceName}</h1>
                    <p>Assessment: {Data.status}</p>
                </div>
            </div>
            <div className="">
                <p>Tags: {Data.tags}</p>
                <p>Shortform: {Data.content}</p>
            </div>
            
        </div>
        })}
        </div>
    </div>
);
    
}