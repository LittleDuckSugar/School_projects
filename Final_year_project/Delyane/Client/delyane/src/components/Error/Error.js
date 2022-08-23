import React from 'react';

import './Error.css';

const Error = () => {
    return (
        <div className='error__main'>
            <div className='error__content'>
                <h1>Error 404 - Page can't be found</h1>
                <p className='error__text'>We're sorry, the page you requested doesn't exist. Please make sure you typed in the correct URL!</p>
                <button className='error__button'>
                    <a href='/' className='error__redirection'>
                        Return to Home
                    </a>
                </button>
            </div>
        </div>
    );
}

export default Error;