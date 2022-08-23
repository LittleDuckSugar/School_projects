package com.ynov.productapi.service;

import java.util.List;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.ynov.productapi.model.Category;
import com.ynov.productapi.repository.CategoryRepository;
import com.ynov.productapi.transformer.category.CategoryFull;
import com.ynov.productapi.transformer.category.CategoryTransformer;

@Service
public class CategoryService {

	@Autowired
	private CategoryTransformer categoryTransformer;

	@Autowired
	private CategoryRepository categoryRepository;

	public List<CategoryFull> getCategories() {
		return categoryTransformer.transform(categoryRepository.findAll());
	}

	public Optional<Category> getCategory(Integer idCategory) {
		return categoryRepository.findById(idCategory);
	}
	
	public void saveCategory(Category category) {
		categoryRepository.save(category);
	}

}