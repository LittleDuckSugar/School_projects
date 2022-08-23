package com.ynov.productapi.transformer.product;

import java.util.ArrayList;
import java.util.List;

import com.ynov.productapi.model.Comment;
import com.ynov.productapi.transformer.category.CategoryLight;

public class ProductFull extends ProductLight {

	private List<Comment> comments = new ArrayList<>();
	private List<CategoryLight> categories = new ArrayList<>();

	public List<Comment> getComments() {
		return comments;
	}

	public void setComments(List<Comment> comments) {
		this.comments = comments;
	}

	public List<CategoryLight> getCategories() {
		return categories;
	}

	public void setCategories(List<CategoryLight> categories) {
		this.categories = categories;
	}

}