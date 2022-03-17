import React from 'react'

export default function Source(props) {

    // Missing changing colour according to assessment, and tags as sub components?
return (
    <div className='bg-white border-2 m-2 border-gray-400 rounded-lg p-1 text-left'>
        <div className='flex'>
            <div className='border-r-2 grid place-items-center pr-1'>
                <div className={`rounded-full ${props.BG} w-10 h-10`}></div>
            </div>
            <div className="p-1"> 
                <h1 className='font-bold'>Source: {props.Data.sourceName}</h1>
                <p>Assessment: {props.Data.status}</p>
            </div>
        </div>
        <div className="">
            <p>Tags: {props.Data.tags}</p>
            <p>Shortform: {props.Data.content}</p>
        </div>   
    </div>
);
    
}