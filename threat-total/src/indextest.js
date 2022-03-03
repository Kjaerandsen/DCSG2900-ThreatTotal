import React, { Component } from 'react';
import Navbar from './navbar';
import MainInput from './mainInput';
import ttLogo from './img/TT.png'
import App from './App';

class Indextest extends React.Component {
//const Indextest = () => {


  search(e){
      e.preventDefault();
    var formData = new FormData(e.target.form);
    var object={};

    formData.forEach((value, key) => object[key] = value);
    const data = JSON.stringify(object)
      //console.log(data)

    fetch("http://localhost:8081/upload",{
        method:'POST',
        headers:{
            'Content-Type':'application/json'
        },
        body:data
        }).then(res => res.json() //Venter på svar
        ).then(data => {    //Sjekker hva data er
            
           window.alert(data)
        });
    }

render(){
    return (
    <>
	<div className="grid place-items-center">
		
		<Navbar />
		
        <div>
            <a href="/" className="w-full">
                <img src={ttLogo} className="h-auto" alt="Threat Total Logo"/>
            </a>
        </div>
		
		<MainInput />
		
        <form className='w-full container'>
        <div className= "container w-full mt-3 mb-1.5 pl-2 pr-2 sm:pl-36 sm:pr-36 overflow-hidden">
            <div className="p-0 m-2 border-2 border-gray-400 rounded-lg">
            <input  className="w-full rounded h-14 p-2 m-0 hover:bg-blue-200" placeholder="https://www.ntnu.no" type="text" name="inputText"></input>
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
  );
    }
//};
}

export default Indextest;

