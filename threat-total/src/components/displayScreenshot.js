import React from 'react';


export default function Screenshot(screenshot) {
    
if (screenshot.screenshot !== undefined && screenshot.screenshot !== null) {

var base64Image = 'data:image/png;base64,'+screenshot.screenshot;
    
    return(
        <div>
            <div className='flex'>
                <div className='grid place-items-center'>
                <img src={base64Image} className="w-full h-auto border-2 border-gray-400 rounded-lg " alt="Screenshot" />
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
    
}else{
    return(
        <></> 
    )
}
}