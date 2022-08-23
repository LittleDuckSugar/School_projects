package com.ynov.productapi.transformer.product;

import java.util.ArrayList;
import java.util.List;

import org.springframework.stereotype.Component;

import com.ynov.productapi.model.Category;
import com.ynov.productapi.model.Product;
import com.ynov.productapi.transformer.category.CategoryLight;

@Component
public class ProductTransformer {

	public ProductFull transform(Product product) {
		ProductFull productFull = new ProductFull();
		productFull.setId(product.getId());
		productFull.setName(product.getName());
		productFull.setDescription(product.getDescription());
		productFull.setCost(product.getCost());

		for (Category category : product.getCategories()) {
			CategoryLight categoryLight = new CategoryLight();
			categoryLight.setCategoryId(category.getCategoryId());
			categoryLight.setName(category.getName());
			productFull.getCategories().add(categoryLight);
		}

		productFull.setComments(product.getComments());

		return productFull;
	}

	public List<ProductFull> transform(Iterable<Product> products) {
		List<ProductFull> productsFull = new ArrayList<>();
		for (Product product : products) {
			productsFull.add(transform(product));
		}
		return productsFull;
	}

}