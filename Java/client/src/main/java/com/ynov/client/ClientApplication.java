package com.ynov.client;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import com.ynov.client.model.Product;
import com.ynov.client.service.ProductService;

@SpringBootApplication
public class ClientApplication implements CommandLineRunner {

	@Autowired
	private ProductService productService;

	public static void main(String[] args) {
		SpringApplication.run(ClientApplication.class, args);
	}

	@Override
	public void run(String... args) throws Exception {
		List<Product> products = productService.getProducts();
		products.stream().forEach((product) -> System.out.println(product.getName()));
	}

}
