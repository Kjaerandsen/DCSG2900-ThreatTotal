import React from "react";
import { useTranslation } from 'react-i18next';

function SubNavbar() {
    const { t } = useTranslation();

    return (
        <div className="container h-auto pl-2 pr-2 sm:pl-18 sm:pr-18 md:pl-36 md:pr-36">
            <div className="flex h-8 sm:h-10 p-1">
			<div className="h-full ml-3 w-full flex items-center">
				<a href="/" className="text-gray-500 hover:underline">
                    {t("Home")} 
				</a>
                <p className="text-gray-500 pl-1 pr-1"> &#8250; </p> <p className="text-gray-500"> Somewhere </p>
			</div>
			<div className="float float-right w-full h-8 sm:h-10 sm:pr-3">
				<div className="float-right h-10 p-1">
					<a href="/login" className="text-gray-500 underline">{t("Login")}</a>
				</div>
			</div>
		</div>
        </div>
    )
}

export default SubNavbar;