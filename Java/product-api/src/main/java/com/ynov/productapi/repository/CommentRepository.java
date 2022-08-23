package com.ynov.productapi.repository;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import com.ynov.productapi.model.Comment;

@Repository
public interface CommentRepository extends CrudRepository<Comment, Integer> {

}