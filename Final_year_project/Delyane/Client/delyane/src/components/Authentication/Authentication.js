import React, { useState } from 'react';
import axios from 'axios';
import Header from '../Layout/Header/Header';
import Footer from '../Layout/Footer/Footer';
import  { Redirect } from 'react-router-dom'
import './Authentication.css';

const Authentication = () => {
    const [user, setUser] = useState({});

    const submitRegister = async (e) => {
        console.log(user);
        e.preventDefault();
        const url = 'http://90.22.250.124:8080/user/login';
        try {
            const result = await axios.post(url, user);
            const uuid= result.data.user.uuid;
            localStorage.setItem('uuid',uuid);
            console.log("log in ok", user);
            window.location = "/painting";
        } catch (err) {
            console.log(err)
        }
    };

    const bannerStyle = {
        backgroundImage: 'url(https://desenio.fr/bilder/inspiration/5a9559030b21a.jpg)',
        backgroundPosition: 'center',
        backgroundSize: 'cover'
    }

    return (
        <div>
            <Header />
            <div className='authentication__main' style={bannerStyle}>
                <div className='authentication__content'>
                    <form className='authentication__form' onSubmit={(e) => submitRegister(e)}>
                        <h1 className='authentication__title'>Delyane</h1>
                        <h2 className='authentication__subtitle'>The Art is yours</h2>
                        <input
                            label='Username'
                            name='username'
                            id='username'
                            type='text'
                            placeholder='Enter your username'
                            className='authentication__input'
                            onChange={(e) => setUser({ ...user, identifier: e.target.value })}
                        />
                        <input
                            label='Password'
                            name='password'
                            id='password'
                            type='password'
                            placeholder='Enter your password'
                            className='authentication__input'
                            onChange={(e) => setUser({ ...user, password: e.target.value })}
                        />
                        <p className='authentication__password'>Forgot your password?</p>
                        <button className='authentication__button' type='submit'>Log in</button>
                        <p className='authentication__text'>New customer? <a className='authentication__login' href='/register'>Register</a></p>
                    </form>
                </div>
            </div>
            <Footer />
        </div>
    );
}

export default Authentication;
