import React, { useState } from 'react';
import Navbar from '../components/navbar';
import SubNavbar from '../components/subNavbar'
import ntnuLogo from '../img/ntnuLogoUtenSlagOrd.svg';
import CookieDisclosure from '../components/cookieDisclosure';
import { useTranslation } from 'react-i18next';

function About() {

// The consts are used for the open / close functionality of the menus
const { t } = useTranslation();

const [q1, setQ1] = useState(false);
const [q2, setQ2] = useState(false);
const [q3, setQ3] = useState(false);
const [q4, setQ4] = useState(false);

function toggleQ1 () {
    setQ1(!q1);
};
function toggleQ2 () {
    setQ2(!q2);
};
function toggleQ3 () {
    setQ3(!q3);
};
function toggleQ4 () {
    setQ4(!q4);
};

    return (
    <>
	<div className="grid place-items-center">
		
		<Navbar />
        <SubNavbar page="aboutPage"/>
        
        <div className='flex justify-center mt-6 sm:mt-8'>
            <img src={ntnuLogo} className="h-20 sm:h-35 md:h-40 w-auto" alt="NTNU Logo"/>
            <h1 className="text-4xl sm:text-6xl md:text-8xl font-bold sm:ml-4 ml-2 pt-2 sm:pt-4 w-auto"> Threat Total </h1>
        </div>
		
        <div className='container pt-6 pb-4 sm:pt-12 sm:pb-6 pl-2 pr-2 sm:pl-16 sm:pr-16 xl:pl-36 xl:pr-36'>
            <h1 className='text-center'> {t("about:about")} </h1>
            <br></br>

            <p>
                {t("about:text1")}
                <br></br>
                {t("about:text2")}
                <a href="https://git.gvk.idi.ntnu.no/Johannesb/dcsg2900-threattotal" className='text-blue-500'> 
                    {t("about:textUrl")} 
                </a>.
                <br></br>
                <br id='cookie'></br>
                {t("about:text3")}
                <br></br>
                <br id="login"></br>
                {t("about:text4")}
            </p>          
        </div>

        <div className='container p-6 sm:p-12 text-center'>
            <h1> {t("about:faq")} </h1>    
            <br></br>
            <ul className='w-full'>
                <li className='w-full pl-4 pr-4 m-0'>
                    <button className='w-full border-2 border-gray-400 rounded-lg p-2' onClick={toggleQ1}>{t("about:q1")}</button>
                </li>
                {q1 ? (<li className='w-full  pl-4 pr-4'>
                    <p className='border border-gray-400 rounded-lg p-2 sm:break-normal'>
                    {t("about:q1text1")}<i>"{t("about:q1text2")}"</i>. {t("about:q1text3")} 
                    <i>"{t("about:q1text4")}"</i> {t("about:q1text5")} <i>"{t("about:q1text6")}"</i>.
                    </p>
                </li>) : <div></div> }
                <li className='w-full pl-4 pr-4 m-0'>
                    <button className='w-full border-2 border-gray-400 rounded-lg p-2' onClick={toggleQ2}>{t("about:q2")}</button>
                </li>
                {q2 && (<li className='w-full  pl-4 pr-4'>
                    <p className='border border-gray-400 rounded-lg p-2 sm:break-normal'>
                    {t("about:q2text")}
                    </p>
                </li>) }
                <li className='w-full pl-4 pr-4 m-0'>
                    <button className='w-full border-2 border-gray-400 rounded-lg p-2' onClick={toggleQ3}>{t("about:q3")}</button>
                </li>
                {q3 && (<li className='w-full  pl-4 pr-4'>
                    <p className='border border-gray-400 rounded-lg p-2 sm:break-normal'>
                    {t("about:q3text")}
                    </p>
                </li>) }
                <li className='w-full pl-4 pr-4 m-0'>
                    <button className='w-full border-2 border-gray-400 rounded-lg p-2' onClick={toggleQ4}>{t("about:q4")}</button>
                </li>
                {q4 && (<li className='w-full pl-4 pr-4'>
                    <div className="border border-gray-400 rounded-lg">
                    <p className='p-2 sm:break-normal'>
                    {t("about:q4text")}
                    </p>
                    <div className='flex flex-col sm:pl-16 sm:pr-16 xl:pl-36 xl:pr-36 mb-2'>
                        <table className="table-auto border-2">
                            <thead>
                                <tr className='bg-gray-100'>
                                <th className='border-r-2 p-1'>
                                    {t("about:q4table1")}
                                </th>
                                <th>
                                    {t("about:q4table2")}
                                </th>
                                </tr>
                            </thead>
                            <tbody>
                            <tr className=''>
                                <td className='border-r-2 p-1'>{t("about:q4table3")}</td>
                                <td>{t("about:q4table4")}</td>
                            </tr>
                            <tr className='bg-gray-100'>
                                <td className='border-r-2 p-1'>{t("about:q4table5")}</td>
                                <td>{t("about:q4table6")}</td>
                            </tr>
                            <tr>
                                <td className='border-r-2 p-1'>{t("about:q4table7")}</td>
                                <td>{t("about:q4table8")}</td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                    </div>
                </li>) }
            </ul>
        </div>

        <CookieDisclosure />

    </div >
	</>
    );}

export default About
