package com.ynov.client.service;

import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.ynov.client.model.Product;
import com.ynov.client.repository.ProductProxy;

@Service
public class ProductService {

	@Autowired
	private ProductProxy productProxy;

	public List<Product> getProducts() {
		return productProxy.getProducts();
	}

	public Product getProductById(Integer id) {
		return productProxy.getProductById(id);
	}
	
	public void save(Product product) {
		productProxy.save(product);
	}

}
