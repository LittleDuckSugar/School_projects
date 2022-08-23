package com.ynov.productapi.service;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.ynov.productapi.model.Comment;
import com.ynov.productapi.repository.CommentRepository;

@Service
public class CommentService {

	@Autowired
	private CommentRepository commentRepository;

	public void deleteComment(int id) {
		commentRepository.deleteById(id);
	}

	public Optional<Comment> getComment(Integer id_comment) {
		return commentRepository.findById(id_comment);
	}

}