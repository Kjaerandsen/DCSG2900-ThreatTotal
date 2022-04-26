import React from 'react';
import Navbar from '../components/navbar';
import ntnuLogo from '../img/ntnuLogoUtenSlagOrd.svg';
import CookieDisclosure from '../components/cookieDisclosure';

class About extends React.Component {

constructor() {
    super();
}

state = { 
    q1: false,
    q2: false,
    q3: false,
    q4: false
};

toggleQ1 = () => {
    this.setState({ q1: !this.state.q1 });  
};

toggleQ2 = () => {
    this.setState({ q2: !this.state.q2 });  
};

toggleQ3 = () => {
    this.setState({ q3: !this.state.q3 });  
};

toggleQ4 = () => {
    this.setState({ q4: !this.state.q4 });  
};

render(){
    return (
    <>
	<div className="grid place-items-center">
		
		<Navbar />
        
        <div className='flex justify-center mt-6 sm:mt-8'>
            <img src={ntnuLogo} className="h-20 sm:h-35 md:h-40 w-auto" alt="NTNU Logo"/>
            <h1 className="text-4xl sm:text-6xl md:text-8xl font-bold sm:ml-4 ml-2 pt-2 sm:pt-4 w-auto"> Threat Total </h1>
        </div>
		
        <div className='container pt-6 pb-6 sm:pt-12 sm:pb-8 pl-2 pr-2 sm:pl-16 sm:pr-16 xl:pl-36 xl:pr-36'>
            <h1 className='text-center'> About: </h1>
            <br></br>

            <p>
                Threat total is a threat intelligence service which allows you to get a quick overlook over the safety of 
                using a particular website, domain or application.
                We retrieve data from the NTNU soc, as well as external sources described below in the questions and answers. 
                <br></br>
                The threat total application is open source software written in Golang and Reactjs. The source code freely is available at 
                <a href="https://git.gvk.idi.ntnu.no/Johannesb/dcsg2900-threattotal" className='text-blue-500'> the NTNU in Gjøvik gitlab instance</a>.
                <br></br>
                <br id='cookie'></br>We use cookies. Which is text stored in the browser, we use this to store authentication data and language choices.
                Specifically cookies are used to save your login information to keep you logged in, if the cookie prompt has been closed, as well as the language selected.
                For more information on the cookies used view the questions and answers below.
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
                    <button className='w-full border-2 border-gray-400 rounded-lg p-2' onClick={this.toggleQ2}>Q: What is a filehash?</button>
                </li>
                {this.state.q2 && (<li className='w-full  pl-4 pr-4'>
                    <p className='border border-gray-400 rounded-lg p-2 sm:break-normal'>
                    A filehash is a hash of a file. A hash is a function which turns an input into a unique output of a defined length. This makes
                    it possible to uniquely identify files for searches without uploading the whole file. Which saves time and resources if the
                    file is already known.
                    </p>
                </li>) }
                <li className='w-full pl-4 pr-4 m-0'>
                    <button className='w-full border-2 border-gray-400 rounded-lg p-2' onClick={this.toggleQ3}>Q: Which information sources do you use?</button>
                </li>
                {this.state.q3 && (<li className='w-full  pl-4 pr-4'>
                    <p className='border border-gray-400 rounded-lg p-2 sm:break-normal'>
                    Our main information source is the NTNU soc, but we also retrieve information from google safebrowsing,
                    Alienvault open threat exchange and Hybrid Analysis Falcon Public API.
                    </p>
                </li>) }
                <li className='w-full pl-4 pr-4 m-0'>
                    <button className='w-full border-2 border-gray-400 rounded-lg p-2' onClick={this.toggleQ4}>Q: What do you use cookies for?</button>
                </li>
                {this.state.q4 && (<li className='w-full pl-4 pr-4'>
                    <div className="border border-gray-400 rounded-lg">
                    <p className='p-2 sm:break-normal'>
                    Cookies are used to save your login information to keep you logged in, if the cookie prompt has been closed, as well as the language selected.
                    The cookies we use are listed in the table below:
                    </p>
                    <div className='flex flex-col sm:pl-16 sm:pr-16 xl:pl-36 xl:pr-36 mb-2'>
                        <table className="table-auto border-2">
                            <thead>
                                <tr className='bg-gray-100'>
                                <th className='border-r-2 p-1'>
                                    Name
                                </th>
                                <th>
                                    Description
                                </th>
                                </tr>
                            </thead>
                            <tbody>
                            <tr className=''>
                                <td className='border-r-2 p-1'>i18nextLng </td>
                                <td>Language selection, only valid for norwegian and english languages</td>
                            </tr>
                            <tr className='bg-gray-100'>
                                <td className='border-r-2 p-1'>cookiesEnabled </td>
                                <td>Cookie for the cookie prompt, saves the prompt as closed.</td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                    </div>
                </li>) }
            </ul>
        </div>

        <CookieDisclosure />

    </div >
	</>
    );}
}

export default About
