package com.ynov.client.repository;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

import com.ynov.client.ApiProperties;
import com.ynov.client.model.Product;

@Component
public class ProductProxy {

	@Autowired
	private ApiProperties props;

	public List<Product> getProducts() {

		RestTemplate restTemplate = new RestTemplate();

		ResponseEntity<List<Product>> response = restTemplate.exchange(props.getUrl() + "/product", HttpMethod.GET,
				null, new ParameterizedTypeReference<List<Product>>() {
				});
		return response.getBody();
	}

	public Product getProductById(Integer id) {

		RestTemplate restTemplate = new RestTemplate();

		ResponseEntity<Product> response = restTemplate.exchange(props.getUrl() + "/product/" + id, HttpMethod.GET,
				null, Product.class);
		return response.getBody();
	}

	public void save(Product product) {
		RestTemplate restTemplate = new RestTemplate();

		HttpEntity<Product> request = new HttpEntity<Product>(product);

		ResponseEntity<Product> response = restTemplate.exchange(props.getUrl() + "/product", HttpMethod.POST, request,
				Product.class);
	}

}
