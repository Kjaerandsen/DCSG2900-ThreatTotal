import React, { Component } from 'react';
import Navbar from './Navbar';
import MainInput from './MainInput';
import ttLogo from './img/TT.png'
import App from './App';
import { createHashHistory } from 'history'

//const history = createHashHistory()

class Indextest extends React.Component {

  search(e){
      e.preventDefault();
    var formData = new FormData(e.target.form);
    var object={};
    const re = /[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()!@:%_\+.~#?&\/\/=]*)/;

    formData.forEach((value, key) => object[key] = value);
    // Also check if empty and return an error, or invalid on both hash and url
    if (object["inputText"].match(re) ) {
        // if user input matches regex, send to url
        window.location.href= "/result?url="+encodeURIComponent(object["inputText"])
    } else {
        // if user input does not match regex, send to backend and do checks there
        window.location.href= "/result?hash="+encodeURIComponent(object["inputText"])
    }
  }

render(){
    return (
    <>
	<div className="grid place-items-center">
		
		<Navbar />
		
        <div>
                <img src={ttLogo} className="" alt="NTNU Logo"/>
        </div>
		
		<MainInput />

        <form className='w-full container'>
        <div className= "container w-full mt-3 mb-1.5 pl-2 pr-2 sm:pl-36 sm:pr-36 overflow-hidden">
            <p className="text-left ml-3">
                Input a <u title="A domain, for example: ntnu.edu">domain</u>,&nbsp;
                <u title="A webpage, for example: https://www.ntnu.edu/contact">url</u>
                &nbsp;or a&nbsp;
                <u title="sha256 hash of a file">file hash</u>:
            </p>
            <div className="p-0 m-2 border-2 border-gray-400 rounded-lg">
            <input className="w-full rounded h-14 p-2 m-0 hover:bg-blue-200" placeholder="https://www.ntnu.no"
                    type="text" name="inputText">
            </input>
            </div>
        </div>

        <div className= "container w-full mt-1.5 mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
            <a href="/investigate">
                <button onClick={this.search} className="bg-orange-500 p-2 rounded justify-center">Investigate</button>
            </a>
        </div>
        </form>
    </div>
	</>
    );}
}

export default Indextest;

