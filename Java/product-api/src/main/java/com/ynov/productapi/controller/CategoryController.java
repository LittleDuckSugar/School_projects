package com.ynov.productapi.controller;

import java.util.List;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;

import com.ynov.productapi.model.Category;
import com.ynov.productapi.model.Product;
import com.ynov.productapi.service.CategoryService;
import com.ynov.productapi.service.ProductService;
import com.ynov.productapi.transformer.category.CategoryFull;

@RestController
public class CategoryController {

	@Autowired
	private CategoryService categoryService;

	@Autowired
	private ProductService productService;

	@GetMapping("/category")
	public List<CategoryFull> getCategories() {
		return categoryService.getCategories();
	}

	@PostMapping("/category/{idCategory}/{idProduct}")
	public void addProductToCategory(@PathVariable(name = "idCategory") Integer idCategory,
			@PathVariable(name = "idProduct") Integer idProduct) {
		Category category = categoryService.getCategory(idCategory).get();
		Product product = productService.getProduct(idProduct).get();
		
		category.addProduct(product);
		
		categoryService.saveCategory(category);
		
	}
	

}