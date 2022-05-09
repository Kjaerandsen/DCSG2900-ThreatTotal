import React from 'react';


export default function Screenshot(screenshot) {

    //i18next.init({ resources: {} });
    //i18next.addResourceBundle('en', 'sourceEn', translationData.en);
    //i18next.addResourceBundle('no', 'sourceNo', translationData.no);
    
    // Add the translation data from the backend
if (screenshot.length !== 0) {

    var blob = new Blob([screenshot], { type: "image/jpeg" });
    var imageUrl = URL.createObjectURL(blob);

var base64Image = 'data:image/png;base64,'+imageUrl;

    
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
}
}