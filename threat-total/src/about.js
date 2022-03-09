import React from 'react';
import Navbar from './navbar';
import MainInput from './mainInput';
import ttLogo from './img/TT.png'

//const history = createHashHistory()

class About extends React.Component {

render(){
    return (
    <>
	<div className="grid place-items-center">
		
		<Navbar />
        
        <div>
            <div className="w-full">
                <img src={ttLogo} className="h-auto" alt="Threat Total Logo"/>
            </div>
        </div>
		
        <div className='container p-6 sm:p-12'>
            <h1 className='text-center'> About: </h1>
            <br></br>

            <p>
                Threat total is a threat intelligence service which allows you to get a quick overlook over the safety of using a particular website, domain or application.

                Tincidunt eget nullam non nisi est sit amet facilisis. 
                Eu turpis egestas pretium aenean. Suspendisse potenti nullam ac tortor vitae purus faucibus ornare. 
                Varius vel pharetra vel turpis. Sit amet luctus venenatis lectus magna fringilla urna. 
                Amet consectetur adipiscing elit ut aliquam purus. Nunc pulvinar sapien et ligula. 
                Diam donec adipiscing tristique risus nec feugiat in fermentum posuere. Odio morbi quis commodo odio aenean. 
                Commodo viverra maecenas accumsan lacus vel facilisis volutpat est velit. Eleifend mi in nulla posuere sollicitudin aliquam. 
                Pulvinar pellentesque habitant morbi tristique senectus. In eu mi bibendum neque egestas congue. Aliquet enim tortor at auctor. 
                At quis risus sed vulputate odio ut. Purus ut faucibus pulvinar elementum. Blandit libero volutpat sed cras ornare arcu dui vivamus. 
                Vel risus commodo viverra maecenas accumsan.
            </p>
        </div>

        <div className='container p-6 sm:p-12'>
            <h1> Frequently asked questions: </h1>    

            <ul>
                <li>
                    Q: What is the difference between a domain and a url?
                </li>
                <li>
                    Q: What information sources do you use?
                </li>
                <li>
                    
                </li>
            </ul>
        </div>

    </div >
	</>
    );}
}

export default About
