import React from 'react';

import './Admin.css';

const Admin = () => {
    return (
        <div className='admin__main'>
            <div className='admin__container'>
                <div className='admin__image'></div>
                <div className='test2'>
                    <form className='admin__form'>
                        <h1 className='admin__title'>Delyane admin</h1>
                        <h2 className='admin__subtitle'>The Art is yours</h2>
                        <input
                            label='Username'
                            name='username'
                            id='username'
                            type='text'
                            placeholder='Enter your username'
                            className='admin__input'
                        />
                        <input
                            label='Password'
                            name='password'
                            id='password'
                            type='password'
                            placeholder='Enter your password'
                            className='admin__input'
                        />
                        <button className='admin__button' type='submit'>Log in</button>
                    </form>

                </div>
            </div>
        </div>
    );
}

export default Admin;
