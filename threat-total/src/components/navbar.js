import React from "react";
import logo from '../img/ntnulogo.svg'
import { useTranslation } from 'react-i18next';
import i18n from '../i18next'

const Navbar = () => {
	const { t } = useTranslation();

	const changeLanguage = () => {
		if (i18n.language === "no") {
			localStorage.setItem('i18nextLng', 'en');
			i18n.changeLanguage("en")
		} else {
			localStorage.setItem('i18nextLng', 'no');
			i18n.changeLanguage("no")
		}
	  };
	
	return (
	<>
	<nav className="container h-auto pl-2 pr-2 sm:pl-18 sm:pr-18 md:pl-36 md:pr-36">
		<div className="flex h-12 sm:h-14 p-1">
			<div className="h-full ml-3 w-full flex items-center">
				<a href="./">
				<img src={logo} className="sm:h-6 h-4 w-auto" alt="NTNU Logo" />
				</a>
			</div>
			<div className="h-full ml-3 w-full flex items-center">
				<a href="/about" className="p-2 text-gray-500 font-semibold m-auto">About</a>
			</div>
			<div className="float float-right w-full h-12 sm:h-full sm:pr-3">
				<div className="float-right place-items-center h-12 mt-2 sm:mt-3">
					&#127760; <button onClick={() => changeLanguage()} className="hover:underline" title={t("languageDescr")}> {t("language")} </button>
				</div>
			</div>
		</div>
	</nav>
	<div className="border-b-2 border-gray-200 h-1 w-full"> </div>
	</>
	)
};

export default Navbar;