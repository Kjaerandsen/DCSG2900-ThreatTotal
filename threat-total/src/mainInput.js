import React from "react";

const MainInput = () => {
	return (
	<>
	<div class = "bg-yellow-300 container w-full h-auto">
		<div className="bg-gray-200 grid grid-cols-1 sm:grid-cols-2 sm:mr-36 sm:ml-36 mt-12 mb-6">
		<div className="flex justify-center place-items-center w-full h-full"> 
			<a href="/" className="w-full h-full
			{{ if .isSelected }} bg-blue-400 {{ else }} bg-white {{ end }}
			hover:bg-blue-500 rounded">
				<button className=" w-full p-2 ">Search</button></a>
		</div>
		<div className="flex justify-center place-items-center w-full h-full "> 
			<a href="/upload" className="w-full h-full">
			<button className="w-full p-2 rounded hover:bg-blue-500
					{{ if .isSelected }} bg-white {{ else }} bg-blue-400 {{ end }}">
					Upload File
				</button></a>
		</div>
		</div>
	</div>
	</>
	)
};

export default MainInput;