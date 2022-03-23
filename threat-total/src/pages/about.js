import React from 'react';
import Navbar from '../components/navbar';
import ttLogo from '../img/TT.png'

class About extends React.Component {

constructor() {
    super();
}

state = { 
    q1: false,
    q2: false
};

toggleQ1 = () => {
    this.setState({ q1: !this.state.q1 });  
};

toggleQ2 = () => {
    this.setState({ q2: !this.state.q2 });  
};

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
                We retrieve data from the NTNU soc database, as well as external sources such as: ....

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

        <div className='container p-6 sm:p-12 text-center'>
            <h1> Frequently asked questions: </h1>    
            <br></br>
            <ul className='w-full'>
                <li className='w-full pl-4 pr-4 m-0'>
                    <button className='w-full border-2 border-gray-400 rounded-lg p-2' onClick={this.toggleQ1}>Q: What is the difference between a domain and a url?</button>
                </li>
                {this.state.q1 && (<li className='w-full  pl-4 pr-4'>
                    <p className='border border-gray-400 rounded-lg p-2 sm:break-normal'>
                    A url is a specific webpage, for example this page <i>"threat-total.edu/about"</i>. While a domain covers the broader website 
                    <i>"threat-total.edu"</i> and all pages under the domain, such as <i>"threat-total.edu/about"</i>.
                    </p>
                </li>) }
                <li className='w-full pl-4 pr-4 m-0'>
                    <button className='w-full border-2 border-gray-400 rounded-lg p-2' onClick={this.toggleQ2}>Q: Which information sources do you use?</button>
                </li>
                {this.state.q2 && (<li className='w-full  pl-4 pr-4'>
                    <p className='border border-gray-400 rounded-lg p-2 sm:break-normal'>
                    Our main information source is the NTNU soc, but we also retrieve information from ....
                    </p>
                </li>) }
            </ul>
        </div>

    </div >
	</>
    );}
}

export default About
