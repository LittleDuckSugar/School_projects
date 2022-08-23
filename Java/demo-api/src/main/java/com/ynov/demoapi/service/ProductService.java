package com.ynov.demoapi.service;

import org.springframework.stereotype.Service;

import com.ynov.demoapi.model.Product;

@Service
public class ProductService {

	public Product createNewRandomProduct() {
		Product p = new Product();
		p.setName("random product");
		p.setDescription("Description of the product");
		p.setCost(10);

		return p;
	}
}
