import React from 'react';


export default function Screenshot(screenshot) {
    
if (screenshot.screenshot !== undefined && screenshot !== null) {

var base64Image = 'data:image/png;base64,'+screenshot.screenshot;
    
    return(
        <div className='bg-white border-2 m-2 border-gray-400 rounded-lg p-1 text-left'>
            <div className='flex'>
                <div className='border-r-2 grid place-items-center pr-1'>
                <img src={base64Image} alt="Screenshot" />
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