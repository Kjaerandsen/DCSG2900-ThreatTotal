import React from "react";

const MainInput = props => {
	return (
	<>
	<div className= "container w-full h-auto p-2">
		
		<div className="bg-gray-200 grid grid-cols-1 sm:grid-cols-2 sm:mr-36 sm:ml-36 mt-12 mb-6 border-2 border-gray-400 rounded-lg">
			<div className="flex justify-center place-items-center w-full h-full overflow-hidden"> 
				<a href="/" className={`w-full h-full
				${props.IsSearch ? "bg-blue-300" : "bg-white"}
				hover:bg-blue-500 rounded`}>
				<button className=" w-full p-2 ">Search</button></a>
			</div>
			<div className="flex justify-center place-items-center w-full h-full overflow-hidden"> 
				<a href="/upload" className="w-full h-full"> 
					<button className={`w-full p-2 rounded 
					${props.IsSearch ? "bg-white" : "bg-blue-300"}
			  	  hover:bg-blue-500`}>
							Upload File
					</button>
				</a>
			</div>
		</div>
	</div>
	</>
	)
};
/*
figure out proper button highlightning when selected search or upload
comment
*/
export default MainInput;