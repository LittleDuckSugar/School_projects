import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Header from '../Layout/Header/Header';
import Footer from '../Layout/Footer/Footer';
import Button from '../Button/Button';
import './Profil.css';

const Profil = () => {
    const [user, setUser] = useState({});
    const uuid = localStorage.getItem('uuid');
    useEffect(() => {
        const getData = async () => {
            try {
                const result = await axios.get(`http://90.22.250.124:8080/user/${uuid}`)
                setUser(result.data);
                console.log(user);
            } catch (err) {
                console.log(err);
            }
        }
        getData();
    }, [uuid]);
    const logout=() =>{
        localStorage.removeItem('uuid');
        window.location = "/";
    }
    return (
        <div>
            <Header />
            <div className='profil__main'>
                <h1>My account</h1>
                <div className='profil__content'>
                    <div className='profil__image'>
                        <img src={`http://90.22.250.124:8080${user.image}`} alt={user.image} /> 
                    </div>
                    <div className='profil__info'>
                        <p className='profil__username'>{user.username}</p>
                        <p className='profil__email'>{user.email}</p>
                        <Button title="Log out" className="profil__button" function={(e)=> logout(e)}/>
                    </div>
                </div>
            </div>
            <Footer />
        </div>
    );
}

export default Profil;
