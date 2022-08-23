import React, { useState, useEffect } from 'react';
import { Link } from "react-router-dom";
import axios from 'axios';
import FavoriteIcon from '@mui/icons-material/Favorite';
import FavoriteBorderIcon from '@mui/icons-material/FavoriteBorder';

import './ListItem.css';

const ListItem = (props) => {
    const [favorite, setFavorite] = useState(false);
    const uuid_user = localStorage.getItem('uuid');
    useEffect(() => {
        const getData = async () => {
            try {
                const result = await axios.get(`http://90.22.250.124:8080/user/${uuid_user}/wishlist`)
                localStorage.setItem('wishlist',result.data.products[0].uuid);
            } catch (err) {
                console.log(err);
            }
        }
        getData();
    }, [uuid_user]);
    const addToList = (idProduct) => { 
        setFavorite(!favorite);
        console.log("id here : "+idProduct);
        localStorage.setItem('wishlist',idProduct);
        const wishlist = localStorage.getItem('wishlist');
        const newProduct={
            products : [wishlist]
        }
        if(idProduct){
        const getData = async () => {
            try {
                const result = await axios.put(`http://90.22.250.124:8080/user/${uuid_user}/wishlist`,newProduct)
            } catch (err) {
                console.log(err)
            }
        }
        getData();
        }
    };
    return (
        <div>
            <div className='listItem__main' key={props.uuid}>
                <div className='listItem__top'>
                    <button className='listItem__favorite' onClick={()=> addToList(props.uuid) }>
                        {favorite ? <FavoriteIcon className='listItem__icon' /> : <FavoriteBorderIcon className='listItem__icon' />}
                    </button>
                    <Link to={`/painting/${props.uuid}`}>
                        <img className='listItem__image' src={`http://90.22.250.124:8080${props.image}`} alt={props.title} />
                    </Link>
                </div>
                <div className='listItem__content'>
                    <p className='listItem__text'>{props.title}</p>
                    <p className='listItem__text listItem__style'>{props.price}€</p>
                    <p className='listItem__text'>{props.description}</p>
çç                </div>
  
  3          </div>
        </div>
    );
}

export default ListItem;
