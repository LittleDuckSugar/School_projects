import React, { useState, useRef, useEffect } from 'react';
import axios from 'axios';
import ListItem from "../ListItem/ListItem";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import './List.css';

const List = () => {
    const [isMoved, setIsMoved] = useState(false);
    const [slideNumber, setSlideNumber] = useState(0);

    const listRef = useRef();

    const handleClick = (direction) => {
        setIsMoved(true);
        let distance = listRef.current.getBoundingClientRect().x - 50;
        if (direction === "left" && slideNumber > 0) {
            setSlideNumber(slideNumber - 1);
            listRef.current.style.transform = `translateX(${230 + distance}px)`;
        }
        if (direction === "right" && slideNumber < 5) {
            setSlideNumber(slideNumber + 1);
            listRef.current.style.transform = `translateX(${-230 + distance}px)`;
        }
    };

    const [products, setProducts] = useState([]);

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

    return (
        <div className='list__main'>
            <div className="list__wrapper">
                <FontAwesomeIcon icon="fa-solid fa-chevron-left" className="sliderArrow left" onClick={() => handleClick("left")} style={{ display: !isMoved && "none" }} />
                <div className="list__container" ref={listRef}>
                    {products.map(product => {
                                return (
                                    <ListItem title={product.title} description={product.description} category={product.category} price={product.price} image={product.image} uuid={product.uuid} />
                                )
                            })}
                </div>
                <FontAwesomeIcon icon="fa-solid fa-chevron-right" className="sliderArrow right" onClick={() => handleClick("right")} />
            </div>
        </div>
    );
}

export default List;
