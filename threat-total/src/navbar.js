import React from "react";
import logo from './img/logo.png'

const Navbar = () => {
	return (
	<>
	<nav className="container h-auto ">
		<div className="grid grid-cols-3 h-auto">
			<div className= "sm:flex sm:place-items-center h-full ml-3"> 
				<a href="./">
				<img src={logo} className="h-12 mt-1" /> 
				</a>
				</div>
				<div className="bg-white flex col-span-2 sm:col-span-1 sm:justify-center sm:place-items-center h-12 sm:h-full">
					Threat Total
				</div>
			<div className="grid col-span-3 sm:col-span-1 sm:flex sm:justify-end place-items-center w-full h-12 sm:h-full sm:pr-3"> 
				<div>  
					&#127760; <a href="./en" className="hover:underline" title="English version"> English </a> 
				</div>
			</div>  
		</div> 
	</nav>
	</>
	)
};

export default Navbar;