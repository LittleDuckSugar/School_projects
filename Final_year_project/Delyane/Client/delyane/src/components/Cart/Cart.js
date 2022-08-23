import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Header from '../Layout/Header/Header';
import Footer from '../Layout/Footer/Footer';
import ListItem from '../Caroussel/ListItem/ListItem';
import './Cart.css';
const Cart = () => {
    const [products, setProducts] = useState([]);
    const uuid_user = localStorage.getItem('uuid');
    useEffect(() => {
        const getData = async () => {
            try {
                const result = await axios.get(`http://90.22.250.124:8080/user/${uuid_user}/cart`)
                setProducts(result.data.products);
            } catch (err) {
                console.log(err);
            }
        }
        getData();
    }, [uuid_user]);
    return (
        <>
        <div className='cart__main'>
            <Header />
            <div className='cart__content'>
                <h1 className='cart__title'>My carts</h1>
                <p className='cart__subtitle'>Find the works, artists and galleries you have followed. A completed and bigger cart will allow our experts to send you personalised suggestions.</p>

                <div className='cart__diviser'></div>
                {products ? (
                <>
                    <li className='cart__art'>
                        {products.map(product => {
                            return (
                                <ListItem title={product.title} description={product.description} category={product.category} price={product.price} image={product.image} uuid={product.uuid} />
                            )
                        })}
                    </li>
                </>
                ) : (
                <div className='cart__container'>
                    <img src='../images/canape.png' alt='' className='container__picture' />
                    <h2 className='container__title'>Add artworks to carts</h2>
                    <p className='container__subtitle'>You have not yet added any artworks to your carts. To find your next cart, explore our catalog.</p>
                    <button className='container__button'><a href="/painting">See all artworks</a></button>
                </div>
                )}
            </div>

            <Footer />
        </div>
        </>
    );
}

export default Cart;
