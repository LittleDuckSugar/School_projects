import React from 'react';
import AwesomeSlider from 'react-awesome-slider';
import 'react-awesome-slider/dist/styles.css';

const Slider = ()=>{
    return(
    <AwesomeSlider className='swiper'>
        <div>
            <img className='swiper__picture' src="https://wallpaperaccess.com/full/6264342.jpg" alt="" />
        </div>
        <div>
            <img className='swiper__picture' src="https://wallpaperaccess.com/full/2010066.jpg" alt="" />
        </div>
        <div>
            <img className='swiper__picture' src="https://wallpaperaccess.com/full/524125.jpg" alt="" />
        </div>
    </AwesomeSlider>
    );
}
export default Slider;