package com.ynov.demoapi.controller;

import static org.hamcrest.CoreMatchers.is;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.web.servlet.MockMvc;

@SpringBootTest
@AutoConfigureMockMvc
public class ProductControllerTest {

	@Autowired
	private MockMvc mockMvc;

	@Test
	public void testgetProduct() throws Exception {
		mockMvc.perform(get("/product/iphone")).andExpect(status().isOk()).andExpect(jsonPath("$.name", is("iphone")))
				.andExpect(jsonPath("$.description", is("Description of the product")))
				.andExpect(jsonPath("$.cost", is(Integer.valueOf(10))));
	}
}
