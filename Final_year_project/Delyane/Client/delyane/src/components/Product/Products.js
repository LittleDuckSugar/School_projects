import React, { useState, useEffect } from 'react';
import axios from 'axios';

import Header from '../Layout/Header/Header';
import Footer from '../Layout/Footer/Footer';
import ListItem from '../Caroussel/ListItem/ListItem';

import './Products.css';

const Product = () => {
    const [products, setProducts] = useState([]);
    const handleValue = (value) => {
        getDataCategory(value);
    };
    useEffect(() => {
        const getData = async () => {
            try {
                const result = await axios.get('http://90.22.250.124:8080/products')
                setProducts(result.data);
                console.log(result.data);
            } catch (err) {
                console.log(err)
            }
        }
        getData();
    }, []);
    const getDataCategory = (idCategory) => {
        if(idCategory){
            const getData = async () => {
                try {
                    const result = await axios.get(`http://90.22.250.124:8080/products?category=${idCategory}`)
                    setProducts(result.data);
                    console.log("here "+result.data);
                } catch (err) {
                    console.log(err);
                }
            }
            getData();
        }
    };
    return (
        <>
            <Header />
            <div className='product__main'>
                <div className='product__content'>
                    <h1 className='product__title'>Painting</h1>

                    <ul className='product__list'> 
                        <li className='product__filter'>
                            <div className='product__category'>
                                <select name="genre" id="genre" onChange={(e)=> handleValue(e.target.value) }>
                                    <option >Category</option>
                                    <option value="899a55e4-fdb8-4ebf-84c5-cf1781086f53" >Painting</option>
                                    <option value="4bdcbc09-3634-4a51-9979-88ba8da3c6b9" >Photography</option>
                                    <option value="460462ab-4446-454b-8249-edabdd55c2db" >Editing</option>
                                    <option value="a7c6488c-6c58-4b25-aabe-303121df72d7" >Sculpture</option>
                                    <option value="ae596731-14b9-466f-9509-d4be54fd8ef8" >Design</option>
                                </select>
                            </div>
                        </li>
                        {products ? (
                        <>
                            <li className='product__art'>
                                {products.map(product => {
                                    return (
                                        <ListItem title={product.title} description={product.description} category={product.category} price={product.price} image={product.image} uuid={product.uuid} />
                                    )
                                })}
                            </li>
                        </>
                        ) : (
                        ""
                        )}
                    </ul>
                </div>

            </div>
            <Footer />
        </>
    );
}

export default Product;

