import React, { useEffect, useState } from "react";
import Button from "../../components/Button";

const Index = () => {
  const [cart, setCart] = useState();

  const deleteCart = () => {
    localStorage.removeItem("cart");
    setCart(null);
  };

  useEffect(() => {
    setCart(JSON.parse(localStorage.getItem("cart")) || []);
  }, []);

  const decrementQty = (product) => {
    const indexOfExistingProduct = cart.findIndex((el) => el.id === product.id);
    if (indexOfExistingProduct !== -1 && cart[indexOfExistingProduct].quantity > 1) {
      cart[indexOfExistingProduct].quantity -= 1;
      localStorage.setItem("cart", JSON.stringify(cart));
      setCart(JSON.parse(localStorage.getItem('cart')));
    }
  };
  const incrementQty = (product) => {
    const indexOfExistingProduct = cart.findIndex((el) => el.id === product.id);
    if (indexOfExistingProduct !== -1) {
      cart[indexOfExistingProduct].quantity += 1;
    }
    localStorage.setItem("cart", JSON.stringify(cart));
    setCart(JSON.parse(localStorage.getItem('cart')));
  };

  const deleteProduct = (product) => {
    const filteredCart = cart.filter((item) => item.id != product.id);
    localStorage.setItem("cart", JSON.stringify(filteredCart));
    setCart(filteredCart);
  };

  const renderTotalAmount = () => {
    return (
      <p>Montant total : {cart.reduce((total, product) => total + (product.quantity * product.price), 0)} €</p>
    )
  }

  const renderTotalQty = () => {
    return cart.reduce((total, product) => total + product.quantity, 0)
  }

  return (
    <div className="page__cart">
      {cart ? (
        <>
          <p>Vous avez {renderTotalQty()} produits dans votre panier</p>
          <table>
            <thead>
              <tr>
                <th>Titre</th>
                <th>Prix</th>
                <th>Quantité</th>
                <th>Total</th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              {cart.map((cartItem) => (
                <tr key={cartItem.id}>
                  <td>{cartItem.title}</td>
                  <td>{cartItem.price}</td>
                  <td>
                    <button onClick={() => decrementQty(cartItem)}>-</button>
                    {cartItem.quantity}
                    <button onClick={() => incrementQty(cartItem)}>+</button>
                  </td>
                  <td>{(cartItem.price * cartItem.quantity).toFixed(2)}</td>
                  {/* .Filter() */}
                  <td>
                    <button onClick={() => deleteProduct(cartItem)}>Supprimer</button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
          <Button
            title="Supprimer le panier"
            classes="btn btn__color-white"
            type="button"
            function={deleteCart}
          />
          {renderTotalAmount()}
        </>
      ) : (
        <p className="text__center">Votre panier est vide</p>
      )}
    </div>
  );
};

export default Index;
