package com.ynov.productapi.repository;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import com.ynov.productapi.model.Category;

@Repository
public interface CategoryRepository extends CrudRepository<Category, Integer> {

}