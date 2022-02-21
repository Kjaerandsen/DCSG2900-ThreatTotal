import React, { Component } from 'react';
import Navbar from './navbar';
import MainInput from './mainInput';
import App from './App';

class Indextest extends React.Component {
//const Indextest = () => {


  search(e){
      e.preventDefault();
    var formData = new FormData(e.target.form);
    var object={};

    formData.forEach((value, key) => object[key] = value);
    const data = JSON.stringify(object)
      console.log(data[2])
  }

render(){
    return (
    <>
	<div className="bg-gray-200 grid place-items-center">
		
		<Navbar />
		
        <div class = "bg-red-300 container text-center break-words sm:justify-center">
            <h1 className="text-3xl font-bold p-0 mt-8 mb-8 sm:mt-12 sm:mb-12 w-auto">
                Threat Total
            </h1>
        </div>
		
		<MainInput />
		
        <form className="bg-gray-200 w-full grid place-items-center">

        <div class = "bg-green-300 container w-full mt-3 mb-1.5 pl-2 pr-2 sm:pl-36 sm:pr-36 overflow-hidden">
            <label for="inputText">Input url, filehash or domain:</label>
                <input className="w-full rounded h-8 box-border p-2 m-0"
                       placeholder="https://www.ntnu.no"
                       type="text" name="inputText" id="inputText"/>
        </div>

        <div class = "bg-green-300 container w-full mt-1.5 mb-3 sm:pl-36 sm:pr-36 flex justify-center overflow-hidden">
            <a href="./investigate">
                <button onClick={this.search} className="bg-orange-500 p-2 rounded justify-center" type="submit">Investigate</button>
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

