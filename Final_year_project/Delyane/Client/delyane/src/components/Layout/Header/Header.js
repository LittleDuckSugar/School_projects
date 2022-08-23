import React, { useState } from 'react';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import './Header.css';

const Header = () => {
    const [links, setLinks] = useState(false)

    const handleClick = () => {
        setLinks(!links)
    }
    const handleLink=()=>{
        if(localStorage.getItem('uuid')){
            window.location = "/profil";
        }
        else{
            window.location = "/authentication";
        }
    }
    return (
        <div className={`header__main ${links ? 'show__nav' : 'hide__nav'}`}>
            <div className='header__logo'>DELYANE</div>

            <ul className='header__nav'>
                <li className='nav__item'>
                    <a href='/' className='nav__link'>Home</a>
                </li>
                <li className='nav__item'>
                    <a href='/painting' className='nav__link'>Painting</a>
                </li>
                <li className='nav__item'>
                    <a onClick={handleLink} className='nav__link'><FontAwesomeIcon className='nav__icon' icon="fa-solid fa-user" /></a>
                </li>
                <li className='nav__item'>
                    <a href='/favorite' className='nav__link'><FontAwesomeIcon className='nav__icon' icon="fa-solid fa-heart" /></a>
                </li>
                <li className='nav__item'>
                    <a href='/cart' className='nav__link'><FontAwesomeIcon className='nav__icon' icon="fa-solid fa-basket-shopping" /></a>
                </li>
            </ul>
            <button className='nav__burger' onClick={handleClick}>
                <span className='burger__bar'></span>
            </button>
        </div>
    );
}

export default Header;