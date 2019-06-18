package com.sample.api;

import java.io.BufferedInputStream;
import java.io.ByteArrayInputStream;

import org.springframework.core.io.InputStreamResource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@RestController
class HelloWorldController {

  static final String HELLO_WORLD = "Hello, world.";
  private static final int BUFFER_SIZE = 1024;

  @RequestMapping(value = "/string", method = RequestMethod.GET)
  String byString() {
    return HELLO_WORLD;
  }

  @RequestMapping(value = "/byte", method = RequestMethod.GET)
  byte[] byBytes() {
    return HELLO_WORLD.getBytes();
  }

  @RequestMapping(value = "/stream", method = RequestMethod.GET)
  ResponseEntity<InputStreamResource> byStream() {
    final InputStreamResource inputStreamResource = new InputStreamResource(GeneratePdfReport.citiesReport());

    GeneratePdfReport.citiesReport();

    final HttpHeaders headers = new HttpHeaders();
    headers.setContentType(MediaType.APPLICATION_PDF);

    return new ResponseEntity<>(inputStreamResource, headers, HttpStatus.OK);
  }

}