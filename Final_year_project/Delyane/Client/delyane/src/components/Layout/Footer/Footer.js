import React, { useState } from 'react';
import axios from 'axios';

import './Footer.css';
import { Link } from 'react-router-dom';

const Footer = () => {
    const [newsletter, setNewsletter] = useState({});

    const submitNewsletter = async (e) => {
        console.log(newsletter);
        e.preventDefault();
        const url = 'http://localhost:8080/newsletter';
        try {
            await axios.post(url, newsletter);
        } catch (err) {
            console.log(err)
        }
    };

    return (
        <div className='footer__main'>
            <ul className='footer__list'>
                <li className='footer__items'>
                    <Link href='/painting'>
                        <h3 className='footer__item'>Artworks</h3>
                    </Link>
                </li>
                <li className='footer__items'>
                    <Link href='/'>
                        <h3 className='footer__item'>Artists</h3>
                    </Link>
                </li>
                <li className='footer__items'>
                    <Link href='/'>
                        <h3 className='footer__item'>Contacts us</h3>
                    </Link>
                </li>
                <li className='footer__items'>
                    <Link href='/'>
                        <h3 className='footer__item'>My account</h3>
                    </Link>
                </li>
                <li className='footer__items'>
                    <h3 className='footer__item'>Secure payment</h3>
                </li>
            </ul>

            <div className='footer__newsletter'>
                <form onSubmit={(e) => submitNewsletter(e)}>
                    <h3 className='newsletter__title'>Let's stay in touch! We'll let you know about the latest sales and new releases!</h3>
                    <div className='newsletter__content'>
                        <input
                            label='Email'
                            name='email'
                            id='email'
                            type='mail'
                            placeholder='Enter your email'
                            className='newsletter__input'
                            onChange={(e) => setNewsletter({ ...newsletter, email: e.target.value })}
                        />
                        <button className='newsletter__button' type='submit'>Sign up</button>
                    </div>
                </form>
            </div>

            <div className='footer__copyright'>
                <p>Copyright ©2022. All Rights Reserved.</p>
            </div>
        </div>
    );
}

export default Footer;
