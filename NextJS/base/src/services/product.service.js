/* eslint-disable import/no-anonymous-default-export */
const apiUrl = "http://localhost:1337/api";
export default {
    getProducts() {
        return fetch(`${apiUrl}/products?populate=image`)
            .then((res) => res.json())
    },
    
    getProduct(id) {
        return fetch(`${apiUrl}/products/${id}?populate=image`)
            .then((res) => res.json())
    }
}