import React, { useState } from 'react';
import axios from 'axios';
import Header from '../Layout/Header/Header';
import Footer from '../Layout/Footer/Footer';
import './Register.css';

const Register = () => {
    const [user, setUser] = useState({});

    const submitRegister = async (e) => {
        console.log(user);
        e.preventDefault();
        const url = 'http://90.22.250.124:8080/user';
        try {
            const result = await axios.post(url, user);
            const uuid= result.data.user.uuid;
            localStorage.setItem('uuid',uuid);
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
            <div className='register__main' style={bannerStyle}>
                <div className='register__content'>
                    <form className='register__form' onSubmit={(e) => submitRegister(e)}>
                        <h1 className='register__title'>Delyane</h1>
                        <h2 className='register__subtitle'>The Art is yours</h2>
                        <input
                            label='Username'
                            name='username'
                            id='username'
                            type='text'
                            placeholder='Enter your username'
                            className='register__input'
                            onChange={(e) => setUser({ ...user, username: e.target.value })}
                        />
                        <input
                            label='Email'
                            name='email'
                            id='email'
                            type='mail'
                            placeholder='Enter your email'
                            className='register__input'
                            onChange={(e) => setUser({ ...user, email: e.target.value })}
                        />
                        <input
                            label='Password'
                            name='password'
                            id='password'
                            type='password'
                            placeholder='Enter your password'
                            className='register__input'
                            onChange={(e) => setUser({ ...user, password: e.target.value })}

                        />
                        <button className='register__button' type='submit'>Register</button>
                        <p className='register__text'>Already customer? <a className='register__login' href='/authentication'>Log in</a></p>
                    </form>
                </div>
            </div>
            <Footer />
        </div>
    );
}

export default Register;