import React from "react";
import logo from './img/logo.png'

const Navbar = () => {
	return (
	<>
	<nav class="container h-auto ">
		<div class="grid grid-cols-3 h-auto p-1">
			<div class= "sm:flex sm:place-items-center h-full ml-3">
				<a href="./">
				<img src={logo} class="h-10 w-auto mt-1" alt="NTNU Logo" />
				</a>
			</div>
			<div class="bg-white flex col-span-2 sm:col-span-1 sm:justify-center sm:place-items-center h-12 sm:h-full">
				Threat Total
			</div>
			<div class="grid col-span-3 sm:col-span-1 sm:flex sm:justify-end place-items-center w-full h-12 sm:h-full sm:pr-3">
				<div>
					&#127760; <a href="./" class="hover:underline" title="Norsk versjon her: https://www.url.domene/"> Norsk </a>
				</div>
			</div>
		</div>
	</nav>
	<div class="border-b-2 border-gray-400 h-1 w-full"> </div>
	</>
	)
};

export default Navbar;