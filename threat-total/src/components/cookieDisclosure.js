import React, { useEffect } from "react";
import { useTranslation } from 'react-i18next';





function CookieDisclosure () {
    const { t } = useTranslation();

    // Hide the cookie disclosure if the cookiesEnabled cookie is present
    useEffect(() => {
        if (localStorage.getItem("cookiesEnabled") !== "true") {
            console.log("True")
            document.getElementById("cookieDisclosure").style.display = "block"
        }
    }, []);
    
    function hideCookieDisclosure() {
        localStorage.setItem('cookiesEnabled', 'true')
        document.getElementById("cookieDisclosure").style.display = "none"
    }

    return (
        <>
        <footer id="cookieDisclosure" className="w-full h-auto place-items-center fixed bottom-0 hidden">
        <div className="border-b-2 border-gray-400 h-1 w-full"> </div>
		<div className="flex h-auto sm:h-10 p-1 bg-blue-200">
			<div className="h-full ml-3 w-full flex items-center">
                <p> {t("cookieDisclosure")}
                <a href="/about#cookie" className="underline">{t("cookieDisclosureLink")}</a>
                </p>
            </div>

			<div className="flex float-right h-full pr-2 sm:pr-3">
				<div className="flex items-center float-right">
                <button onClick={() => hideCookieDisclosure()}><p className="text-2xl">&#10006;</p></button>
				</div>
			</div>
		</div>
	</footer>
        </>
    )
}

export default CookieDisclosure;