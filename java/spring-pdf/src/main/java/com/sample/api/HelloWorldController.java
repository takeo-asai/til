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

  @RequestMapping(value = "/openpdf", method = RequestMethod.GET)
  ResponseEntity<InputStreamResource> openpdf() {
    final HttpHeaders headers = new HttpHeaders();
    headers.setContentType(MediaType.APPLICATION_PDF);
    final InputStreamResource inputStreamResource = new InputStreamResource(OpenPdf.generate());
    return new ResponseEntity<>(inputStreamResource, headers, HttpStatus.OK);
  }

  @RequestMapping(value = "/pdfbox", method = RequestMethod.GET)
  ResponseEntity<InputStreamResource> pdfbox() {
    final HttpHeaders headers = new HttpHeaders();
    headers.setContentType(MediaType.APPLICATION_PDF);
    final InputStreamResource inputStreamResource = new InputStreamResource(PdfBox.generate());
    return new ResponseEntity<>(inputStreamResource, headers, HttpStatus.OK);
  }
}