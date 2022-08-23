import React from 'react';

const ProductPrice = (props) => {
    return (
        <div className="product__price">
            {props.price} {props.currency}
        </div>
    );
}

export default ProductPrice;
