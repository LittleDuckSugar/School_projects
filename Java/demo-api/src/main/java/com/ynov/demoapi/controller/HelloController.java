package com.ynov.demoapi.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HelloController {

	private Logger logger = LoggerFactory.getLogger(HelloController.class);

	@GetMapping("/helloworld")
	public String getHelloWorld() {
		return "Hello World!";
	}

	@GetMapping("/helloworld/{str}")
	public String getHelloWorldWithParam(@PathVariable("str") String str) {
		return "Get the value " + str;
	}

	@PostMapping("/helloworld")
	public void postHelloWorld(@RequestBody String str) {
		logger.info("Received : " + str);
	}

	@DeleteMapping("/helloworld/{id}")
	public void deleteHelloWorld(@PathVariable("id") Integer id) {
		logger.info("deleteHelloWorld endpoint called for id : " + id);
	}
}
