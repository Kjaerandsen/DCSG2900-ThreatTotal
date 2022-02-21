import React from "react";

const MainInput = () => {
	return (
	<>
	<div class = "container w-full h-auto p-2">
		<div class="bg-gray-200 grid grid-cols-1 sm:grid-cols-2 sm:mr-36 sm:ml-36 mt-12 mb-6 border-2 border-gray-400 rounded-lg">
			<div class="flex justify-center place-items-center w-full h-full overflow-hidden"> 
				<a href="/" class="w-full h-full
				bg-white
				hover:bg-blue-500 rounded">
				<button class=" w-full p-2 ">Search</button></a>
			</div>
			<div class="flex justify-center place-items-center w-full h-full overflow-hidden"> 
				<a href="/upload" class="w-full h-full">
					<button class="w-full p-2 rounded hover:bg-blue-500
							bg-blue-400">
							Upload File
					</button>
				</a>
			</div>
		</div>
	</div>
	</>
	)
};

export default MainInput;