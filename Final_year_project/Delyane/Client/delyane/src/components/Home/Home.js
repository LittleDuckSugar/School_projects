import React from 'react';

import Header from '../Layout/Header/Header';
import Footer from '../Layout/Footer/Footer';
import List from '../Caroussel/List/List';
import Slider from '../Caroussel/Slider/Slider';

import './Home.css';

const Home = () => {
    return (
        <div className='home__main'>
            <Header />
            {/* ----------- To do : Carrousel ------------  */}
            <div className='home__swiper'>
                <Slider/>
            </div>

            <div className='home__content'>
                <div className='home__bestsellers'>
                    <h1 className='home__title'>Bestsellers</h1>
                    <p className='home__subtitle'>Discover artworks our collectors love</p>
                    {/* ----------- To do : Carrousel ------------  */}
                    <div className='bestsellers__swiper'>
                        <List />
                    </div>
                </div>
                <div className='home__artists'>
                    <h1 className='home__title'>Featured Artists</h1>
                    <p className='home__subtitle'>The artists you should be keeping an eyes on</p>
                    <div className='main__product'>
                        <img className='first__picture' src='https://desenio.fr/bilder/artiklar/zoom/17002_2.jpg?imgwidth=435&qt=Pivoines%20roses%20abstraites' alt='' />
                        <img className='second__picture' src='https://desenio.fr/bilder/artiklar/zoom/16040_2.jpg?imgwidth=435&qt=Vase%20en%20or' alt='' />
                        <img className='third__picture' src='https://desenio.fr/bilder/artiklar/zoom/3802_2.jpg?imgwidth=435&qt=Femme%20%C3%A0%20Paris' alt='' />
                    </div>
                </div>
            </div>
            <Footer />
        </div>
    );
}

export default Home;
