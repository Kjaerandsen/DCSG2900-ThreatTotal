import React from "react";
import logo from './img/logo.png'

const Navbar = () => {
	return (
	<>
	<nav className="container h-auto ">
		<div className="flex h-14 p-1">
			<div className="h-full ml-3">
				<a href="./">
				<img src={logo} className="h-10 w-auto mt-1" alt="NTNU Logo" />
				</a>
			</div>
			<div className="float float-right w-full h-12 sm:h-full sm:pr-3">
				<div className="float-right place-items-center h-12 mt-3">
					&#127760; <a href="./" className="hover:underline" title="Norsk versjon her: https://www.url.domene/"> Norsk </a>
				</div>
			</div>
		</div>
	</nav>
	<div className="border-b-2 border-gray-400 h-1 w-full"> </div>
	</>
	)
};

export default Navbar;