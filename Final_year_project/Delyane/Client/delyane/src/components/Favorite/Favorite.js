import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Header from '../Layout/Header/Header';
import Footer from '../Layout/Footer/Footer';
import ListItem from '../Caroussel/ListItem/ListItem';
import './Favorite.css';
const Favorite = () => {
    const [products, setProducts] = useState([]);
    const uuid_user = localStorage.getItem('uuid');
    useEffect(() => {
        const getData = async () => {
            try {
                const result = await axios.get(`http://90.22.250.124:8080/user/${uuid_user}/wishlist`)
                setProducts(result.data.products);
            } catch (err) {
                console.log(err);
            }
        }
        getData();
    }, [uuid_user]);
    return (
        <>
        <div className='favorite__main'>
            <Header />
            <div className='favorite__content'>
                <h1 className='favorite__title'>My favorites</h1>
                <p className='favorite__subtitle'>Find the works, artists and galleries you have followed. A completed and bigger wishlist will allow our experts to send you personalised suggestions.</p>

                <div className='favorite__diviser'></div>
                {products ? (
                <>
                    <li className='favorite__art'>
                        {products.map(product => {
                            return (
                                <ListItem title={product.title} description={product.description} category={product.category} price={product.price} image={product.image} uuid={product.uuid} />
                            )
                        })}
                    </li>
                </>
                ) : (
                <div className='favorite__container'>
                    <img src='../images/canape.png' alt='' className='container__picture' />
                    <h2 className='container__title'>Add artworks to favorites</h2>
                    <p className='container__subtitle'>You have not yet added any artworks to your favorites. To find your next favorite, explore our catalog.</p>
                    <button className='container__button'><a href="/painting">See all artworks</a></button>
                </div>
                )}
            </div>

            <Footer />
        </div>
        </>
    );
}

export default Favorite;
